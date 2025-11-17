package services

import (
	"codxis-obras/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type MaterialDiarioService struct {
	connection *sql.DB
}

func NewMaterialDiarioService(connection *sql.DB) MaterialDiarioService {
	return MaterialDiarioService{
		connection: connection,
	}
}

func (s *MaterialDiarioService) Create(material models.MaterialDiario) (int64, error) {
	var id int64

	query := `INSERT INTO material_diario (obra_id, data, codigo, descricao, quantidade, unidade, fornecedor, valor_unitario, valor_total, observacoes) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
              RETURNING id`

	err := s.connection.QueryRow(query,
		material.ObraID.Int64,
		material.Data.String,
		material.Codigo.String,
		material.Descricao.String,
		material.Quantidade.Float64,
		material.Unidade.String,
		material.Fornecedor.String,
		material.ValorUnitario.Float64,
		material.ValorTotal.Float64,
		material.Observacoes.String).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("erro ao criar material: %w", err)
	}

	return id, nil
}

func (s *MaterialDiarioService) GetByDiarioId(diarioId int64) ([]models.MaterialDiario, error) {
	query := `SELECT id, obra_id, data, codigo, descricao, quantidade, unidade, fornecedor, valor_unitario, valor_total, observacoes, created_at, updated_at 
	          FROM material_diario WHERE obra_id = $1`

	rows, err := s.connection.Query(query, diarioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materiais []models.MaterialDiario
	for rows.Next() {
		var material models.MaterialDiario
		err = rows.Scan(
			&material.ID,
			&material.ObraID,
			&material.Data,
			&material.Codigo,
			&material.Descricao,
			&material.Quantidade,
			&material.Unidade,
			&material.Fornecedor,
			&material.ValorUnitario,
			&material.ValorTotal,
			&material.Observacoes,
			&material.CreatedAt,
			&material.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		materiais = append(materiais, material)
	}

	return materiais, nil
}

func (s *MaterialDiarioService) Update(id int, material models.MaterialDiario) (models.MaterialDiario, error) {
	query := `UPDATE material_diario 
              SET codigo = $1, descricao = $2, quantidade = $3, unidade = $4, fornecedor = $5,
                  valor_unitario = $6, valor_total = $7, observacoes = $8, updated_at = $9
              WHERE id = $10
              RETURNING id, obra_id, data, codigo, descricao, quantidade, unidade, fornecedor, valor_unitario, valor_total, observacoes, created_at, updated_at`

	var updated models.MaterialDiario

	err := s.connection.QueryRowContext(context.Background(), query,
		material.Codigo.String,
		material.Descricao.String,
		material.Quantidade.Float64,
		material.Unidade.String,
		material.Fornecedor.String,
		material.ValorUnitario.Float64,
		material.ValorTotal.Float64,
		material.Observacoes.String,
		time.Now(),
		id,
	).Scan(
		&updated.ID,
		&updated.ObraID,
		&updated.Data,
		&updated.Codigo,
		&updated.Descricao,
		&updated.Quantidade,
		&updated.Unidade,
		&updated.Fornecedor,
		&updated.ValorUnitario,
		&updated.ValorTotal,
		&updated.Observacoes,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if err != nil {
		return models.MaterialDiario{}, err
	}

	return updated, nil
}

func (s *MaterialDiarioService) Delete(id int) error {
	query := "DELETE FROM material_diario WHERE id = $1"
	result, err := s.connection.ExecContext(context.Background(), query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhum material encontrado com o ID fornecido")
	}

	return nil
}
