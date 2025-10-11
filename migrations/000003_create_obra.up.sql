CREATE TABLE obra (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(500) NOT NULL,
    contrato_numero VARCHAR(100) UNIQUE NOT NULL,
    contratante_id INT REFERENCES pessoa(id) ON DELETE SET NULL,
    responsavel_id INT REFERENCES pessoa(id) ON DELETE SET NULL,
    data_inicio DATE NOT NULL,
    prazo_dias INT NOT NULL CHECK (prazo_dias > 0),
    data_fim_prevista DATE,
    orcamento DECIMAL(15,2),
    status VARCHAR(20) DEFAULT 'planejamento' CHECK (
        status IN ('planejamento', 'em_andamento', 'pausada', 'concluida', 'cancelada')
    ),
    endereco_rua VARCHAR(200),
    endereco_numero VARCHAR(20),
    endereco_bairro VARCHAR(100),
    endereco_cidade VARCHAR(100),
    endereco_estado CHAR(2),
    endereco_cep VARCHAR(10),
    observacoes TEXT,
    ativo BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_obra_status ON obra(status);
CREATE INDEX idx_obra_contratante ON obra(contratante_id);
CREATE INDEX idx_obra_responsavel ON obra(responsavel_id);
CREATE INDEX idx_obra_ativo ON obra(ativo);