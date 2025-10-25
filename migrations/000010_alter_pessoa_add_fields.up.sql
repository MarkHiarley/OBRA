-- Adicionar campos para empresas e endereço completo em pessoa
ALTER TABLE pessoa
    ADD COLUMN IF NOT EXISTS nome_fantasia VARCHAR(255),
    ADD COLUMN IF NOT EXISTS eh_pessoa_juridica BOOLEAN DEFAULT false,
    ADD COLUMN IF NOT EXISTS funcao VARCHAR(100),
    ADD COLUMN IF NOT EXISTS endereco_rua VARCHAR(255),
    ADD COLUMN IF NOT EXISTS endereco_numero VARCHAR(20),
    ADD COLUMN IF NOT EXISTS endereco_complemento VARCHAR(100),
    ADD COLUMN IF NOT EXISTS endereco_bairro VARCHAR(100),
    ADD COLUMN IF NOT EXISTS endereco_cidade VARCHAR(100),
    ADD COLUMN IF NOT EXISTS endereco_estado VARCHAR(2),
    ADD COLUMN IF NOT EXISTS endereco_cep VARCHAR(10),
    ADD COLUMN IF NOT EXISTS observacao TEXT;

CREATE INDEX IF NOT EXISTS idx_pessoa_juridica ON pessoa(eh_pessoa_juridica);
CREATE INDEX IF NOT EXISTS idx_pessoa_funcao ON pessoa(funcao);
