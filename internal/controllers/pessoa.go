package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pessoaController struct {
	pessoaUseCase usecases.PessoaUseCase
}

func NewPessoaController(usecase usecases.PessoaUseCase) pessoaController {
	return pessoaController{

		pessoaUseCase: usecase,
	}
}

func (p *pessoaController) CreatePessoa(ctx *gin.Context) {
	var pessoa models.Pessoa

	if err := ctx.ShouldBindJSON(&pessoa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	createdPessoa, err := p.pessoaUseCase.CreatePessoa(pessoa)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao criar pessoa",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Pessoa criada com sucesso",
		"data":    createdPessoa,
	})
}

func (p *pessoaController) GetPessoas(ctx *gin.Context) {
	pessoas, err := p.pessoaUseCase.GetPessoas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": pessoas,
	})
}

func (p *pessoaController) GetPessoaById(ctx *gin.Context) {
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
	pessoa, err := p.pessoaUseCase.GetPessoaById(validId)
	if err != nil {
		if err.Error() == "Pessoa não encontrada" {
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

	ctx.JSON(http.StatusOK, pessoa)
}
