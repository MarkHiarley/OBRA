package controller

import (
	"codxis-obras/internal/models"
	"net/http"
	"strconv"

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

func (p *DiarioController) GetDiarioById(ctx *gin.Context) {
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
	validNum := int64(idNumero)
	usuario, err := p.DiarioUseCase.GetDiarioById(validNum)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
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

	ctx.JSON(http.StatusOK, usuario)
}

func (p *DiarioController) GetDiariosByObraId(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id não pode ser nulo",
		})
		return
	}

	idInt, err := strconv.Atoi(id)

	if err != nil {
		message := models.Response{Messagem: "tem que ser número"}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": message.Messagem,
		})
		return
	}

	validId := int64(idInt)

	diarios, err := p.DiarioUseCase.GetDiariosByObraId(validId)
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
