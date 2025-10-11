package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ObraController struct {
	ObraUseCase usecases.ObraUseCase
}

func NewObraController(usecase usecases.ObraUseCase) ObraController {
	return ObraController{

		ObraUseCase: usecase,
	}
}

func (p *ObraController) CreateObra(ctx *gin.Context) {
	var obra models.Obra

	if err := ctx.ShouldBindJSON(&obra); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	createdObra, err := p.ObraUseCase.CreateObra(obra)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao criar Obra",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Obra criada com sucesso",
		"data":    createdObra,
	})
}

func (p *ObraController) GetObras(ctx *gin.Context) {
	obras, err := p.ObraUseCase.GetObras()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": obras,
	})
}
