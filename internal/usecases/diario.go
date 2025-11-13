package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"database/sql"
	"fmt"
	"strings"
)

type DiarioUseCase struct {
	services         services.DiarioServices
	relatorioService services.RelatorioService
	obraService      services.ObraServices
	pessoaService    services.PessoaServices
}

func NewDiarioUsecase(services services.DiarioServices, relatorioService services.RelatorioService, obraService services.ObraServices, pessoaService services.PessoaServices) DiarioUseCase {
	return DiarioUseCase{
		services:         services,
		relatorioService: relatorioService,
		obraService:      obraService,
		pessoaService:    pessoaService,
	}
}

func (pu *DiarioUseCase) CreateDiario(newDiario models.DiarioObra) (models.DiarioObra, error) {

	// Validação: relação entre status e aprovado_por_id
	status := ""
	if newDiario.StatusAprovacao.Valid {
		status = strings.ToUpper(newDiario.StatusAprovacao.String)
	}

	if status == "APROVADO" {
		if !newDiario.AprovadoPorID.Valid || newDiario.AprovadoPorID.Int64 == 0 {
			return models.DiarioObra{}, fmt.Errorf("aprovado_por_id é obrigatório quando status_aprovacao = APROVADO")
		}
	}

	if status == "PENDENTE" {
		if newDiario.AprovadoPorID.Valid && newDiario.AprovadoPorID.Int64 != 0 {
			return models.DiarioObra{}, fmt.Errorf("aprovado_por_id deve ser nulo quando status_aprovacao = PENDENTE")
		}
	}

	diarioId, err := pu.services.CreateDiario(newDiario)

	if err != nil {
		return models.DiarioObra{}, err
	}

	fmt.Println(diarioId, newDiario)

	newDiario.ID.Int64 = diarioId
	newDiario.ID.Valid = true

	return newDiario, nil
}

func (pu *DiarioUseCase) GetDiarios() ([]models.DiarioObra, error) {
	return pu.services.GetDiarios()
}

func (pu *DiarioUseCase) GetDiarioById(id int64) (models.DiarioObra, error) {

	diario, err := pu.services.GetDiarioById(id)

	if err != nil {
		return models.DiarioObra{}, fmt.Errorf("diario não encontrado")

	}

	return diario, err
}

func (pu *DiarioUseCase) GetDiariosByObraId(id int64) ([]models.DiarioObra, error) {
	return pu.services.GetDiarioByObraId(id)
}

func (pu *DiarioUseCase) PutDiario(id int, updatedDiario models.DiarioObra) (models.DiarioObra, error) {
	// Validação: relação entre status e aprovado_por_id
	status := ""
	if updatedDiario.StatusAprovacao.Valid {
		status = strings.ToUpper(updatedDiario.StatusAprovacao.String)
	}

	if status == "APROVADO" {
		if !updatedDiario.AprovadoPorID.Valid || updatedDiario.AprovadoPorID.Int64 == 0 {
			return models.DiarioObra{}, fmt.Errorf("aprovado_por_id é obrigatório quando status_aprovacao = APROVADO")
		}
	}

	if status == "PENDENTE" {
		if updatedDiario.AprovadoPorID.Valid && updatedDiario.AprovadoPorID.Int64 != 0 {
			return models.DiarioObra{}, fmt.Errorf("aprovado_por_id deve ser nulo quando status_aprovacao = PENDENTE")
		}
	}

	updatedDiario, err := pu.services.PutDiarios(id, updatedDiario)
	if err != nil {
		if err == sql.ErrNoRows {

			return models.DiarioObra{}, ErrUserNotFound
		}
		return models.DiarioObra{}, err
	}

	return updatedDiario, nil
}

