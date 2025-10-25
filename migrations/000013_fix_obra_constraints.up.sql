-- Corrigir constraints duplicados e conflitantes na tabela obra

-- Remove duplicate constraints from obra table

-- Drop old status constraint (minúsculo)
ALTER TABLE obra DROP CONSTRAINT IF EXISTS obra_status_check;

-- Drop duplicate FK constraints
ALTER TABLE obra DROP CONSTRAINT IF EXISTS obra_responsavel_id_fkey1;
ALTER TABLE obra DROP CONSTRAINT IF EXISTS obra_responsavel_id_fkey2;

-- Ensure the correct constraint exists
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint 
        WHERE conname = 'obra_status_check1' 
        AND conrelid = 'obra'::regclass
    ) THEN
        ALTER TABLE obra ADD CONSTRAINT obra_status_check1 
        CHECK (status IN ('EM_ANDAMENTO', 'CONCLUIDA', 'PARALISADA', 'CANCELADA', 'PLANEJAMENTO'));
    END IF;
END $$;
