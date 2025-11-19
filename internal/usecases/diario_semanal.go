package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"fmt"
)

type DiarioSemanalUseCase struct {
	service services.DiarioSemanalService
}

func NewDiarioSemanalUseCase(service services.DiarioSemanalService) DiarioSemanalUseCase {
	return DiarioSemanalUseCase{
		service: service,
	}
}

// GetDiarioSemanal retorna o diário de obras agrupado por semana
func (uc *DiarioSemanalUseCase) GetDiarioSemanal(request models.DiarioSemanalRequest) (models.DiarioSemanalResponse, error) {
	// Validar datas
	if request.DataInicio == "" || request.DataFim == "" {
		return models.DiarioSemanalResponse{}, fmt.Errorf("data_inicio e data_fim são obrigatórios")
	}

	if request.ObraID <= 0 {
		return models.DiarioSemanalResponse{}, fmt.Errorf("obra_id inválido")
	}

	return uc.service.GetDiarioSemanal(request.ObraID, request.DataInicio, request.DataFim)
}
