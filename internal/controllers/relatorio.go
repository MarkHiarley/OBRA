package controller

import (
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type relatorioController struct {
	relatorioUseCase usecases.RelatorioUseCase
}

func NewRelatorioController(relatorioUseCase usecases.RelatorioUseCase) relatorioController {
	return relatorioController{
		relatorioUseCase: relatorioUseCase,
	}
}

func (rc *relatorioController) GetRelatorioObra(ctx *gin.Context) {
	obraIdParam := ctx.Param("obra_id")
	obraId, err := strconv.Atoi(obraIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da obra inválido",
		})
		return
	}

	relatorio, err := rc.relatorioUseCase.GetRelatorioObra(obraId)
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
		"data": relatorio,
	})
}

func (rc *relatorioController) GetRelatorioDespesasPorCategoria(ctx *gin.Context) {
	obraIdParam := ctx.Param("obra_id")
	obraId, err := strconv.Atoi(obraIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da obra inválido",
		})
		return
	}

	relatorio, err := rc.relatorioUseCase.GetRelatorioDespesasPorCategoria(obraId)
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
		"data": relatorio,
	})
}

func (rc *relatorioController) GetRelatorioPagamentos(ctx *gin.Context) {
	obraIdParam := ctx.Param("obra_id")
	obraId, err := strconv.Atoi(obraIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da obra inválido",
		})
		return
	}

	// Status é opcional via query parameter
	status := ctx.Query("status")

	relatorio, err := rc.relatorioUseCase.GetRelatorioPagamentos(obraId, status)
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

	ctx.JSON(http.StatusOK, gin.H{
		"data": relatorio,
	})
}

func (rc *relatorioController) GetRelatorioMateriais(ctx *gin.Context) {
	obraIdParam := ctx.Param("obra_id")
	obraId, err := strconv.Atoi(obraIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da obra inválido",
		})
		return
	}

	relatorio, err := rc.relatorioUseCase.GetRelatorioMateriais(obraId)
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
		"data": relatorio,
	})
}

func (rc *relatorioController) GetRelatorioProfissionais(ctx *gin.Context) {
	obraIdParam := ctx.Param("obra_id")
	obraId, err := strconv.Atoi(obraIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID da obra inválido",
		})
		return
	}

	relatorio, err := rc.relatorioUseCase.GetRelatorioProfissionais(obraId)
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
		"data": relatorio,
	})
}