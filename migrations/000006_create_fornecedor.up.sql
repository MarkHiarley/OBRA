-- Tabela de Fornecedores
CREATE TABLE IF NOT EXISTS fornecedor (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    tipo_documento VARCHAR(10) CHECK (tipo_documento IN ('CPF', 'CNPJ')),
    documento VARCHAR(18) UNIQUE,
    email VARCHAR(255),
    telefone VARCHAR(20),
    endereco TEXT,
    cidade VARCHAR(100),
    estado VARCHAR(2),
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_fornecedor_documento ON fornecedor(documento);
CREATE INDEX idx_fornecedor_ativo ON fornecedor(ativo);
