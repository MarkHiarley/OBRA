package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"database/sql"
	"fmt"
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
		return models.DiarioObra{}, fmt.Errorf("usuário não encontrado")

	}

	return diario, err
}

func (pu *DiarioUseCase) GetDiariosByObraId(id int64) ([]models.DiarioObra, error) {
	return pu.services.GetDiarioByObraId(id)
}

func (pu *DiarioUseCase) PutDiario(id int, updatedDiario models.DiarioObra) (models.DiarioObra, error) {

	updatedDiario, err := pu.services.PutDiarios(id, updatedDiario)
	if err != nil {
		if err == sql.ErrNoRows {

			return models.DiarioObra{}, ErrUserNotFound
		}
		return models.DiarioObra{}, err
	}

	return updatedDiario, nil
}
