package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"fmt"

	"gopkg.in/guregu/null.v4"
)

type DiarioConsolidadoUseCase struct {
	consolidadoService services.DiarioConsolidadoService
	metadadosService   services.DiarioMetadadosService
	obraService        services.ObraServices
}

func NewDiarioConsolidadoUseCase(
	consolidadoService services.DiarioConsolidadoService,
	metadadosService services.DiarioMetadadosService,
	obraService services.ObraServices,
) DiarioConsolidadoUseCase {
	return DiarioConsolidadoUseCase{
		consolidadoService: consolidadoService,
		metadadosService:   metadadosService,
		obraService:        obraService,
	}
}

// GetDiarioConsolidado retorna o diário consolidado
func (dcu *DiarioConsolidadoUseCase) GetDiarioConsolidado() ([]models.DiarioConsolidado, error) {
	return dcu.consolidadoService.GetDiarioConsolidado()
}

// GetDiarioConsolidadoByObra retorna diário consolidado filtrado por obra
func (dcu *DiarioConsolidadoUseCase) GetDiarioConsolidadoByObra(obraID int) ([]models.DiarioConsolidado, error) {
	if obraID <= 0 {
		return nil, fmt.Errorf("ID da obra inválido")
	}

	// Verificar se obra existe
	_, err := dcu.obraService.GetObraById(int64(obraID))
	if err != nil {
		return nil, fmt.Errorf("obra não encontrada")
	}

	return dcu.consolidadoService.GetDiarioConsolidadoByObra(obraID)
}

// GetDiarioConsolidadoByData retorna diário consolidado filtrado por data
func (dcu *DiarioConsolidadoUseCase) GetDiarioConsolidadoByData(data string) ([]models.DiarioConsolidado, error) {
	if data == "" {
		return nil, fmt.Errorf("data é obrigatória")
	}

	return dcu.consolidadoService.GetDiarioConsolidadoByData(data)
}

// CreateOrUpdateMetadados cria ou atualiza metadados do diário
func (dcu *DiarioConsolidadoUseCase) CreateOrUpdateMetadados(metadados models.DiarioMetadados) (models.DiarioMetadados, error) {
	// Validações
	if !metadados.ObraID.Valid || metadados.ObraID.Int64 <= 0 {
		return models.DiarioMetadados{}, fmt.Errorf("obra_id é obrigatório")
	}

	if !metadados.Data.Valid || metadados.Data.String == "" {
		return models.DiarioMetadados{}, fmt.Errorf("data é obrigatória")
	}

	// Verificar se obra existe
	_, err := dcu.obraService.GetObraById(metadados.ObraID.Int64)
	if err != nil {
		return models.DiarioMetadados{}, fmt.Errorf("obra não encontrada")
	}

	// Definir valores padrão
	if !metadados.Periodo.Valid {
		metadados.Periodo = null.StringFrom("integral")
	}

	if !metadados.StatusAprovacao.Valid {
		metadados.StatusAprovacao = null.StringFrom("pendente")
	}

	return dcu.metadadosService.CreateMetadados(metadados)
}
