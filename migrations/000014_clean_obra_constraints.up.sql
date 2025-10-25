-- Limpar todas as foreign keys duplicadas na tabela obra para deploy limpo

-- Remover FKs duplicadas de cliente_id
ALTER TABLE obra DROP CONSTRAINT IF EXISTS obra_cliente_id_fkey1;
ALTER TABLE obra DROP CONSTRAINT IF EXISTS obra_cliente_id_fkey2;

-- Remover FKs duplicadas de parceiro_id  
ALTER TABLE obra DROP CONSTRAINT IF EXISTS obra_parceiro_id_fkey1;
ALTER TABLE obra DROP CONSTRAINT IF EXISTS obra_parceiro_id_fkey2;

-- Garantir que apenas as FKs principais existam
DO $$ 
BEGIN
    -- Verificar se FK principal de cliente existe
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'obra_cliente_id_fkey') THEN
        ALTER TABLE obra ADD CONSTRAINT obra_cliente_id_fkey 
        FOREIGN KEY (cliente_id) REFERENCES pessoa(id) ON DELETE SET NULL;
    END IF;
    
    -- Verificar se FK principal de parceiro existe
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'obra_parceiro_id_fkey') THEN
        ALTER TABLE obra ADD CONSTRAINT obra_parceiro_id_fkey 
        FOREIGN KEY (parceiro_id) REFERENCES pessoa(id) ON DELETE SET NULL;
    END IF;
END $$;