-- Add status column to receitas table
ALTER TABLE receitas
ADD COLUMN IF NOT EXISTS status VARCHAR(50) DEFAULT 'a_receber';
