package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type fornecedorController struct {
	fornecedorUseCase usecases.FornecedorUseCase
}

func NewFornecedorController(usecase usecases.FornecedorUseCase) fornecedorController {
	return fornecedorController{
		fornecedorUseCase: usecase,
	}
}

func (fc *fornecedorController) CreateFornecedor(ctx *gin.Context) {
	var fornecedor models.Fornecedor

	if err := ctx.ShouldBindJSON(&fornecedor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	createdFornecedor, err := fc.fornecedorUseCase.CreateFornecedor(fornecedor)
	if err != nil {
		// Validação de negócio
		if err.Error() == "já existe um fornecedor cadastrado com este documento" ||
			err.Error() == "nome é obrigatório" ||
			err.Error() == "tipo de documento é obrigatório" ||
			err.Error() == "tipo de documento deve ser CPF ou CNPJ" ||
			err.Error() == "documento é obrigatório" ||
			err.Error() == "CPF deve ter 11 dígitos" ||
			err.Error() == "CNPJ deve ter 14 dígitos" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar fornecedor",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Fornecedor criado com sucesso",
		"data":    createdFornecedor,
	})
}

func (fc *fornecedorController) GetFornecedores(ctx *gin.Context) {
	// Verifica se quer apenas ativos
	apenasAtivos := ctx.Query("ativos")

	var fornecedores []models.Fornecedor
	var err error

	if apenasAtivos == "true" {
		fornecedores, err = fc.fornecedorUseCase.GetFornecedoresAtivos()
	} else {
		fornecedores, err = fc.fornecedorUseCase.GetFornecedores()
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": fornecedores,
	})
}

func (fc *fornecedorController) GetFornecedorById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não pode ser nulo",
		})
		return
	}

	idNumero, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID deve ser um número válido",
		})
		return
	}

	validId := int64(idNumero)
	fornecedor, err := fc.fornecedorUseCase.GetFornecedorById(validId)
	if err != nil {
		if err.Error() == "fornecedor não encontrado" {
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

	ctx.JSON(http.StatusOK, fornecedor)
}

func (fc *fornecedorController) PutFornecedorById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID é obrigatório"})
		return
	}

	var updatedFornecedor models.Fornecedor
	if err := ctx.ShouldBindJSON(&updatedFornecedor); err != nil {
		log.Printf("Erro no binding do JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados JSON inválidos: " + err.Error()})
		return
	}

	idNumero, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID deve ser um número válido",
		})
		return
	}

	fornecedor, err := fc.fornecedorUseCase.PutFornecedor(idNumero, updatedFornecedor)
	if err != nil {
		// Erros de validação de negócio
		if err.Error() == "fornecedor não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err.Error() == "já existe um fornecedor cadastrado com este documento" ||
			err.Error() == "nome é obrigatório" ||
			err.Error() == "tipo de documento é obrigatório" ||
			err.Error() == "tipo de documento deve ser CPF ou CNPJ" ||
			err.Error() == "documento é obrigatório" ||
			err.Error() == "CPF deve ter 11 dígitos" ||
			err.Error() == "CNPJ deve ter 14 dígitos" {
			ctx.JSON(http.StatusBadRequest, gin.H{
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
		"message": "Fornecedor atualizado com sucesso",
		"data":    fornecedor,
	})
}

func (fc *fornecedorController) DeleteFornecedorById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID é obrigatório",
		})
		return
	}

	idNumero, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID deve ser um número válido",
		})
		return
	}

	err = fc.fornecedorUseCase.DeleteFornecedorById(idNumero)
	if err != nil {
		if err.Error() == "fornecedor não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Fornecedor não encontrado",
			})
			return
		}

		log.Printf("Erro ao deletar fornecedor ID %d: %v", idNumero, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao deletar fornecedor",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
