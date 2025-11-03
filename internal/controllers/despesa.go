package controller

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type despesaController struct {
	despesaUseCase usecases.DespesaUseCase
}

func NewDespesaController(usecase usecases.DespesaUseCase) despesaController {
	return despesaController{
		despesaUseCase: usecase,
	}
}

func (dc *despesaController) CreateDespesa(ctx *gin.Context) {
	// Parse raw JSON to handle custom date formats
	var rawData map[string]interface{}
	if err := ctx.ShouldBindJSON(&rawData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	var despesa models.Despesa

	// Parse standard fields
	if val, ok := rawData["obra_id"]; ok {
		if floatVal, ok := val.(float64); ok {
			despesa.ObraID = null.IntFrom(int64(floatVal))
		}
	}
	if val, ok := rawData["fornecedor_id"]; ok {
		if floatVal, ok := val.(float64); ok {
			despesa.FornecedorID = null.IntFrom(int64(floatVal))
		}
	}
	if val, ok := rawData["descricao"]; ok {
		if strVal, ok := val.(string); ok {
			despesa.Descricao = null.StringFrom(strVal)
		}
	}
	if val, ok := rawData["categoria"]; ok {
		if strVal, ok := val.(string); ok {
			despesa.Categoria = null.StringFrom(strings.ToUpper(strVal))
		}
	}
	if val, ok := rawData["valor"]; ok {
		if floatVal, ok := val.(float64); ok {
			despesa.Valor = null.FloatFrom(floatVal)
		}
	}
	if val, ok := rawData["forma_pagamento"]; ok {
		if strVal, ok := val.(string); ok {
			despesa.FormaPagamento = null.StringFrom(strings.ToUpper(strVal))
		}
	}
	if val, ok := rawData["status_pagamento"]; ok {
		if strVal, ok := val.(string); ok {
			despesa.StatusPagamento = null.StringFrom(strings.ToUpper(strVal))
		}
	}
	if val, ok := rawData["observacao"]; ok {
		if strVal, ok := val.(string); ok {
			despesa.Observacao = null.StringFrom(strVal)
		}
	}

	// Parse dates with flexible format support
	parseDate := func(dateStr string) (time.Time, error) {
		formats := []string{
			"2006-01-02T15:04:05Z07:00", // ISO format
			"2006-01-02T15:04:05Z",      // ISO format (Z)
			"2006-01-02T15:04:05",       // ISO format (no timezone)
			"2006-01-02",                // Date only (Pablo's format)
		}

		for _, format := range formats {
			if t, err := time.Parse(format, dateStr); err == nil {
				return t, nil
			}
		}
		return time.Time{}, json.Unmarshal([]byte(`"`+dateStr+`"`), new(time.Time))
	}

	if val, ok := rawData["data"]; ok {
		if strVal, ok := val.(string); ok && strVal != "" {
			if t, err := parseDate(strVal); err == nil {
				despesa.Data = null.TimeFrom(t)
			}
		}
	}

	if val, ok := rawData["data_vencimento"]; ok {
		if strVal, ok := val.(string); ok && strVal != "" {
			if t, err := parseDate(strVal); err == nil {
				despesa.DataVencimento = null.TimeFrom(t)
			}
		}
	}

	if val, ok := rawData["data_pagamento"]; ok {
		if strVal, ok := val.(string); ok && strVal != "" {
			if t, err := parseDate(strVal); err == nil {
				despesa.DataPagamento = null.TimeFrom(t)
			}
		}
	}

	createdDespesa, err := dc.despesaUseCase.CreateDespesa(despesa)
	if err != nil {
		// Erros de validação de negócio
		if err.Error() == "obra não encontrada" ||
			err.Error() == "fornecedor não encontrado" ||
			err.Error() == "obra_id é obrigatório" ||
			err.Error() == "data da despesa é obrigatória" ||
			err.Error() == "descrição é obrigatória" ||
			err.Error() == "categoria é obrigatória" ||
			err.Error() == "valor deve ser maior ou igual a zero" ||
			err.Error() == "forma de pagamento é obrigatória" ||
			err.Error() == "data de pagamento é obrigatória quando status é PAGO" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Validação de categoria/forma pagamento/status
		if err.Error() == "categoria inválida. Valores permitidos: MATERIAL, MAO_DE_OBRA, COMBUSTIVEL, ALIMENTACAO, MATERIAL_ELETRICO, ALUGUEL_EQUIPAMENTO, TRANSPORTE, OUTROS" ||
			err.Error() == "forma de pagamento inválida. Valores permitidos: PIX, BOLETO, CARTAO_CREDITO, CARTAO_DEBITO, TRANSFERENCIA, ESPECIE, CHEQUE" ||
			err.Error() == "status de pagamento inválido. Valores permitidos: PENDENTE, PAGO, CANCELADO" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar despesa",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Despesa criada com sucesso",
		"data":    createdDespesa,
	})
}

