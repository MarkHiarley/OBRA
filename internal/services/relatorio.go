package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
)

type RelatorioService struct {
	connection *sql.DB
}

func NewRelatorioService(connection *sql.DB) RelatorioService {
	return RelatorioService{
		connection: connection,
	}
}

// GetRelatorioObra retorna relatório financeiro completo de uma obra
func (rs *RelatorioService) GetRelatorioObra(obraId int64) (models.RelatorioObra, error) {
	var relatorio models.RelatorioObra

	query := `
		SELECT 
			o.id,
			o.nome,
			COALESCE(o.orcamento, 0) as orcamento,
			COALESCE((SELECT SUM(valor) FROM despesa WHERE obra_id = o.id AND status_pagamento = 'PAGO'), 0) as gasto_realizado,
			COALESCE((SELECT SUM(valor) FROM receitas WHERE obra_id = o.id), 0) as receita_total,
			COALESCE((SELECT SUM(valor) FROM receitas WHERE obra_id = o.id), 0) - 
			COALESCE((SELECT SUM(valor) FROM despesa WHERE obra_id = o.id AND status_pagamento = 'PAGO'), 0) as saldo_atual,
			COALESCE((SELECT SUM(valor) FROM despesa WHERE obra_id = o.id AND status_pagamento = 'PENDENTE'), 0) as pagamento_pendente,
			COALESCE(o.status, 'PLANEJADA') as status,
			CASE 
				WHEN COALESCE(o.orcamento, 0) > 0 THEN 
					(COALESCE((SELECT SUM(valor) FROM despesa WHERE obra_id = o.id AND status_pagamento = 'PAGO'), 0) / COALESCE(o.orcamento, 1)) * 100 
				ELSE 0 
			END as percentual_executado,
			CASE 
				WHEN COALESCE((SELECT SUM(valor) FROM receitas WHERE obra_id = o.id), 0) > 0 THEN 
					((COALESCE((SELECT SUM(valor) FROM receitas WHERE obra_id = o.id), 0) - COALESCE((SELECT SUM(valor) FROM despesa WHERE obra_id = o.id AND status_pagamento = 'PAGO'), 0)) / COALESCE((SELECT SUM(valor) FROM receitas WHERE obra_id = o.id), 1)) * 100 
				ELSE 0 
			END as percentual_lucro,
			COALESCE((SELECT COUNT(*) FROM despesa WHERE obra_id = o.id), 0) as total_despesas,
			COALESCE((SELECT COUNT(*) FROM receitas WHERE obra_id = o.id), 0) as total_receitas
		FROM obra o
		WHERE o.id = $1
	`

	err := rs.connection.QueryRow(query, obraId).Scan(
		&relatorio.ObraID,
		&relatorio.ObraNome,
		&relatorio.OrcamentoPrevisto,
		&relatorio.GastoRealizado,
		&relatorio.ReceitaTotal,
		&relatorio.SaldoAtual,
		&relatorio.PagamentoPendente,
		&relatorio.Status,
		&relatorio.PercentualExecutado,
		&relatorio.PercentualLucro,
		&relatorio.TotalDespesas,
		&relatorio.TotalReceitas,
	)

	if err != nil {
		return models.RelatorioObra{}, fmt.Errorf("erro ao gerar relatório da obra: %v", err)
	}

	return relatorio, nil
}

// GetRelatorioDespesasPorCategoria retorna despesas agrupadas por categoria
func (rs *RelatorioService) GetRelatorioDespesasPorCategoria(obraId int64) ([]models.RelatorioFinanceiroPorCategoria, error) {
	query := `
		SELECT 
			d.obra_id,
			o.nome as obra_nome,
			d.categoria,
			COALESCE(SUM(d.valor), 0) as total_gasto,
			COUNT(*) as quantidade_itens,
			CASE 
				WHEN (SELECT SUM(valor) FROM despesa WHERE obra_id = $1) > 0 THEN 
					(COALESCE(SUM(d.valor), 0) / (SELECT SUM(valor) FROM despesa WHERE obra_id = $1)) * 100 
				ELSE 0 
			END as percentual_total
		FROM despesa d
		LEFT JOIN obra o ON d.obra_id = o.id
		WHERE d.obra_id = $1
		GROUP BY d.obra_id, o.nome, d.categoria
		ORDER BY total_gasto DESC
	`

	rows, err := rs.connection.Query(query, obraId)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar relatório de despesas por categoria: %v", err)
	}
	defer rows.Close()

	var relatorios []models.RelatorioFinanceiroPorCategoria
	for rows.Next() {
		var relatorio models.RelatorioFinanceiroPorCategoria
		err := rows.Scan(
			&relatorio.ObraID,
			&relatorio.ObraNome,
			&relatorio.Categoria,
			&relatorio.TotalGasto,
			&relatorio.QuantidadeItens,
			&relatorio.PercentualTotal,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear relatório de categoria: %v", err)
		}
		relatorios = append(relatorios, relatorio)
	}

	return relatorios, nil
}

