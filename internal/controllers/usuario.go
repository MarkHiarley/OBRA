package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usuarioController struct {
	usuarioUseCase usecases.UsuarioUseCase
}

func NewUsuarioController(usecase usecases.UsuarioUseCase) usuarioController {
	return usuarioController{
		usuarioUseCase: usecase,
	}
}

func (p *usuarioController) CreateUsuario(ctx *gin.Context) {
	var usuario models.Usuario

	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	createdUsuario, err := p.usuarioUseCase.CreateUsuario(usuario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao criar pessoa",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Usuario criado com sucesso",
		"data":    createdUsuario,
	})
}

func (p *usuarioController) GetUsuarios(ctx *gin.Context) {
	users, err := p.usuarioUseCase.GetUsuarios()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
