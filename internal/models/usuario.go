package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Usuario struct {
	ID            int         `json:"id"`
	Email         null.String `json:"email,omitempty"`
	Nome          null.String `json:"nome" binding:"required"`
	Senha         null.String `json:"senha" binding:"required"`
	TipoDocumento null.String `json:"tipo_documento"`
	Documento     null.String `json:"documento"`
	Telefone      null.String `json:"telefone"`
	PerfilAcesso  null.String `json:"perfil_acesso" binding:"required"`
	Ativo         null.Bool   `json:"ativo"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
