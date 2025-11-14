package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
	"time"
)

type DiarioMetadadosService struct {
	connection *sql.DB
}

func NewDiarioMetadadosService(connection *sql.DB) DiarioMetadadosService {
	return DiarioMetadadosService{
		connection: connection,
	}
}

// CreateMetadados cria ou atualiza metadados do diário
func (dms *DiarioMetadadosService) CreateMetadados(metadados models.DiarioMetadados) (models.DiarioMetadados, error) {
	query := `
		INSERT INTO diario_metadados (obra_id, data, periodo, foto, observacoes, responsavel_id, aprovado_por_id, status_aprovacao)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (obra_id, data, periodo) 
		DO UPDATE SET 
			foto = EXCLUDED.foto,
			observacoes = EXCLUDED.observacoes,
			responsavel_id = EXCLUDED.responsavel_id,
			aprovado_por_id = EXCLUDED.aprovado_por_id,
			status_aprovacao = EXCLUDED.status_aprovacao,
			updated_at = NOW()
		RETURNING id, created_at, updated_at
	`

	err := dms.connection.QueryRow(
		query,
		metadados.ObraID,
		metadados.Data,
		metadados.Periodo,
		metadados.Foto,
		metadados.Observacoes,
		metadados.ResponsavelID,
		metadados.AprovadoPorID,
		metadados.StatusAprovacao,
	).Scan(&metadados.ID, &metadados.CreatedAt, &metadados.UpdatedAt)

	if err != nil {
		return models.DiarioMetadados{}, fmt.Errorf("erro ao criar/atualizar metadados do diário: %v", err)
	}

	return metadados, nil
}

// GetMetadadosByObraData busca metadados por obra e data
func (dms *DiarioMetadadosService) GetMetadadosByObraData(obraID int, data string) ([]models.DiarioMetadados, error) {
	query := `
		SELECT id, obra_id, data, periodo, foto, observacoes, responsavel_id, aprovado_por_id, status_aprovacao, created_at, updated_at
		FROM diario_metadados
		WHERE obra_id = $1 AND data = $2
		ORDER BY periodo
	`

	rows, err := dms.connection.Query(query, obraID, data)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar metadados: %v", err)
	}
	defer rows.Close()

	var metadados []models.DiarioMetadados
	for rows.Next() {
		var m models.DiarioMetadados
		err := rows.Scan(
			&m.ID, &m.ObraID, &m.Data, &m.Periodo, &m.Foto, &m.Observacoes,
			&m.ResponsavelID, &m.AprovadoPorID, &m.StatusAprovacao,
			&m.CreatedAt, &m.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear metadados: %v", err)
		}
		metadados = append(metadados, m)
	}

	return metadados, nil
}

// UpdateMetadados atualiza metadados existentes
func (dms *DiarioMetadadosService) UpdateMetadados(id int, metadados models.DiarioMetadados) error {
	query := `
		UPDATE diario_metadados 
		SET foto = $2, observacoes = $3, responsavel_id = $4, aprovado_por_id = $5, 
		    status_aprovacao = $6, updated_at = $7
		WHERE id = $1
	`

	result, err := dms.connection.Exec(
		query, id, metadados.Foto, metadados.Observacoes,
		metadados.ResponsavelID, metadados.AprovadoPorID,
		metadados.StatusAprovacao, time.Now(),
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar metadados: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("metadados não encontrados")
	}

	return nil
}

// DeleteMetadados remove metadados
func (dms *DiarioMetadadosService) DeleteMetadados(id int) error {
	query := `DELETE FROM diario_metadados WHERE id = $1`

	result, err := dms.connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar metadados: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("metadados não encontrados")
	}

	return nil
}
