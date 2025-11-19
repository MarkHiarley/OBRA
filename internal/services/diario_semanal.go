package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/guregu/null.v4"
)

type DiarioSemanalService struct {
	connection *sql.DB
}

// DiarioData representa dados temporários de um diário para agrupamento
type DiarioData struct {
	Data                 string
	AtividadesRealizadas null.String
	Observacoes          null.String
}

func NewDiarioSemanalService(connection *sql.DB) DiarioSemanalService {
	return DiarioSemanalService{
		connection: connection,
	}
}

// GetDiarioSemanal busca os diários de uma obra no período especificado e agrupa por semana
func (ds *DiarioSemanalService) GetDiarioSemanal(obraID int64, dataInicio, dataFim string) (models.DiarioSemanalResponse, error) {
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
			COALESCE(p_contratante.nome, 'Não informado') as contratante,
			COALESCE(o.contratada, 'Não informado') as contratada
		FROM obra o
		LEFT JOIN pessoa p_contratante ON o.contratante_id = p_contratante.id
		WHERE o.id = $1
	`

	var nomeObra, localizacao, contratante, contratada string
	var contratoNumero null.String

	err := ds.connection.QueryRow(queryObra, obraID).Scan(
		&nomeObra,
		&localizacao,
		&contratoNumero,
		&contratante,
		&contratada,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.DiarioSemanalResponse{}, fmt.Errorf("obra não encontrada")
		}
		return models.DiarioSemanalResponse{}, fmt.Errorf("erro ao buscar dados da obra: %v", err)
	}

	// 2. Buscar diários no período
	queryDiarios := `
		SELECT 
			data,
			atividades_realizadas,
			observacoes
		FROM diario_obra
		WHERE obra_id = $1 
		  AND data >= $2 
		  AND data <= $3
		ORDER BY data ASC
	`

	rows, err := ds.connection.Query(queryDiarios, obraID, dataInicio, dataFim)
	if err != nil {
		return models.DiarioSemanalResponse{}, fmt.Errorf("erro ao buscar diários: %v", err)
	}
	defer rows.Close()

	// Coletar diários do banco
	var diarios []DiarioData
	for rows.Next() {
		var d DiarioData
		err := rows.Scan(&d.Data, &d.AtividadesRealizadas, &d.Observacoes)
		if err != nil {
			continue
		}
		diarios = append(diarios, d)
	}

	// 3. Agrupar por semana
	semanas := ds.agruparPorSemana(diarios, dataInicio, dataFim)

	// 4. Montar resposta
	response := models.DiarioSemanalResponse{
		DadosObra: models.DadosObraDiario{
			NomeObra:       nomeObra,
			Localizacao:    localizacao,
			ContratoNumero: contratoNumero,
			Contratante:    null.StringFrom(contratante),
			Contratada:     null.StringFrom(contratada),
		},
		Semanas: semanas,
	}

	return response, nil
}

// agruparPorSemana agrupa os diários por semana
func (ds *DiarioSemanalService) agruparPorSemana(diarios []DiarioData, dataInicio, dataFim string) []models.SemanaDiario {
	// Parse das datas de início e fim
	inicio, err := time.Parse("2006-01-02", dataInicio)
	if err != nil {
		return []models.SemanaDiario{}
	}

	fim, err := time.Parse("2006-01-02", dataFim)
	if err != nil {
		return []models.SemanaDiario{}
	}

	var semanas []models.SemanaDiario
	numeroSemana := 1

	// Iterar semana por semana
	for inicio.Before(fim) || inicio.Equal(fim) {
		// Calcular fim da semana (domingo)
		fimSemana := inicio.AddDate(0, 0, 6)
		if fimSemana.After(fim) {
			fimSemana = fim
		}

		// Buscar diários desta semana
		var diasTrabalho []string

		for _, diario := range diarios {
			dataDiario, err := time.Parse("2006-01-02", diario.Data)
			if err != nil {
				continue
			}

			// Se o diário está dentro desta semana
			if (dataDiario.After(inicio) || dataDiario.Equal(inicio)) &&
				(dataDiario.Before(fimSemana) || dataDiario.Equal(fimSemana)) {

				diasTrabalho = append(diasTrabalho, diario.Data)
			}
		}

		// Criar objeto da semana
		semana := models.SemanaDiario{
			Numero:       numeroSemana,
			DataInicio:   inicio.Format("2006-01-02"),
			DataFim:      fimSemana.Format("2006-01-02"),
			DiasTrabalho: diasTrabalho,
			Descricao:    null.String{}, // Campo vazio para o usuário preencher
		}

		semanas = append(semanas, semana)

		// Avançar para próxima semana
		inicio = fimSemana.AddDate(0, 0, 1)
		numeroSemana++
	}

	return semanas
}
