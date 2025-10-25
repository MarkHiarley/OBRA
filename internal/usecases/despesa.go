package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"database/sql"
	"fmt"
	"strings"
)

type DespesaUseCase struct {
	services services.DespesaServices
}

func NewDespesaUsecase(services services.DespesaServices) DespesaUseCase {
	return DespesaUseCase{
		services: services,
	}
}

func (du *DespesaUseCase) CreateDespesa(newDespesa models.Despesa) (models.Despesa, error) {
	// Validações de negócio
	if err := du.validateDespesa(newDespesa); err != nil {
		return models.Despesa{}, err
	}

	// Define status padrão como PENDENTE se não informado
	if !newDespesa.StatusPagamento.Valid || strings.TrimSpace(newDespesa.StatusPagamento.String) == "" {
		newDespesa.StatusPagamento.String = models.StatusPagamentoPendente
		newDespesa.StatusPagamento.Valid = true
	}

	despesaId, err := du.services.CreateDespesa(newDespesa)
	if err != nil {
		// Verifica se é erro de FK (obra ou fornecedor não existe)
		if strings.Contains(err.Error(), "foreign key") {
			if strings.Contains(err.Error(), "obra_id") {
				return models.Despesa{}, fmt.Errorf("obra não encontrada")
			}
			if strings.Contains(err.Error(), "fornecedor_id") {
				return models.Despesa{}, fmt.Errorf("fornecedor não encontrado")
			}
		}
		return models.Despesa{}, err
	}

	newDespesa.ID.Int64 = despesaId
	newDespesa.ID.Valid = true

	return newDespesa, nil
}

func (du *DespesaUseCase) GetDespesas() ([]models.DespesaComRelacionamentos, error) {
	return du.services.GetDespesas()
}

func (du *DespesaUseCase) GetDespesaById(id int64) (models.DespesaComRelacionamentos, error) {
	return du.services.GetDespesaById(id)
}

func (du *DespesaUseCase) GetDespesasByObraId(obraId int64) ([]models.DespesaComRelacionamentos, error) {
	return du.services.GetDespesasByObraId(obraId)
}

func (du *DespesaUseCase) GetDespesasByFornecedorId(fornecedorId int64) ([]models.DespesaComRelacionamentos, error) {
	return du.services.GetDespesasByFornecedorId(fornecedorId)
}

func (du *DespesaUseCase) GetRelatorioPorObra(obraId int64) ([]models.RelatorioDespesas, error) {
	return du.services.GetRelatorioPorObra(obraId)
}

func (du *DespesaUseCase) PutDespesa(id int, updatedDespesa models.Despesa) (models.Despesa, error) {
	// Validações de negócio
	if err := du.validateDespesa(updatedDespesa); err != nil {
		return models.Despesa{}, err
	}

	updatedDespesa, err := du.services.PutDespesa(id, updatedDespesa)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Despesa{}, fmt.Errorf("despesa não encontrada")
		}
		// Verifica se é erro de FK
		if strings.Contains(err.Error(), "foreign key") {
			if strings.Contains(err.Error(), "obra_id") {
				return models.Despesa{}, fmt.Errorf("obra não encontrada")
			}
			if strings.Contains(err.Error(), "fornecedor_id") {
				return models.Despesa{}, fmt.Errorf("fornecedor não encontrado")
			}
		}
		return models.Despesa{}, err
	}

	return updatedDespesa, nil
}

func (du *DespesaUseCase) DeleteDespesaById(id int) error {
	err := du.services.DeleteDespesaById(id)
	if err != nil {
		switch err.Error() {
		case "nenhuma despesa encontrada com o ID fornecido":
			return fmt.Errorf("despesa não encontrada")
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

// validateDespesa valida os dados da despesa
func (du *DespesaUseCase) validateDespesa(despesa models.Despesa) error {
	// Valida obra_id
	if !despesa.ObraID.Valid || despesa.ObraID.Int64 <= 0 {
		return fmt.Errorf("obra_id é obrigatório")
	}

	// Valida data_despesa
	if !despesa.DataDespesa.Valid {
		return fmt.Errorf("data da despesa é obrigatória")
	}

	// Valida descrição
	if !despesa.Descricao.Valid || strings.TrimSpace(despesa.Descricao.String) == "" {
		return fmt.Errorf("descrição é obrigatória")
	}

	// Valida categoria
	if !despesa.Categoria.Valid || strings.TrimSpace(despesa.Categoria.String) == "" {
		return fmt.Errorf("categoria é obrigatória")
	}

	categoria := strings.ToUpper(despesa.Categoria.String)
	categoriesValidas := []string{
		models.CategoriaMaterial,
		models.CategoriaMaoDeObra,
		models.CategoriaCombustivel,
		models.CategoriaAlimentacao,
		models.CategoriaMaterialEletrico,
		models.CategoriaAluguelEquipamento,
		models.CategoriaTransporte,
		models.CategoriaOutros,
	}

	categoriaValida := false
	for _, cat := range categoriesValidas {
		if categoria == cat {
			categoriaValida = true
			break
		}
	}

	if !categoriaValida {
		return fmt.Errorf("categoria inválida. Valores permitidos: MATERIAL, MAO_DE_OBRA, COMBUSTIVEL, ALIMENTACAO, MATERIAL_ELETRICO, ALUGUEL_EQUIPAMENTO, TRANSPORTE, OUTROS")
	}

	// Valida valor
	if !despesa.Valor.Valid || despesa.Valor.Float64 < 0 {
		return fmt.Errorf("valor deve ser maior ou igual a zero")
	}

	// Valida forma de pagamento
	if !despesa.FormaPagamento.Valid || strings.TrimSpace(despesa.FormaPagamento.String) == "" {
		return fmt.Errorf("forma de pagamento é obrigatória")
	}

	formaPagamento := strings.ToUpper(despesa.FormaPagamento.String)
	formasPagamentoValidas := []string{
		models.FormaPagamentoPix,
		models.FormaPagamentoBoleto,
		models.FormaPagamentoCartaoCredito,
		models.FormaPagamentoCartaoDebito,
		models.FormaPagamentoTransferencia,
		models.FormaPagamentoEspecie,
		models.FormaPagamentoCheque,
	}

	formaPagamentoValida := false
	for _, forma := range formasPagamentoValidas {
		if formaPagamento == forma {
			formaPagamentoValida = true
			break
		}
	}

	if !formaPagamentoValida {
		return fmt.Errorf("forma de pagamento inválida. Valores permitidos: PIX, BOLETO, CARTAO_CREDITO, CARTAO_DEBITO, TRANSFERENCIA, ESPECIE, CHEQUE")
	}

	// Valida status de pagamento se informado
	if despesa.StatusPagamento.Valid && strings.TrimSpace(despesa.StatusPagamento.String) != "" {
		statusPagamento := strings.ToUpper(despesa.StatusPagamento.String)
		statusValidos := []string{
			models.StatusPagamentoPendente,
			models.StatusPagamentoPago,
			models.StatusPagamentoCancelado,
		}

		statusValido := false
		for _, status := range statusValidos {
			if statusPagamento == status {
				statusValido = true
				break
			}
		}

		if !statusValido {
			return fmt.Errorf("status de pagamento inválido. Valores permitidos: PENDENTE, PAGO, CANCELADO")
		}
	}

	// Se status é PAGO, data_pagamento deve estar preenchida
	if despesa.StatusPagamento.Valid && despesa.StatusPagamento.String == models.StatusPagamentoPago {
		if !despesa.DataPagamento.Valid {
			return fmt.Errorf("data de pagamento é obrigatória quando status é PAGO")
		}
	}

	return nil
}
