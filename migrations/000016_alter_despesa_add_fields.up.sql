-- Renomear data_despesa para data e adicionar data_vencimento
-- Adicionar novas categorias IMPOSTO e PARCEIRO

-- 1. Adicionar nova coluna data (temporária)
ALTER TABLE despesa ADD COLUMN IF NOT EXISTS data DATE;

-- 2. Copiar dados de data_despesa para data
UPDATE despesa SET data = data_despesa WHERE data IS NULL;

-- 3. Adicionar coluna data_vencimento
ALTER TABLE despesa ADD COLUMN IF NOT EXISTS data_vencimento DATE;

-- 4. Tornar data NOT NULL após copiar os dados
ALTER TABLE despesa ALTER COLUMN data SET NOT NULL;

-- 5. Remover a coluna antiga data_despesa
ALTER TABLE despesa DROP COLUMN IF EXISTS data_despesa;

-- 6. Atualizar constraint de categoria para incluir IMPOSTO e PARCEIRO
ALTER TABLE despesa DROP CONSTRAINT IF EXISTS despesa_categoria_check;
ALTER TABLE despesa ADD CONSTRAINT despesa_categoria_check CHECK (categoria IN (
    'MATERIAL', 
    'MAO_DE_OBRA', 
    'COMBUSTIVEL', 
    'ALIMENTACAO',
    'MATERIAL_ELETRICO', 
    'ALUGUEL_EQUIPAMENTO',
    'TRANSPORTE',
    'IMPOSTO',
    'PARCEIRO',
    'OUTROS'
));

-- 7. Recriar índice com o nome correto da coluna
DROP INDEX IF EXISTS idx_despesa_data;
CREATE INDEX IF NOT EXISTS idx_despesa_data ON despesa(data);
CREATE INDEX IF NOT EXISTS idx_despesa_data_vencimento ON despesa(data_vencimento);
