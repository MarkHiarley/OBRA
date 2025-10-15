package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"log"
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

func (p *usuarioController) PutUsuarioById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be null"})
		return
	}
	var updatedUsuario models.Usuario
	if err := ctx.ShouldBindJSON(&updatedUsuario); err != nil {
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
	if !updatedUsuario.Nome.Valid || updatedUsuario.Nome.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Nome' é obrigatório."})
		return // Stop processing
	}
	if !updatedUsuario.Email.Valid || updatedUsuario.Email.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Email' é obrigatório."})
		return // Stop processing
	}
	if !updatedUsuario.TipoDocumento.Valid || updatedUsuario.TipoDocumento.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'TipoDocumento' é obrigatório."})
		return // Stop processing
	}
	if !updatedUsuario.Documento.Valid || updatedUsuario.Documento.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Documento' é obrigatório."})
		return // Stop processing
	}
	if !updatedUsuario.Telefone.Valid || updatedUsuario.Telefone.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'telefone' é obrigatório."})
		return // Stop processing
	}
	if !updatedUsuario.PerfilAcesso.Valid || updatedUsuario.PerfilAcesso.String == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'Perfil acesso' é obrigatório."})
		return // Stop processing
	}
	if !updatedUsuario.Ativo.Valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'ativo' é obrigatório."})
		return // Para o processamento
	}

	usuario, err := p.usuarioUseCase.PutUsuario(idNumero, updatedUsuario)
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
