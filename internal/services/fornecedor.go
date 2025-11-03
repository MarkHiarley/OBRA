package services

import (
	"codxis-obras/internal/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type FornecedorServices struct {
	connection *sql.DB
}

func NewFornecedorService(connection *sql.DB) FornecedorServices {
	return FornecedorServices{
		connection: connection,
	}
}

func (fs *FornecedorServices) CreateFornecedor(fornecedor models.Fornecedor) (int64, error) {
	var id int64

	// Persist contact fields (contato_nome, contato_telefone, contato_email) added to model
	query := `INSERT INTO fornecedor (nome, tipo_documento, documento, email, telefone, endereco, cidade, estado, contato_nome, contato_telefone, contato_email, ativo) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
			  RETURNING id`

	err := fs.connection.QueryRow(query,
		fornecedor.Nome.String,
		fornecedor.TipoDocumento.String,
		fornecedor.Documento.String,
		fornecedor.Email.String,
		fornecedor.Telefone.String,
		fornecedor.Endereco.String,
		fornecedor.Cidade.String,
		fornecedor.Estado.String,
		fornecedor.ContatoNome.String,
		fornecedor.ContatoTelefone.String,
		fornecedor.ContatoEmail.String,
		fornecedor.Ativo.Bool).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar fornecedor: %v\n", err)
		return 0, err
	}

	return id, nil
}

func (fs *FornecedorServices) GetFornecedores() ([]models.Fornecedor, error) {
	query := `SELECT id, nome, tipo_documento, documento, email, telefone, endereco, cidade, estado, contato_nome, contato_telefone, contato_email, ativo, created_at, updated_at 
			  FROM fornecedor 
			  ORDER BY nome`

	rows, err := fs.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Fornecedor{}, err
	}
	defer rows.Close()

	var fornecedoresList []models.Fornecedor

	for rows.Next() {
		var fornecedor models.Fornecedor

		err = rows.Scan(
			&fornecedor.ID,
			&fornecedor.Nome,
			&fornecedor.TipoDocumento,
			&fornecedor.Documento,
			&fornecedor.Email,
			&fornecedor.Telefone,
			&fornecedor.Endereco,
			&fornecedor.Cidade,
			&fornecedor.Estado,
			&fornecedor.ContatoNome,
			&fornecedor.ContatoTelefone,
			&fornecedor.ContatoEmail,
			&fornecedor.Ativo,
			&fornecedor.CreatedAt,
			&fornecedor.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []models.Fornecedor{}, err
		}

		fornecedoresList = append(fornecedoresList, fornecedor)
	}

	return fornecedoresList, nil
}

func (fs *FornecedorServices) GetFornecedorById(id int64) (models.Fornecedor, error) {
	query := `SELECT id, nome, tipo_documento, documento, email, telefone, endereco, cidade, estado, contato_nome, contato_telefone, contato_email, ativo, created_at, updated_at 
			  FROM fornecedor 
			  WHERE id = $1`

	row := fs.connection.QueryRow(query, id)

	var fornecedor models.Fornecedor


	err := row.Scan(
		&fornecedor.ID,
		&fornecedor.Nome,
		&fornecedor.TipoDocumento,
		&fornecedor.Documento,
		&fornecedor.Email,
		&fornecedor.Telefone,
		&fornecedor.Endereco,
		&fornecedor.Cidade,
		&fornecedor.Estado,
		&fornecedor.ContatoNome,
		&fornecedor.ContatoTelefone,
		&fornecedor.ContatoEmail,
		&fornecedor.Ativo,
		&fornecedor.CreatedAt,
		&fornecedor.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Fornecedor{}, fmt.Errorf("fornecedor não encontrado")
		}
		return models.Fornecedor{}, err
	}

	return fornecedor, nil
}

func (fs *FornecedorServices) PutFornecedor(id int, fornecedorToUpdate models.Fornecedor) (models.Fornecedor, error) {
	query := `
		UPDATE fornecedor 
		SET 
			nome = $1,
			tipo_documento = $2, 
			documento = $3, 
			email = $4, 
			telefone = $5, 
			endereco = $6, 
			cidade = $7,
			estado = $8,
			contato_nome = $9,
			contato_telefone = $10,
			contato_email = $11,
			ativo = $12,
			updated_at = $13
		WHERE id = $14
		RETURNING id, nome, tipo_documento, documento, email, telefone, endereco, cidade, estado, contato_nome, contato_telefone, contato_email, ativo, created_at, updated_at`

	var updatedFornecedor models.Fornecedor


	err := fs.connection.QueryRowContext(context.Background(), query,
		fornecedorToUpdate.Nome.String,
		fornecedorToUpdate.TipoDocumento.String,
		fornecedorToUpdate.Documento.String,
		fornecedorToUpdate.Email.String,
		fornecedorToUpdate.Telefone.String,
		fornecedorToUpdate.Endereco.String,
		fornecedorToUpdate.Cidade.String,
		fornecedorToUpdate.Estado.String,
		fornecedorToUpdate.ContatoNome.String,
		fornecedorToUpdate.ContatoTelefone.String,
		fornecedorToUpdate.ContatoEmail.String,
		fornecedorToUpdate.Ativo.Bool,
		time.Now(),
		id,
	).Scan(
		&updatedFornecedor.ID,
		&updatedFornecedor.Nome,
		&updatedFornecedor.TipoDocumento,
		&updatedFornecedor.Documento,
		&updatedFornecedor.Email,
		&updatedFornecedor.Telefone,
		&updatedFornecedor.Endereco,
		&updatedFornecedor.Cidade,
		&updatedFornecedor.Estado,
		&updatedFornecedor.ContatoNome,
		&updatedFornecedor.ContatoTelefone,
		&updatedFornecedor.ContatoEmail,
		&updatedFornecedor.Ativo,
		&updatedFornecedor.CreatedAt,
		&updatedFornecedor.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Fornecedor{}, err
		}
		log.Printf("Error updating fornecedor: %v\n", err)
		return models.Fornecedor{}, fmt.Errorf("não foi possível atualizar esse fornecedor: %w", err)
	}

	return updatedFornecedor, nil
}

func (fs *FornecedorServices) DeleteFornecedorById(id int) error {
	query := "DELETE FROM fornecedor WHERE id = $1"

	result, err := fs.connection.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("erro ao executar a query de delete")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhum fornecedor encontrado com o ID fornecido")
	}

	return nil
}

// GetFornecedoresAtivos retorna apenas fornecedores ativos
func (fs *FornecedorServices) GetFornecedoresAtivos() ([]models.Fornecedor, error) {
	query := `SELECT id, nome, tipo_documento, documento, email, telefone, endereco, cidade, estado, ativo, created_at, updated_at 
	          FROM fornecedor 
	          WHERE ativo = true
	          ORDER BY nome`

	rows, err := fs.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Fornecedor{}, err
	}
	defer rows.Close()

	var fornecedoresList []models.Fornecedor

	for rows.Next() {
		var fornecedor models.Fornecedor
		err = rows.Scan(
			&fornecedor.ID,
			&fornecedor.Nome,
			&fornecedor.TipoDocumento,
			&fornecedor.Documento,
			&fornecedor.Email,
			&fornecedor.Telefone,
			&fornecedor.Endereco,
			&fornecedor.Cidade,
			&fornecedor.Estado,
			&fornecedor.Ativo,
			&fornecedor.CreatedAt,
			&fornecedor.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []models.Fornecedor{}, err
		}

		fornecedoresList = append(fornecedoresList, fornecedor)
	}

	return fornecedoresList, nil
}
