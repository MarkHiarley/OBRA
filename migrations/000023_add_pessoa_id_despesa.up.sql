-- Adiciona campo pessoa_id para associar despesas de mão de obra a pessoas específicas
ALTER TABLE despesa 
ADD COLUMN IF NOT EXISTS pessoa_id INTEGER REFERENCES pessoa(id);

-- Criar índice para melhor performance
CREATE INDEX IF NOT EXISTS idx_despesa_pessoa ON despesa(pessoa_id);

-- Comentário explicativo
COMMENT ON COLUMN despesa.pessoa_id IS 'Referência à pessoa responsável (usado principalmente para despesas de mão de obra)';