func (pu *DiarioUseCase) DeleteDiariosById(id int) error {
	err := pu.services.DeleteDiarioById(id)
	if err != nil {
		// ✅ CORRETO: Comparar a MENSAGEM do erro, não o objeto
		switch err.Error() {
		case "nenhum diario encontrado com o ID fornecido":
			return fmt.Errorf("diario não encontrado")
		case "erro ao executar a query de delete":
			return fmt.Errorf("erro ao executar operação de delete: %w", err)
		case "erro ao obter linhas afetadas":
			return fmt.Errorf("erro ao verificar resultado: %w", err)
		default:
			return err
		}
	}

	return nil
}

// GetRelatorioDiarioCompleto retorna relatório completo de um diário
func (pu *DiarioUseCase) GetRelatorioDiarioCompleto(diarioId int64) (models.RelatorioDiarioCompleto, error) {
	return pu.relatorioService.GetRelatorioDiarioCompleto(diarioId)
}

// GetDiarioRelatorioFormatado retorna relatório de diário de obra formatado para impressão
func (pu *DiarioUseCase) GetDiarioRelatorioFormatado(obraId int64) (models.DiarioRelatorioCompleto, error) {
	// Buscar dados da obra
	obra, err := pu.obraService.GetObraById(obraId)
	if err != nil {
		return models.DiarioRelatorioCompleto{}, fmt.Errorf("erro ao buscar dados da obra: %w", err)
	}

	// Buscar diários da obra
	diarios, err := pu.services.GetDiarioByObraId(obraId)
	if err != nil {
		return models.DiarioRelatorioCompleto{}, fmt.Errorf("erro ao buscar diários: %w", err)
	}

	if len(diarios) == 0 {
		return models.DiarioRelatorioCompleto{}, fmt.Errorf("nenhum diário encontrado para a obra")
	}

	// Buscar dados do contratante (se disponível)
	var contratanteNome string = "N/A"
	if obra.ContratanteID.Valid {
		contratante, err := pu.pessoaService.GetPessoaById(obra.ContratanteID.Int64)
		if err == nil && contratante.Nome.Valid {
			contratanteNome = contratante.Nome.String
		}
	}

	// Buscar dados do responsável técnico (se disponível)
	var responsavelTecnico string = "N/A"
	var registroProfissional string = "N/A"
	if obra.ResponsavelID.Valid {
		responsavel, err := pu.pessoaService.GetPessoaById(obra.ResponsavelID.Int64)
		if err == nil {
			if responsavel.Nome.Valid {
				responsavelTecnico = responsavel.Nome.String
			}
			if responsavel.Documento.Valid {
				registroProfissional = "Registro: " + responsavel.Documento.String
			}
		}
	}

	// Calcular tempo decorrido
	tempoDecorrido := pu.calcularTempoDecorrido(obra.DataInicio.String)

	// Buscar equipe, equipamentos e materiais do banco de dados
	equipeData, _ := pu.services.GetEquipeByObraId(obraId)
	equipamentosData, _ := pu.services.GetEquipamentosByObraId(obraId)
	materiaisData, _ := pu.services.GetMateriaisByObraId(obraId)

	// Montar relatório formatado baseado nos dados reais do banco
	relatorio := models.DiarioRelatorioCompleto{
		InformacoesObra: models.InformacoesObra{
			Titulo:               obra.Nome.String,
			NumeroContrato:       obra.ContratoNumero.String,
			Contratante:          contratanteNome,
			PrazoObra:            fmt.Sprintf("%d DIAS", obra.PrazoDias.Int64),
			TempoDecorrido:       fmt.Sprintf("%d DIAS", tempoDecorrido),
			Contratada:           "N/A", // Pode ser adicionado posteriormente se houver campo para contratada
			ResponsavelTecnico:   responsavelTecnico,
			RegistroProfissional: registroProfissional,
		},
		TarefasRealizadas:      pu.formatarTarefasDoDiario(diarios),
		Ocorrencias:            pu.formatarOcorrenciasDoDiario(diarios),
		Fotos:                  pu.extrairFotosDoDiario(diarios),
		EquipeEnvolvida:        pu.converterEquipe(equipeData),
		EquipamentosUtilizados: pu.converterEquipamentos(equipamentosData),
		MateriaisUtilizados:    pu.converterMateriais(materiaisData),
		ResponsavelEmpresa: models.ResponsavelInfo{
			Nome:      responsavelTecnico,
			Cargo:     "Responsável Técnico",
			Documento: "",
			Empresa:   "N/A",
		},
		ResponsavelPrefeitura: models.ResponsavelInfo{
			Nome:      contratanteNome,
			Cargo:     "Fiscal da Obra",
			Documento: "",
			Empresa:   contratanteNome,
		},
	}

	return relatorio, nil
}

