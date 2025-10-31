-- Rollback: Reverter alterações da despesa

-- 1. Remover índices novos
DROP INDEX IF EXISTS idx_despesa_data_vencimento;
DROP INDEX IF EXISTS idx_despesa_data;

-- 2. Adicionar coluna data_despesa de volta
ALTER TABLE despesa ADD COLUMN IF NOT EXISTS data_despesa DATE;

-- 3. Copiar dados de data para data_despesa
UPDATE despesa SET data_despesa = data WHERE data_despesa IS NULL;

-- 4. Tornar data_despesa NOT NULL
ALTER TABLE despesa ALTER COLUMN data_despesa SET NOT NULL;

-- 5. Remover colunas novas
ALTER TABLE despesa DROP COLUMN IF EXISTS data;
ALTER TABLE despesa DROP COLUMN IF EXISTS data_vencimento;

-- 6. Restaurar constraint antiga de categoria
ALTER TABLE despesa DROP CONSTRAINT IF EXISTS despesa_categoria_check;
ALTER TABLE despesa ADD CONSTRAINT despesa_categoria_check CHECK (categoria IN (
    'MATERIAL', 
    'MAO_DE_OBRA', 
    'COMBUSTIVEL', 
    'ALIMENTACAO',
    'MATERIAL_ELETRICO', 
    'ALUGUEL_EQUIPAMENTO',
    'TRANSPORTE',
    'OUTROS'
));

-- 7. Recriar índice original
CREATE INDEX IF NOT EXISTS idx_despesa_data ON despesa(data_despesa);
