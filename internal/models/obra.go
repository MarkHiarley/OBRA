package models

import (
	"gopkg.in/guregu/null.v4"
)

type Obra struct {
	ID              null.Int    `json:"id"`
	Nome            null.String `json:"nome"`
	ContratoNumero  null.String `json:"contrato_numero"`
	ContratanteID   null.Int    `json:"contratante_id"`
	ResponsavelID   null.Int    `json:"responsavel_id,omitempty"`
	DataInicio      null.String `json:"data_inicio"`
	PrazoDias       null.Int    `json:"prazo_dias"`
	DataFimPrevista null.String `json:"data_fim_prevista,omitempty"`
	Orcamento       null.Float  `json:"orcamento,omitempty"`
	Status          null.String `json:"status"`
	Art             null.String `json:"art,omitempty"`
	EnderecoRua     null.String `json:"endereco_rua,omitempty"`
	EnderecoNumero  null.String `json:"endereco_numero,omitempty"`
	EnderecoBairro  null.String `json:"endereco_bairro,omitempty"`
	EnderecoCidade  null.String `json:"endereco_cidade,omitempty"`
	EnderecoEstado  null.String `json:"endereco_estado,omitempty"`
	EnderecoCep     null.String `json:"endereco_cep,omitempty"`
	Observacoes     null.String `json:"observacoes,omitempty"`
	Ativo           null.Bool   `json:"ativo"`
	CreatedAt       null.Time   `json:"created_at"`
	UpdatedAt       null.Time   `json:"updated_at"`
}