// Métodos auxiliares para formatação
func (pu *DiarioUseCase) formatarTarefasDoDiario(diarios []models.DiarioObra) []models.TarefaRealizada {
	var tarefas []models.TarefaRealizada

	for _, diario := range diarios {
		if diario.AtividadesRealizadas.Valid && diario.Data.Valid {
			// Separar múltiplas atividades por linha
			atividades := strings.Split(diario.AtividadesRealizadas.String, "\n")
			for _, atividade := range atividades {
				if strings.TrimSpace(atividade) != "" {
					tarefa := models.TarefaRealizada{
						Descricao: strings.TrimSpace(atividade),
						Data:      diario.Data.String,
					}
					tarefas = append(tarefas, tarefa)
				}
			}
		}
	}

	return tarefas
}

func (pu *DiarioUseCase) formatarOcorrenciasDoDiario(diarios []models.DiarioObra) []models.Ocorrencia {
	var ocorrencias []models.Ocorrencia

	for _, diario := range diarios {
		if diario.Ocorrencias.Valid && diario.Ocorrencias.String != "" && strings.ToLower(diario.Ocorrencias.String) != "não houve ocorrências" {
			ocorrencia := models.Ocorrencia{
				Descricao: diario.Ocorrencias.String,
				Tipo:      "OBSERVACAO",
			}
			ocorrencias = append(ocorrencias, ocorrencia)
		}
	}

	// Se não houver ocorrências, adicionar a padrão
	if len(ocorrencias) == 0 {
		ocorrencias = append(ocorrencias, models.Ocorrencia{
			Descricao: "Não houve ocorrências",
			Tipo:      "OBSERVACAO",
		})
	}

	return ocorrencias
}

func (pu *DiarioUseCase) converterEquipe(data []map[string]interface{}) []models.EquipeMembro {
	var equipe []models.EquipeMembro
	for _, item := range data {
		membro := models.EquipeMembro{
			Codigo:              item["codigo"].(string),
			Descricao:           item["descricao"].(string),
			QuantidadeUtilizada: item["quantidade_utilizada"].(int),
		}
		equipe = append(equipe, membro)
	}
	return equipe
}

func (pu *DiarioUseCase) converterEquipamentos(data []map[string]interface{}) []models.EquipamentoUtilizado {
	var equipamentos []models.EquipamentoUtilizado
	for _, item := range data {
		equipamento := models.EquipamentoUtilizado{
			Codigo:              item["codigo"].(string),
			Descricao:           item["descricao"].(string),
			QuantidadeUtilizada: item["quantidade_utilizada"].(int),
		}
		equipamentos = append(equipamentos, equipamento)
	}
	return equipamentos
}

func (pu *DiarioUseCase) converterMateriais(data []map[string]interface{}) []models.MaterialUtilizado {
	var materiais []models.MaterialUtilizado
	for _, item := range data {
		material := models.MaterialUtilizado{
			Codigo:        item["codigo"].(string),
			Descricao:     item["descricao"].(string),
			Quantidade:    item["quantidade"].(float64),
			Unidade:       item["unidade"].(string),
			Fornecedor:    item["fornecedor"].(string),
			ValorUnitario: item["valor_unitario"].(float64),
			ValorTotal:    item["valor_total"].(float64),
		}
		materiais = append(materiais, material)
	}
	return materiais
}

