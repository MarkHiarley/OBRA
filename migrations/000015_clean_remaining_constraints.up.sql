-- Remove remaining duplicate foreign key constraints
ALTER TABLE obra DROP CONSTRAINT IF EXISTS obra_cliente_id_fkey3;
ALTER TABLE obra DROP CONSTRAINT IF EXISTS obra_parceiro_id_fkey3;

-- Ensure we only have the correct constraints
DO $$
BEGIN
    -- Ensure obra_cliente_id_fkey exists and is correct
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'obra_cliente_id_fkey') THEN
        ALTER TABLE obra ADD CONSTRAINT obra_cliente_id_fkey 
            FOREIGN KEY (cliente_id) REFERENCES pessoa(id) ON DELETE SET NULL;
    END IF;
    
    -- Ensure obra_parceiro_id_fkey exists and is correct  
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'obra_parceiro_id_fkey') THEN
        ALTER TABLE obra ADD CONSTRAINT obra_parceiro_id_fkey 
            FOREIGN KEY (parceiro_id) REFERENCES pessoa(id) ON DELETE SET NULL;
    END IF;
END
$$;