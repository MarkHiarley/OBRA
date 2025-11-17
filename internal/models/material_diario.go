package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type MaterialDiario struct {
	ID            null.Int    `json:"id"`
	ObraID        null.Int    `json:"obra_id" binding:"required"`
	Data          null.String `json:"data" binding:"required"` // Formato: YYYY-MM-DD
	Codigo        null.String `json:"codigo"`
	Descricao     null.String `json:"descricao" binding:"required"`
	Quantidade    null.Float  `json:"quantidade" binding:"required"`
	Unidade       null.String `json:"unidade" binding:"required"`
	Fornecedor    null.String `json:"fornecedor"`
	ValorUnitario null.Float  `json:"valor_unitario"`
	ValorTotal    null.Float  `json:"valor_total"`
	Observacoes   null.String `json:"observacoes"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     null.Time   `json:"updated_at"`
}
