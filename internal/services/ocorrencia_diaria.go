package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
	"time"
)

type OcorrenciaDiariaService struct {
	connection  *sql.DB
	fotoService *FotoDiarioService
}

func NewOcorrenciaDiariaService(connection *sql.DB) OcorrenciaDiariaService {
	return OcorrenciaDiariaService{
		connection:  connection,
		fotoService: &FotoDiarioService{connection: connection},
	}
}

// CreateOcorrencia cria uma nova ocorrência diária com suas fotos
func (ods *OcorrenciaDiariaService) CreateOcorrencia(ocorrencia models.OcorrenciaDiaria) (models.OcorrenciaDiaria, error) {
	query := `
		INSERT INTO ocorrencia_diaria (obra_id, data, periodo, tipo, gravidade, descricao, responsavel_id, status_resolucao, acao_tomada)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`

	err := ods.connection.QueryRow(
		query,
		ocorrencia.ObraID,
		ocorrencia.Data,
		ocorrencia.Periodo,
		ocorrencia.Tipo,
		ocorrencia.Gravidade,
		ocorrencia.Descricao,
		ocorrencia.ResponsavelID,
		ocorrencia.StatusResolucao,
		ocorrencia.AcaoTomada,
	).Scan(&ocorrencia.ID, &ocorrencia.CreatedAt, &ocorrencia.UpdatedAt)

	if err != nil {
		return models.OcorrenciaDiaria{}, fmt.Errorf("erro ao criar ocorrência: %v", err)
	}

	// Salvar fotos relacionadas se houver
	if len(ocorrencia.Fotos) > 0 && ocorrencia.ID.Valid {
		err = ods.fotoService.CreateFotos(ocorrencia.Fotos, "ocorrencia", ocorrencia.ID.Int64)
		if err != nil {
			return models.OcorrenciaDiaria{}, fmt.Errorf("erro ao salvar fotos da ocorrência: %v", err)
		}
	}

	return ocorrencia, nil
}

// GetOcorrencias retorna todas as ocorrências com relacionamentos
func (ods *OcorrenciaDiariaService) GetOcorrencias() ([]models.OcorrenciaDiariaComRelacionamentos, error) {
	query := `
		SELECT oc.id, oc.obra_id, oc.data, oc.periodo, oc.tipo, oc.gravidade, 
		       oc.descricao, oc.responsavel_id, oc.status_resolucao, oc.acao_tomada,
		       oc.created_at, oc.updated_at,
		       o.nome as obra_nome, p.nome as responsavel_nome
		FROM ocorrencia_diaria oc
		LEFT JOIN obra o ON oc.obra_id = o.id
		LEFT JOIN pessoa p ON oc.responsavel_id = p.id
		ORDER BY oc.data DESC, oc.gravidade DESC, oc.created_at DESC
	`

	rows, err := ods.connection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar ocorrências: %v", err)
	}
	defer rows.Close()

	var ocorrencias []models.OcorrenciaDiariaComRelacionamentos
	for rows.Next() {
		var oc models.OcorrenciaDiariaComRelacionamentos
		err := rows.Scan(
			&oc.ID, &oc.ObraID, &oc.Data, &oc.Periodo, &oc.Tipo, &oc.Gravidade,
			&oc.Descricao, &oc.ResponsavelID, &oc.StatusResolucao, &oc.AcaoTomada,
			&oc.CreatedAt, &oc.UpdatedAt, &oc.ObraNome, &oc.ResponsavelNome,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear ocorrência: %v", err)
		}
		ocorrencias = append(ocorrencias, oc)
	}

	return ocorrencias, nil
}

