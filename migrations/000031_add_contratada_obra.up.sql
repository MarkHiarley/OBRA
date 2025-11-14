-- Migration 000031: Adiciona campo contratada na tabela obra

-- Adicionar coluna contratada (empresa executora da obra)
ALTER TABLE obra ADD COLUMN IF NOT EXISTS contratada VARCHAR(255);

COMMENT ON COLUMN obra.contratada IS 'Empresa contratada para executar a obra';
