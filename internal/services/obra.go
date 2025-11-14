package services

import (
	"codxis-obras/internal/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type ObraServices struct {
	connection *sql.DB
}

func NewObraService(connection *sql.DB) ObraServices {
	return ObraServices{
		connection: connection,
	}
}

func (pr *ObraServices) CreateObra(obra models.Obra) (int64, error) {
	var id int64

	query := `INSERT INTO obra (nome, contrato_numero, contratante_id, contratada, responsavel_id, data_inicio, prazo_dias, data_fim_prevista, orcamento, status, art, foto, endereco_rua, endereco_numero, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, observacoes, ativo ) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20) 
			  RETURNING id`

	err := pr.connection.QueryRow(query,
		obra.Nome.String,
		obra.ContratoNumero.String,
		obra.ContratanteID.Int64,
		obra.Contratada.String,
		obra.ResponsavelID.Int64,
		obra.DataInicio.String,
		obra.PrazoDias.Int64,
		obra.DataFimPrevista.String,
		obra.Orcamento.Float64,
		obra.Status.String,
		obra.Art.String,
		obra.Foto.String,
		obra.EnderecoRua.String,
		obra.EnderecoNumero.String,
		obra.EnderecoBairro.String,
		obra.EnderecoCidade.String,
		obra.EnderecoEstado.String,
		obra.EnderecoCep.String,
		obra.Observacoes.String,
		obra.Ativo).Scan(&id)

	if err != nil {
		fmt.Printf("Erro ao criar Obra: %v\n", err)
		return 0, err
	}

	return id, nil
}

func (pr *ObraServices) GetObras() ([]models.Obra, error) {
	query := "select id, nome, contrato_numero, contratante_id, contratada, responsavel_id, data_inicio, prazo_dias, data_fim_prevista, orcamento, status, art, foto, endereco_rua, endereco_numero, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, observacoes, ativo, created_at, updated_at from obra"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Obra{}, err
	}

	var obraList []models.Obra
	var obraObj models.Obra

	for rows.Next() {

		err = rows.Scan(
			&obraObj.ID,
			&obraObj.Nome,
			&obraObj.ContratoNumero,
			&obraObj.ContratanteID,
			&obraObj.Contratada,
			&obraObj.ResponsavelID,
			&obraObj.DataInicio,
			&obraObj.PrazoDias,
			&obraObj.DataFimPrevista,
			&obraObj.Orcamento,
			&obraObj.Status,
			&obraObj.Art,
			&obraObj.Foto,
			&obraObj.EnderecoRua,
			&obraObj.EnderecoNumero,
			&obraObj.EnderecoBairro,
			&obraObj.EnderecoCidade,
			&obraObj.EnderecoEstado,
			&obraObj.EnderecoCep,
			&obraObj.Observacoes,
			&obraObj.Ativo,
			&obraObj.CreatedAt,
			&obraObj.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			return []models.Obra{}, err
		}

		obraList = append(obraList, obraObj)

	}

	rows.Close()
	return obraList, nil
}

