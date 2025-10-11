    
CREATE TABLE usuario (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    senha_hash VARCHAR(255) NOT NULL,
    tipo_documento VARCHAR(2) CHECK (tipo_documento IN ('PF', 'PJ')),
    documento VARCHAR(20) UNIQUE,
    telefone VARCHAR(20),
    perfil_acesso VARCHAR(50)[] DEFAULT ARRAY['usuario'],
    ativo BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE INDEX idx_usuario_email ON usuario(email);
CREATE INDEX idx_usuario_ativo ON usuario(ativo);
CREATE INDEX idx_usuario_perfil_acesso ON usuario USING GIN(perfil_acesso);