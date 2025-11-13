package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"database/sql"
	"fmt"
)

type PessoaUseCase struct {
	services services.PessoaServices
}

func NewPessoaUsecase(services services.PessoaServices) PessoaUseCase {
	return PessoaUseCase{
		services: services,
	}
}

func (pu *PessoaUseCase) CreatePessoa(newPessoa models.Pessoa) (models.Pessoa, error) {

	userId, err := pu.services.CreatePessoa(newPessoa)

	if err != nil {
		return models.Pessoa{}, err
	}
	newPessoa.ID.Int64 = userId
	newPessoa.ID.Valid = true

	return newPessoa, nil
}

func (pu *PessoaUseCase) GetPessoas() ([]models.Pessoa, error) {
	return pu.services.GetPessoas()
}

func (pu *PessoaUseCase) GetPessoaById(id int64) (models.Pessoa, error) {

	return pu.services.GetPessoaById(id)
}

func (pu *PessoaUseCase) PutPessoa(id int, updatedPessoa models.Pessoa) (models.Pessoa, error) {

	updatedPessoa, err := pu.services.PutPessoa(id, updatedPessoa)
	if err != nil {
		if err == sql.ErrNoRows {

			return models.Pessoa{}, ErrUserNotFound
		}
		return models.Pessoa{}, err
	}

	return updatedPessoa, nil
}

func (pu *PessoaUseCase) DeletePessoaById(id int) error {
	err := pu.services.DeletePessoaById(id)
	if err != nil {
		// ✅ CORRETO: Comparar a MENSAGEM do erro, não o objeto
		switch err.Error() {
		case "nenhuma pessoa encontrada com o ID fornecido":
			return fmt.Errorf("pessoa não encontrada")
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
