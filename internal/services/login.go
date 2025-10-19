package services

import (
	"database/sql"
	"fmt"
)

type LoginServices struct {
	connection *sql.DB
}

func NewLoginService(connection *sql.DB) LoginServices {
	return LoginServices{
		connection: connection,
	}
}

func (pr *LoginServices) CheckUser(email string) (senha_hash string, error error) {
	//'sase@gmadsddsdsl.com'

	query := "SELECT senha_hash FROM usuario WHERE email = $1"

	row := pr.connection.QueryRow(query, email)
	err := row.Scan(&senha_hash)

	if err != nil {
		if err == sql.ErrNoRows {

			return "", fmt.Errorf("usuário não encontrado")
		}

		return "", err
	}
	return senha_hash, nil

}
