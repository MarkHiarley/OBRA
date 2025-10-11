package models

import (
	"database/sql"
	"time"
)

type Pessoa struct {
	ID        int            `json:"id"`
	Nome      string         `json:"nome" binding:"required"`
	Tipo      string         `json:"tipo" binding:"required,oneof=PF PJ"`
	Documento string         `json:"documento" binding:"required"`
	Email     sql.NullString `json:"email,omitempty"`
	Telefone  sql.NullString `json:"telefone,omitempty"`
	Cargo     sql.NullString `json:"cargo,omitempty"`
	Ativo     sql.NullBool   `json:"ativo"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
