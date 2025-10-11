package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"
	"strconv"

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
func (p *usuarioController) GetUsuarioById(ctx *gin.Context) {
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

	usuario, err := p.usuarioUseCase.GetUsuariosById(idNumero)
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
