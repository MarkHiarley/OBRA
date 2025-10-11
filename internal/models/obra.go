package models

import (
	"database/sql"
	"time"
)

type Obra struct {
	ID              int             `json:"id"`
	Nome            string          `json:"nome" binding:"required"`
	ContratoNumero  string          `json:"contrato_numero" binding:"required"`
	ContratanteID   sql.NullInt64   `json:"contratante_id" binding:"required"`
	ResponsavelID   sql.NullInt64   `json:"responsavel_id,omitempty"`
	DataInicio      string          `json:"data_inicio" binding:"required"` // "2024-01-23"
	PrazoDias       int             `json:"prazo_dias" binding:"required,gt=0"`
	DataFimPrevista sql.NullString  `json:"data_fim_prevista,omitempty"`
	Orcamento       sql.NullFloat64 `json:"orcamento,omitempty"`
	Status          sql.NullString  `json:"status"`
	EnderecoRua     sql.NullString  `json:"endereco_rua,omitempty"`
	EnderecoNumero  sql.NullString  `json:"endereco_numero,omitempty"`
	EnderecoBairro  sql.NullString  `json:"endereco_bairro,omitempty"`
	EnderecoCidade  sql.NullString  `json:"endereco_cidade,omitempty"`
	EnderecoEstado  sql.NullString  `json:"endereco_estado,omitempty"`
	EnderecoCep     sql.NullString  `json:"endereco_cep,omitempty"`
	Observacoes     sql.NullString  `json:"observacoes,omitempty"`
	Ativo           sql.NullBool    `json:"ativo"`
	CreatedAt       time.Time       `json:"createdAt"`
}
