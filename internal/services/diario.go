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

func (pr *DiarioServices) CreateDiario(diario models.DiarioObra) (int64, error) {
	var id int

	query := `INSERT INTO diario_obra (obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id, aprovado_por_id, status_aprovacao) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
              RETURNING id`

	err := pr.connection.QueryRow(query,
		diario.ObraID.Int64,
		diario.Data.String,
		diario.Periodo.String,
		diario.AtividadesRealizadas.String,
		diario.Ocorrencias.String,
		diario.Observacoes.String,
		diario.ResponsavelID.Int64,
		diario.AprovadoPorID.Int64,
		diario.StatusAprovacao.String).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar diario: %v\n", err)
		return 0, err
	}

	return int64(id), nil
}

func (pr *DiarioServices) GetDiarios() ([]models.DiarioObra, error) {

	query := "select id, obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id, aprovado_por_id, status_aprovacao, created_at, update_at from diario_obra"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.DiarioObra{}, err
	}

	var diarioList []models.DiarioObra
	var diariosObj models.DiarioObra

	for rows.Next() {
		err = rows.Scan(
			&diariosObj.ID,
			&diariosObj.ObraID,
			&diariosObj.Data,
			&diariosObj.Periodo,
			&diariosObj.AtividadesRealizadas,
			&diariosObj.Ocorrencias,
			&diariosObj.Observacoes,
			&diariosObj.ResponsavelID,
			&diariosObj.AprovadoPorID,
			&diariosObj.StatusAprovacao,
			&diariosObj.CreatedAt,
			&diariosObj.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []models.DiarioObra{}, err
		}

		diarioList = append(diarioList, diariosObj)

	}

	rows.Close()
	return diarioList, nil
}

func (pr DiarioServices) GetDiarioById(id int64) (models.DiarioObra, error) {

	//id, nome, email, tipo_documento, documento, telefone, perfil_acesso, ativo, created_at, updated_at
	query := "select id, obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id,aprovado_por_id, status_aprovacao, created_at, update_at from diario_obra where id = $1"

	row := pr.connection.QueryRow(query, id)

	var diario models.DiarioObra

	err := row.Scan(
		&diario.ID,
		&diario.ObraID,
		&diario.Data,
		&diario.Periodo,
		&diario.AtividadesRealizadas,
		&diario.Ocorrencias,
		&diario.Observacoes,
		&diario.ResponsavelID,
		&diario.AprovadoPorID,
		&diario.StatusAprovacao,
		&diario.CreatedAt,
		&diario.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {

			return models.DiarioObra{}, fmt.Errorf("diário não encontrado")
		}

		return models.DiarioObra{}, err
	}

	return diario, nil

}

func (pr DiarioServices) GetDiarioByObraId(id int64) ([]models.DiarioObra, error) {

	query := "select id, obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id, aprovado_por_id, status_aprovacao, created_at, update_at from diario_obra where obra_id = $1"
	rows, err := pr.connection.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return []models.DiarioObra{}, err
	}

	var diarioList []models.DiarioObra
	var diariosObj models.DiarioObra

	for rows.Next() {
		err = rows.Scan(
			&diariosObj.ID,
			&diariosObj.ObraID,
			&diariosObj.Data,
			&diariosObj.Periodo,
			&diariosObj.AtividadesRealizadas,
			&diariosObj.Ocorrencias,
			&diariosObj.Observacoes,
			&diariosObj.ResponsavelID,
			&diariosObj.AprovadoPorID,
			&diariosObj.StatusAprovacao,
			&diariosObj.CreatedAt,
			&diariosObj.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []models.DiarioObra{}, err
		}

		diarioList = append(diarioList, diariosObj)

	}

	rows.Close()
	return diarioList, nil

}
