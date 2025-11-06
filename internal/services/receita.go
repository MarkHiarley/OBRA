package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
)

type ReceitaService struct {
	connection *sql.DB
}

func NewReceitaService(connection *sql.DB) ReceitaService {
	return ReceitaService{
		connection: connection,
	}
}

func (rs *ReceitaService) CreateReceita(receita models.Receita) (models.Receita, error) {
	query := `
		INSERT INTO receitas (obra_id, descricao, valor, data, fonte_receita, numero_documento, responsavel_id, observacao)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

	err := rs.connection.QueryRow(
		query,
		receita.ObraID,
		receita.Descricao,
		receita.Valor,
		receita.Data,
		receita.FonteReceita,
		receita.NumeroDocumento,
		receita.ResponsavelID,
		receita.Observacao,
	).Scan(&receita.ID, &receita.CreatedAt, &receita.UpdatedAt)

	if err != nil {
		return models.Receita{}, fmt.Errorf("erro ao criar receita: %v", err)
	}

	return receita, nil
}

func (rs *ReceitaService) GetReceitas() ([]models.ReceitaComRelacionamentos, error) {
	query := `
		SELECT r.id, r.obra_id, r.descricao, r.valor, r.data, r.fonte_receita, 
		       r.numero_documento, r.responsavel_id, r.observacao, r.created_at, r.updated_at,
		       o.nome as obra_nome, p.nome as responsavel_nome
		FROM receitas r
		LEFT JOIN obra o ON r.obra_id = o.id
		LEFT JOIN pessoa p ON r.responsavel_id = p.id
		ORDER BY r.created_at DESC
	`

	rows, err := rs.connection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar receitas: %v", err)
	}
	defer rows.Close()

	var receitas []models.ReceitaComRelacionamentos
	for rows.Next() {
		var receita models.ReceitaComRelacionamentos
		err := rows.Scan(
			&receita.ID, &receita.ObraID, &receita.Descricao, &receita.Valor,
			&receita.Data, &receita.FonteReceita, &receita.NumeroDocumento,
			&receita.ResponsavelID, &receita.Observacao, &receita.CreatedAt,
			&receita.UpdatedAt, &receita.ObraNome, &receita.ResponsavelNome,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear receita: %v", err)
		}
		receitas = append(receitas, receita)
	}

	return receitas, nil
}

func (rs *ReceitaService) GetReceitaById(id int) (models.ReceitaComRelacionamentos, error) {
	query := `
		SELECT r.id, r.obra_id, r.descricao, r.valor, r.data, r.fonte_receita, 
		       r.numero_documento, r.responsavel_id, r.observacao, r.created_at, r.updated_at,
		       o.nome as obra_nome, p.nome as responsavel_nome
		FROM receitas r
		LEFT JOIN obra o ON r.obra_id = o.id
		LEFT JOIN pessoa p ON r.responsavel_id = p.id
		WHERE r.id = $1
	`

	var receita models.ReceitaComRelacionamentos
	err := rs.connection.QueryRow(query, id).Scan(
		&receita.ID, &receita.ObraID, &receita.Descricao, &receita.Valor,
		&receita.Data, &receita.FonteReceita, &receita.NumeroDocumento,
		&receita.ResponsavelID, &receita.Observacao, &receita.CreatedAt,
		&receita.UpdatedAt, &receita.ObraNome, &receita.ResponsavelNome,
	)

	if err != nil {
		return models.ReceitaComRelacionamentos{}, fmt.Errorf("erro ao buscar receita: %v", err)
	}

	return receita, nil
}

func (rs *ReceitaService) UpdateReceita(id int, receita models.Receita) error {
	query := `
		UPDATE receitas 
		SET obra_id = $2, descricao = $3, valor = $4, data = $5, fonte_receita = $6,
		    numero_documento = $7, responsavel_id = $8, observacao = $9, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`

	result, err := rs.connection.Exec(
		query, id, receita.ObraID, receita.Descricao, receita.Valor, receita.Data,
		receita.FonteReceita, receita.NumeroDocumento, receita.ResponsavelID, receita.Observacao,
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar receita: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("receita não encontrada")
	}

	return nil
}

func (rs *ReceitaService) DeleteReceita(id int) error {
	query := `DELETE FROM receitas WHERE id = $1`

	result, err := rs.connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar receita: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("receita não encontrada")
	}

	return nil
}

// GetReceitasByObra retorna receitas filtradas por obra
func (rs *ReceitaService) GetReceitasByObra(obraId int) ([]models.ReceitaComRelacionamentos, error) {
	query := `
		SELECT r.id, r.obra_id, r.descricao, r.valor, r.data, r.fonte_receita, 
		       r.numero_documento, r.responsavel_id, r.observacao, r.created_at, r.updated_at,
		       o.nome as obra_nome, p.nome as responsavel_nome
		FROM receitas r
		LEFT JOIN obra o ON r.obra_id = o.id
		LEFT JOIN pessoa p ON r.responsavel_id = p.id
		WHERE r.obra_id = $1
		ORDER BY r.data DESC
	`

	rows, err := rs.connection.Query(query, obraId)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar receitas por obra: %v", err)
	}
	defer rows.Close()

	var receitas []models.ReceitaComRelacionamentos
	for rows.Next() {
		var receita models.ReceitaComRelacionamentos
		err := rows.Scan(
			&receita.ID, &receita.ObraID, &receita.Descricao, &receita.Valor,
			&receita.Data, &receita.FonteReceita, &receita.NumeroDocumento,
			&receita.ResponsavelID, &receita.Observacao, &receita.CreatedAt,
			&receita.UpdatedAt, &receita.ObraNome, &receita.ResponsavelNome,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear receita: %v", err)
		}
		receitas = append(receitas, receita)
	}

	return receitas, nil
}
