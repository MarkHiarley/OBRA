package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
)

type DiarioConsolidadoService struct {
	connection *sql.DB
}

func NewDiarioConsolidadoService(connection *sql.DB) DiarioConsolidadoService {
	return DiarioConsolidadoService{
		connection: connection,
	}
}

// GetDiarioConsolidado retorna o diário consolidado (view agregada)
func (dcs *DiarioConsolidadoService) GetDiarioConsolidado() ([]models.DiarioConsolidado, error) {
	query := `
		SELECT diario_id, obra_id, obra_nome, data, periodo, 
		       atividades, ocorrencias, foto, observacoes,
		       responsavel_id, responsavel_nome, aprovado_por_id, aprovado_por_nome,
		       status_aprovacao, qtd_atividades, qtd_ocorrencias, qtd_equipe,
		       qtd_equipamentos, qtd_materiais, created_at, updated_at
		FROM vw_diario_consolidado
		ORDER BY data DESC, periodo
	`

	rows, err := dcs.connection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar diário consolidado: %v", err)
	}
	defer rows.Close()

	var diarios []models.DiarioConsolidado
	for rows.Next() {
		var d models.DiarioConsolidado
		err := rows.Scan(
			&d.DiarioID, &d.ObraID, &d.ObraNome, &d.Data, &d.Periodo,
			&d.AtividadesRealizadas, &d.Ocorrencias, &d.Foto, &d.Observacoes,
			&d.ResponsavelID, &d.ResponsavelNome, &d.AprovadoPorID, &d.AprovadoPorNome,
			&d.StatusAprovacao, &d.QtdAtividades, &d.QtdOcorrencias, &d.QtdEquipe,
			&d.QtdEquipamentos, &d.QtdMateriais, &d.CreatedAt, &d.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear diário consolidado: %v", err)
		}
		diarios = append(diarios, d)
	}

	return diarios, nil
}

// GetDiarioConsolidadoByObra retorna diário consolidado filtrado por obra
func (dcs *DiarioConsolidadoService) GetDiarioConsolidadoByObra(obraID int) ([]models.DiarioConsolidado, error) {
	query := `
		SELECT diario_id, obra_id, obra_nome, data, periodo, 
		       atividades, ocorrencias, foto, observacoes,
		       responsavel_id, responsavel_nome, aprovado_por_id, aprovado_por_nome,
		       status_aprovacao, qtd_atividades, qtd_ocorrencias, qtd_equipe,
		       qtd_equipamentos, qtd_materiais, created_at, updated_at
		FROM vw_diario_consolidado
		WHERE obra_id = $1
		ORDER BY data DESC, periodo
	`

	rows, err := dcs.connection.Query(query, obraID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar diário consolidado: %v", err)
	}
	defer rows.Close()

	var diarios []models.DiarioConsolidado
	for rows.Next() {
		var d models.DiarioConsolidado
		err := rows.Scan(
			&d.DiarioID, &d.ObraID, &d.ObraNome, &d.Data, &d.Periodo,
			&d.AtividadesRealizadas, &d.Ocorrencias, &d.Foto, &d.Observacoes,
			&d.ResponsavelID, &d.ResponsavelNome, &d.AprovadoPorID, &d.AprovadoPorNome,
			&d.StatusAprovacao, &d.QtdAtividades, &d.QtdOcorrencias, &d.QtdEquipe,
			&d.QtdEquipamentos, &d.QtdMateriais, &d.CreatedAt, &d.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear diário consolidado: %v", err)
		}
		diarios = append(diarios, d)
	}

	return diarios, nil
}

// GetDiarioConsolidadoByData retorna diário consolidado filtrado por data
func (dcs *DiarioConsolidadoService) GetDiarioConsolidadoByData(data string) ([]models.DiarioConsolidado, error) {
	query := `
		SELECT diario_id, obra_id, obra_nome, data, periodo, 
		       atividades, ocorrencias, foto, observacoes,
		       responsavel_id, responsavel_nome, aprovado_por_id, aprovado_por_nome,
		       status_aprovacao, qtd_atividades, qtd_ocorrencias, qtd_equipe,
		       qtd_equipamentos, qtd_materiais, created_at, updated_at
		FROM vw_diario_consolidado
		WHERE data = $1
		ORDER BY obra_id, periodo
	`

	rows, err := dcs.connection.Query(query, data)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar diário consolidado: %v", err)
	}
	defer rows.Close()

	var diarios []models.DiarioConsolidado
	for rows.Next() {
		var d models.DiarioConsolidado
		err := rows.Scan(
			&d.DiarioID, &d.ObraID, &d.ObraNome, &d.Data, &d.Periodo,
			&d.AtividadesRealizadas, &d.Ocorrencias, &d.Foto, &d.Observacoes,
			&d.ResponsavelID, &d.ResponsavelNome, &d.AprovadoPorID, &d.AprovadoPorNome,
			&d.StatusAprovacao, &d.QtdAtividades, &d.QtdOcorrencias, &d.QtdEquipe,
			&d.QtdEquipamentos, &d.QtdMateriais, &d.CreatedAt, &d.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear diário consolidado: %v", err)
		}
		diarios = append(diarios, d)
	}

	return diarios, nil
}
