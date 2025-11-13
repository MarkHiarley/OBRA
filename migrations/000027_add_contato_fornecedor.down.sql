-- Remove contact fields from fornecedor table
ALTER TABLE fornecedor 
DROP COLUMN IF EXISTS contato_nome,
DROP COLUMN IF EXISTS contato_telefone,
DROP COLUMN IF EXISTS contato_email;
