package usecases

import (
	"codxis-obras/internal/auth"
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	services services.LoginServices
}

func NewLoginUsecase(services services.LoginServices) LoginUseCase {
	return LoginUseCase{
		services: services,
	}
}

func (pu *LoginUseCase) LoginUseCase(bodyLogin models.LoginUser) (accessToken string, refreshToken string, error error) {

	email := bodyLogin.Email.String
	senha_hash, err := pu.services.CheckUser(email)
	if err != nil {
		return "", "", err
	}

	comparar := bcrypt.CompareHashAndPassword([]byte(senha_hash), []byte(bodyLogin.Senha.String))
	if comparar != nil {
		return "", "", comparar
	}
	accessToken1, err := auth.GenerateAccessToken(email)
	if err != nil {
		return "", "", err
	}
	refreshToken1, err := auth.GenerateRefreshToken(email)
	if err != nil {
		return "", "", err
	}

	return accessToken1, refreshToken1, nil
}
