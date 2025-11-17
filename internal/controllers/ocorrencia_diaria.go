package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OcorrenciaDiariaController struct {
	ocorrenciaUseCase usecases.OcorrenciaDiariaUseCase
}

func NewOcorrenciaDiariaController(ocorrenciaUseCase usecases.OcorrenciaDiariaUseCase) OcorrenciaDiariaController {
	return OcorrenciaDiariaController{
		ocorrenciaUseCase: ocorrenciaUseCase,
	}
}

// CreateOcorrencia cria uma nova ocorrência diária
func (odc *OcorrenciaDiariaController) CreateOcorrencia(ctx *gin.Context) {
	var ocorrencia models.OcorrenciaDiaria

	if err := ctx.ShouldBindJSON(&ocorrencia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	created, err := odc.ocorrenciaUseCase.CreateOcorrencia(ocorrencia)
	if err != nil {
		if err.Error() == "obra não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Ocorrência criada com sucesso",
		"data":    created,
	})
}

// GetOcorrencias retorna todas as ocorrências
func (odc *OcorrenciaDiariaController) GetOcorrencias(ctx *gin.Context) {
	ocorrencias, err := odc.ocorrenciaUseCase.GetOcorrencias()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ocorrencias})
}

// GetOcorrenciasByObra retorna todas as ocorrências de uma obra (todas as datas)
func (odc *OcorrenciaDiariaController) GetOcorrenciasByObra(ctx *gin.Context) {
	obraIDParam := ctx.Param("obra_id")
	obraID, err := strconv.Atoi(obraIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID da obra inválido"})
		return
	}

	ocorrencias, err := odc.ocorrenciaUseCase.GetOcorrenciasByObra(obraID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ocorrencias})
}

// GetOcorrenciasByObraData retorna ocorrências filtradas por obra e data
func (odc *OcorrenciaDiariaController) GetOcorrenciasByObraData(ctx *gin.Context) {
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

	ocorrencias, err := odc.ocorrenciaUseCase.GetOcorrenciasByObraData(obraID, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ocorrencias})
}

// GetOcorrenciasByGravidade retorna ocorrências filtradas por gravidade
func (odc *OcorrenciaDiariaController) GetOcorrenciasByGravidade(ctx *gin.Context) {
	gravidade := ctx.Param("gravidade")
	if gravidade == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Gravidade é obrigatória"})
		return
	}

	ocorrencias, err := odc.ocorrenciaUseCase.GetOcorrenciasByGravidade(gravidade)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ocorrencias})
}

// UpdateOcorrencia atualiza uma ocorrência existente
func (odc *OcorrenciaDiariaController) UpdateOcorrencia(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var ocorrencia models.OcorrenciaDiaria
	if err := ctx.ShouldBindJSON(&ocorrencia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	err = odc.ocorrenciaUseCase.UpdateOcorrencia(id, ocorrencia)
	if err != nil {
		if err.Error() == "ocorrência não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ocorrência atualizada com sucesso"})
}

// DeleteOcorrencia remove uma ocorrência
func (odc *OcorrenciaDiariaController) DeleteOcorrencia(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = odc.ocorrenciaUseCase.DeleteOcorrencia(id)
	if err != nil {
		if err.Error() == "ocorrência não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ocorrência deletada com sucesso"})
}
