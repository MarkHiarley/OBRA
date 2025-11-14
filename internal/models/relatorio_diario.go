package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// RelatorioDiarioCompleto representa o relatório completo de um diário de obra
type RelatorioDiarioCompleto struct {
	// Informações da obra
	Obra InfoObra `json:"obra"`

	// Informações do diário
	Diario DiarioInfo `json:"diario"`

	Clima ClimaInfo `json:"clima"`

	Equipe []EquipeInfo `json:"equipe"`

	Equipamentos []EquipamentoInfo `json:"equipamentos"`

	Materiais []MaterialInfo `json:"materiais"`

	Fotos FotosInfo `json:"fotos"`

	Progresso ProgressoInfo `json:"progresso"`

	Ocorrencias []OcorrenciaInfo `json:"ocorrencias"`

	Observacoes ObservacoesInfo `json:"observacoes_detalhadas"`
}

type InfoObra struct {
	ID                  int64       `json:"id"`
	Nome                string      `json:"nome"`
	ContratoNumero      null.String `json:"contrato_numero"`
	ContratanteName     null.String `json:"contratante_nome"`
	ContratadaName      null.String `json:"contratada_nome"`
	ResponsavelTecnico  null.String `json:"responsavel_tecnico"`
	DataInicio          null.Time   `json:"data_inicio"`
	PrazoDias           null.Int    `json:"prazo_dias"`
	TempoDecorrido      null.Int    `json:"tempo_decorrido"`
	Status              string      `json:"status"`
	EnderecoCompleto    null.String `json:"endereco_completo"`
	PercentualConcluido float64     `json:"percentual_concluido"`
}

// DiarioInfo contém informações específicas do diário
type DiarioInfo struct {
	ID                   int64       `json:"id"`
	Data                 string      `json:"data"`
	Periodo              null.String `json:"periodo"`
	AtividadesRealizadas string      `json:"atividades_realizadas"`
	Observacoes          null.String `json:"observacoes"`
	Foto                 null.String `json:"foto"`
	ResponsavelID        null.Int    `json:"responsavel_id"`
	ResponsavelNome      null.String `json:"responsavel_nome"`
	AprovadoPorID        null.Int    `json:"aprovado_por_id"`
	AprovadoPorNome      null.String `json:"aprovado_por_nome"`
	StatusAprovacao      string      `json:"status_aprovacao"`
	CreatedAt            time.Time   `json:"created_at"`
	UpdatedAt            null.Time   `json:"updated_at"`
}

// ClimaInfo contém informações climáticas detalhadas
type ClimaInfo struct {
	Condicao        null.String `json:"condicao"`
	TemperaturaMin  null.Int    `json:"temperatura_min"`
	TemperaturaMax  null.Int    `json:"temperatura_max"`
	Umidade         null.String `json:"umidade"`
	VentoCondicao   null.String `json:"vento"`
	Observacoes     null.String `json:"observacoes"`
	ImpactoTrabalho null.String `json:"impacto_trabalho"`
}

// EquipeInfo representa um membro ou grupo da equipe
type EquipeInfo struct {
	ID                int64       `json:"id"`
	Funcao            string      `json:"funcao"`
	Nome              null.String `json:"nome"`
	Quantidade        int         `json:"quantidade"`
	HorasTrabalhadas  float64     `json:"horas_trabalhadas"`
	PeriodoTrabalho   null.String `json:"periodo_trabalho"`
	Observacoes       null.String `json:"observacoes"`
	ProdutividadeNota null.Int    `json:"produtividade_nota"`
}

// EquipamentoInfo representa equipamento utilizado
type EquipamentoInfo struct {
	ID                int64       `json:"id"`
	Nome              string      `json:"nome"`
	Codigo            null.String `json:"codigo"`
	Tipo              null.String `json:"tipo"`
	HorasUso          float64     `json:"horas_uso"`
	CombustivelGasto  null.Float  `json:"combustivel_gasto"`
	EstadoConservacao null.String `json:"estado_conservacao"`
	Observacoes       null.String `json:"observacoes"`
	ProximaManutencao null.String `json:"proxima_manutencao"`
}

// MaterialInfo representa material utilizado/consumido
type MaterialInfo struct {
	ID             int64       `json:"id"`
	Descricao      string      `json:"descricao"`
	Quantidade     float64     `json:"quantidade"`
	Unidade        string      `json:"unidade"`
	Fornecedor     null.String `json:"fornecedor"`
	NumeroNota     null.String `json:"numero_nota"`
	ValorUnitario  null.Float  `json:"valor_unitario"`
	ValorTotal     null.Float  `json:"valor_total"`
	LocalAplicacao null.String `json:"local_aplicacao"`
	Observacoes    null.String `json:"observacoes"`
}

// FotosInfo organiza as fotos por momento do dia
type FotosInfo struct {
	AntesInicio     []FotoInfo `json:"antes_inicio"`
	DuranteExecucao []FotoInfo `json:"durante_execucao"`
	FimJornada      []FotoInfo `json:"fim_jornada"`
	Detalhes        []FotoInfo `json:"detalhes"`
	Problemas       []FotoInfo `json:"problemas"`
}

// FotoInfo representa uma foto individual
// FotoInfo contém informações sobre fotos anexadas ao diário
type FotoInfo struct {
	ID        int64       `json:"id"`
	URL       string      `json:"url"` // Base64 encoded image (data:image/jpeg;base64,...)
	Descricao null.String `json:"descricao"`
	Timestamp string      `json:"timestamp"`
	LocalFoto null.String `json:"local_foto"`
	Categoria string      `json:"categoria"` // Ex: "DIARIO", "OBRA", "OCORRENCIA"
}

// ProgressoInfo contém informações de progresso e medições
type ProgressoInfo struct {
	PercentualDia       float64     `json:"percentual_dia"`
	PercentualObraGeral float64     `json:"percentual_obra_geral"`
	AreasExecutadas     interface{} `json:"areas_executadas"` // JSON object
	MetaDia             null.String `json:"meta_dia"`
	MetaCumprida        bool        `json:"meta_cumprida"`
	ProximaAtividade    null.String `json:"proxima_atividade"`
	PrazoEstimado       null.String `json:"prazo_estimado"`
	Observacoes         null.String `json:"observacoes"`
}

// OcorrenciaInfo representa uma ocorrência ou problema
type OcorrenciaInfo struct {
	ID              int64       `json:"id"`
	Tipo            string      `json:"tipo"`
	Gravidade       string      `json:"gravidade"`
	Descricao       string      `json:"descricao"`
	AcaoTomada      null.String `json:"acao_tomada"`
	Responsavel     null.String `json:"responsavel"`
	StatusResolucao string      `json:"status_resolucao"`
	PrazoResolucao  null.Time   `json:"prazo_resolucao"`
	Fotos           []string    `json:"fotos"`
	Observacoes     null.String `json:"observacoes"`
}

// ObservacoesInfo contém observações detalhadas organizadas
type ObservacoesInfo struct {
	Geral         null.String `json:"geral"`
	Qualidade     null.String `json:"qualidade"`
	Seguranca     null.String `json:"seguranca"`
	Produtividade null.String `json:"produtividade"`
	Melhorias     null.String `json:"melhorias"`
	ProximoDia    null.String `json:"proximo_dia"`
}
