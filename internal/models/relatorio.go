package models

import "gopkg.in/guregu/null.v4"

// RelatorioObra consolida dados financeiros e status de uma obra
type RelatorioObra struct {
	ObraID              null.Int    `json:"obra_id"`
	ObraNome            null.String `json:"obra_nome"`
	OrcamentoPrevisto   null.Float  `json:"orcamento_previsto"`   // Orçamento inicial da obra
	GastoRealizado      null.Float  `json:"gasto_realizado"`      // Soma de todas as despesas pagas
	ReceitaTotal        null.Float  `json:"receita_total"`        // Soma de todas as receitas
	SaldoAtual          null.Float  `json:"saldo_atual"`          // Receitas - Despesas
	PagamentoPendente   null.Float  `json:"pagamento_pendente"`   // Despesas com status PENDENTE
	Status              null.String `json:"status"`               // Status da obra
	PercentualExecutado null.Float  `json:"percentual_executado"` // (Gasto / Orçamento) * 100
	PercentualLucro     null.Float  `json:"percentual_lucro"`     // (Saldo / Receita) * 100
	TotalDespesas       null.Int    `json:"total_despesas"`       // Quantidade de despesas
	TotalReceitas       null.Int    `json:"total_receitas"`       // Quantidade de receitas
}

// RelatorioFinanceiroPorCategoria agrupa despesas por categoria
type RelatorioFinanceiroPorCategoria struct {
	ObraID          null.Int    `json:"obra_id"`
	ObraNome        null.String `json:"obra_nome"`
	Categoria       null.String `json:"categoria"`
	TotalGasto      null.Float  `json:"total_gasto"`
	QuantidadeItens null.Int    `json:"quantidade_itens"`
	PercentualTotal null.Float  `json:"percentual_total"` // % do total gasto na obra
}

// RelatorioPagamentos detalha status de pagamentos
type RelatorioPagamentos struct {
	ObraID               null.Int    `json:"obra_id"`
	ObraNome             null.String `json:"obra_nome"`
	DespesaID            null.Int    `json:"despesa_id"`
	Descricao            null.String `json:"descricao"`
	Valor                null.Float  `json:"valor"`
	StatusPagamento      null.String `json:"status_pagamento"`
	DataVencimento       null.Time   `json:"data_vencimento"`
	DataPagamento        null.Time   `json:"data_pagamento"`
	DiasAtraso           null.Int    `json:"dias_atraso"` // Para pagamentos em atraso
	FormaPagamento       null.String `json:"forma_pagamento"`
	ResponsavelPagamento null.String `json:"responsavel_pagamento"`
}

// RelatorioMateriais agrupa materiais consumidos
type RelatorioMateriais struct {
	ObraID          null.Int    `json:"obra_id"`
	ObraNome        null.String `json:"obra_nome"`
	TotalMateriais  null.Float  `json:"total_materiais"` // Soma categoria MATERIAL
	QuantidadeItens null.Int    `json:"quantidade_itens"`
	MaiorGasto      null.String `json:"maior_gasto_descricao"` // Descrição do maior gasto
	MaiorGastoValor null.Float  `json:"maior_gasto_valor"`
}

// RelatorioProfissionais agrupa custos de mão de obra
type RelatorioProfissionais struct {
	ObraID               null.Int    `json:"obra_id"`
	ObraNome             null.String `json:"obra_nome"`
	TotalMaoDeObra       null.Float  `json:"total_mao_de_obra"` // Soma categoria MAO_DE_OBRA
	QuantidadePagamentos null.Int    `json:"quantidade_pagamentos"`
	MaiorPagamento       null.String `json:"maior_pagamento_descricao"`
	MaiorPagamentoValor  null.Float  `json:"maior_pagamento_valor"`
}
