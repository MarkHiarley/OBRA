package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"fmt"
)

type MaterialDiarioUseCase struct {
	service services.MaterialDiarioService
}

func NewMaterialDiarioUseCase(service services.MaterialDiarioService) MaterialDiarioUseCase {
	return MaterialDiarioUseCase{service: service}
}

func (uc *MaterialDiarioUseCase) Create(material models.MaterialDiario) (models.MaterialDiario, error) {
	id, err := uc.service.Create(material)
	if err != nil {
		return models.MaterialDiario{}, err
	}
	material.ID.Int64 = id
	material.ID.Valid = true
	return material, nil
}

func (uc *MaterialDiarioUseCase) GetByDiarioId(diarioId int64) ([]models.MaterialDiario, error) {
	return uc.service.GetByDiarioId(diarioId)
}

func (uc *MaterialDiarioUseCase) Update(id int, material models.MaterialDiario) (models.MaterialDiario, error) {
	return uc.service.Update(id, material)
}

func (uc *MaterialDiarioUseCase) Delete(id int) error {
	err := uc.service.Delete(id)
	if err != nil {
		if err.Error() == "nenhum material encontrado com o ID fornecido" {
			return fmt.Errorf("material não encontrado")
		}
		return err
	}
	return nil
}
