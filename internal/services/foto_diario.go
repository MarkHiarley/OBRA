package services

import (
	"codxis-obras/internal/models"
	"database/sql"
	"fmt"
)

type FotoDiarioService struct {
	connection *sql.DB
}

func NewFotoDiarioService(connection *sql.DB) FotoDiarioService {
	return FotoDiarioService{
		connection: connection,
	}
}

// CreateFoto cria uma nova foto relacionada a uma entidade (metadados, atividade ou ocorrencia)
func (fds *FotoDiarioService) CreateFoto(foto models.FotoDiario) (models.FotoDiario, error) {
	query := `
		INSERT INTO foto_diario (entidade_tipo, entidade_id, foto, descricao, ordem, categoria, largura, altura, tamanho_bytes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`

	err := fds.connection.QueryRow(
		query,
		foto.EntidadeTipo,
		foto.EntidadeID,
		foto.Foto,
		foto.Descricao,
		foto.Ordem,
		foto.Categoria,
		foto.Largura,
		foto.Altura,
		foto.TamanhoBytes,
	).Scan(&foto.ID, &foto.CreatedAt, &foto.UpdatedAt)

	if err != nil {
		return models.FotoDiario{}, fmt.Errorf("erro ao criar foto: %v", err)
	}

	return foto, nil
}

// CreateFotos cria múltiplas fotos de uma vez (batch insert)
func (fds *FotoDiarioService) CreateFotos(fotos []models.FotoDiario, entidadeTipo string, entidadeID int64) error {
	if len(fotos) == 0 {
		return nil // Não há fotos para inserir
	}

	tx, err := fds.connection.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %v", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO foto_diario (entidade_tipo, entidade_id, foto, descricao, ordem, categoria, largura, altura, tamanho_bytes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`)
	if err != nil {
		return fmt.Errorf("erro ao preparar statement: %v", err)
	}
	defer stmt.Close()

	for i, foto := range fotos {
		// Se ordem não foi especificada, usar índice do array
		ordem := foto.Ordem
		if !ordem.Valid {
			ordem.Int64 = int64(i)
			ordem.Valid = true
		}

		// Se categoria não foi especificada, usar padrão baseado no tipo
		categoria := foto.Categoria
		if !categoria.Valid || categoria.String == "" {
			switch entidadeTipo {
			case "atividade":
				categoria.String = "ATIVIDADE"
			case "ocorrencia":
				categoria.String = "OCORRENCIA"
			case "metadados":
				categoria.String = "DIARIO"
			default:
				categoria.String = "DIARIO"
			}
			categoria.Valid = true
		}

		_, err = stmt.Exec(
			entidadeTipo,
			entidadeID,
			foto.Foto,
			foto.Descricao,
			ordem,
			categoria,
			foto.Largura,
			foto.Altura,
			foto.TamanhoBytes,
		)
		if err != nil {
			return fmt.Errorf("erro ao inserir foto %d: %v", i, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("erro ao commitar transação: %v", err)
	}

	return nil
}

// GetFotosByEntidade retorna todas as fotos de uma entidade específica
func (fds *FotoDiarioService) GetFotosByEntidade(entidadeTipo string, entidadeID int64) ([]models.FotoDiario, error) {
	query := `
		SELECT id, entidade_tipo, entidade_id, foto, descricao, ordem, categoria, largura, altura, tamanho_bytes, created_at, updated_at
		FROM foto_diario
		WHERE entidade_tipo = $1 AND entidade_id = $2
		ORDER BY ordem ASC, created_at ASC
	`

	rows, err := fds.connection.Query(query, entidadeTipo, entidadeID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar fotos: %v", err)
	}
	defer rows.Close()

	var fotos []models.FotoDiario
	for rows.Next() {
		var foto models.FotoDiario
		err := rows.Scan(
			&foto.ID,
			&foto.EntidadeTipo,
			&foto.EntidadeID,
			&foto.Foto,
			&foto.Descricao,
			&foto.Ordem,
			&foto.Categoria,
			&foto.Largura,
			&foto.Altura,
			&foto.TamanhoBytes,
			&foto.CreatedAt,
			&foto.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear foto: %v", err)
		}
		fotos = append(fotos, foto)
	}

	return fotos, nil
}

// DeleteFotosByEntidade deleta todas as fotos de uma entidade
func (fds *FotoDiarioService) DeleteFotosByEntidade(entidadeTipo string, entidadeID int64) error {
	query := `DELETE FROM foto_diario WHERE entidade_tipo = $1 AND entidade_id = $2`

	_, err := fds.connection.Exec(query, entidadeTipo, entidadeID)
	if err != nil {
		return fmt.Errorf("erro ao deletar fotos: %v", err)
	}

	return nil
}

// DeleteFoto deleta uma foto específica por ID
func (fds *FotoDiarioService) DeleteFoto(id int64) error {
	query := `DELETE FROM foto_diario WHERE id = $1`

	result, err := fds.connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar foto: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("foto não encontrada")
	}

	return nil
}