func (pr ObraServices) GetObraById(id int64) (models.Obra, error) {

	//id, nome, email, tipo_documento, documento, telefone, perfil_acesso, ativo, created_at, updated_at
	query := "select id, nome, contrato_numero, contratante_id, contratada, responsavel_id, data_inicio, prazo_dias, data_fim_prevista, orcamento, status, art, foto, endereco_rua, endereco_numero, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, observacoes, ativo, created_at, updated_at from obra where id = $1"

	row := pr.connection.QueryRow(query, id)
	fmt.Println(query, id)
	var obra models.Obra

	err := row.Scan(
		&obra.ID,
		&obra.Nome,
		&obra.ContratoNumero,
		&obra.ContratanteID,
		&obra.Contratada,
		&obra.ResponsavelID,
		&obra.DataInicio,
		&obra.PrazoDias,
		&obra.DataFimPrevista,
		&obra.Orcamento,
		&obra.Status,
		&obra.Art,
		&obra.Foto,
		&obra.EnderecoRua,
		&obra.EnderecoNumero,
		&obra.EnderecoBairro,
		&obra.EnderecoCidade,
		&obra.EnderecoEstado,
		&obra.EnderecoCep,
		&obra.Observacoes,
		&obra.Ativo,
		&obra.CreatedAt,
		&obra.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {

			return models.Obra{}, fmt.Errorf("obra não encontrado")
		}

		return models.Obra{}, err
	}

	return obra, nil

}
func (pr ObraServices) PutObra(id int, ObraToUpdate models.Obra) (models.Obra, error) {

	query := `
		UPDATE obra 
		SET 
			nome = $1,
			contrato_numero = $2, 
			contratante_id = $3,
			contratada = $4, 
			responsavel_id = $5, 
			data_inicio = $6, 
			prazo_dias = $7, 
			data_fim_prevista = $8,
			orcamento = $9,
			status = $10,
			art = $11,
			foto = $12,
			endereco_rua = $13,
			endereco_numero = $14,
			endereco_bairro = $15,
			endereco_cidade = $16,
			endereco_estado = $17,
			endereco_cep = $18,
			observacoes = $19,
			ativo = $20,
			updated_at = $21
		WHERE id = $22
		RETURNING id, nome, contrato_numero, contratante_id, contratada, responsavel_id, data_inicio, prazo_dias, data_fim_prevista, orcamento, status, art, foto, endereco_rua, endereco_numero, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, observacoes, ativo, created_at, updated_at`
	var updatedObra models.Obra

	err := pr.connection.QueryRowContext(context.Background(), query,
		ObraToUpdate.Nome.String,
		ObraToUpdate.ContratoNumero.String,
		ObraToUpdate.ContratanteID.Int64,
		ObraToUpdate.Contratada.String,
		ObraToUpdate.ResponsavelID.Int64,
		ObraToUpdate.DataInicio.String,
		ObraToUpdate.PrazoDias.Int64,
		ObraToUpdate.DataFimPrevista.String,
		ObraToUpdate.Orcamento.Float64,
		ObraToUpdate.Status.String,
		ObraToUpdate.Art.String,
		ObraToUpdate.Foto.String,
		ObraToUpdate.EnderecoRua.String,
		ObraToUpdate.EnderecoNumero.String,
		ObraToUpdate.EnderecoBairro.String,
		ObraToUpdate.EnderecoCidade.String,
		ObraToUpdate.EnderecoEstado.String,
		ObraToUpdate.EnderecoCep.String,
		ObraToUpdate.Observacoes.String,
		ObraToUpdate.Ativo.Bool,
		time.Now(),
		id,
	).Scan(
		&updatedObra.ID,
		&updatedObra.Nome,
		&updatedObra.ContratoNumero,
		&updatedObra.ContratanteID,
		&updatedObra.Contratada,
		&updatedObra.ResponsavelID,
		&updatedObra.DataInicio,
		&updatedObra.PrazoDias,
		&updatedObra.DataFimPrevista,
		&updatedObra.Orcamento,
		&updatedObra.Status,
		&updatedObra.Art,
		&updatedObra.Foto,
		&updatedObra.EnderecoRua,
		&updatedObra.EnderecoNumero,
		&updatedObra.EnderecoBairro,
		&updatedObra.EnderecoCidade,
		&updatedObra.EnderecoEstado,
		&updatedObra.EnderecoCep,
		&updatedObra.Observacoes,
		&updatedObra.Ativo,
		&updatedObra.CreatedAt,
		&updatedObra.UpdatedAt,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return models.Obra{}, err
		}

		log.Printf("Error updating user: %v\n", err)
		return models.Obra{}, fmt.Errorf("não foi possivel atualizar essa obra: %w", err)
	}

	return updatedObra, nil
}

func (pr *ObraServices) DeleteObraById(id int) error {
	query := "DELETE FROM obra WHERE id = $1"

	result, err := pr.connection.ExecContext(context.Background(), query, id)
	if err != nil {

		return fmt.Errorf("erro ao executar a query de delete")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao obter linhas afetadas")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhuma obra encontrada com o ID fornecido")
	}

	return nil
}
