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
// AGORA USA A NOVA ARQUITETURA NORMALIZADA (atividade_diaria, ocorrencia_diaria, diario_metadados)
func (pu *DiarioUseCase) GetDiarioRelatorioFormatado(obraId int64) (models.DiarioRelatorioCompleto, error) {
	// Buscar dados da obra
	obra, err := pu.obraService.GetObraById(obraId)
	if err != nil {
		return models.DiarioRelatorioCompleto{}, fmt.Errorf("erro ao buscar dados da obra: %w", err)
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

	// NOVA ARQUITETURA: Buscar do diário consolidado ao invés de diario_obra legado
	diariosConsolidados, err := pu.services.GetDiariosConsolidadosByObra(obraId)
	if err != nil {
		return models.DiarioRelatorioCompleto{}, fmt.Errorf("erro ao buscar diários consolidados: %w", err)
	}

	if len(diariosConsolidados) == 0 {
		return models.DiarioRelatorioCompleto{}, fmt.Errorf("nenhum diário encontrado para a obra")
	}

	// Buscar equipe, equipamentos e materiais do banco de dados
	equipeData, _ := pu.services.GetEquipeByObraId(obraId)
	equipamentosData, _ := pu.services.GetEquipamentosByObraId(obraId)
	materiaisData, _ := pu.services.GetMateriaisByObraId(obraId)

	// Definir contratada
	contratada := "N/A"
	if obra.Contratada.Valid && obra.Contratada.String != "" {
		contratada = obra.Contratada.String
	}

	// Montar relatório formatado baseado nos dados da NOVA ARQUITETURA
	relatorio := models.DiarioRelatorioCompleto{
		InformacoesObra: models.InformacoesObra{
			Titulo:               obra.Nome.String,
			NumeroContrato:       obra.ContratoNumero.String,
			Contratante:          contratanteNome,
			PrazoObra:            fmt.Sprintf("%d DIAS", obra.PrazoDias.Int64),
			TempoDecorrido:       fmt.Sprintf("%d DIAS", tempoDecorrido),
			Contratada:           contratada,
			ResponsavelTecnico:   responsavelTecnico,
			RegistroProfissional: registroProfissional,
		},
		TarefasRealizadas:      pu.formatarTarefasDoDiarioConsolidado(diariosConsolidados),
		Ocorrencias:            pu.formatarOcorrenciasDoDiarioConsolidado(diariosConsolidados),
		Fotos:                  pu.extrairFotosDoDiarioConsolidado(diariosConsolidados),
		EquipeEnvolvida:        pu.converterEquipe(equipeData),
		EquipamentosUtilizados: pu.converterEquipamentos(equipamentosData),
		MateriaisUtilizados:    pu.converterMateriais(materiaisData),
		ResponsavelEmpresa: models.ResponsavelInfo{
			Nome:      responsavelTecnico,
			Cargo:     "Responsável Técnico",
			Documento: "",
			Empresa:   contratada,
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

// Métodos auxiliares para formatação usando NOVA ARQUITETURA

// formatarTarefasDoDiarioConsolidado extrai atividades da view consolidada
func (pu *DiarioUseCase) formatarTarefasDoDiarioConsolidado(diarios []models.DiarioConsolidado) []models.TarefaRealizada {
	var tarefas []models.TarefaRealizada

	for _, diario := range diarios {
		// A view já agrega as atividades como string
		if diario.AtividadesRealizadas.Valid && diario.AtividadesRealizadas.String != "" {
			// Separar múltiplas atividades (formato: "desc1 (status - 50%); desc2 (status - 100%)")
			atividades := strings.Split(diario.AtividadesRealizadas.String, ";")
			for _, atividade := range atividades {
				atividadeTrimmed := strings.TrimSpace(atividade)
				if atividadeTrimmed != "" {
					tarefa := models.TarefaRealizada{
						Descricao: atividadeTrimmed,
						Data:      diario.Data.String,
					}
					tarefas = append(tarefas, tarefa)
				}
			}
		}
	}

	return tarefas
}

// formatarOcorrenciasDoDiarioConsolidado extrai ocorrências da view consolidada
func (pu *DiarioUseCase) formatarOcorrenciasDoDiarioConsolidado(diarios []models.DiarioConsolidado) []models.Ocorrencia {
	var ocorrencias []models.Ocorrencia

	for _, diario := range diarios {
		// A view já agrega as ocorrências como string (formato: "[ALTA] desc - status; [MEDIA] desc2 - status2")
		if diario.Ocorrencias.Valid && diario.Ocorrencias.String != "" {
			// Separar múltiplas ocorrências
			ocorrenciasList := strings.Split(diario.Ocorrencias.String, ";")
			for _, ocorrencia := range ocorrenciasList {
				ocorrenciaTrimmed := strings.TrimSpace(ocorrencia)
				if ocorrenciaTrimmed != "" {
					// Extrair tipo/gravidade se estiver no formato [GRAVIDADE] descrição
					tipo := "OBSERVACAO"
					descricao := ocorrenciaTrimmed

					if strings.HasPrefix(ocorrenciaTrimmed, "[") {
						endBracket := strings.Index(ocorrenciaTrimmed, "]")
						if endBracket > 0 {
							gravidade := ocorrenciaTrimmed[1:endBracket]
							descricao = strings.TrimSpace(ocorrenciaTrimmed[endBracket+1:])

							// Mapear gravidade para tipo
							switch strings.ToUpper(gravidade) {
							case "ALTA", "CRITICA":
								tipo = "CRITICO"
							case "MEDIA":
								tipo = "IMPORTANTE"
							case "BAIXA":
								tipo = "OBSERVACAO"
							}
						}
					}

					ocorrenciaObj := models.Ocorrencia{
						Descricao: descricao,
						Tipo:      tipo,
					}
					ocorrencias = append(ocorrencias, ocorrenciaObj)
				}
			}
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

// extrairFotosDoDiarioConsolidado extrai fotos dos metadados
func (pu *DiarioUseCase) extrairFotosDoDiarioConsolidado(diarios []models.DiarioConsolidado) []models.FotoInfo {
	var fotos []models.FotoInfo

	for _, diario := range diarios {
		if diario.Foto.Valid && diario.Foto.String != "" {
			foto := models.FotoInfo{
				ID:        diario.ObraID.Int64, // Usar obra_id como ID da foto
				URL:       diario.Foto.String,
				Timestamp: diario.Data.String,
				Categoria: "DIARIO",
			}
			fotos = append(fotos, foto)
		}
	}

	return fotos
}

// Métodos auxiliares para conversores (ainda usados para equipe, equipamentos, materiais)

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

// Métodos auxiliares para cálculos e formatação
func (pu *DiarioUseCase) calcularTempoDecorrido(dataInicio string) int64 {
	// Implementação simples - pode ser melhorada com cálculos reais de data
	// Para agora, retorna um valor padrão baseado na quantidade de diários
	return 30 // Valor padrão, pode ser calculado baseado em diários ou data atual vs data de início
}
