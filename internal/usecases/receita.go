package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"fmt"
	"time"

	"gopkg.in/guregu/null.v4"
)

type ReceitaUseCase struct {
	receitaService services.ReceitaService
	obraService    services.ObraServices
}

func NewReceitaUseCase(receitaService services.ReceitaService, obraService services.ObraServices) ReceitaUseCase {
	return ReceitaUseCase{
		receitaService: receitaService,
		obraService:    obraService,
	}
}

func (ruc *ReceitaUseCase) CreateReceita(receita models.Receita) (models.Receita, error) {
	// Validações básicas
	if !receita.ObraID.Valid || receita.ObraID.Int64 <= 0 {
		return models.Receita{}, fmt.Errorf("obra_id é obrigatório")
	}

	if !receita.Descricao.Valid || receita.Descricao.String == "" {
		return models.Receita{}, fmt.Errorf("descrição é obrigatória")
	}

	if !receita.Valor.Valid || receita.Valor.Float64 <= 0 {
		return models.Receita{}, fmt.Errorf("valor deve ser maior que zero")
	}

	// Verificar se a obra existe
	_, err := ruc.obraService.GetObraById(receita.ObraID.Int64)
	if err != nil {
		return models.Receita{}, fmt.Errorf("obra não encontrada")
	}

	// Definir valores padrão
	if !receita.FonteReceita.Valid {
		receita.FonteReceita = null.StringFrom(models.FonteReceitaOutros)
	}

	if !receita.Data.Valid {
		receita.Data = null.TimeFrom(time.Now())
	}

	// Definir status padrão
	if !receita.Status.Valid {
		receita.Status = null.StringFrom("a_receber")
	}

	// Criar a receita
	receitaCriada, err := ruc.receitaService.CreateReceita(receita)
	if err != nil {
		return models.Receita{}, fmt.Errorf("erro ao criar receita: %v", err)
	}

	return receitaCriada, nil
}

func (ruc *ReceitaUseCase) GetReceitas() ([]models.ReceitaComRelacionamentos, error) {
	receitas, err := ruc.receitaService.GetReceitas()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar receitas: %v", err)
	}

	return receitas, nil
}

func (ruc *ReceitaUseCase) GetReceitaById(id int) (models.ReceitaComRelacionamentos, error) {
	if id <= 0 {
		return models.ReceitaComRelacionamentos{}, fmt.Errorf("ID inválido")
	}

	receita, err := ruc.receitaService.GetReceitaById(id)
	if err != nil {
		return models.ReceitaComRelacionamentos{}, fmt.Errorf("receita não encontrada")
	}

	return receita, nil
}

func (ruc *ReceitaUseCase) UpdateReceita(id int, receita models.Receita) error {
	if id <= 0 {
		return fmt.Errorf("ID inválido")
	}

	// Verificar se a receita existe
	_, err := ruc.receitaService.GetReceitaById(id)
	if err != nil {
		return fmt.Errorf("receita não encontrada")
	}

	// Validações
	if receita.ObraID.Valid {
		_, err := ruc.obraService.GetObraById(receita.ObraID.Int64)
		if err != nil {
			return fmt.Errorf("obra não encontrada")
		}
	}

	if receita.Valor.Valid && receita.Valor.Float64 <= 0 {
		return fmt.Errorf("valor deve ser maior que zero")
	}

	err = ruc.receitaService.UpdateReceita(id, receita)
	if err != nil {
		return fmt.Errorf("erro ao atualizar receita: %v", err)
	}

	return nil
}

func (ruc *ReceitaUseCase) DeleteReceita(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID inválido")
	}

	// Verificar se a receita existe
	_, err := ruc.receitaService.GetReceitaById(id)
	if err != nil {
		return fmt.Errorf("receita não encontrada")
	}

	err = ruc.receitaService.DeleteReceita(id)
	if err != nil {
		return fmt.Errorf("erro ao deletar receita: %v", err)
	}

	return nil
}

func (ruc *ReceitaUseCase) GetReceitasByObra(obraId int) ([]models.ReceitaComRelacionamentos, error) {
	if obraId <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}

	// Verificar se a obra existe
	_, err := ruc.obraService.GetObraById(int64(obraId))
	if err != nil {
		return nil, fmt.Errorf("obra não encontrada")
	}

	receitas, err := ruc.receitaService.GetReceitasByObra(obraId)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar receitas da obra: %v", err)
	}

	return receitas, nil
}
