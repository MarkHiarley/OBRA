package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"

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
