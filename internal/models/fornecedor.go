package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Fornecedor struct {
	ID            null.Int    `json:"id"`
	Nome          null.String `json:"nome" binding:"required"`
	TipoDocumento null.String `json:"tipo_documento" binding:"required"` // CPF ou CNPJ
	Documento     null.String `json:"documento" binding:"required"`
	Email         null.String `json:"email,omitempty"`
	Telefone      null.String `json:"telefone,omitempty"`
	Endereco      null.String `json:"endereco,omitempty"`
	Cidade        null.String `json:"cidade,omitempty"`
	Estado        null.String `json:"estado,omitempty"`
	Ativo         null.Bool   `json:"ativo"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
