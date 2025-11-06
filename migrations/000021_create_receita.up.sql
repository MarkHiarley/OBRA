CREATE TABLE IF NOT EXISTS receitas (
    id SERIAL PRIMARY KEY,
    obra_id INTEGER REFERENCES obra(id) ON DELETE CASCADE,
    descricao TEXT NOT NULL,
    valor NUMERIC(15,2) NOT NULL,
    data TIMESTAMP NOT NULL,
    fonte_receita VARCHAR(50) DEFAULT 'OUTROS',
    numero_documento VARCHAR(100),
    responsavel_id INTEGER REFERENCES pessoa(id),
    observacao TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);