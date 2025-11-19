package models

import "gopkg.in/guregu/null.v4"

// RelatorioFotografico representa o relatório fotográfico simplificado da obra
// Contém apenas: cabeçalho + resumo da obra + fotos
type RelatorioFotografico struct {
	CabecalhoEmpresa CabecalhoEmpresa `json:"cabecalho_empresa"`
	ResumoObra       ResumoObra       `json:"resumo_obra"`
	Fotos            []FotoDetalhada  `json:"fotos"`
}

// CabecalhoEmpresa contém os dados da empresa para o cabeçalho do relatório
type CabecalhoEmpresa struct {
	NomeEmpresa string      `json:"nome_empresa"`
	Logotipo    null.String `json:"logotipo"` // URL ou base64 do logo
}

// ResumoObra contém informações básicas da obra (SEM valores financeiros)
type ResumoObra struct {
	NomeObra          string      `json:"nome_obra"`
	Localizacao       string      `json:"localizacao"`
	ContratoNumero    null.String `json:"contrato_numero"`
	Lote              null.String `json:"lote"`
	DescricaoBreve    null.String `json:"descricao_breve"`    // O que é a obra
	FotoObra          null.String `json:"foto_obra"`          // Foto principal da obra
	InformacoesGerais null.String `json:"informacoes_gerais"` // Informações gerais sem valores
}

// FotoDetalhada representa uma foto com todas as informações necessárias
type FotoDetalhada struct {
	ID            int64       `json:"id"`
	URL           string      `json:"url"`            // Base64 ou URL da imagem
	TituloLegenda null.String `json:"titulo_legenda"` // Título ou legenda curta
	Data          null.String `json:"data"`           // Data da foto
	Observacao    null.String `json:"observacao"`     // Observação opcional
	Categoria     null.String `json:"categoria"`      // Ex: "DIARIO", "OBRA", etc.
}
