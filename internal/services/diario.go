package services

import (
	"codxis-obras/internal/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
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

	query := `INSERT INTO diario_obra (obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id, aprovado_por_id, status_aprovacao, foto) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
              RETURNING id`

	// Converter null.Int para interface{} tratando NULL corretamente
	var responsavelID interface{}
	if diario.ResponsavelID.Valid {
		responsavelID = diario.ResponsavelID.Int64
	} else {
		responsavelID = nil
	}

	var aprovadoPorID interface{}
	if diario.AprovadoPorID.Valid {
		aprovadoPorID = diario.AprovadoPorID.Int64
	} else {
		aprovadoPorID = nil
	}

	var foto interface{}
	if diario.Foto.Valid {
		foto = diario.Foto.String
	} else {
		foto = nil
	}

	err := pr.connection.QueryRow(query,
		diario.ObraID.Int64,
		diario.Data.String,
		diario.Periodo.String,
		diario.AtividadesRealizadas.String,
		diario.Ocorrencias.String,
		diario.Observacoes.String,
		responsavelID,
		aprovadoPorID,
		diario.StatusAprovacao.String,
		foto).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar diario: %v\n", err)
		return 0, err
	}

	return int64(id), nil
}

func (pr *DiarioServices) GetDiarios() ([]models.DiarioObra, error) {

	query := "select id, obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id, aprovado_por_id, status_aprovacao, created_at, update_at, foto from diario_obra"
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
			&diariosObj.Foto,
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
	query := "select id, obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id,aprovado_por_id, status_aprovacao, created_at, update_at, foto from diario_obra where id = $1"

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
		&diario.Foto,
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

	query := "select id, obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id, aprovado_por_id, status_aprovacao, created_at, update_at, foto from diario_obra where obra_id = $1"
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
			&diariosObj.Foto,
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

func (pr DiarioServices) PutDiarios(id int, diarioToUpdate models.DiarioObra) (models.DiarioObra, error) {

	query := `
        UPDATE diario_obra 
        SET 
            obra_id = $1,
            data = $2,
            periodo = $3,
            atividades_realizadas = $4,
            ocorrencias = $5,
            observacoes = $6,
            responsavel_id = $7,
            aprovado_por_id = $8,
            status_aprovacao = $9,
			update_at = $10,
			foto = $11
		WHERE id = $12
		RETURNING id, obra_id, data, periodo, atividades_realizadas, ocorrencias, observacoes, responsavel_id, aprovado_por_id, status_aprovacao, created_at, update_at, foto`

	var updatedDiario models.DiarioObra

	// Converter null.Int para interface{} tratando NULL corretamente
	var responsavelID interface{}
	if diarioToUpdate.ResponsavelID.Valid {
		responsavelID = diarioToUpdate.ResponsavelID.Int64
	} else {
		responsavelID = nil
	}

	var aprovadoPorID interface{}
	if diarioToUpdate.AprovadoPorID.Valid {
		aprovadoPorID = diarioToUpdate.AprovadoPorID.Int64
	} else {
		aprovadoPorID = nil
	}

	var foto interface{}
	if diarioToUpdate.Foto.Valid {
		foto = diarioToUpdate.Foto.String
	} else {
		foto = nil
	}

	err := pr.connection.QueryRowContext(context.Background(), query,
		diarioToUpdate.ObraID,
		diarioToUpdate.Data,
		diarioToUpdate.Periodo,
		diarioToUpdate.AtividadesRealizadas,
		diarioToUpdate.Ocorrencias,
		diarioToUpdate.Observacoes,
		responsavelID,
		aprovadoPorID,
		diarioToUpdate.StatusAprovacao,
		time.Now(),
		foto,
		id, // The ID for the WHERE clause
	).Scan(

		&updatedDiario.ID,
		&updatedDiario.ObraID,
		&updatedDiario.Data,
		&updatedDiario.Periodo,
		&updatedDiario.AtividadesRealizadas,
		&updatedDiario.Ocorrencias,
		&updatedDiario.Observacoes,
		&updatedDiario.ResponsavelID,
		&updatedDiario.AprovadoPorID,
		&updatedDiario.StatusAprovacao,
		&updatedDiario.CreatedAt,
		&updatedDiario.UpdatedAt,
		&updatedDiario.Foto,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return models.DiarioObra{}, err
		}

		log.Printf("Error updating user: %v\n", err)
		return models.DiarioObra{}, fmt.Errorf("não foi possivel atualizar esse diario: %w", err)
	}

	return updatedDiario, nil
}

func (pr *DiarioServices) DeleteDiarioById(id int) error {
	query := "DELETE FROM diario_obra WHERE id = $1"

	result, err := pr.connection.ExecContext(context.Background(), query, id)
	if err != nil {

		return fmt.Errorf("erro ao executar a query de delete")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhum diario encontrado com o ID fornecido")
	}

	return nil
}
