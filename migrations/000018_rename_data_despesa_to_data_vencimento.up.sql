-- 000018_rename_data_despesa_to_data_vencimento.up.sql
-- Renomeia a coluna data_despesa para data_vencimento de forma idempotente
DO $$
BEGIN
  IF EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_name='despesa' AND column_name='data_despesa'
  ) THEN
    ALTER TABLE despesa RENAME COLUMN data_despesa TO data_vencimento;
  END IF;
END$$;
