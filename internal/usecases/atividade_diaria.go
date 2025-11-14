package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"fmt"

	"gopkg.in/guregu/null.v4"
)

type AtividadeDiariaUseCase struct {
	atividadeService services.AtividadeDiariaService
	obraService      services.ObraServices
}

func NewAtividadeDiariaUseCase(atividadeService services.AtividadeDiariaService, obraService services.ObraServices) AtividadeDiariaUseCase {
	return AtividadeDiariaUseCase{
		atividadeService: atividadeService,
		obraService:      obraService,
	}
}

func (auc *AtividadeDiariaUseCase) CreateAtividade(atividade models.AtividadeDiaria) (models.AtividadeDiaria, error) {
	// Validações
	if !atividade.ObraID.Valid || atividade.ObraID.Int64 <= 0 {
		return models.AtividadeDiaria{}, fmt.Errorf("obra_id é obrigatório")
	}

	if !atividade.Descricao.Valid || atividade.Descricao.String == "" {
		return models.AtividadeDiaria{}, fmt.Errorf("descrição é obrigatória")
	}

	// Verificar se obra existe
	_, err := auc.obraService.GetObraById(atividade.ObraID.Int64)
	if err != nil {
		return models.AtividadeDiaria{}, fmt.Errorf("obra não encontrada")
	}

	// Definir valores padrão
	if !atividade.Status.Valid {
		atividade.Status = null.StringFrom(models.StatusAtividadeEmAndamento)
	}

	if !atividade.Periodo.Valid {
		atividade.Periodo = null.StringFrom("integral")
	}

	if !atividade.PercentualConclusao.Valid {
		atividade.PercentualConclusao = null.IntFrom(0)
	}

	return auc.atividadeService.CreateAtividade(atividade)
}

func (auc *AtividadeDiariaUseCase) GetAtividades() ([]models.AtividadeDiariaComRelacionamentos, error) {
	return auc.atividadeService.GetAtividades()
}

func (auc *AtividadeDiariaUseCase) GetAtividadesByObraData(obraID int, data string) ([]models.AtividadeDiariaComRelacionamentos, error) {
	if obraID <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}

	if data == "" {
		return nil, fmt.Errorf("data é obrigatória")
	}

	return auc.atividadeService.GetAtividadesByObraData(obraID, data)
}

func (auc *AtividadeDiariaUseCase) UpdateAtividade(id int, atividade models.AtividadeDiaria) error {
	if id <= 0 {
		return fmt.Errorf("ID inválido")
	}

	if atividade.PercentualConclusao.Valid && (atividade.PercentualConclusao.Int64 < 0 || atividade.PercentualConclusao.Int64 > 100) {
		return fmt.Errorf("percentual de conclusão deve estar entre 0 e 100")
	}

	return auc.atividadeService.UpdateAtividade(id, atividade)
}

func (auc *AtividadeDiariaUseCase) DeleteAtividade(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID inválido")
	}

	return auc.atividadeService.DeleteAtividade(id)
}