// GetRelatorioPagamentos retorna status detalhado dos pagamentos
func (rs *RelatorioService) GetRelatorioPagamentos(obraId int64, statusFiltro string) ([]models.RelatorioPagamentos, error) {
	whereClause := "WHERE d.obra_id = $1"
	args := []interface{}{obraId}

	if statusFiltro != "" {
		whereClause += " AND d.status_pagamento = $2"
		args = append(args, statusFiltro)
	}

	query := fmt.Sprintf(`
		SELECT 
			d.obra_id,
			o.nome as obra_nome,
			d.id as despesa_id,
			d.descricao,
			d.valor,
			d.status_pagamento,
			d.data_vencimento,
			d.data_pagamento,
			CASE 
				WHEN d.data_vencimento IS NOT NULL AND d.status_pagamento = 'PENDENTE' 
					AND d.data_vencimento::DATE < CURRENT_DATE 
				THEN (CURRENT_DATE - d.data_vencimento::DATE)
				ELSE 0
			END as dias_atraso,
			d.forma_pagamento,
			d.responsavel_pagamento
		FROM despesa d
		LEFT JOIN obra o ON d.obra_id = o.id
		%s
		ORDER BY d.data_vencimento ASC
	`, whereClause)

	rows, err := rs.connection.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar relatório de pagamentos: %v", err)
	}
	defer rows.Close()

	var relatorios []models.RelatorioPagamentos
	for rows.Next() {
		var relatorio models.RelatorioPagamentos
		err := rows.Scan(
			&relatorio.ObraID,
			&relatorio.ObraNome,
			&relatorio.DespesaID,
			&relatorio.Descricao,
			&relatorio.Valor,
			&relatorio.StatusPagamento,
			&relatorio.DataVencimento,
			&relatorio.DataPagamento,
			&relatorio.DiasAtraso,
			&relatorio.FormaPagamento,
			&relatorio.ResponsavelPagamento,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear relatório de pagamentos: %v", err)
		}
		relatorios = append(relatorios, relatorio)
	}

	return relatorios, nil
}

// GetRelatorioMateriais retorna relatório específico de materiais
func (rs *RelatorioService) GetRelatorioMateriais(obraId int64) (models.RelatorioMateriais, error) {
	var relatorio models.RelatorioMateriais

	query := `
		SELECT 
			o.id as obra_id,
			o.nome as obra_nome,
			COALESCE((SELECT SUM(valor) FROM despesa WHERE obra_id = o.id AND categoria = 'MATERIAL'), 0) as total_materiais,
			COALESCE((SELECT COUNT(*) FROM despesa WHERE obra_id = o.id AND categoria = 'MATERIAL'), 0) as quantidade_itens,
			COALESCE((SELECT descricao FROM despesa WHERE obra_id = o.id AND categoria = 'MATERIAL' ORDER BY valor DESC LIMIT 1), '') as maior_gasto_descricao,
			COALESCE((SELECT valor FROM despesa WHERE obra_id = o.id AND categoria = 'MATERIAL' ORDER BY valor DESC LIMIT 1), 0) as maior_gasto_valor
		FROM obra o
		WHERE o.id = $1
	`

	err := rs.connection.QueryRow(query, obraId).Scan(
		&relatorio.ObraID,
		&relatorio.ObraNome,
		&relatorio.TotalMateriais,
		&relatorio.QuantidadeItens,
		&relatorio.MaiorGasto,
		&relatorio.MaiorGastoValor,
	)

	if err != nil {
		return models.RelatorioMateriais{}, fmt.Errorf("erro ao gerar relatório de materiais: %v", err)
	}

	return relatorio, nil
}

// GetRelatorioProfissionais retorna relatório específico de mão de obra
func (rs *RelatorioService) GetRelatorioProfissionais(obraId int64) (models.RelatorioProfissionais, error) {
	var relatorio models.RelatorioProfissionais

	query := `
		SELECT 
			o.id as obra_id,
			o.nome as obra_nome,
			COALESCE((SELECT SUM(valor) FROM despesa WHERE obra_id = o.id AND categoria = 'MAO_DE_OBRA'), 0) as total_mao_de_obra,
			COALESCE((SELECT COUNT(*) FROM despesa WHERE obra_id = o.id AND categoria = 'MAO_DE_OBRA'), 0) as quantidade_pagamentos,
			COALESCE((SELECT descricao FROM despesa WHERE obra_id = o.id AND categoria = 'MAO_DE_OBRA' ORDER BY valor DESC LIMIT 1), '') as maior_pagamento_descricao,
			COALESCE((SELECT valor FROM despesa WHERE obra_id = o.id AND categoria = 'MAO_DE_OBRA' ORDER BY valor DESC LIMIT 1), 0) as maior_pagamento_valor
		FROM obra o
		WHERE o.id = $1
	`

	err := rs.connection.QueryRow(query, obraId).Scan(
		&relatorio.ObraID,
		&relatorio.ObraNome,
		&relatorio.TotalMaoDeObra,
		&relatorio.QuantidadePagamentos,
		&relatorio.MaiorPagamento,
		&relatorio.MaiorPagamentoValor,
	)

	if err != nil {
		return models.RelatorioProfissionais{}, fmt.Errorf("erro ao gerar relatório de profissionais: %v", err)
	}

	return relatorio, nil
}
