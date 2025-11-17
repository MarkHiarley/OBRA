package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
	"time"
)

type AtividadeDiariaService struct {
	connection  *sql.DB
	fotoService *FotoDiarioService
}

func NewAtividadeDiariaService(connection *sql.DB) AtividadeDiariaService {
	return AtividadeDiariaService{
		connection:  connection,
		fotoService: &FotoDiarioService{connection: connection},
	}
}

// CreateAtividade cria uma nova atividade diária com suas fotos
func (ads *AtividadeDiariaService) CreateAtividade(atividade models.AtividadeDiaria) (models.AtividadeDiaria, error) {
	query := `
		INSERT INTO atividade_diaria (obra_id, data, periodo, descricao, responsavel_id, status, percentual_conclusao, observacao)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

	err := ads.connection.QueryRow(
		query,
		atividade.ObraID,
		atividade.Data,
		atividade.Periodo,
		atividade.Descricao,
		atividade.ResponsavelID,
		atividade.Status,
		atividade.PercentualConclusao,
		atividade.Observacao,
	).Scan(&atividade.ID, &atividade.CreatedAt, &atividade.UpdatedAt)

	if err != nil {
		return models.AtividadeDiaria{}, fmt.Errorf("erro ao criar atividade: %v", err)
	}

	// Salvar fotos relacionadas se houver
	if len(atividade.Fotos) > 0 && atividade.ID.Valid {
		err = ads.fotoService.CreateFotos(atividade.Fotos, "atividade", atividade.ID.Int64)
		if err != nil {
			return models.AtividadeDiaria{}, fmt.Errorf("erro ao salvar fotos da atividade: %v", err)
		}
	}

	return atividade, nil
}

// GetAtividades retorna todas as atividades com relacionamentos e fotos
func (ads *AtividadeDiariaService) GetAtividades() ([]models.AtividadeDiariaComRelacionamentos, error) {
	query := `
		SELECT a.id, a.obra_id, a.data, a.periodo, a.descricao, a.responsavel_id, 
		       a.status, a.percentual_conclusao, a.observacao, a.created_at, a.updated_at,
		       o.nome as obra_nome, p.nome as responsavel_nome
		FROM atividade_diaria a
		LEFT JOIN obra o ON a.obra_id = o.id
		LEFT JOIN pessoa p ON a.responsavel_id = p.id
		ORDER BY a.data DESC, a.created_at DESC
	`

	rows, err := ads.connection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar atividades: %v", err)
	}
	defer rows.Close()

	var atividades []models.AtividadeDiariaComRelacionamentos
	for rows.Next() {
		var ativ models.AtividadeDiariaComRelacionamentos
		err := rows.Scan(
			&ativ.ID, &ativ.ObraID, &ativ.Data, &ativ.Periodo, &ativ.Descricao,
			&ativ.ResponsavelID, &ativ.Status, &ativ.PercentualConclusao, &ativ.Observacao,
			&ativ.CreatedAt, &ativ.UpdatedAt, &ativ.ObraNome, &ativ.ResponsavelNome,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear atividade: %v", err)
		}

		// Buscar fotos da atividade
		if ativ.ID.Valid {
			fotos, err := ads.fotoService.GetFotosByEntidade("atividade", ativ.ID.Int64)
			if err == nil {
				ativ.Fotos = fotos
			}
		}

		atividades = append(atividades, ativ)
	}

	return atividades, nil
}

// GetAtividadesByObra retorna todas as atividades de uma obra (todas as datas)
func (ads *AtividadeDiariaService) GetAtividadesByObra(obraID int) ([]models.AtividadeDiariaComRelacionamentos, error) {
	query := `
		SELECT a.id, a.obra_id, a.data, a.periodo, a.descricao, a.responsavel_id, 
		       a.status, a.percentual_conclusao, a.observacao, a.created_at, a.updated_at,
		       o.nome as obra_nome, p.nome as responsavel_nome
		FROM atividade_diaria a
		LEFT JOIN obra o ON a.obra_id = o.id
		LEFT JOIN pessoa p ON a.responsavel_id = p.id
		WHERE a.obra_id = $1
		ORDER BY a.data DESC, a.created_at DESC
	`

	rows, err := ads.connection.Query(query, obraID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar atividades da obra: %v", err)
	}
	defer rows.Close()

	var atividades []models.AtividadeDiariaComRelacionamentos
	for rows.Next() {
		var ativ models.AtividadeDiariaComRelacionamentos
		err := rows.Scan(
			&ativ.ID, &ativ.ObraID, &ativ.Data, &ativ.Periodo, &ativ.Descricao,
			&ativ.ResponsavelID, &ativ.Status, &ativ.PercentualConclusao, &ativ.Observacao,
			&ativ.CreatedAt, &ativ.UpdatedAt, &ativ.ObraNome, &ativ.ResponsavelNome,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear atividade: %v", err)
		}

		// Buscar fotos da atividade
		if ativ.ID.Valid {
			fotos, err := ads.fotoService.GetFotosByEntidade("atividade", ativ.ID.Int64)
			if err == nil {
				ativ.Fotos = fotos
			}
		}

		atividades = append(atividades, ativ)
	}

	return atividades, nil
}

// GetAtividadesByObraData retorna atividades filtradas por obra e data
func (ads *AtividadeDiariaService) GetAtividadesByObraData(obraID int, data string) ([]models.AtividadeDiariaComRelacionamentos, error) {
	query := `
		SELECT a.id, a.obra_id, a.data, a.periodo, a.descricao, a.responsavel_id, 
		       a.status, a.percentual_conclusao, a.observacao, a.created_at, a.updated_at,
		       o.nome as obra_nome, p.nome as responsavel_nome
		FROM atividade_diaria a
		LEFT JOIN obra o ON a.obra_id = o.id
		LEFT JOIN pessoa p ON a.responsavel_id = p.id
		WHERE a.obra_id = $1 AND a.data = $2
		ORDER BY a.periodo, a.created_at
	`

	rows, err := ads.connection.Query(query, obraID, data)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar atividades: %v", err)
	}
	defer rows.Close()

	var atividades []models.AtividadeDiariaComRelacionamentos
	for rows.Next() {
		var ativ models.AtividadeDiariaComRelacionamentos
		err := rows.Scan(
			&ativ.ID, &ativ.ObraID, &ativ.Data, &ativ.Periodo, &ativ.Descricao,
			&ativ.ResponsavelID, &ativ.Status, &ativ.PercentualConclusao, &ativ.Observacao,
			&ativ.CreatedAt, &ativ.UpdatedAt, &ativ.ObraNome, &ativ.ResponsavelNome,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear atividade: %v", err)
		}
		atividades = append(atividades, ativ)
	}

	return atividades, nil
}

// UpdateAtividade atualiza uma atividade existente
func (ads *AtividadeDiariaService) UpdateAtividade(id int, atividade models.AtividadeDiaria) error {
	query := `
		UPDATE atividade_diaria 
		SET data = $2, periodo = $3, descricao = $4, responsavel_id = $5, 
		    status = $6, percentual_conclusao = $7, observacao = $8, updated_at = $9
		WHERE id = $1
	`

	result, err := ads.connection.Exec(
		query, id, atividade.Data, atividade.Periodo, atividade.Descricao,
		atividade.ResponsavelID, atividade.Status, atividade.PercentualConclusao,
		atividade.Observacao, time.Now(),
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar atividade: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("atividade não encontrada")
	}

	return nil
}

// DeleteAtividade remove uma atividade
func (ads *AtividadeDiariaService) DeleteAtividade(id int) error {
	query := `DELETE FROM atividade_diaria WHERE id = $1`

	result, err := ads.connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar atividade: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("atividade não encontrada")
	}

	return nil
}