func (pu *DiarioUseCase) obterEquipePadrao() []models.EquipeMembro {
	return []models.EquipeMembro{
		{Codigo: "MESTRE", Descricao: "MESTRE DE OBRA", QuantidadeUtilizada: 1},
		{Codigo: "PEDREIRO", Descricao: "PEDREIRO", QuantidadeUtilizada: 5},
		{Codigo: "SERVENTE", Descricao: "SERVENTE", QuantidadeUtilizada: 8},
		{Codigo: "ELETRICISTA", Descricao: "ELETRICISTA", QuantidadeUtilizada: 1},
	}
}

func (pu *DiarioUseCase) obterEquipamentosPadrao() []models.EquipamentoUtilizado {
	return []models.EquipamentoUtilizado{
		{Codigo: "PA", Descricao: "PÁ", QuantidadeUtilizada: 4},
		{Codigo: "INCHADA", Descricao: "INCHADA", QuantidadeUtilizada: 2},
		{Codigo: "MARRETA", Descricao: "MARRETA", QuantidadeUtilizada: 3},
		{Codigo: "RETRO", Descricao: "RETRO ESCAVADEIRA", QuantidadeUtilizada: 1},
		{Codigo: "MAKITA", Descricao: "MAKITA", QuantidadeUtilizada: 3},
		{Codigo: "CACAMBA", Descricao: "CAÇAMBA", QuantidadeUtilizada: 1},
		{Codigo: "BETONEIRA", Descricao: "BETONEIRA", QuantidadeUtilizada: 1},
		{Codigo: "ENCERADEIRA", Descricao: "ENCERADEIRA DE PISO", QuantidadeUtilizada: 1},
		{Codigo: "CARRINHO", Descricao: "CARRINHO DE MÃO", QuantidadeUtilizada: 5},
	}
}

// Métodos auxiliares para cálculos e formatação
func (pu *DiarioUseCase) calcularTempoDecorrido(dataInicio string) int64 {
	// Implementação simples - pode ser melhorada com cálculos reais de data
	// Para agora, retorna um valor padrão baseado na quantidade de diários
	return 30 // Valor padrão, pode ser calculado baseado em diários ou data atual vs data de início
}

func (pu *DiarioUseCase) formatarEndereco(obra models.Obra) string {
	var endereco strings.Builder

	if obra.EnderecoRua.Valid && obra.EnderecoRua.String != "" {
		endereco.WriteString(obra.EnderecoRua.String)
		if obra.EnderecoNumero.Valid && obra.EnderecoNumero.String != "" {
			endereco.WriteString(", " + obra.EnderecoNumero.String)
		}
	}

	if obra.EnderecoBairro.Valid && obra.EnderecoBairro.String != "" {
		if endereco.Len() > 0 {
			endereco.WriteString(" - ")
		}
		endereco.WriteString(obra.EnderecoBairro.String)
	}

	if obra.EnderecoCidade.Valid && obra.EnderecoCidade.String != "" {
		if endereco.Len() > 0 {
			endereco.WriteString(", ")
		}
		endereco.WriteString(obra.EnderecoCidade.String)
	}

	if obra.EnderecoEstado.Valid && obra.EnderecoEstado.String != "" {
		if endereco.Len() > 0 {
			endereco.WriteString("/")
		}
		endereco.WriteString(obra.EnderecoEstado.String)
	}

	return endereco.String()
}

// extrairFotosDoDiario extrai todas as fotos dos diários
func (pu *DiarioUseCase) extrairFotosDoDiario(diarios []models.DiarioObra) []models.FotoInfo {
	var fotos []models.FotoInfo

	for _, diario := range diarios {
		if diario.Foto.Valid && diario.Foto.String != "" {
			foto := models.FotoInfo{
				ID:        diario.ID.Int64,
				URL:       diario.Foto.String,
				Timestamp: diario.Data.String,
				Categoria: "DIARIO",
			}
			fotos = append(fotos, foto)
		}
	}

	return fotos
}
