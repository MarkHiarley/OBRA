-- Tabela para registrar equipamentos utilizados em cada diário
CREATE TABLE IF NOT EXISTS equipamento_diario (
    id SERIAL PRIMARY KEY,
    diario_id INTEGER NOT NULL REFERENCES diario_obra(id) ON DELETE CASCADE,
    codigo VARCHAR(50),
    descricao VARCHAR(255) NOT NULL,
    quantidade_utilizada INTEGER NOT NULL DEFAULT 1,
    horas_uso DECIMAL(5,2),
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_equipamento_diario FOREIGN KEY (diario_id) REFERENCES diario_obra(id)
);

CREATE INDEX idx_equipamento_diario_id ON equipamento_diario(diario_id);
