package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// =============== EQUIPE DIARIO CONTROLLER ===============

type EquipeDiarioController struct {
	useCase usecases.EquipeDiarioUseCase
}

func NewEquipeDiarioController(useCase usecases.EquipeDiarioUseCase) EquipeDiarioController {
	return EquipeDiarioController{useCase: useCase}
}

func (c *EquipeDiarioController) Create(ctx *gin.Context) {
	var equipe models.EquipeDiario
	if err := ctx.ShouldBindJSON(&equipe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	created, err := c.useCase.Create(equipe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao criar equipe", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Equipe criada com sucesso", "data": created})
}

func (c *EquipeDiarioController) GetByDiarioId(ctx *gin.Context) {
	diarioIdStr := ctx.Param("diario_id")
	diarioId, err := strconv.ParseInt(diarioIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	equipes, err := c.useCase.GetByDiarioId(diarioId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": equipes})
}

func (c *EquipeDiarioController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var equipe models.EquipeDiario
	if err := ctx.ShouldBindJSON(&equipe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	updated, err := c.useCase.Update(id, equipe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updated})
}

func (c *EquipeDiarioController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.useCase.Delete(id)
	if err != nil {
		if err.Error() == "equipe não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// =============== EQUIPAMENTO DIARIO CONTROLLER ===============

type EquipamentoDiarioController struct {
	useCase usecases.EquipamentoDiarioUseCase
}

func NewEquipamentoDiarioController(useCase usecases.EquipamentoDiarioUseCase) EquipamentoDiarioController {
	return EquipamentoDiarioController{useCase: useCase}
}

func (c *EquipamentoDiarioController) Create(ctx *gin.Context) {
	var equipamento models.EquipamentoDiario
	if err := ctx.ShouldBindJSON(&equipamento); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	created, err := c.useCase.Create(equipamento)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao criar equipamento", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Equipamento criado com sucesso", "data": created})
}

func (c *EquipamentoDiarioController) GetByDiarioId(ctx *gin.Context) {
	diarioIdStr := ctx.Param("diario_id")
	diarioId, err := strconv.ParseInt(diarioIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	equipamentos, err := c.useCase.GetByDiarioId(diarioId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": equipamentos})
}

func (c *EquipamentoDiarioController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var equipamento models.EquipamentoDiario
	if err := ctx.ShouldBindJSON(&equipamento); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	updated, err := c.useCase.Update(id, equipamento)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updated})
}

func (c *EquipamentoDiarioController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.useCase.Delete(id)
	if err != nil {
		if err.Error() == "equipamento não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// =============== MATERIAL DIARIO CONTROLLER ===============

type MaterialDiarioController struct {
	useCase usecases.MaterialDiarioUseCase
}

func NewMaterialDiarioController(useCase usecases.MaterialDiarioUseCase) MaterialDiarioController {
	return MaterialDiarioController{useCase: useCase}
}

func (c *MaterialDiarioController) Create(ctx *gin.Context) {
	var material models.MaterialDiario
	if err := ctx.ShouldBindJSON(&material); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	created, err := c.useCase.Create(material)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao criar material", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Material criado com sucesso", "data": created})
}

func (c *MaterialDiarioController) GetByDiarioId(ctx *gin.Context) {
	diarioIdStr := ctx.Param("diario_id")
	diarioId, err := strconv.ParseInt(diarioIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	materiais, err := c.useCase.GetByDiarioId(diarioId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": materiais})
}

func (c *MaterialDiarioController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var material models.MaterialDiario
	if err := ctx.ShouldBindJSON(&material); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	updated, err := c.useCase.Update(id, material)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updated})
}

func (c *MaterialDiarioController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.useCase.Delete(id)
	if err != nil {
		if err.Error() == "material não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
