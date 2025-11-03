-- 000018_rename_data_despesa_to_data_vencimento.down.sql
-- Reverte a renomeação de data_vencimento para data_despesa se aplicável
DO $$
BEGIN
  IF EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_name='despesa' AND column_name='data_vencimento'
  ) AND NOT EXISTS (
    SELECT 1 FROM information_schema.columns
    WHERE table_name='despesa' AND column_name='data_despesa'
  ) THEN
    ALTER TABLE despesa RENAME COLUMN data_vencimento TO data_despesa;
  END IF;
END$$;
