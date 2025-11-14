-- Remover trigger
DROP TRIGGER IF EXISTS trigger_update_foto_diario_updated_at ON foto_diario;
DROP FUNCTION IF EXISTS update_foto_diario_updated_at();

-- Remover índices
DROP INDEX IF EXISTS idx_foto_diario_ordem;
DROP INDEX IF EXISTS idx_foto_diario_categoria;
DROP INDEX IF EXISTS idx_foto_diario_entidade;

-- Remover tabela
DROP TABLE IF EXISTS foto_diario;
