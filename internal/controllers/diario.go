package controller

import (
	"codxis-obras/internal/models"
	"net/http"

	"codxis-obras/internal/usecases"

	"github.com/gin-gonic/gin"
)

type DiarioController struct {
	DiarioUseCase usecases.DiarioUseCase
}

func NewDiarioController(usecase usecases.DiarioUseCase) DiarioController {
	return DiarioController{

		DiarioUseCase: usecase,
	}
}

func (p *DiarioController) CreateDiario(ctx *gin.Context) {
	var diario models.DiarioObra

	if err := ctx.ShouldBindJSON(&diario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	createdDiario, err := p.DiarioUseCase.CreateDiario(diario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao criar Diario",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Diario criada com sucesso",
		"data":    createdDiario,
	})
}

func (p *DiarioController) GetDiarios(ctx *gin.Context) {
	diarios, err := p.DiarioUseCase.GetDiarios()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": diarios,
	})
}
