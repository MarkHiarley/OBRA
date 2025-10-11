package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
)

type ObraUseCase struct {
	services services.ObraServices
}

func NewObraUsecase(services services.ObraServices) ObraUseCase {
	return ObraUseCase{
		services: services,
	}
}

func (pu *ObraUseCase) CreateObra(newObra models.Obra) (models.Obra, error) {

	obraId, err := pu.services.CreateObra(newObra)

	if err != nil {
		return models.Obra{}, err
	}
	newObra.ID = obraId

	return newObra, nil
}
