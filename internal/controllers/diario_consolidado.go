package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DiarioConsolidadoController struct {
	diarioUseCase usecases.DiarioConsolidadoUseCase
}

func NewDiarioConsolidadoController(diarioUseCase usecases.DiarioConsolidadoUseCase) DiarioConsolidadoController {
	return DiarioConsolidadoController{
		diarioUseCase: diarioUseCase,
	}
}

// GetDiarioConsolidado retorna o diário consolidado de todas as obras
func (dcc *DiarioConsolidadoController) GetDiarioConsolidado(ctx *gin.Context) {
	diarios, err := dcc.diarioUseCase.GetDiarioConsolidado()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": diarios})
}

// GetDiarioConsolidadoByObra retorna o diário consolidado de uma obra específica
func (dcc *DiarioConsolidadoController) GetDiarioConsolidadoByObra(ctx *gin.Context) {
	obraIDParam := ctx.Param("obra_id")
	obraID, err := strconv.Atoi(obraIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID da obra inválido"})
		return
	}

	diarios, err := dcc.diarioUseCase.GetDiarioConsolidadoByObra(obraID)
	if err != nil {
		if err.Error() == "obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": diarios})
}

// GetDiarioConsolidadoByData retorna o diário consolidado de uma data específica
func (dcc *DiarioConsolidadoController) GetDiarioConsolidadoByData(ctx *gin.Context) {
	data := ctx.Param("data")
	if data == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data é obrigatória"})
		return
	}

	diarios, err := dcc.diarioUseCase.GetDiarioConsolidadoByData(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": diarios})
}

// CreateOrUpdateMetadados cria ou atualiza metadados do diário (foto, observações, aprovação)
func (dcc *DiarioConsolidadoController) CreateOrUpdateMetadados(ctx *gin.Context) {
	var metadados models.DiarioMetadados

	if err := ctx.ShouldBindJSON(&metadados); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	created, err := dcc.diarioUseCase.CreateOrUpdateMetadados(metadados)
	if err != nil {
		if err.Error() == "obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Metadados salvos com sucesso",
		"data":    created,
	})
}
