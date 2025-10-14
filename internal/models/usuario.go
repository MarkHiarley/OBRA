package models

import (
	"gopkg.in/guregu/null.v4"
)

type Usuario struct {
	ID            null.Int    `json:"id"`
	Email         null.String `json:"email"` // ✅ Sem omitempty
	Nome          null.String `json:"nome"`  // ✅ Sem omitempty
	Senha         null.String `json:"senha"`
	TipoDocumento null.String `json:"tipo_documento"`
	Documento     null.String `json:"documento"`
	Telefone      null.String `json:"telefone"`
	PerfilAcesso  null.String `json:"perfil_acesso"`
	Ativo         null.Bool   `json:"ativo"`
	CreatedAt     null.Time   `json:"createdAt"`
	UpdatedAt     null.Time   `json:"updatedAt"`
}
