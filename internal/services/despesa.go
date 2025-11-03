package services

import (
	"codxis-obras/internal/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type DespesaServices struct {
	connection *sql.DB
}

func NewDespesaService(connection *sql.DB) DespesaServices {
	return DespesaServices{
		connection: connection,
	}
}

func (ds *DespesaServices) CreateDespesa(despesa models.Despesa) (int64, error) {
	var id int64

	query := `INSERT INTO despesa (obra_id, fornecedor_id, data, data_vencimento, descricao, categoria, valor, forma_pagamento, status_pagamento, data_pagamento, responsavel_pagamento, observacao) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
              RETURNING id`

	// Tratar campos nullable
	var fornecedorID, dataVencimento, dataPagamento interface{}

	if despesa.FornecedorID.Valid {
		fornecedorID = despesa.FornecedorID.Int64
	} else {
		fornecedorID = nil
	}

	if despesa.DataVencimento.Valid {
		dataVencimento = despesa.DataVencimento.Time
	} else {
		dataVencimento = nil
	}

	if despesa.DataPagamento.Valid {
		dataPagamento = despesa.DataPagamento.Time
	} else {
		dataPagamento = nil
	}

	err := ds.connection.QueryRow(query,
		despesa.ObraID.Int64,
		fornecedorID,
		despesa.Data.Time,
		dataVencimento,
		despesa.Descricao.String,
		despesa.Categoria.String,
		despesa.Valor.Float64,
		despesa.FormaPagamento.String,
		despesa.StatusPagamento.String,
		dataPagamento,
		despesa.ResponsavelPagamento.String,
		despesa.Observacao.String).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar despesa: %v\n", err)
		return 0, err
	}

	return id, nil
}

func (ds *DespesaServices) GetDespesas() ([]models.DespesaComRelacionamentos, error) {
	query := `
		SELECT 
			d.id, d.obra_id, d.fornecedor_id, d.data, d.data_vencimento, d.descricao, 
			d.categoria, d.valor, d.forma_pagamento, d.status_pagamento, 
			d.data_pagamento, d.responsavel_pagamento, d.observacao, 
			d.created_at, d.updated_at,
			f.nome as fornecedor_nome,
			o.nome as obra_nome
		FROM despesa d
		LEFT JOIN fornecedor f ON d.fornecedor_id = f.id
		LEFT JOIN obra o ON d.obra_id = o.id
		ORDER BY d.data DESC, d.created_at DESC`

	rows, err := ds.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.DespesaComRelacionamentos{}, err
	}
	defer rows.Close()

	var despesasList []models.DespesaComRelacionamentos

	for rows.Next() {
		var despesa models.DespesaComRelacionamentos
		err = rows.Scan(
			&despesa.ID,
			&despesa.ObraID,
			&despesa.FornecedorID,
			&despesa.Data,
			&despesa.DataVencimento,
			&despesa.Descricao,
			&despesa.Categoria,
			&despesa.Valor,
			&despesa.FormaPagamento,
			&despesa.StatusPagamento,
			&despesa.DataPagamento,
			&despesa.ResponsavelPagamento,
			&despesa.Observacao,
			&despesa.CreatedAt,
			&despesa.UpdatedAt,
			&despesa.FornecedorNome,
			&despesa.ObraNome,
		)

		if err != nil {
			fmt.Println(err)
			return []models.DespesaComRelacionamentos{}, err
		}

		despesasList = append(despesasList, despesa)
	}

	return despesasList, nil
}

func (ds *DespesaServices) GetDespesaById(id int64) (models.DespesaComRelacionamentos, error) {
	query := `
		SELECT 
			d.id, d.obra_id, d.fornecedor_id, d.data, d.data_vencimento, d.descricao, 
			d.categoria, d.valor, d.forma_pagamento, d.status_pagamento, 
			d.data_pagamento, d.responsavel_pagamento, d.observacao, 
			d.created_at, d.updated_at,
			f.nome as fornecedor_nome,
			o.nome as obra_nome
		FROM despesa d
		LEFT JOIN fornecedor f ON d.fornecedor_id = f.id
		LEFT JOIN obra o ON d.obra_id = o.id
		WHERE d.id = $1`

	row := ds.connection.QueryRow(query, id)

	var despesa models.DespesaComRelacionamentos

	err := row.Scan(
		&despesa.ID,
		&despesa.ObraID,
		&despesa.FornecedorID,
		&despesa.Data,
		&despesa.DataVencimento,
		&despesa.Descricao,
		&despesa.Categoria,
		&despesa.Valor,
		&despesa.FormaPagamento,
		&despesa.StatusPagamento,
		&despesa.DataPagamento,
		&despesa.ResponsavelPagamento,
		&despesa.Observacao,
		&despesa.CreatedAt,
		&despesa.UpdatedAt,
		&despesa.FornecedorNome,
		&despesa.ObraNome,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.DespesaComRelacionamentos{}, fmt.Errorf("despesa não encontrada")
		}
		return models.DespesaComRelacionamentos{}, err
	}

	return despesa, nil
}

