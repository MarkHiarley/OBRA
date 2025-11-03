-- 000020_add_art_obra.down.sql
-- Reverte a adição do campo art na tabela obra
DO $$
BEGIN
  IF EXISTS (
    SELECT 1 FROM information_schema.columns WHERE table_name='obra' AND column_name='art'
  ) THEN
    ALTER TABLE obra DROP COLUMN IF EXISTS art;
  END IF;
END$$;
