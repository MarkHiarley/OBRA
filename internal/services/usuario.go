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

	perfilArray := fmt.Sprintf("{%s}", user.PerfilAcesso.String)
	err := pr.connection.QueryRow(
		query,
		user.Nome.NullString,
		user.Email.NullString,
		string(senhaHash),
		user.TipoDocumento.NullString,
		user.Documento.NullString,
		user.Telefone.NullString,
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
			&usuarioObj.TipoDocumento,
			&usuarioObj.Documento,
			&usuarioObj.Telefone,
			&usuarioObj.PerfilAcesso,
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

func (pr UsuarioServices) GetUsuarioById(id int) (models.Usuario, error) {

	//id, nome, email, tipo_documento, documento, telefone, perfil_acesso, ativo, created_at, updated_at
	query := "select id, nome, email, tipo_documento, documento, telefone, perfil_acesso, ativo, created_at, updated_at from usuario where id = $1"

	row := pr.connection.QueryRow(query, id)

	var usuario models.Usuario

	err := row.Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Email,
		&usuario.TipoDocumento,
		&usuario.Documento,
		&usuario.Telefone,
		&usuario.PerfilAcesso,
		&usuario.Ativo,
		&usuario.CreatedAt,
		&usuario.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {

			return models.Usuario{}, fmt.Errorf("usuário não encontrado")
		}

		return models.Usuario{}, err
	}

	return usuario, nil

}
