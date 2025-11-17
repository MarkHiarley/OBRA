package services

import (
	"codxis-obras/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type EquipamentoDiarioService struct {
	connection *sql.DB
}

func NewEquipamentoDiarioService(connection *sql.DB) EquipamentoDiarioService {
	return EquipamentoDiarioService{
		connection: connection,
	}
}

func (s *EquipamentoDiarioService) Create(equipamento models.EquipamentoDiario) (int64, error) {
	var id int64

	query := `INSERT INTO equipamento_diario (obra_id, data, codigo, descricao, quantidade_utilizada, horas_uso, observacoes) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) 
              RETURNING id`

	err := s.connection.QueryRow(query,
		equipamento.ObraID.Int64,
		equipamento.Data.String,
		equipamento.Codigo.String,
		equipamento.Descricao.String,
		equipamento.QuantidadeUtilizada.Int64,
		equipamento.HorasUso.Float64,
		equipamento.Observacoes.String).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("erro ao criar equipamento: %w", err)
	}

	return id, nil
}

func (s *EquipamentoDiarioService) GetByDiarioId(diarioId int64) ([]models.EquipamentoDiario, error) {
	query := `SELECT id, obra_id, data, codigo, descricao, quantidade_utilizada, horas_uso, observacoes, created_at, updated_at 
	          FROM equipamento_diario WHERE obra_id = $1`

	rows, err := s.connection.Query(query, diarioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipamentos []models.EquipamentoDiario
	for rows.Next() {
		var equipamento models.EquipamentoDiario
		err = rows.Scan(
			&equipamento.ID,
			&equipamento.ObraID,
			&equipamento.Data,
			&equipamento.Codigo,
			&equipamento.Descricao,
			&equipamento.QuantidadeUtilizada,
			&equipamento.HorasUso,
			&equipamento.Observacoes,
			&equipamento.CreatedAt,
			&equipamento.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		equipamentos = append(equipamentos, equipamento)
	}

	return equipamentos, nil
}

func (s *EquipamentoDiarioService) Update(id int, equipamento models.EquipamentoDiario) (models.EquipamentoDiario, error) {
	query := `UPDATE equipamento_diario 
              SET codigo = $1, descricao = $2, quantidade_utilizada = $3, horas_uso = $4, 
                  observacoes = $5, updated_at = $6
              WHERE id = $7
              RETURNING id, obra_id, data, codigo, descricao, quantidade_utilizada, horas_uso, observacoes, created_at, updated_at`

	var updated models.EquipamentoDiario

	err := s.connection.QueryRowContext(context.Background(), query,
		equipamento.Codigo.String,
		equipamento.Descricao.String,
		equipamento.QuantidadeUtilizada.Int64,
		equipamento.HorasUso.Float64,
		equipamento.Observacoes.String,
		time.Now(),
		id,
	).Scan(
		&updated.ID,
		&updated.ObraID,
		&updated.Data,
		&updated.Codigo,
		&updated.Descricao,
		&updated.QuantidadeUtilizada,
		&updated.HorasUso,
		&updated.Observacoes,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if err != nil {
		return models.EquipamentoDiario{}, err
	}

	return updated, nil
}

func (s *EquipamentoDiarioService) Delete(id int) error {
	query := "DELETE FROM equipamento_diario WHERE id = $1"
	result, err := s.connection.ExecContext(context.Background(), query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhum equipamento encontrado com o ID fornecido")
	}

	return nil
}
