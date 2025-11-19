package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
)

type RelatorioFotograficoUseCase struct {
	service services.RelatorioFotograficoService
}

func NewRelatorioFotograficoUseCase(service services.RelatorioFotograficoService) RelatorioFotograficoUseCase {
	return RelatorioFotograficoUseCase{
		service: service,
	}
}

// GetRelatorioFotografico retorna o relatório fotográfico da obra
// Contém apenas: cabeçalho + resumo + fotos
func (uc *RelatorioFotograficoUseCase) GetRelatorioFotografico(obraID int64) (models.RelatorioFotografico, error) {
	return uc.service.GetRelatorioFotografico(obraID)
}
