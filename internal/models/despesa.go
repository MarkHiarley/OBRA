package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// Categorias de despesa
const (
	CategoriaMaterial           = "MATERIAL"
	CategoriaMaoDeObra          = "MAO_DE_OBRA"
	CategoriaCombustivel        = "COMBUSTIVEL"
	CategoriaAlimentacao        = "ALIMENTACAO"
	CategoriaMaterialEletrico   = "MATERIAL_ELETRICO"
	CategoriaAluguelEquipamento = "ALUGUEL_EQUIPAMENTO"
	CategoriaTransporte         = "TRANSPORTE"
	CategoriaImposto            = "IMPOSTO"
	CategoriaParceiro           = "PARCEIRO"
	CategoriaOutros             = "OUTROS"
)

// Formas de pagamento
const (
	FormaPagamentoPix           = "PIX"
	FormaPagamentoBoleto        = "BOLETO"
	FormaPagamentoCartaoCredito = "CARTAO_CREDITO"
	FormaPagamentoCartaoDebito  = "CARTAO_DEBITO"
	FormaPagamentoTransferencia = "TRANSFERENCIA"
	FormaPagamentoEspecie       = "ESPECIE"
	FormaPagamentoCheque        = "CHEQUE"
)

// Status de pagamento
const (
	StatusPagamentoPendente  = "PENDENTE"
	StatusPagamentoPago      = "PAGO"
	StatusPagamentoCancelado = "CANCELADO"
)

type Despesa struct {
	ID                   null.Int    `json:"id"`
	ObraID               null.Int    `json:"obra_id" binding:"required"`
	FornecedorID         null.Int    `json:"fornecedor_id,omitempty"`
	Data                 null.Time   `json:"data,omitempty"`            // Data da despesa/compra (aceita também data_vencimento como fallback)
	DataVencimento       null.Time   `json:"data_vencimento,omitempty"` // Data de vencimento do pagamento
	Descricao            null.String `json:"descricao" binding:"required"`
	Categoria            null.String `json:"categoria,omitempty"` // MATERIAL, MAO_DE_OBRA, COMBUSTIVEL, etc (opcional - padrão OUTROS)
	Valor                null.Float  `json:"valor" binding:"required"`
	FormaPagamento       null.String `json:"forma_pagamento,omitempty"` // PIX, BOLETO, CARTAO_CREDITO, etc (opcional - padrão PIX)
	StatusPagamento      null.String `json:"status_pagamento"`          // PENDENTE, PAGO, CANCELADO
	DataPagamento        null.Time   `json:"data_pagamento,omitempty"`
	ResponsavelPagamento null.String `json:"responsavel_pagamento,omitempty"`
	Observacao           null.String `json:"observacao,omitempty"`
	CreatedAt            time.Time   `json:"created_at"`
	UpdatedAt            time.Time   `json:"updated_at"`
}

// DespesaComRelacionamentos inclui dados do fornecedor e obra
type DespesaComRelacionamentos struct {
	Despesa
	FornecedorNome null.String `json:"fornecedor_nome,omitempty"`
	ObraNome       null.String `json:"obra_nome,omitempty"`
}

// RelatorioDespesas para agrupamentos e totalizações
type RelatorioDespesas struct {
	ObraID          null.Int    `json:"obra_id,omitempty"`
	ObraNome        null.String `json:"obra_nome,omitempty"`
	Categoria       null.String `json:"categoria,omitempty"`
	TotalDespesas   null.Float  `json:"total_despesas"`
	QuantidadeItens null.Int    `json:"quantidade_itens"`
}
