package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
)

type PessoaServices struct {
	connection *sql.DB
}

func NewPessoaService(connection *sql.DB) PessoaServices {
	return PessoaServices{
		connection: connection,
	}
}

func (pr *PessoaServices) CreatePessoa(pessoa models.Pessoa) (int64, error) {
	var id int64

	query := `INSERT INTO pessoa (nome, tipo, documento, email, telefone, cargo) 
              VALUES ($1, $2, $3, $4, $5, $6) 
              RETURNING id`

	err := pr.connection.QueryRow(query,
		pessoa.Nome.String,
		pessoa.TipoDocumento.String,
		pessoa.Documento.String,
		pessoa.Email.String,
		pessoa.Telefone.String,
		pessoa.Cargo.String).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar pessoa: %v\n", err)
		return 0, err
	}

	return id, nil
}

func (pr *PessoaServices) GetPessoas() ([]models.Pessoa, error) {
	query := "select id, nome, tipo, documento, email, telefone, cargo, ativo, created_at, updated_at from pessoa"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Pessoa{}, err
	}

	var pessoasList []models.Pessoa
	var pessoaObj models.Pessoa

	for rows.Next() {
		err = rows.Scan(
			&pessoaObj.ID,
			&pessoaObj.Nome,
			&pessoaObj.TipoDocumento,
			&pessoaObj.Documento,
			&pessoaObj.Email,
			&pessoaObj.Telefone,
			&pessoaObj.Cargo,
			&pessoaObj.Ativo,
			&pessoaObj.CreatedAt,
			&pessoaObj.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []models.Pessoa{}, err
		}

		pessoasList = append(pessoasList, pessoaObj)

	}

	rows.Close()
	return pessoasList, nil
}
