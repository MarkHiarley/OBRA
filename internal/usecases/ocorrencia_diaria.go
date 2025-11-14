package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"fmt"

	"gopkg.in/guregu/null.v4"
)

type OcorrenciaDiariaUseCase struct {
	ocorrenciaService services.OcorrenciaDiariaService
	obraService       services.ObraServices
}

func NewOcorrenciaDiariaUseCase(ocorrenciaService services.OcorrenciaDiariaService, obraService services.ObraServices) OcorrenciaDiariaUseCase {
	return OcorrenciaDiariaUseCase{
		ocorrenciaService: ocorrenciaService,
		obraService:       obraService,
	}
}

func (ouc *OcorrenciaDiariaUseCase) CreateOcorrencia(ocorrencia models.OcorrenciaDiaria) (models.OcorrenciaDiaria, error) {
	// Validações
	if !ocorrencia.ObraID.Valid || ocorrencia.ObraID.Int64 <= 0 {
		return models.OcorrenciaDiaria{}, fmt.Errorf("obra_id é obrigatório")
	}

	if !ocorrencia.Descricao.Valid || ocorrencia.Descricao.String == "" {
		return models.OcorrenciaDiaria{}, fmt.Errorf("descrição é obrigatória")
	}

	// Verificar se obra existe
	_, err := ouc.obraService.GetObraById(ocorrencia.ObraID.Int64)
	if err != nil {
		return models.OcorrenciaDiaria{}, fmt.Errorf("obra não encontrada")
	}

	// Definir valores padrão
	if !ocorrencia.Tipo.Valid {
		ocorrencia.Tipo = null.StringFrom(models.TipoOcorrenciaGeral)
	}

	if !ocorrencia.Gravidade.Valid {
		ocorrencia.Gravidade = null.StringFrom(models.GravidadeBaixa)
	}

	if !ocorrencia.Periodo.Valid {
		ocorrencia.Periodo = null.StringFrom("integral")
	}

	if !ocorrencia.StatusResolucao.Valid {
		ocorrencia.StatusResolucao = null.StringFrom(models.StatusResolucaoPendente)
	}

	return ouc.ocorrenciaService.CreateOcorrencia(ocorrencia)
}

func (ouc *OcorrenciaDiariaUseCase) GetOcorrencias() ([]models.OcorrenciaDiariaComRelacionamentos, error) {
	return ouc.ocorrenciaService.GetOcorrencias()
}

func (ouc *OcorrenciaDiariaUseCase) GetOcorrenciasByObraData(obraID int, data string) ([]models.OcorrenciaDiariaComRelacionamentos, error) {
	if obraID <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}

	if data == "" {
		return nil, fmt.Errorf("data é obrigatória")
	}

	return ouc.ocorrenciaService.GetOcorrenciasByObraData(obraID, data)
}

func (ouc *OcorrenciaDiariaUseCase) GetOcorrenciasByGravidade(gravidade string) ([]models.OcorrenciaDiariaComRelacionamentos, error) {
	if gravidade == "" {
		return nil, fmt.Errorf("gravidade é obrigatória")
	}

	return ouc.ocorrenciaService.GetOcorrenciasByGravidade(gravidade)
}

func (ouc *OcorrenciaDiariaUseCase) UpdateOcorrencia(id int, ocorrencia models.OcorrenciaDiaria) error {
	if id <= 0 {
		return fmt.Errorf("ID inválido")
	}

	return ouc.ocorrenciaService.UpdateOcorrencia(id, ocorrencia)
}

func (ouc *OcorrenciaDiariaUseCase) DeleteOcorrencia(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID inválido")
	}

	return ouc.ocorrenciaService.DeleteOcorrencia(id)
}
