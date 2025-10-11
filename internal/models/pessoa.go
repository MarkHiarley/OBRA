package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Pessoa struct {
	ID            null.Int    `json:"id"`
	Nome          null.String `json:"nome" binding:"required"`
	TipoDocumento null.String `json:"tipo" binding:"required"`
	Documento     null.String `json:"documento" binding:"required"`
	Email         null.String `json:"email,omitempty"`
	Telefone      null.String `json:"telefone,omitempty"`
	Cargo         null.String `json:"cargo,omitempty"`
	Ativo         null.Bool   `json:"ativo"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}
