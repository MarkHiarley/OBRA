package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

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

func (p *ObraController) GetObraById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id não pode ser nulo",
		})
		return
	}

	idNumero, err := strconv.Atoi(id)
	if err != nil {
		message := models.Response{Messagem: "tem que ser número"}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": message.Messagem,
		})
		return
	}
	validId := int64(idNumero)
	obra, err := p.ObraUseCase.GetObraById(validId)
	if err != nil {
		if err.Error() == "Obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, obra)
}
