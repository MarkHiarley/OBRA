package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type DiarioObra struct {
	ID                   null.Int    `json:"id"`
	ObraID               null.Int    `json:"obra_id" binding:"required"`
	Data                 null.String `json:"data" binding:"required"` // "2024-10-08"
	Periodo              null.String `json:"periodo"`
	AtividadesRealizadas null.String `json:"atividades_realizadas" binding:"required"`
	Ocorrencias          null.String `json:"ocorrencias,omitempty"`
	Observacoes          null.String `json:"observacoes,omitempty"`
	Foto                 null.String `json:"foto,omitempty"` // Base64 encoded image
	ResponsavelID        null.Int    `json:"responsavel_id,omitempty"`
	AprovadoPorID        null.Int    `json:"aprovado_por_id,omitempty"`
	StatusAprovacao      null.String `json:"status_aprovacao"`
	CreatedAt            time.Time   `json:"createdAt"`
	UpdatedAt            null.Time   `json:"updatedAt"`
}
