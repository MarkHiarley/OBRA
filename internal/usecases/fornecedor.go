package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"database/sql"
	"fmt"
	"strings"
)

type FornecedorUseCase struct {
	services services.FornecedorServices
}

func NewFornecedorUsecase(services services.FornecedorServices) FornecedorUseCase {
	return FornecedorUseCase{
		services: services,
	}
}

func (fu *FornecedorUseCase) CreateFornecedor(newFornecedor models.Fornecedor) (models.Fornecedor, error) {
	// Validações de negócio
	if err := fu.validateFornecedor(newFornecedor); err != nil {
		return models.Fornecedor{}, err
	}

	fornecedorId, err := fu.services.CreateFornecedor(newFornecedor)
	if err != nil {
		// Verifica se é erro de documento duplicado
		if strings.Contains(err.Error(), "duplicate key") && strings.Contains(err.Error(), "documento") {
			return models.Fornecedor{}, fmt.Errorf("já existe um fornecedor cadastrado com este documento")
		}
		return models.Fornecedor{}, err
	}

	newFornecedor.ID.Int64 = fornecedorId
	newFornecedor.ID.Valid = true

	return newFornecedor, nil
}

func (fu *FornecedorUseCase) GetFornecedores() ([]models.Fornecedor, error) {
	return fu.services.GetFornecedores()
}

func (fu *FornecedorUseCase) GetFornecedoresAtivos() ([]models.Fornecedor, error) {
	return fu.services.GetFornecedoresAtivos()
}

func (fu *FornecedorUseCase) GetFornecedorById(id int64) (models.Fornecedor, error) {
	return fu.services.GetFornecedorById(id)
}

func (fu *FornecedorUseCase) PutFornecedor(id int, updatedFornecedor models.Fornecedor) (models.Fornecedor, error) {
	// Validações de negócio
	if err := fu.validateFornecedor(updatedFornecedor); err != nil {
		return models.Fornecedor{}, err
	}

	updatedFornecedor, err := fu.services.PutFornecedor(id, updatedFornecedor)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Fornecedor{}, fmt.Errorf("fornecedor não encontrado")
		}
		// Verifica se é erro de documento duplicado
		if strings.Contains(err.Error(), "duplicate key") && strings.Contains(err.Error(), "documento") {
			return models.Fornecedor{}, fmt.Errorf("já existe um fornecedor cadastrado com este documento")
		}
		return models.Fornecedor{}, err
	}

	return updatedFornecedor, nil
}

func (fu *FornecedorUseCase) DeleteFornecedorById(id int) error {
	err := fu.services.DeleteFornecedorById(id)
	if err != nil {
		switch err.Error() {
		case "nenhum fornecedor encontrado com o ID fornecido":
			return fmt.Errorf("fornecedor não encontrado")
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

// validateFornecedor valida os dados do fornecedor
func (fu *FornecedorUseCase) validateFornecedor(fornecedor models.Fornecedor) error {
	// Valida nome
	if !fornecedor.Nome.Valid || strings.TrimSpace(fornecedor.Nome.String) == "" {
		return fmt.Errorf("nome é obrigatório")
	}

	// Valida tipo de documento
	if !fornecedor.TipoDocumento.Valid || strings.TrimSpace(fornecedor.TipoDocumento.String) == "" {
		return fmt.Errorf("tipo de documento é obrigatório")
	}

	tipoDoc := strings.ToUpper(fornecedor.TipoDocumento.String)
	if tipoDoc != "CPF" && tipoDoc != "CNPJ" {
		return fmt.Errorf("tipo de documento deve ser CPF ou CNPJ")
	}

	// Valida documento
	if !fornecedor.Documento.Valid || strings.TrimSpace(fornecedor.Documento.String) == "" {
		return fmt.Errorf("documento é obrigatório")
	}

	// Validação básica de CPF (11 dígitos) ou CNPJ (14 dígitos)
	documentoLimpo := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fornecedor.Documento.String, ".", ""), "-", ""), "/", "")
	
	if tipoDoc == "CPF" && len(documentoLimpo) != 11 {
		return fmt.Errorf("CPF deve ter 11 dígitos")
	}
	
	if tipoDoc == "CNPJ" && len(documentoLimpo) != 14 {
		return fmt.Errorf("CNPJ deve ter 14 dígitos")
	}

	return nil
}
