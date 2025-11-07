-- Remove o campo pessoa_id da tabela despesa
DROP INDEX IF EXISTS idx_despesa_pessoa;
ALTER TABLE despesa DROP COLUMN IF EXISTS pessoa_id;