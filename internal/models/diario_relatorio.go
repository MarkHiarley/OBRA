package models

import (
	"gopkg.in/guregu/null.v4"
)

// DiarioRelatorioCompleto representa o relatório de diário de obra formatado
type DiarioRelatorioCompleto struct {
	InformacoesObra        InformacoesObra        `json:"informacoes_obra"`
	TarefasRealizadas      []TarefaRealizada      `json:"tarefas_realizadas"`
	Ocorrencias            []Ocorrencia           `json:"ocorrencias"`
	EquipeEnvolvida        []EquipeMembro         `json:"equipe_envolvida"`
	EquipamentosUtilizados []EquipamentoUtilizado `json:"equipamentos_utilizados"`
	MateriaisUtilizados    []MaterialUtilizado    `json:"materiais_utilizados"`
	Fotos                  []FotoInfo             `json:"fotos"`
	ResponsavelEmpresa     ResponsavelInfo        `json:"responsavel_empresa"`
	ResponsavelPrefeitura  ResponsavelInfo        `json:"responsavel_prefeitura"`
}

// InformacoesObra contém as informações principais da obra
type InformacoesObra struct {
	Titulo               string `json:"titulo"`
	NumeroContrato       string `json:"numero_contrato"`
	Contratante          string `json:"contratante"`
	PrazoObra            string `json:"prazo_obra"`
	TempoDecorrido       string `json:"tempo_decorrido"`
	Contratada           string `json:"contratada"`
	ResponsavelTecnico   string `json:"responsavel_tecnico"`
	RegistroProfissional string `json:"registro_profissional"`
}

// TarefaRealizada representa uma tarefa executada em um dia específico
type TarefaRealizada struct {
	Descricao string `json:"descricao"`
	Data      string `json:"data"`
}

// Ocorrencia representa uma ocorrência no canteiro de obras
type Ocorrencia struct {
	Descricao string `json:"descricao"`
	Tipo      string `json:"tipo"` // "INCIDENTE", "OBSERVACAO", "PROBLEMA", etc.
}

// EquipeMembro representa um membro ou função da equipe
type EquipeMembro struct {
	Codigo              string `json:"codigo"`
	Descricao           string `json:"descricao"`
	QuantidadeUtilizada int    `json:"quantidade_utilizada"`
}

// EquipamentoUtilizado representa um equipamento utilizado na obra
type EquipamentoUtilizado struct {
	Codigo              string `json:"codigo"`
	Descricao           string `json:"descricao"`
	QuantidadeUtilizada int    `json:"quantidade_utilizada"`
}

// MaterialUtilizado representa um material utilizado na obra
type MaterialUtilizado struct {
	Codigo        string  `json:"codigo"`
	Descricao     string  `json:"descricao"`
	Quantidade    float64 `json:"quantidade"`
	Unidade       string  `json:"unidade"`
	Fornecedor    string  `json:"fornecedor,omitempty"`
	ValorUnitario float64 `json:"valor_unitario,omitempty"`
	ValorTotal    float64 `json:"valor_total,omitempty"`
}

// ResponsavelInfo representa informações do responsável
type ResponsavelInfo struct {
	Nome      string `json:"nome"`
	Cargo     string `json:"cargo"`
	Documento string `json:"documento"`
	Empresa   string `json:"empresa"`
}

// DiarioObraRequest representa a estrutura para criação do relatório
type DiarioObraRequest struct {
	ObraID                  int64    `json:"obra_id" binding:"required"`
	DataInicio              string   `json:"data_inicio" binding:"required"`
	DataFim                 string   `json:"data_fim" binding:"required"`
	ResponsavelEmpresaID    null.Int `json:"responsavel_empresa_id"`
	ResponsavelPrefeituraID null.Int `json:"responsavel_prefeitura_id"`
}
