-- Migration 000034: Adiciona campos lote e descricao na tabela obra
-- Para uso nos relatórios fotográfico e diário semanal

-- Adicionar coluna lote (identificação do lote da obra)
ALTER TABLE obra ADD COLUMN IF NOT EXISTS lote VARCHAR(100);

-- Adicionar coluna descricao (descrição breve da obra)
ALTER TABLE obra ADD COLUMN IF NOT EXISTS descricao TEXT;

-- Comentários
COMMENT ON COLUMN obra.lote IS 'Identificação do lote da obra (ex: LOTE-A, LOTE-01)';
COMMENT ON COLUMN obra.descricao IS 'Descrição breve e objetiva do que é a obra';

-- Criar índice para facilitar buscas por lote
CREATE INDEX IF NOT EXISTS idx_obra_lote ON obra(lote);
