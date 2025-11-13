package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"fmt"
)

type EquipamentoDiarioUseCase struct {
	service services.EquipamentoDiarioService
}

func NewEquipamentoDiarioUseCase(service services.EquipamentoDiarioService) EquipamentoDiarioUseCase {
	return EquipamentoDiarioUseCase{service: service}
}

func (uc *EquipamentoDiarioUseCase) Create(equipamento models.EquipamentoDiario) (models.EquipamentoDiario, error) {
	id, err := uc.service.Create(equipamento)
	if err != nil {
		return models.EquipamentoDiario{}, err
	}
	equipamento.ID.Int64 = id
	equipamento.ID.Valid = true
	return equipamento, nil
}

func (uc *EquipamentoDiarioUseCase) GetByDiarioId(diarioId int64) ([]models.EquipamentoDiario, error) {
	return uc.service.GetByDiarioId(diarioId)
}

func (uc *EquipamentoDiarioUseCase) Update(id int, equipamento models.EquipamentoDiario) (models.EquipamentoDiario, error) {
	return uc.service.Update(id, equipamento)
}

func (uc *EquipamentoDiarioUseCase) Delete(id int) error {
	err := uc.service.Delete(id)
	if err != nil {
		if err.Error() == "nenhum equipamento encontrado com o ID fornecido" {
			return fmt.Errorf("equipamento não encontrado")
		}
		return err
	}
	return nil
}