func (dc *despesaController) GetDespesas(ctx *gin.Context) {
	// Verifica se há filtro por obra ou fornecedor
	obraIdStr := ctx.Query("obra_id")
	fornecedorIdStr := ctx.Query("fornecedor_id")

	var despesas []models.DespesaComRelacionamentos
	var err error

	if obraIdStr != "" {
		obraId, errConv := strconv.ParseInt(obraIdStr, 10, 64)
		if errConv != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "obra_id deve ser um número válido",
			})
			return
		}
		despesas, err = dc.despesaUseCase.GetDespesasByObraId(obraId)
	} else if fornecedorIdStr != "" {
		fornecedorId, errConv := strconv.ParseInt(fornecedorIdStr, 10, 64)
		if errConv != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "fornecedor_id deve ser um número válido",
			})
			return
		}
		despesas, err = dc.despesaUseCase.GetDespesasByFornecedorId(fornecedorId)
	} else {
		despesas, err = dc.despesaUseCase.GetDespesas()
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": despesas,
	})
}

func (dc *despesaController) GetDespesaById(ctx *gin.Context) {
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
	despesa, err := dc.despesaUseCase.GetDespesaById(validId)
	if err != nil {
		if err.Error() == "despesa não encontrada" {
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

	ctx.JSON(http.StatusOK, despesa)
}

func (dc *despesaController) GetRelatorioPorObra(ctx *gin.Context) {
	obraIdStr := ctx.Param("obra_id")

	if obraIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "obra_id é obrigatório",
		})
		return
	}

	obraId, err := strconv.ParseInt(obraIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "obra_id deve ser um número válido",
		})
		return
	}

	relatorio, err := dc.despesaUseCase.GetRelatorioPorObra(obraId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"obra_id": obraId,
		"data":    relatorio,
	})
}

func (dc *despesaController) PutDespesaById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID é obrigatório"})
		return
	}

	var updatedDespesa models.Despesa
	if err := ctx.ShouldBindJSON(&updatedDespesa); err != nil {
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

	despesa, err := dc.despesaUseCase.PutDespesa(idNumero, updatedDespesa)
	if err != nil {
		// Erros de validação de negócio
		if err.Error() == "despesa não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err.Error() == "obra não encontrada" ||
			err.Error() == "fornecedor não encontrado" ||
			err.Error() == "obra_id é obrigatório" ||
			err.Error() == "data da despesa é obrigatória" ||
			err.Error() == "descrição é obrigatória" ||
			err.Error() == "categoria é obrigatória" ||
			err.Error() == "valor deve ser maior ou igual a zero" ||
			err.Error() == "forma de pagamento é obrigatória" ||
			err.Error() == "data de pagamento é obrigatória quando status é PAGO" {
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
		"message": "Despesa atualizada com sucesso",
		"data":    despesa,
	})
}

func (dc *despesaController) DeleteDespesaById(ctx *gin.Context) {
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

	err = dc.despesaUseCase.DeleteDespesaById(idNumero)
	if err != nil {
		if err.Error() == "despesa não encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Despesa não encontrada",
			})
			return
		}

		log.Printf("Erro ao deletar despesa ID %d: %v", idNumero, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao deletar despesa",
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