// GetOcorrenciasByObraData retorna ocorrências filtradas por obra e data
func (ods *OcorrenciaDiariaService) GetOcorrenciasByObraData(obraID int, data string) ([]models.OcorrenciaDiariaComRelacionamentos, error) {
	query := `
		SELECT oc.id, oc.obra_id, oc.data, oc.periodo, oc.tipo, oc.gravidade, 
		       oc.descricao, oc.responsavel_id, oc.status_resolucao, oc.acao_tomada,
		       oc.created_at, oc.updated_at,
		       o.nome as obra_nome, p.nome as responsavel_nome
		FROM ocorrencia_diaria oc
		LEFT JOIN obra o ON oc.obra_id = o.id
		LEFT JOIN pessoa p ON oc.responsavel_id = p.id
		WHERE oc.obra_id = $1 AND oc.data = $2
		ORDER BY oc.gravidade DESC, oc.created_at
	`

	rows, err := ods.connection.Query(query, obraID, data)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar ocorrências: %v", err)
	}
	defer rows.Close()

	var ocorrencias []models.OcorrenciaDiariaComRelacionamentos
	for rows.Next() {
		var oc models.OcorrenciaDiariaComRelacionamentos
		err := rows.Scan(
			&oc.ID, &oc.ObraID, &oc.Data, &oc.Periodo, &oc.Tipo, &oc.Gravidade,
			&oc.Descricao, &oc.ResponsavelID, &oc.StatusResolucao, &oc.AcaoTomada,
			&oc.CreatedAt, &oc.UpdatedAt, &oc.ObraNome, &oc.ResponsavelNome,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear ocorrência: %v", err)
		}
		ocorrencias = append(ocorrencias, oc)
	}

	return ocorrencias, nil
}

// GetOcorrenciasByGravidade filtra ocorrências por gravidade
func (ods *OcorrenciaDiariaService) GetOcorrenciasByGravidade(gravidade string) ([]models.OcorrenciaDiariaComRelacionamentos, error) {
	query := `
		SELECT oc.id, oc.obra_id, oc.data, oc.periodo, oc.tipo, oc.gravidade, 
		       oc.descricao, oc.responsavel_id, oc.status_resolucao, oc.acao_tomada,
		       oc.created_at, oc.updated_at,
		       o.nome as obra_nome, p.nome as responsavel_nome
		FROM ocorrencia_diaria oc
		LEFT JOIN obra o ON oc.obra_id = o.id
		LEFT JOIN pessoa p ON oc.responsavel_id = p.id
		WHERE oc.gravidade = $1
		ORDER BY oc.data DESC, oc.created_at DESC
	`

	rows, err := ods.connection.Query(query, gravidade)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar ocorrências: %v", err)
	}
	defer rows.Close()

	var ocorrencias []models.OcorrenciaDiariaComRelacionamentos
	for rows.Next() {
		var oc models.OcorrenciaDiariaComRelacionamentos
		err := rows.Scan(
			&oc.ID, &oc.ObraID, &oc.Data, &oc.Periodo, &oc.Tipo, &oc.Gravidade,
			&oc.Descricao, &oc.ResponsavelID, &oc.StatusResolucao, &oc.AcaoTomada,
			&oc.CreatedAt, &oc.UpdatedAt, &oc.ObraNome, &oc.ResponsavelNome,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear ocorrência: %v", err)
		}
		ocorrencias = append(ocorrencias, oc)
	}

	return ocorrencias, nil
}

// UpdateOcorrencia atualiza uma ocorrência existente
func (ods *OcorrenciaDiariaService) UpdateOcorrencia(id int, ocorrencia models.OcorrenciaDiaria) error {
	query := `
		UPDATE ocorrencia_diaria 
		SET data = $2, periodo = $3, tipo = $4, gravidade = $5, descricao = $6, 
		    responsavel_id = $7, status_resolucao = $8, acao_tomada = $9, updated_at = $10
		WHERE id = $1
	`

	result, err := ods.connection.Exec(
		query, id, ocorrencia.Data, ocorrencia.Periodo, ocorrencia.Tipo,
		ocorrencia.Gravidade, ocorrencia.Descricao, ocorrencia.ResponsavelID,
		ocorrencia.StatusResolucao, ocorrencia.AcaoTomada, time.Now(),
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar ocorrência: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("ocorrência não encontrada")
	}

	return nil
}

// DeleteOcorrencia remove uma ocorrência
func (ods *OcorrenciaDiariaService) DeleteOcorrencia(id int) error {
	query := `DELETE FROM ocorrencia_diaria WHERE id = $1`

	result, err := ods.connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar ocorrência: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("ocorrência não encontrada")
	}

	return nil
}
