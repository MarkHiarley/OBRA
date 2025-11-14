package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// AtividadeDiaria representa uma atividade individual realizada em um dia específico
type AtividadeDiaria struct {
	ID                  null.Int     `json:"id"`
	ObraID              null.Int     `json:"obra_id" binding:"required"`
	Data                null.String  `json:"data" binding:"required"` // "2024-10-08"
	Periodo             null.String  `json:"periodo"`                 // manha, tarde, integral, noite
	Descricao           null.String  `json:"descricao" binding:"required"`
	ResponsavelID       null.Int     `json:"responsavel_id,omitempty"`
	Status              null.String  `json:"status"`               // planejada, em_andamento, concluida, cancelada
	PercentualConclusao null.Int     `json:"percentual_conclusao"` // 0-100
	Observacao          null.String  `json:"observacao,omitempty"`
	Fotos               []FotoDiario `json:"fotos,omitempty"` // Array de fotos relacionadas à atividade
	CreatedAt           time.Time    `json:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at"`
}

// AtividadeDiariaComRelacionamentos inclui dados relacionados
type AtividadeDiariaComRelacionamentos struct {
	AtividadeDiaria
	ObraNome        null.String `json:"obra_nome,omitempty"`
	ResponsavelNome null.String `json:"responsavel_nome,omitempty"`
}

// OcorrenciaDiaria representa uma ocorrência/problema registrado em um dia específico
type OcorrenciaDiaria struct {
	ID              null.Int     `json:"id"`
	ObraID          null.Int     `json:"obra_id" binding:"required"`
	Data            null.String  `json:"data" binding:"required"` // "2024-10-08"
	Periodo         null.String  `json:"periodo"`                 // manha, tarde, integral, noite
	Tipo            null.String  `json:"tipo"`                    // seguranca, qualidade, prazo, custo, clima, equipamento, material, geral
	Gravidade       null.String  `json:"gravidade"`               // baixa, media, alta, critica
	Descricao       null.String  `json:"descricao" binding:"required"`
	ResponsavelID   null.Int     `json:"responsavel_id,omitempty"`
	StatusResolucao null.String  `json:"status_resolucao"` // pendente, em_analise, resolvida, nao_aplicavel
	AcaoTomada      null.String  `json:"acao_tomada,omitempty"`
	Fotos           []FotoDiario `json:"fotos,omitempty"` // Array de fotos relacionadas à ocorrência
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

// OcorrenciaDiariaComRelacionamentos inclui dados relacionados
type OcorrenciaDiariaComRelacionamentos struct {
	OcorrenciaDiaria
	ObraNome        null.String `json:"obra_nome,omitempty"`
	ResponsavelNome null.String `json:"responsavel_nome,omitempty"`
}

// FotoDiario armazena múltiplas fotos relacionadas a diários, atividades ou ocorrências
type FotoDiario struct {
	ID           null.Int    `json:"id"`
	EntidadeTipo string      `json:"entidade_tipo"`           // metadados, atividade, ocorrencia
	EntidadeID   null.Int    `json:"entidade_id"`             // ID da entidade relacionada
	Foto         string      `json:"foto" binding:"required"` // Base64 encoded image (data:image/jpeg;base64,...)
	Descricao    null.String `json:"descricao,omitempty"`     // Descrição da foto
	Ordem        null.Int    `json:"ordem"`                   // Ordem de exibição (0 = primeira)
	Categoria    null.String `json:"categoria"`               // DIARIO, OBRA, OCORRENCIA, ATIVIDADE, SEGURANCA
	Largura      null.Int    `json:"largura,omitempty"`       // Largura da imagem em pixels
	Altura       null.Int    `json:"altura,omitempty"`        // Altura da imagem em pixels
	TamanhoBytes null.Int    `json:"tamanho_bytes,omitempty"` // Tamanho em bytes
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

// DiarioMetadados armazena informações complementares do diário (foto, aprovação, observações)
type DiarioMetadados struct {
	ID              null.Int     `json:"id"`
	ObraID          null.Int     `json:"obra_id" binding:"required"`
	Data            null.String  `json:"data" binding:"required"` // "2024-10-08"
	Periodo         null.String  `json:"periodo"`                 // manha, tarde, integral, noite
	Foto            null.String  `json:"foto,omitempty"`          // DEPRECATED: Use Fotos[]
	Fotos           []FotoDiario `json:"fotos,omitempty"`         // Array de fotos em Base64
	Observacoes     null.String  `json:"observacoes,omitempty"`
	ResponsavelID   null.Int     `json:"responsavel_id,omitempty"`
	AprovadoPorID   null.Int     `json:"aprovado_por_id,omitempty"`
	StatusAprovacao null.String  `json:"status_aprovacao"` // pendente, aprovado, rejeitado
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

// DiarioConsolidado representa a view agregada (gerada dinamicamente)
type DiarioConsolidado struct {
	DiarioID             null.Int    `json:"diario_id"`
	ObraID               null.Int    `json:"obra_id"`
	ObraNome             null.String `json:"obra_nome"`
	Data                 null.String `json:"data"`
	Periodo              null.String `json:"periodo"`
	AtividadesRealizadas null.String `json:"atividades"`
	Ocorrencias          null.String `json:"ocorrencias"`
	Foto                 null.String `json:"foto,omitempty"`
	Observacoes          null.String `json:"observacoes,omitempty"`
	ResponsavelID        null.Int    `json:"responsavel_id,omitempty"`
	ResponsavelNome      null.String `json:"responsavel_nome,omitempty"`
	AprovadoPorID        null.Int    `json:"aprovado_por_id,omitempty"`
	AprovadoPorNome      null.String `json:"aprovado_por_nome,omitempty"`
	StatusAprovacao      null.String `json:"status_aprovacao"`
	QtdAtividades        null.Int    `json:"qtd_atividades"`
	QtdOcorrencias       null.Int    `json:"qtd_ocorrencias"`
	QtdEquipe            null.Int    `json:"qtd_equipe"`
	QtdEquipamentos      null.Int    `json:"qtd_equipamentos"`
	QtdMateriais         null.Int    `json:"qtd_materiais"`
	CreatedAt            time.Time   `json:"created_at"`
	UpdatedAt            null.Time   `json:"updated_at"`
}

// Constantes para tipos de ocorrências
const (
	TipoOcorrenciaSeguranca   = "seguranca"
	TipoOcorrenciaQualidade   = "qualidade"
	TipoOcorrenciaPrazo       = "prazo"
	TipoOcorrenciaCusto       = "custo"
	TipoOcorrenciaClima       = "clima"
	TipoOcorrenciaEquipamento = "equipamento"
	TipoOcorrenciaMaterial    = "material"
	TipoOcorrenciaGeral       = "geral"
)

// Constantes para gravidade de ocorrências
const (
	GravidadeBaixa   = "baixa"
	GravidadeMedia   = "media"
	GravidadeAlta    = "alta"
	GravidadeCritica = "critica"
)

// Constantes para status de atividades
const (
	StatusAtividadePlanejada   = "planejada"
	StatusAtividadeEmAndamento = "em_andamento"
	StatusAtividadeConcluida   = "concluida"
	StatusAtividadeCancelada   = "cancelada"
)

// Constantes para status de resolução de ocorrências
const (
	StatusResolucaoPendente     = "pendente"
	StatusResolucaoEmAnalise    = "em_analise"
	StatusResolucaoResolvida    = "resolvida"
	StatusResolucaoNaoAplicavel = "nao_aplicavel"
)
