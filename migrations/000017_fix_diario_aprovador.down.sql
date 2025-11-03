-- 000017_fix_diario_aprovador.down.sql
-- Reverte as alterações feitas por 000017_fix_diario_aprovador.up.sql
DO $$
BEGIN
  -- Remove constraint se existir
  IF EXISTS (
    SELECT 1 FROM pg_constraint c
    JOIN pg_class t ON c.conrelid = t.oid
    WHERE t.relname = 'diario_obra' AND c.conname = 'ck_diario_aprovador_status'
  ) THEN
    ALTER TABLE diario_obra DROP CONSTRAINT ck_diario_aprovador_status;
  END IF;

  -- Nota: não remove a coluna aprovado_por_id para evitar perda de dados; se quiser remover, faça manualmente
END$$;
