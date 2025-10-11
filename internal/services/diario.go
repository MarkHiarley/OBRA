package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
)

type DiarioServices struct {
	connection *sql.DB
}

func NewDiarioService(connection *sql.DB) DiarioServices {
	return DiarioServices{
		connection: connection,
	}
}

func (pr *DiarioServices) CreateDiario(diario models.DiarioObra) (int, error) {
	var id int

	query := `INSERT INTO diario_obra (obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id, aprovado_por_id, status_aprovacao) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
              RETURNING id`

	err := pr.connection.QueryRow(query,
		diario.ObraID,
		diario.Data,
		diario.Periodo,
		diario.AtividadesRealizadas,
		diario.Ocorrencias,
		diario.Observacoes,
		diario.ResponsavelID,
		diario.AprovadoPorID,
		diario.StatusAprovacao).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar diario: %v\n", err)
		return 0, err
	}

	return id, nil
}
