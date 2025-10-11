package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
)

type UsuarioServices struct {
	connection *sql.DB
}

func NewUsuarioService(connection *sql.DB) UsuarioServices {
	return UsuarioServices{
		connection: connection,
	}
}

func (pr *UsuarioServices) CreateUsuario(user models.Usuario, senhaHash []byte) (int, error) {
	var id int

	query := `
		INSERT INTO usuario (
			nome,
			email, 
			senha_hash,
			tipo_documento,
			documento,
			telefone, 
			perfil_acesso,
			ativo
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	perfilArray := fmt.Sprintf("{%s}", user.Perfil_acesso)
	err := pr.connection.QueryRow(
		query,
		user.Nome,
		user.Email,
		string(senhaHash),
		user.Tipo_documento,
		user.Documento,
		user.Telefone,
		perfilArray,
		true,
	).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar usuário: %v\n", err)
		return 0, err
	}

	return id, nil
}

func (pr *UsuarioServices) GetUsuarios() ([]models.Usuario, error) {
	query := "select id, nome, email, tipo_documento, documento, telefone, perfil_acesso, ativo, created_at, updated_at from usuario"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Usuario{}, err
	}

	var usuarioList []models.Usuario
	var usuarioObj models.Usuario

	for rows.Next() {
		err = rows.Scan(
			&usuarioObj.ID,
			&usuarioObj.Nome,
			&usuarioObj.Email,
			&usuarioObj.Tipo_documento,
			&usuarioObj.Documento,
			&usuarioObj.Telefone,
			&usuarioObj.Perfil_acesso,
			&usuarioObj.Ativo,
			&usuarioObj.CreatedAt,
			&usuarioObj.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return []models.Usuario{}, err
		}

		usuarioList = append(usuarioList, usuarioObj)

	}

	rows.Close()
	return usuarioList, nil
}
