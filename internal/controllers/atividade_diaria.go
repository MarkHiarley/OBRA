package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AtividadeDiariaController struct {
	atividadeUseCase usecases.AtividadeDiariaUseCase
}

func NewAtividadeDiariaController(atividadeUseCase usecases.AtividadeDiariaUseCase) AtividadeDiariaController {
	return AtividadeDiariaController{
		atividadeUseCase: atividadeUseCase,
	}
}

// CreateAtividade cria uma nova atividade diária
func (adc *AtividadeDiariaController) CreateAtividade(ctx *gin.Context) {
	var atividade models.AtividadeDiaria

	if err := ctx.ShouldBindJSON(&atividade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	created, err := adc.atividadeUseCase.CreateAtividade(atividade)
	if err != nil {
		if err.Error() == "obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Atividade criada com sucesso",
		"data":    created,
	})
}

// GetAtividades retorna todas as atividades
func (adc *AtividadeDiariaController) GetAtividades(ctx *gin.Context) {
	atividades, err := adc.atividadeUseCase.GetAtividades()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": atividades})
}

// GetAtividadesByObra retorna todas as atividades de uma obra (todas as datas)
func (adc *AtividadeDiariaController) GetAtividadesByObra(ctx *gin.Context) {
	obraIDParam := ctx.Param("obra_id")
	obraID, err := strconv.Atoi(obraIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID da obra inválido"})
		return
	}

	atividades, err := adc.atividadeUseCase.GetAtividadesByObra(obraID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": atividades})
}

// GetAtividadesByObraData retorna atividades filtradas por obra e data
func (adc *AtividadeDiariaController) GetAtividadesByObraData(ctx *gin.Context) {
	obraIDParam := ctx.Param("obra_id")
	obraID, err := strconv.Atoi(obraIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID da obra inválido"})
		return
	}

	data := ctx.Param("data")
	if data == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data é obrigatória"})
		return
	}

	atividades, err := adc.atividadeUseCase.GetAtividadesByObraData(obraID, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": atividades})
}

// UpdateAtividade atualiza uma atividade existente
func (adc *AtividadeDiariaController) UpdateAtividade(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var atividade models.AtividadeDiaria
	if err := ctx.ShouldBindJSON(&atividade); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	err = adc.atividadeUseCase.UpdateAtividade(id, atividade)
	if err != nil {
		if err.Error() == "atividade não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Atividade atualizada com sucesso"})
}

// DeleteAtividade remove uma atividade
func (adc *AtividadeDiariaController) DeleteAtividade(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = adc.atividadeUseCase.DeleteAtividade(id)
	if err != nil {
		if err.Error() == "atividade não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Atividade deletada com sucesso"})
}
