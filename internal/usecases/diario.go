package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
)

type DiarioUseCase struct {
	services services.DiarioServices
}

func NewDiarioUsecase(services services.DiarioServices) DiarioUseCase {
	return DiarioUseCase{
		services: services,
	}
}

func (pu *DiarioUseCase) CreateDiario(newDiario models.DiarioObra) (models.DiarioObra, error) {

	diarioId, err := pu.services.CreateDiario(newDiario)

	if err != nil {
		return models.DiarioObra{}, err
	}
	newDiario.ID = diarioId

	return newDiario, nil
}

func (pu *DiarioUseCase) GetDiarios() ([]models.DiarioObra, error) {
	return pu.services.GetDiarios()
}
