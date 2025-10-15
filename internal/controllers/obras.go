package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"log"
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

func (p *ObraController) PutObraById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be null"})
		return
	}
	var updatedObra models.Obra
	if err := ctx.ShouldBindJSON(&updatedObra); err != nil {
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
	if !updatedObra.Nome.Valid || updatedObra.Nome.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Nome' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.ContratoNumero.Valid || updatedObra.ContratoNumero.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Email' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.ContratanteID.Valid || updatedObra.ContratanteID.Int64 == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'TipoDocumento' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.ResponsavelID.Valid || updatedObra.ResponsavelID.Int64 == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Documento' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.DataInicio.Valid || updatedObra.DataInicio.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'telefone' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.PrazoDias.Valid || updatedObra.PrazoDias.Int64 == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.DataFimPrevista.Valid || updatedObra.DataFimPrevista.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.Orcamento.Valid || updatedObra.Orcamento.Float64 == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.Status.Valid || updatedObra.Status.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.EnderecoRua.Valid || updatedObra.EnderecoRua.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.EnderecoNumero.Valid || updatedObra.EnderecoNumero.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.EnderecoBairro.Valid || updatedObra.EnderecoBairro.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.EnderecoCidade.Valid || updatedObra.EnderecoCidade.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.EnderecoEstado.Valid || updatedObra.EnderecoEstado.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.EnderecoCep.Valid || updatedObra.EnderecoCep.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.Observacoes.Valid || updatedObra.Observacoes.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedObra.Ativo.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'ativo' é obrigatório."})
		return // Para o processamento
	}

	obra, err := p.ObraUseCase.PutObra(idNumero, updatedObra)
	if err != nil {
		if err.Error() == "obra não encontrado" {
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
