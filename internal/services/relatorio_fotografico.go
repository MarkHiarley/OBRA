package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"

	"gopkg.in/guregu/null.v4"
)

type RelatorioFotograficoService struct {
	connection *sql.DB
}

func NewRelatorioFotograficoService(connection *sql.DB) RelatorioFotograficoService {
	return RelatorioFotograficoService{
		connection: connection,
	}
}

// GetRelatorioFotografico busca todas as fotos da obra com informações básicas
// Retorna apenas: cabeçalho + resumo da obra + lista de fotos
func (rs *RelatorioFotograficoService) GetRelatorioFotografico(obraID int64) (models.RelatorioFotografico, error) {
	// 1. Buscar informações da obra
	queryObra := `
		SELECT 
			o.nome,
			COALESCE(
				CONCAT(o.endereco_rua, ', ', o.endereco_numero, ' - ', o.endereco_bairro, ' - ', o.endereco_cidade, ' - ', o.endereco_estado),
				CONCAT(o.endereco_cidade, ' - ', o.endereco_estado),
				'Localização não informada'
			) as localizacao,
			o.contrato_numero,
			o.lote,
			o.descricao,
			o.foto
		FROM obra o
		WHERE o.id = $1
	`

	var nomeObra, localizacao string
	var contratoNumero, lote, descricao, fotoObra null.String

	err := rs.connection.QueryRow(queryObra, obraID).Scan(
		&nomeObra,
		&localizacao,
		&contratoNumero,
		&lote,
		&descricao,
		&fotoObra,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.RelatorioFotografico{}, fmt.Errorf("obra não encontrada")
		}
		return models.RelatorioFotografico{}, fmt.Errorf("erro ao buscar dados da obra: %v", err)
	}

	// 2. Buscar todas as fotos da obra (da tabela foto_diario)
	// Busca fotos associadas a diários da obra através de JOIN
	queryFotos := `
		SELECT 
			fd.id,
			fd.foto,
			fd.descricao,
			fd.categoria,
			fd.ordem,
			d.data,
			d.periodo
		FROM foto_diario fd
		INNER JOIN diario_obra d ON fd.entidade_id = d.id
		WHERE d.obra_id = $1 
		  AND fd.entidade_tipo = 'atividade'
		  AND fd.foto IS NOT NULL 
		  AND fd.foto != ''
		ORDER BY d.data DESC, fd.ordem ASC, fd.id DESC
	`

	rows, err := rs.connection.Query(queryFotos, obraID)
	if err != nil {
		return models.RelatorioFotografico{}, fmt.Errorf("erro ao buscar fotos: %v", err)
	}
	defer rows.Close()

	var fotos []models.FotoDetalhada
	for rows.Next() {
		var id int64
		var foto string
		var descricao, categoria null.String
		var ordem null.Int
		var data, periodo null.String

		err := rows.Scan(&id, &foto, &descricao, &categoria, &ordem, &data, &periodo)
		if err != nil {
			continue // Ignora erros de scan individual
		}

		// Usar a descrição da foto, ou criar título baseado no período
		var tituloLegenda null.String
		if descricao.Valid && descricao.String != "" {
			tituloLegenda = descricao
		} else if periodo.Valid {
			tituloLegenda = null.StringFrom(fmt.Sprintf("Foto do período: %s", periodo.String))
		} else {
			tituloLegenda = null.StringFrom("Foto da obra")
		}

		fotoDetalhada := models.FotoDetalhada{
			ID:            id,
			URL:           foto,
			TituloLegenda: tituloLegenda,
			Data:          data,
			Observacao:    null.String{}, // foto_diario não tem observação separada
			Categoria:     categoria,
		}

		fotos = append(fotos, fotoDetalhada)
	}

	// 3. Montar o relatório fotográfico
	relatorio := models.RelatorioFotografico{
		CabecalhoEmpresa: models.CabecalhoEmpresa{
			NomeEmpresa: "EMPRESA CONSTRUTORA", // Pode ser configurável
			Logotipo:    null.String{},         // Pode ser buscado de configuração
		},
		ResumoObra: models.ResumoObra{
			NomeObra:          nomeObra,
			Localizacao:       localizacao,
			ContratoNumero:    contratoNumero,
			Lote:              lote,
			DescricaoBreve:    descricao,
			FotoObra:          fotoObra,
			InformacoesGerais: null.StringFrom("Relatório fotográfico da execução da obra"),
		},
		Fotos: fotos,
	}

	return relatorio, nil
}
