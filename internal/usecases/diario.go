package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"database/sql"
	"fmt"
	"strings"
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

	// Validação: relação entre status e aprovado_por_id
	status := ""
	if newDiario.StatusAprovacao.Valid {
		status = strings.ToUpper(newDiario.StatusAprovacao.String)
	}

	if status == "APROVADO" {
		if !newDiario.AprovadoPorID.Valid || newDiario.AprovadoPorID.Int64 == 0 {
			return models.DiarioObra{}, fmt.Errorf("aprovado_por_id é obrigatório quando status_aprovacao = APROVADO")
		}
	}

	if status == "PENDENTE" {
		if newDiario.AprovadoPorID.Valid && newDiario.AprovadoPorID.Int64 != 0 {
			return models.DiarioObra{}, fmt.Errorf("aprovado_por_id deve ser nulo quando status_aprovacao = PENDENTE")
		}
	}

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
		return models.DiarioObra{}, fmt.Errorf("diario não encontrado")

	}

	return diario, err
}

func (pu *DiarioUseCase) GetDiariosByObraId(id int64) ([]models.DiarioObra, error) {
	return pu.services.GetDiarioByObraId(id)
}

func (pu *DiarioUseCase) PutDiario(id int, updatedDiario models.DiarioObra) (models.DiarioObra, error) {
	// Validação: relação entre status e aprovado_por_id
	status := ""
	if updatedDiario.StatusAprovacao.Valid {
		status = strings.ToUpper(updatedDiario.StatusAprovacao.String)
	}

	if status == "APROVADO" {
		if !updatedDiario.AprovadoPorID.Valid || updatedDiario.AprovadoPorID.Int64 == 0 {
			return models.DiarioObra{}, fmt.Errorf("aprovado_por_id é obrigatório quando status_aprovacao = APROVADO")
		}
	}

	if status == "PENDENTE" {
		if updatedDiario.AprovadoPorID.Valid && updatedDiario.AprovadoPorID.Int64 != 0 {
			return models.DiarioObra{}, fmt.Errorf("aprovado_por_id deve ser nulo quando status_aprovacao = PENDENTE")
		}
	}

	updatedDiario, err := pu.services.PutDiarios(id, updatedDiario)
	if err != nil {
		if err == sql.ErrNoRows {

			return models.DiarioObra{}, ErrUserNotFound
		}
		return models.DiarioObra{}, err
	}

	return updatedDiario, nil
}

func (pu *DiarioUseCase) DeleteDiariosById(id int) error {
	err := pu.services.DeleteDiarioById(id)
	if err != nil {
		// ✅ CORRETO: Comparar a MENSAGEM do erro, não o objeto
		switch err.Error() {
		case "nenhum diario encontrado com o ID fornecido":
			return fmt.Errorf("Diario não encontrado")
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
