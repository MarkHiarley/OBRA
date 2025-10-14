CREATE TABLE diario_obra (
    id SERIAL PRIMARY KEY,
    obra_id INT NOT NULL REFERENCES obra(id) ON DELETE CASCADE,
    data DATE NOT NULL,
    periodo VARCHAR(10) DEFAULT 'integral' CHECK (periodo IN ('manha', 'tarde', 'integral', 'noite')),
    atividades_realizadas TEXT NOT NULL,
    ocorrencias TEXT DEFAULT 'Não houve ocorrências',
    observacoes TEXT,
    responsavel_id INT REFERENCES pessoa(id) ON DELETE SET NULL,
    aprovado_por_id INT REFERENCES pessoa(id) ON DELETE SET NULL,
    status_aprovacao VARCHAR(20) DEFAULT 'pendente' CHECK (
        status_aprovacao IN ('pendente', 'aprovado', 'rejeitado')
    ),
    created_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP,
    UNIQUE(obra_id, data, periodo)
);

CREATE INDEX idx_diario_obra_id ON diario_obra(obra_id);
CREATE INDEX idx_diario_data ON diario_obra(data);
CREATE INDEX idx_diario_status ON diario_obra(status_aprovacao);
CREATE INDEX idx_diario_responsavel ON diario_obra(responsavel_id);

