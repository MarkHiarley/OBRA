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

func (uc *MaterialDiarioUseCase) GetByObraId(obraId int) ([]models.MaterialDiario, error) {
	if obraId <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}
	return uc.service.GetByObraId(int64(obraId))
}

func (uc *MaterialDiarioUseCase) GetByObraAndData(obraId int, data string) ([]models.MaterialDiario, error) {
	if obraId <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}
	if data == "" {
		return nil, fmt.Errorf("data inválida")
	}
	return uc.service.GetByObraAndData(int64(obraId), data)
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
