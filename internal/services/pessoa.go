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

	query := `INSERT INTO pessoa (nome, tipo, documento, email, telefone, cargo, endereco_rua, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
              RETURNING id`

	err := pr.connection.QueryRow(query,
		pessoa.Nome.String,
		pessoa.TipoDocumento.String,
		pessoa.Documento.String,
		pessoa.Email.String,
		pessoa.Telefone.String,
		pessoa.Cargo.String,
		pessoa.EnderecoRua.String,
		pessoa.EnderecoNumero.String,
		pessoa.EnderecoComplemento.String,
		pessoa.EnderecoBairro.String,
		pessoa.EnderecoCidade.String,
		pessoa.EnderecoEstado.String,
		pessoa.EnderecoCep.String).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar pessoa: %v\n", err)
		return 0, err
	}

	return id, nil
}

func (pr *PessoaServices) GetPessoas() ([]models.Pessoa, error) {
	query := "select id, nome, tipo, documento, email, telefone, cargo, endereco_rua, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, ativo, created_at, updated_at from pessoa"
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
			&pessoaObj.EnderecoRua,
			&pessoaObj.EnderecoNumero,
			&pessoaObj.EnderecoComplemento,
			&pessoaObj.EnderecoBairro,
			&pessoaObj.EnderecoCidade,
			&pessoaObj.EnderecoEstado,
			&pessoaObj.EnderecoCep,
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
	query := "select id, nome, tipo, documento, email, telefone, cargo, endereco_rua, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, ativo, created_at, updated_at from pessoa where id = $1"

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
		&pessoa.EnderecoRua,
		&pessoa.EnderecoNumero,
		&pessoa.EnderecoComplemento,
		&pessoa.EnderecoBairro,
		&pessoa.EnderecoCidade,
		&pessoa.EnderecoEstado,
		&pessoa.EnderecoCep,
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
            endereco_rua = $7,
            endereco_numero = $8,
            endereco_complemento = $9,
            endereco_bairro = $10,
            endereco_cidade = $11,
            endereco_estado = $12,
            endereco_cep = $13, 
            ativo = $14,
			updated_at = $15
        WHERE id = $16
        RETURNING id, nome, tipo, documento, email, telefone, cargo, endereco_rua, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, ativo, created_at, updated_at`
	var updatedPessoa models.Pessoa

	err := pr.connection.QueryRowContext(context.Background(), query,
		pessoaToUpdate.Nome.String,
		pessoaToUpdate.TipoDocumento.String,
		pessoaToUpdate.Documento.String,
		pessoaToUpdate.Email.String,
		pessoaToUpdate.Telefone.String,
		pessoaToUpdate.Cargo.String,
		pessoaToUpdate.EnderecoRua.String,
		pessoaToUpdate.EnderecoNumero.String,
		pessoaToUpdate.EnderecoComplemento.String,
		pessoaToUpdate.EnderecoBairro.String,
		pessoaToUpdate.EnderecoCidade.String,
		pessoaToUpdate.EnderecoEstado.String,
		pessoaToUpdate.EnderecoCep.String,
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
		&updatedPessoa.EnderecoRua,
		&updatedPessoa.EnderecoNumero,
		&updatedPessoa.EnderecoComplemento,
		&updatedPessoa.EnderecoBairro,
		&updatedPessoa.EnderecoCidade,
		&updatedPessoa.EnderecoEstado,
		&updatedPessoa.EnderecoCep,
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
