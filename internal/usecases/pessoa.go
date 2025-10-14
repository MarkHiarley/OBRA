package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
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