func (ds *DespesaServices) GetDespesasByObraId(obraId int64) ([]models.DespesaComRelacionamentos, error) {
	query := `
		SELECT 
			d.id, d.obra_id, d.fornecedor_id, d.data, d.data_vencimento, d.descricao, 
			d.categoria, d.valor, d.forma_pagamento, d.status_pagamento, 
			d.data_pagamento, d.responsavel_pagamento, d.observacao, 
			d.created_at, d.updated_at,
			f.nome as fornecedor_nome,
			o.nome as obra_nome
		FROM despesa d
		LEFT JOIN fornecedor f ON d.fornecedor_id = f.id
		LEFT JOIN obra o ON d.obra_id = o.id
		WHERE d.obra_id = $1
		ORDER BY d.data DESC`

	rows, err := ds.connection.Query(query, obraId)
	if err != nil {
		fmt.Println(err)
		return []models.DespesaComRelacionamentos{}, err
	}
	defer rows.Close()

	var despesasList []models.DespesaComRelacionamentos

	for rows.Next() {
		var despesa models.DespesaComRelacionamentos
		err = rows.Scan(
			&despesa.ID,
			&despesa.ObraID,
			&despesa.FornecedorID,
			&despesa.Data,
			&despesa.DataVencimento,
			&despesa.Descricao,
			&despesa.Categoria,
			&despesa.Valor,
			&despesa.FormaPagamento,
			&despesa.StatusPagamento,
			&despesa.DataPagamento,
			&despesa.ResponsavelPagamento,
			&despesa.Observacao,
			&despesa.CreatedAt,
			&despesa.UpdatedAt,
			&despesa.FornecedorNome,
			&despesa.ObraNome,
		)

		if err != nil {
			fmt.Println(err)
			return []models.DespesaComRelacionamentos{}, err
		}

		despesasList = append(despesasList, despesa)
	}

	return despesasList, nil
}

func (ds *DespesaServices) GetDespesasByFornecedorId(fornecedorId int64) ([]models.DespesaComRelacionamentos, error) {
	query := `
		SELECT 
			d.id, d.obra_id, d.fornecedor_id, d.data, d.data_vencimento, d.descricao, 
			d.categoria, d.valor, d.forma_pagamento, d.status_pagamento, 
			d.data_pagamento, d.responsavel_pagamento, d.observacao, 
			d.created_at, d.updated_at,
			f.nome as fornecedor_nome,
			o.nome as obra_nome
		FROM despesa d
		LEFT JOIN fornecedor f ON d.fornecedor_id = f.id
		LEFT JOIN obra o ON d.obra_id = o.id
		WHERE d.fornecedor_id = $1
		ORDER BY d.data DESC`

	rows, err := ds.connection.Query(query, fornecedorId)
	if err != nil {
		fmt.Println(err)
		return []models.DespesaComRelacionamentos{}, err
	}
	defer rows.Close()

	var despesasList []models.DespesaComRelacionamentos

	for rows.Next() {
		var despesa models.DespesaComRelacionamentos
		err = rows.Scan(
			&despesa.ID,
			&despesa.ObraID,
			&despesa.FornecedorID,
			&despesa.Data,
			&despesa.DataVencimento,
			&despesa.Descricao,
			&despesa.Categoria,
			&despesa.Valor,
			&despesa.FormaPagamento,
			&despesa.StatusPagamento,
			&despesa.DataPagamento,
			&despesa.ResponsavelPagamento,
			&despesa.Observacao,
			&despesa.CreatedAt,
			&despesa.UpdatedAt,
			&despesa.FornecedorNome,
			&despesa.ObraNome,
		)

		if err != nil {
			fmt.Println(err)
			return []models.DespesaComRelacionamentos{}, err
		}

		despesasList = append(despesasList, despesa)
	}

	return despesasList, nil
}

