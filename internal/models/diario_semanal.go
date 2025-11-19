package models

import "gopkg.in/guregu/null.v4"

// DiarioSemanalRequest representa a requisição para gerar o diário de obras por período
type DiarioSemanalRequest struct {
	ObraID     int64  `json:"obra_id" binding:"required"`
	DataInicio string `json:"data_inicio" binding:"required"` // Formato: "2024-01-01"
	DataFim    string `json:"data_fim" binding:"required"`    // Formato: "2024-01-31"
}

// DiarioSemanalResponse representa o diário de obras estruturado por semanas
type DiarioSemanalResponse struct {
	DadosObra DadosObraDiario `json:"dados_obra"`
	Semanas   []SemanaDiario  `json:"semanas"`
}

// DadosObraDiario contém informações do cabeçalho da obra
type DadosObraDiario struct {
	NomeObra       string      `json:"nome_obra"`
	Localizacao    string      `json:"localizacao"`
	ContratoNumero null.String `json:"contrato_numero"`
	Contratante    null.String `json:"contratante"`
	Contratada     null.String `json:"contratada"`
}

// SemanaDiario representa uma semana do diário de obras
type SemanaDiario struct {
	Numero       int         `json:"numero"`        // Número da semana (1, 2, 3...)
	DataInicio   string      `json:"data_inicio"`   // Formato: "2024-01-01"
	DataFim      string      `json:"data_fim"`      // Formato: "2024-01-07"
	Descricao    null.String `json:"descricao"`     // Descrição do que foi executado (campo editável)
	DiasTrabalho []string    `json:"dias_trabalho"` // Lista de datas que tiveram registro
}
