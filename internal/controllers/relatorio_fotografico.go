package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RelatorioFotograficoController struct {
	useCase usecases.RelatorioFotograficoUseCase
}

func NewRelatorioFotograficoController(useCase usecases.RelatorioFotograficoUseCase) RelatorioFotograficoController {
	return RelatorioFotograficoController{
		useCase: useCase,
	}
}

// GetRelatorioFotografico retorna o relatório fotográfico da obra
// GET /api/relatorios/fotografico/:obra_id
func (c *RelatorioFotograficoController) GetRelatorioFotografico(ctx *gin.Context) {
	obraIDParam := ctx.Param("obra_id")
	obraID, err := strconv.ParseInt(obraIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da obra inválido",
		})
		return
	}

	relatorio, err := c.useCase.GetRelatorioFotografico(obraID)
	if err != nil {
		if err.Error() == "obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Obra não encontrada",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar relatório fotográfico: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": relatorio,
	})
}

type DiarioSemanalController struct {
	useCase usecases.DiarioSemanalUseCase
}

func NewDiarioSemanalController(useCase usecases.DiarioSemanalUseCase) DiarioSemanalController {
	return DiarioSemanalController{
		useCase: useCase,
	}
}

// GetDiarioSemanal retorna o diário de obras agrupado por semana
// POST /api/diarios/semanal
func (c *DiarioSemanalController) GetDiarioSemanal(ctx *gin.Context) {
	var request models.DiarioSemanalRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	diario, err := c.useCase.GetDiarioSemanal(request)
	if err != nil {
		if err.Error() == "obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Obra não encontrada",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": diario,
	})
}
