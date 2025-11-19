-- Migration 000034: Remove campos lote e descricao da tabela obra

-- Remover índice
DROP INDEX IF EXISTS idx_obra_lote;

-- Remover colunas
ALTER TABLE obra DROP COLUMN IF EXISTS descricao;
ALTER TABLE obra DROP COLUMN IF EXISTS lote;
