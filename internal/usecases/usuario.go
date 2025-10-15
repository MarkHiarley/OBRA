package usecases

import (
	"codxis-obras/internal/models"
	"codxis-obras/internal/services"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UsuarioUseCase struct {
	services services.UsuarioServices
}

var ErrUserNotFound = errors.New("usuário não encontrado")

func NewUsuarioUsecase(services services.UsuarioServices) UsuarioUseCase {
	return UsuarioUseCase{
		services: services,
	}
}

func (pu *UsuarioUseCase) CreateUsuario(newUser models.Usuario) (models.Usuario, error) {

	senhaHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Senha.String), bcrypt.DefaultCost)
	if err != nil {
		return models.Usuario{}, fmt.Errorf("erro ao gerar hash: %v", err)
	}

	userId, err := pu.services.CreateUsuario(newUser, senhaHash)

	if err != nil {
		return models.Usuario{}, err
	}

	newUser.ID.Int64 = userId
	newUser.ID.Valid = true
	return newUser, nil
}

func (pu *UsuarioUseCase) GetUsuarios() ([]models.Usuario, error) {
	return pu.services.GetUsuarios()
}

func (pu *UsuarioUseCase) GetUsuariosById(id int) (models.Usuario, error) {

	return pu.services.GetUsuarioById(id)
}

func (pu *UsuarioUseCase) PutUsuario(id int, UpdateUsuario models.Usuario) (models.Usuario, error) {

	UpdateUsuario, err := pu.services.PutUsuario(id, UpdateUsuario)
	if err != nil {
		if err == sql.ErrNoRows {

			return models.Usuario{}, ErrUserNotFound
		}
		return models.Usuario{}, err
	}

	return UpdateUsuario, nil
}
