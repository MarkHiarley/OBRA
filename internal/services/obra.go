package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
)

type ObraServices struct {
	connection *sql.DB
}

func NewObraService(connection *sql.DB) ObraServices {
	return ObraServices{
		connection: connection,
	}
}

func (pr *ObraServices) CreateObra(obra models.Obra) (int, error) {
	var id int

	query := `INSERT INTO obra (nome, contrato_numero, contratante_id, responsavel_id, data_inicio, prazo_dias, data_fim_prevista, orcamento, status, endereco_rua, endereco_numero, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, observacoes) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
              RETURNING id`

	err := pr.connection.QueryRow(query,
		obra.Nome,
		obra.ContratoNumero,
		obra.ContratanteID,
		obra.ResponsavelID,
		obra.DataInicio,
		obra.PrazoDias,
		obra.DataFimPrevista,
		obra.Orcamento,
		obra.Status,
		obra.EnderecoRua,
		obra.EnderecoNumero,
		obra.EnderecoBairro,
		obra.EnderecoCidade,
		obra.EnderecoEstado,
		obra.EnderecoCep,
		obra.Observacoes).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar Obra: %v\n", err)
		return 0, err
	}

	return id, nil
}