func (ds *DespesaServices) GetRelatorioPorObra(obraId int64) ([]models.RelatorioDespesas, error) {
	query := `
		SELECT 
			d.obra_id,
			o.nome as obra_nome,
			d.categoria,
			SUM(d.valor) as total_despesas,
			COUNT(d.id) as quantidade_itens
		FROM despesa d
		LEFT JOIN obra o ON d.obra_id = o.id
		WHERE d.obra_id = $1
		GROUP BY d.obra_id, o.nome, d.categoria
		ORDER BY total_despesas DESC`

	rows, err := ds.connection.Query(query, obraId)
	if err != nil {
		fmt.Println(err)
		return []models.RelatorioDespesas{}, err
	}
	defer rows.Close()

	var relatorio []models.RelatorioDespesas

	for rows.Next() {
		var item models.RelatorioDespesas
		err = rows.Scan(
			&item.ObraID,
			&item.ObraNome,
			&item.Categoria,
			&item.TotalDespesas,
			&item.QuantidadeItens,
		)

		if err != nil {
			fmt.Println(err)
			return []models.RelatorioDespesas{}, err
		}

		relatorio = append(relatorio, item)
	}

	return relatorio, nil
}

func (ds *DespesaServices) PutDespesa(id int, despesaToUpdate models.Despesa) (models.Despesa, error) {
	query := `
        UPDATE despesa 
        SET 
            obra_id = $1,
            fornecedor_id = $2, 
            data = $3,
            data_vencimento = $4, 
            descricao = $5, 
            categoria = $6, 
            valor = $7, 
            forma_pagamento = $8,
            status_pagamento = $9,
            data_pagamento = $10,
            responsavel_pagamento = $11,
            observacao = $12,
			updated_at = $13
        WHERE id = $14
        RETURNING id, obra_id, fornecedor_id, data, data_vencimento, descricao, categoria, valor, 
                  forma_pagamento, status_pagamento, data_pagamento, responsavel_pagamento, 
                  observacao, created_at, updated_at`

	var updatedDespesa models.Despesa

	// Tratar campos nullable
	var fornecedorID, dataVencimento, dataPagamento interface{}

	if despesaToUpdate.FornecedorID.Valid {
		fornecedorID = despesaToUpdate.FornecedorID.Int64
	} else {
		fornecedorID = nil
	}

	if despesaToUpdate.DataVencimento.Valid {
		dataVencimento = despesaToUpdate.DataVencimento.Time
	} else {
		dataVencimento = nil
	}

	if despesaToUpdate.DataPagamento.Valid {
		dataPagamento = despesaToUpdate.DataPagamento.Time
	} else {
		dataPagamento = nil
	}

	err := ds.connection.QueryRowContext(context.Background(), query,
		despesaToUpdate.ObraID.Int64,
		fornecedorID,
		despesaToUpdate.Data.Time,
		dataVencimento,
		despesaToUpdate.Descricao.String,
		despesaToUpdate.Categoria.String,
		despesaToUpdate.Valor.Float64,
		despesaToUpdate.FormaPagamento.String,
		despesaToUpdate.StatusPagamento.String,
		dataPagamento,
		despesaToUpdate.ResponsavelPagamento.String,
		despesaToUpdate.Observacao.String,
		time.Now(),
		id,
	).Scan(
		&updatedDespesa.ID,
		&updatedDespesa.ObraID,
		&updatedDespesa.FornecedorID,
		&updatedDespesa.Data,
		&updatedDespesa.DataVencimento,
		&updatedDespesa.Descricao,
		&updatedDespesa.Categoria,
		&updatedDespesa.Valor,
		&updatedDespesa.FormaPagamento,
		&updatedDespesa.StatusPagamento,
		&updatedDespesa.DataPagamento,
		&updatedDespesa.ResponsavelPagamento,
		&updatedDespesa.Observacao,
		&updatedDespesa.CreatedAt,
		&updatedDespesa.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Despesa{}, err
		}
		log.Printf("Error updating despesa: %v\n", err)
		return models.Despesa{}, fmt.Errorf("não foi possível atualizar essa despesa: %w", err)
	}

	return updatedDespesa, nil
}

func (ds *DespesaServices) DeleteDespesaById(id int) error {
	query := "DELETE FROM despesa WHERE id = $1"

	result, err := ds.connection.ExecContext(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("erro ao executar a query de delete")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhuma despesa encontrada com o ID fornecido")
	}

	return nil
}
