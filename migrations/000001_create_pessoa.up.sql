CREATE TABLE pessoa (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    tipo VARCHAR(2) NOT NULL CHECK (tipo IN ('PF', 'PJ')),
    documento VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(255),
    telefone VARCHAR(20),
    cargo VARCHAR(100),
    ativo BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() 
);


CREATE INDEX idx_pessoa_tipo ON pessoa(tipo);
CREATE INDEX idx_pessoa_documento ON pessoa(documento);
CREATE INDEX idx_pessoa_email ON pessoa(email) WHERE email IS NOT NULL; 
CREATE INDEX idx_pessoa_ativo ON pessoa(ativo);  
CREATE INDEX idx_pessoa_created_at ON pessoa(created_at);
