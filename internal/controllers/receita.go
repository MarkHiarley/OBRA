package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type receitaController struct {
	receitaUseCase usecases.ReceitaUseCase
}

func NewReceitaController(receitaUseCase usecases.ReceitaUseCase) receitaController {
	return receitaController{
		receitaUseCase: receitaUseCase,
	}
}

func (rc *receitaController) CreateReceita(ctx *gin.Context) {
	var receita models.Receita

	if err := ctx.ShouldBindJSON(&receita); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	receitaCriada, err := rc.receitaUseCase.CreateReceita(receita)
	if err != nil {
		if err.Error() == "obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Receita criada com sucesso",
		"data":    receitaCriada,
	})
}

func (rc *receitaController) GetReceitas(ctx *gin.Context) {
	receitas, err := rc.receitaUseCase.GetReceitas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": receitas,
	})
}

func (rc *receitaController) GetReceitaById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	receita, err := rc.receitaUseCase.GetReceitaById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": receita,
	})
}

func (rc *receitaController) PutReceitaById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var receita models.Receita
	if err := ctx.ShouldBindJSON(&receita); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	err = rc.receitaUseCase.UpdateReceita(id, receita)
	if err != nil {
		if err.Error() == "receita não encontrada" || err.Error() == "obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Receita atualizada com sucesso",
	})
}

func (rc *receitaController) DeleteReceitaById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	err = rc.receitaUseCase.DeleteReceita(id)
	if err != nil {
		if err.Error() == "receita não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Receita deletada com sucesso",
	})
}

func (rc *receitaController) GetReceitasByObra(ctx *gin.Context) {
	obraIdParam := ctx.Param("obra_id")
	obraId, err := strconv.Atoi(obraIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da obra inválido",
		})
		return
	}

	receitas, err := rc.receitaUseCase.GetReceitasByObra(obraId)
	if err != nil {
		if err.Error() == "obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": receitas,
	})
}