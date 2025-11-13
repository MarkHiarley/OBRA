-- Tabela para registrar materiais utilizados em cada diário
CREATE TABLE IF NOT EXISTS material_diario (
    id SERIAL PRIMARY KEY,
    diario_id INTEGER NOT NULL REFERENCES diario_obra(id) ON DELETE CASCADE,
    codigo VARCHAR(50),
    descricao VARCHAR(255) NOT NULL,
    quantidade DECIMAL(10,2) NOT NULL,
    unidade VARCHAR(20) NOT NULL,
    fornecedor VARCHAR(255),
    valor_unitario DECIMAL(10,2),
    valor_total DECIMAL(10,2),
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_material_diario FOREIGN KEY (diario_id) REFERENCES diario_obra(id)
);

CREATE INDEX idx_material_diario_id ON material_diario(diario_id);
