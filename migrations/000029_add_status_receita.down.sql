-- Remove status column from receita table
-- Remove status column from receitas table
ALTER TABLE receitas
DROP COLUMN IF EXISTS status;
