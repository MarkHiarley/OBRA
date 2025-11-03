-- 000017_fix_diario_aprovador.up.sql
-- Garante que aprovado_por_id exista como NULLABLE e adiciona um CHECK que relaciona status e aprovado_por_id
DO $$
BEGIN
  -- 1) Adiciona coluna se não existir
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_name='diario_obra' AND column_name='aprovado_por_id'
  ) THEN
    ALTER TABLE diario_obra ADD COLUMN aprovado_por_id bigint;
  END IF;

  -- 2) Permitir NULL caso esteja NOT NULL
  BEGIN
    ALTER TABLE diario_obra ALTER COLUMN aprovado_por_id DROP NOT NULL;
  EXCEPTION WHEN undefined_column THEN
    -- coluna não existe: já tratada acima
    NULL;
  END;

  -- 3) Adiciona constraint de consistência entre status e aprovado_por_id se não existir
  IF NOT EXISTS (
    SELECT 1 FROM pg_constraint c
    JOIN pg_class t ON c.conrelid = t.oid
    WHERE t.relname = 'diario_obra' AND c.conname = 'ck_diario_aprovador_status'
  ) THEN
    ALTER TABLE diario_obra
      ADD CONSTRAINT ck_diario_aprovador_status CHECK (
        (status = 'APROVADO' AND aprovado_por_id IS NOT NULL)
        OR (status = 'PENDENTE' AND aprovado_por_id IS NULL)
        OR (status NOT IN ('APROVADO','PENDENTE'))
      );
  END IF;
END$$;
