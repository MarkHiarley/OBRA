package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"fmt"
)

type EquipeDiarioUseCase struct {
	service services.EquipeDiarioService
}

func NewEquipeDiarioUseCase(service services.EquipeDiarioService) EquipeDiarioUseCase {
	return EquipeDiarioUseCase{service: service}
}

func (uc *EquipeDiarioUseCase) Create(equipe models.EquipeDiario) (models.EquipeDiario, error) {
	id, err := uc.service.Create(equipe)
	if err != nil {
		return models.EquipeDiario{}, err
	}
	equipe.ID.Int64 = id
	equipe.ID.Valid = true
	return equipe, nil
}

func (uc *EquipeDiarioUseCase) GetByDiarioId(diarioId int64) ([]models.EquipeDiario, error) {
	return uc.service.GetByDiarioId(diarioId)
}

func (uc *EquipeDiarioUseCase) GetByObraId(obraId int) ([]models.EquipeDiario, error) {
	if obraId <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}
	return uc.service.GetByObraId(int64(obraId))
}

func (uc *EquipeDiarioUseCase) GetByObraAndData(obraId int, data string) ([]models.EquipeDiario, error) {
	if obraId <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}
	if data == "" {
		return nil, fmt.Errorf("data é obrigatória")
	}
	return uc.service.GetByObraAndData(int64(obraId), data)
}

func (uc *EquipeDiarioUseCase) Update(id int, equipe models.EquipeDiario) (models.EquipeDiario, error) {
	return uc.service.Update(id, equipe)
}

func (uc *EquipeDiarioUseCase) Delete(id int) error {
	err := uc.service.Delete(id)
	if err != nil {
		if err.Error() == "nenhuma equipe encontrada com o ID fornecido" {
			return fmt.Errorf("equipe não encontrada")
		}
		return err
	}
	return nil
}
