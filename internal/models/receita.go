package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// Fontes de receita
const (
	FonteReceitaContrato     = "CONTRATO"
	FonteReceitaPagamento    = "PAGAMENTO_CLIENTE"
	FonteReceitaAdiantamento = "ADIANTAMENTO"
	FonteReceitaFinanciamento = "FINANCIAMENTO"
	FonteReceitaMedicao      = "MEDICAO"
	FonteReceitaOutros       = "OUTROS"
)

type Receita struct {
	ID               null.Int    `json:"id"`
	ObraID           null.Int    `json:"obra_id" binding:"required"`
	Descricao        null.String `json:"descricao" binding:"required"`
	Valor            null.Float  `json:"valor" binding:"required"`
	Data             null.Time   `json:"data" binding:"required"`
	FonteReceita     null.String `json:"fonte_receita"` // CONTRATO, PAGAMENTO_CLIENTE, ADIANTAMENTO, etc
	NumeroDocumento  null.String `json:"numero_documento,omitempty"` // Número do contrato, nota fiscal, etc
	ResponsavelID    null.Int    `json:"responsavel_id,omitempty"`
	Observacao       null.String `json:"observacao,omitempty"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}

// ReceitaComRelacionamentos inclui dados da obra
type ReceitaComRelacionamentos struct {
	Receita
	ObraNome         null.String `json:"obra_nome,omitempty"`
	ResponsavelNome  null.String `json:"responsavel_nome,omitempty"`
}

// RelatorioReceitas para agrupamentos e totalizações
type RelatorioReceitas struct {
	ObraID           null.Int    `json:"obra_id,omitempty"`
	ObraNome         null.String `json:"obra_nome,omitempty"`
	FonteReceita     null.String `json:"fonte_receita,omitempty"`
	TotalReceitas    null.Float  `json:"total_receitas"`
	QuantidadeItens  null.Int    `json:"quantidade_itens"`
}