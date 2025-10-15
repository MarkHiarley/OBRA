package controller

import (
	"codxis-obras/internal/models"
	"log"
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

func (p *DiarioController) PutDiarioById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id é obrigatório"})
		return
	}
	var updatedDiario models.DiarioObra
	if err := ctx.ShouldBindJSON(&updatedDiario); err != nil {
		// LOG 1: Imprime o erro exato do binding no seu terminal
		log.Printf("!!! ERRO NO BINDING DO JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data: " + err.Error()})
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
	if !updatedDiario.ObraID.Valid || updatedDiario.ObraID.Int64 == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'ObraID' é obrigatório."})
		return // Stop processing
	}
	if !updatedDiario.Data.Valid || updatedDiario.Data.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Data' é obrigatório."})
		return // Stop processing
	}
	if !updatedDiario.Periodo.Valid || updatedDiario.Periodo.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Tipo' é obrigatório."})
		return
	}
	if !updatedDiario.AtividadesRealizadas.Valid || updatedDiario.AtividadesRealizadas.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Atividades Realizadas' é obrigatório."})
		return // Stop processing
	}
	if !updatedDiario.Ocorrencias.Valid || updatedDiario.Ocorrencias.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Ocorrências' é obrigatório."})
		return // Stop processing
	}
	if !updatedDiario.Observacoes.Valid || updatedDiario.Observacoes.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Observações' é obrigatório."})
		return // Stop processing
	}
	if !updatedDiario.ResponsavelID.Valid || updatedDiario.ResponsavelID.Int64 == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Responsável' é obrigatório."})
		return // Stop processing
	}
	if !updatedDiario.AprovadoPorID.Valid || updatedDiario.AprovadoPorID.Int64 == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Aprovado Por' é obrigatório."})
		return // Stop processing
	}
	if !updatedDiario.StatusAprovacao.Valid || updatedDiario.StatusAprovacao.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Status de Aprovação' é obrigatório."})
		return // Stop processing
	}

	diario, err := p.DiarioUseCase.PutDiario(idNumero, updatedDiario)
	if err != nil {
		if err.Error() == "Diario não encontrado" {
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

	ctx.JSON(http.StatusOK, diario)
}
