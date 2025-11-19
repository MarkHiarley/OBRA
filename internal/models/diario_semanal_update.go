package models

import "gopkg.in/guregu/null.v4"

// DiarioSemanalUpdateRequest representa a requisição para atualizar descrição de uma semana
type DiarioSemanalUpdateRequest struct {
	ObraID     int64       `json:"obra_id" binding:"required"`
	Semana     int         `json:"semana" binding:"required"`      // Número da semana (1, 2, 3...)
	DataInicio string      `json:"data_inicio" binding:"required"` // Data início da semana
	DataFim    string      `json:"data_fim" binding:"required"`    // Data fim da semana
	Descricao  null.String `json:"descricao"`                      // Descrição editada pelo usuário
}
