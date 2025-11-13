package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"database/sql"
	"fmt"
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
	newObra.ID.Int64 = obraId
	newObra.ID.Valid = true

	return newObra, nil
}

func (pu *ObraUseCase) GetObras() ([]models.Obra, error) {
	return pu.services.GetObras()
}

func (pu *ObraUseCase) GetObraById(id int64) (models.Obra, error) {

	return pu.services.GetObraById(id)
}

func (pu *ObraUseCase) PutObra(id int, updatedObra models.Obra) (models.Obra, error) {

	updatedObra, err := pu.services.PutObra(id, updatedObra)
	if err != nil {
		if err == sql.ErrNoRows {

			return models.Obra{}, ErrUserNotFound
		}
		return models.Obra{}, err
	}

	return updatedObra, nil
}

func (pu *ObraUseCase) DeleteObraById(id int) error {
	err := pu.services.DeleteObraById(id)
	if err != nil {

		switch err.Error() {
		case "nenhuma obra encontrada com o ID fornecido":
			return fmt.Errorf("obra não encontrada")
		case "erro ao executar a query de delete":
			return fmt.Errorf("erro ao executar operação de delete: %w", err)
		case "erro ao obter linhas afetadas":
			return fmt.Errorf("erro ao verificar resultado: %w", err)
		default:
			return err
		}
	}

	return nil
}
