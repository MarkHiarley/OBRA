package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"fmt"
)

type RelatorioUseCase struct {
	relatorioService services.RelatorioService
	obraService      services.ObraServices
}

func NewRelatorioUseCase(relatorioService services.RelatorioService, obraService services.ObraServices) RelatorioUseCase {
	return RelatorioUseCase{
		relatorioService: relatorioService,
		obraService:      obraService,
	}
}

func (ruc *RelatorioUseCase) GetRelatorioObra(obraId int) (models.RelatorioObra, error) {
	if obraId <= 0 {
		return models.RelatorioObra{}, fmt.Errorf("ID da obra inválido")
	}

	// Verificar se a obra existe
	_, err := ruc.obraService.GetObraById(int64(obraId))
	if err != nil {
		return models.RelatorioObra{}, fmt.Errorf("obra não encontrada")
	}

	relatorio, err := ruc.relatorioService.GetRelatorioObra(int64(obraId))
	if err != nil {
		return models.RelatorioObra{}, fmt.Errorf("erro ao gerar relatório da obra: %v", err)
	}

	return relatorio, nil
}

func (ruc *RelatorioUseCase) GetRelatorioDespesasPorCategoria(obraId int) ([]models.RelatorioFinanceiroPorCategoria, error) {
	if obraId <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}

	// Verificar se a obra existe
	_, err := ruc.obraService.GetObraById(int64(obraId))
	if err != nil {
		return nil, fmt.Errorf("obra não encontrada")
	}

	relatorio, err := ruc.relatorioService.GetRelatorioDespesasPorCategoria(int64(obraId))
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar relatório de despesas: %v", err)
	}

	return relatorio, nil
}

func (ruc *RelatorioUseCase) GetRelatorioPagamentos(obraId int, status string) ([]models.RelatorioPagamentos, error) {
	if obraId <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}

	// Verificar se a obra existe
	_, err := ruc.obraService.GetObraById(int64(obraId))
	if err != nil {
		return nil, fmt.Errorf("obra não encontrada")
	}

	// Validar status se fornecido
	if status != "" && status != "PENDENTE" && status != "PAGO" && status != "CANCELADO" {
		return nil, fmt.Errorf("status inválido. Use: PENDENTE, PAGO ou CANCELADO")
	}

	relatorio, err := ruc.relatorioService.GetRelatorioPagamentos(int64(obraId), status)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar relatório de pagamentos: %v", err)
	}

	return relatorio, nil
}

func (ruc *RelatorioUseCase) GetRelatorioMateriais(obraId int) (models.RelatorioMateriais, error) {
	if obraId <= 0 {
		return models.RelatorioMateriais{}, fmt.Errorf("ID da obra inválido")
	}

	// Verificar se a obra existe
	_, err := ruc.obraService.GetObraById(int64(obraId))
	if err != nil {
		return models.RelatorioMateriais{}, fmt.Errorf("obra não encontrada")
	}

	relatorio, err := ruc.relatorioService.GetRelatorioMateriais(int64(obraId))
	if err != nil {
		return models.RelatorioMateriais{}, fmt.Errorf("erro ao gerar relatório de materiais: %v", err)
	}

	return relatorio, nil
}

func (ruc *RelatorioUseCase) GetRelatorioProfissionais(obraId int) (models.RelatorioProfissionais, error) {
	if obraId <= 0 {
		return models.RelatorioProfissionais{}, fmt.Errorf("ID da obra inválido")
	}

	// Verificar se a obra existe
	_, err := ruc.obraService.GetObraById(int64(obraId))
	if err != nil {
		return models.RelatorioProfissionais{}, fmt.Errorf("obra não encontrada")
	}

	relatorio, err := ruc.relatorioService.GetRelatorioProfissionais(int64(obraId))
	if err != nil {
		return models.RelatorioProfissionais{}, fmt.Errorf("erro ao gerar relatório de profissionais: %v", err)
	}

	return relatorio, nil
}