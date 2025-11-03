-- 000020_add_art_obra.up.sql
-- Adiciona campo art em obra (idempotente)
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='obra' AND column_name='art'
  ) THEN
    ALTER TABLE obra ADD COLUMN art varchar;
  END IF;
END$$;
