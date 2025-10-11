package models

import "time"

type DiarioObra struct {
	ID                   int       `json:"id"`
	ObraID               int       `json:"obra_id" binding:"required"`
	Data                 string    `json:"data" binding:"required"` // "2024-10-08"
	Periodo              string    `json:"periodo"`
	AtividadesRealizadas string    `json:"atividades_realizadas" binding:"required"`
	Ocorrencias          string    `json:"ocorrencias,omitempty"`
	Observacoes          string    `json:"observacoes,omitempty"`
	ResponsavelID        int       `json:"responsavel_id,omitempty"`
	AprovadoPorID        int       `json:"aprovado_por_id,omitempty"`
	StatusAprovacao      string    `json:"status_aprovacao"`
	CreatedAt            time.Time `json:"createdAt"`
}
