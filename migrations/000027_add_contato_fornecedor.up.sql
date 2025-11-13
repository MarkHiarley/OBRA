-- Add contact fields to fornecedor table
ALTER TABLE fornecedor 
ADD COLUMN IF NOT EXISTS contato_nome VARCHAR(255),
ADD COLUMN IF NOT EXISTS contato_telefone VARCHAR(20),
ADD COLUMN IF NOT EXISTS contato_email VARCHAR(255);
