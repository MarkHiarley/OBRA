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
			o.descricao
		FROM obra o
		WHERE o.id = $1
	`

	var nomeObra, localizacao string
	var contratoNumero, lote, descricao null.String

	err := rs.connection.QueryRow(queryObra, obraID).Scan(
		&nomeObra,
		&localizacao,
		&contratoNumero,
		&lote,
		&descricao,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.RelatorioFotografico{}, fmt.Errorf("obra não encontrada")
		}
		return models.RelatorioFotografico{}, fmt.Errorf("erro ao buscar dados da obra: %v", err)
	}

	// 2. Buscar todas as fotos da obra (de diários)
	queryFotos := `
		SELECT 
			d.id,
			d.foto,
			d.data,
			d.periodo,
			d.observacoes
		FROM diario_obra d
		WHERE d.obra_id = $1 
		  AND d.foto IS NOT NULL 
		  AND d.foto != ''
		ORDER BY d.data DESC, d.created_at DESC
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
		var data, periodo, observacao null.String

		err := rows.Scan(&id, &foto, &data, &periodo, &observacao)
		if err != nil {
			continue // Ignora erros de scan individual
		}

		// Criar título/legenda baseado no período
		var tituloLegenda null.String
		if periodo.Valid {
			tituloLegenda = null.StringFrom(fmt.Sprintf("Foto do período: %s", periodo.String))
		} else {
			tituloLegenda = null.StringFrom("Foto da obra")
		}

		fotoDetalhada := models.FotoDetalhada{
			ID:            id,
			URL:           foto,
			TituloLegenda: tituloLegenda,
			Data:          data,
			Observacao:    observacao,
			Categoria:     null.StringFrom("DIARIO"),
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
			InformacoesGerais: null.StringFrom("Relatório fotográfico da execução da obra"),
		},
		Fotos: fotos,
	}

	return relatorio, nil
}
