package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type EquipeDiario struct {
	ID                  null.Int    `json:"id"`
	DiarioID            null.Int    `json:"diario_id" binding:"required"`
	Codigo              null.String `json:"codigo"`
	Descricao           null.String `json:"descricao" binding:"required"`
	QuantidadeUtilizada null.Int    `json:"quantidade_utilizada" binding:"required"`
	HorasTrabalhadas    null.Float  `json:"horas_trabalhadas"`
	Observacoes         null.String `json:"observacoes"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           null.Time   `json:"updated_at"`
}
