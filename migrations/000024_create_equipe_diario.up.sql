-- Tabela para registrar equipe utilizada em cada diário
CREATE TABLE IF NOT EXISTS equipe_diario (
    id SERIAL PRIMARY KEY,
    diario_id INTEGER NOT NULL REFERENCES diario_obra(id) ON DELETE CASCADE,
    codigo VARCHAR(50),
    descricao VARCHAR(255) NOT NULL,
    quantidade_utilizada INTEGER NOT NULL DEFAULT 1,
    horas_trabalhadas DECIMAL(5,2),
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_equipe_diario FOREIGN KEY (diario_id) REFERENCES diario_obra(id)
);

CREATE INDEX idx_equipe_diario_id ON equipe_diario(diario_id);
