package controller

import (
	"codxis-obras/internal/auth"
	"codxis-obras/internal/models"
	"codxis-obras/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUseCase usecases.LoginUseCase
}

func NewLoginController(usecase usecases.LoginUseCase) LoginController {
	return LoginController{
		LoginUseCase: usecase,
	}
}

func (p *LoginController) CreateLogin(ctx *gin.Context) {
	var login models.LoginUser

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := p.LoginUseCase.LoginUseCase(login)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (p *LoginController) RefreshToken(ctx *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Refresh token não fornecido",
			"details": err.Error(),
		})
		return
	}

	claims, err := auth.ValidateToken(body.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Refresh token inválido ou expirado",
		})
		return
	}

	newAccessToken, err := auth.GenerateAccessToken(claims.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar novo access token",
		})
		return
	}

	newRefreshToken, err := auth.GenerateRefreshToken(claims.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao gerar novo refresh token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
