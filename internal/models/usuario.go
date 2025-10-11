package models

import (
	"time"

	"database/sql"
)

type Usuario struct {
	ID int `json:"id"`

	Email          string         `json:"email" binding:"required,email"`
	Nome           string         `json:"nome" binding:"required"`
	Senha          string         `json:"senha" binding:"required"`
	Tipo_documento sql.NullString `json:"tipo_documento"`
	Documento      sql.NullString `json:"documento"`
	Telefone       sql.NullInt64  `json:"telefone"`
	Perfil_acesso  string         `json:"perfil_acesso" binding:"required"`
	Ativo          sql.NullBool   `json:"ativo"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
