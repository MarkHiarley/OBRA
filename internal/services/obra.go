package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
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

	query := `INSERT INTO obra (nome, contrato_numero, contratante_id, responsavel_id, data_inicio, prazo_dias, data_fim_prevista, orcamento, status, endereco_rua, endereco_numero, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, observacoes, ativo ) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) 
              RETURNING id`

	err := pr.connection.QueryRow(query,
		obra.Nome.String,
		obra.ContratoNumero.String,
		obra.ContratanteID.Int64,
		obra.ResponsavelID.Int64,
		obra.DataInicio.String,
		obra.PrazoDias.Int64,
		obra.DataFimPrevista.String,
		obra.Orcamento.Float64,
		obra.Status.String,
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
	query := "select id, nome, contrato_numero, contratante_id, responsavel_id, data_inicio, prazo_dias, data_fim_prevista, orcamento, status, endereco_rua, endereco_numero, endereco_bairro, endereco_cidade, endereco_estado, endereco_cep, observacoes, ativo, created_at, updated_at from obra"
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
			&obraObj.ResponsavelID,
			&obraObj.DataInicio,
			&obraObj.PrazoDias,
			&obraObj.DataFimPrevista,
			&obraObj.Orcamento,
			&obraObj.Status,
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
	query := "select id, nome, contrato_numero, contratante_id, responsavel_id, data_inicio, prazo_dias, data_fim_prevista,orcamento, endereco_rua,endereco_numero, endereco_bairro, endereco_cidade,endereco_estado, endereco_cep, observacoes, status, ativo,created_at, updated_at from obra where id = $1"

	row := pr.connection.QueryRow(query, id)
	fmt.Println(query, id)
	var obra models.Obra

	err := row.Scan(
		&obra.ID,
		&obra.Nome,
		&obra.ContratoNumero,
		&obra.ContratanteID,
		&obra.ResponsavelID,
		&obra.DataInicio,
		&obra.PrazoDias,
		&obra.DataFimPrevista,
		&obra.Orcamento,
		&obra.EnderecoRua,
		&obra.EnderecoNumero,
		&obra.EnderecoBairro,
		&obra.EnderecoCidade,
		&obra.EnderecoEstado,
		&obra.EnderecoCep,
		&obra.Observacoes,
		&obra.Status,
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
