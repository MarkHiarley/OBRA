package services

import (
	"codxis-obras/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type EquipeDiarioService struct {
	connection *sql.DB
}

func NewEquipeDiarioService(connection *sql.DB) EquipeDiarioService {
	return EquipeDiarioService{
		connection: connection,
	}
}

func (s *EquipeDiarioService) Create(equipe models.EquipeDiario) (int64, error) {
	var id int64

	query := `INSERT INTO equipe_diario (diario_id, codigo, descricao, quantidade_utilizada, horas_trabalhadas, observacoes) 
              VALUES ($1, $2, $3, $4, $5, $6) 
              RETURNING id`

	err := s.connection.QueryRow(query,
		equipe.DiarioID.Int64,
		equipe.Codigo.String,
		equipe.Descricao.String,
		equipe.QuantidadeUtilizada.Int64,
		equipe.HorasTrabalhadas.Float64,
		equipe.Observacoes.String).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("erro ao criar equipe: %w", err)
	}

	return id, nil
}

func (s *EquipeDiarioService) GetByDiarioId(diarioId int64) ([]models.EquipeDiario, error) {
	query := `SELECT id, diario_id, codigo, descricao, quantidade_utilizada, horas_trabalhadas, observacoes, created_at, updated_at 
	          FROM equipe_diario WHERE diario_id = $1`

	rows, err := s.connection.Query(query, diarioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipes []models.EquipeDiario
	for rows.Next() {
		var equipe models.EquipeDiario
		err = rows.Scan(
			&equipe.ID,
			&equipe.DiarioID,
			&equipe.Codigo,
			&equipe.Descricao,
			&equipe.QuantidadeUtilizada,
			&equipe.HorasTrabalhadas,
			&equipe.Observacoes,
			&equipe.CreatedAt,
			&equipe.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		equipes = append(equipes, equipe)
	}

	return equipes, nil
}

func (s *EquipeDiarioService) Update(id int, equipe models.EquipeDiario) (models.EquipeDiario, error) {
	query := `UPDATE equipe_diario 
              SET codigo = $1, descricao = $2, quantidade_utilizada = $3, horas_trabalhadas = $4, 
                  observacoes = $5, updated_at = $6
              WHERE id = $7
              RETURNING id, diario_id, codigo, descricao, quantidade_utilizada, horas_trabalhadas, observacoes, created_at, updated_at`

	var updated models.EquipeDiario

	err := s.connection.QueryRowContext(context.Background(), query,
		equipe.Codigo.String,
		equipe.Descricao.String,
		equipe.QuantidadeUtilizada.Int64,
		equipe.HorasTrabalhadas.Float64,
		equipe.Observacoes.String,
		time.Now(),
		id,
	).Scan(
		&updated.ID,
		&updated.DiarioID,
		&updated.Codigo,
		&updated.Descricao,
		&updated.QuantidadeUtilizada,
		&updated.HorasTrabalhadas,
		&updated.Observacoes,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if err != nil {
		return models.EquipeDiario{}, err
	}

	return updated, nil
}

func (s *EquipeDiarioService) Delete(id int) error {
	query := "DELETE FROM equipe_diario WHERE id = $1"
	result, err := s.connection.ExecContext(context.Background(), query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhuma equipe encontrada com o ID fornecido")
	}

	return nil
}
