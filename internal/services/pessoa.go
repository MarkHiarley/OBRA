package services

import (
	"codxis-obras/internal/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
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

func (pr PessoaServices) GetPessoaById(id int64) (models.Pessoa, error) {

	//id, nome, email, tipo_documento, documento, telefone, perfil_acesso, ativo, created_at, updated_at
	query := "select id, nome, tipo, documento, email, telefone, cargo, ativo,created_at, updated_at from pessoa where id = $1"

	row := pr.connection.QueryRow(query, id)

	var pessoa models.Pessoa

	err := row.Scan(
		&pessoa.ID,
		&pessoa.Nome,
		&pessoa.TipoDocumento,
		&pessoa.Documento,
		&pessoa.Email,
		&pessoa.Telefone,
		&pessoa.Cargo,
		&pessoa.Ativo,
		&pessoa.CreatedAt,
		&pessoa.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {

			return models.Pessoa{}, fmt.Errorf("pessoa não encontrado")
		}

		return models.Pessoa{}, err
	}

	return pessoa, nil

}

func (pr PessoaServices) PutPessoa(id int, pessoaToUpdate models.Pessoa) (models.Pessoa, error) {

	query := `
        UPDATE pessoa 
        SET 
            nome = $1,
            tipo = $2, 
            documento = $3, 
            email = $4, 
            telefone = $5, 
            cargo = $6, 
            ativo = $7,
			updated_at =$8
        WHERE id = $9
        RETURNING id, nome, tipo, documento, email, telefone, cargo, ativo, created_at, updated_at`
	var updatedPessoa models.Pessoa

	err := pr.connection.QueryRowContext(context.Background(), query,
		pessoaToUpdate.Nome.String,
		pessoaToUpdate.TipoDocumento.String,
		pessoaToUpdate.Documento.String,
		pessoaToUpdate.Email.String,
		pessoaToUpdate.Telefone.String,
		pessoaToUpdate.Cargo.String,
		pessoaToUpdate.Ativo.Bool,
		time.Now(),
		id, // The ID for the WHERE clause
	).Scan(

		&updatedPessoa.ID,
		&updatedPessoa.Nome,
		&updatedPessoa.TipoDocumento,
		&updatedPessoa.Documento,
		&updatedPessoa.Email,
		&updatedPessoa.Telefone,
		&updatedPessoa.Cargo,
		&updatedPessoa.Ativo,
		&updatedPessoa.CreatedAt,
		&updatedPessoa.UpdatedAt,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return models.Pessoa{}, err
		}

		log.Printf("Error updating user: %v\n", err)
		return models.Pessoa{}, fmt.Errorf("não foi possivel atualizar essa pessoa: %w", err)
	}

	return updatedPessoa, nil
}

func (pr *PessoaServices) DeletePessoaById(id int) error {
	query := "DELETE FROM pessoa WHERE id = $1"

	result, err := pr.connection.ExecContext(context.Background(), query, id)
	if err != nil {

		return fmt.Errorf("erro ao executar a query de delete")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhuma pessoa encontrada com o ID fornecido")
	}

	return nil
}
