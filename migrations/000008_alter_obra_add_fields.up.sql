-- Adicionar campos financeiros e de controle na tabela obra
ALTER TABLE obra 
    ADD COLUMN IF NOT EXISTS art VARCHAR(50),
    ADD COLUMN IF NOT EXISTS cliente_id INTEGER REFERENCES pessoa(id) ON DELETE SET NULL,
    ADD COLUMN IF NOT EXISTS responsavel_id INTEGER REFERENCES pessoa(id) ON DELETE SET NULL,
    ADD COLUMN IF NOT EXISTS parceiro_id INTEGER REFERENCES pessoa(id) ON DELETE SET NULL,
    ADD COLUMN IF NOT EXISTS data_inicio DATE,
    ADD COLUMN IF NOT EXISTS data_previsao_termino DATE,
    ADD COLUMN IF NOT EXISTS data_termino_real DATE,
    ADD COLUMN IF NOT EXISTS orcamento_inicial DECIMAL(15, 2) DEFAULT 0.00,
    ADD COLUMN IF NOT EXISTS valor_total DECIMAL(15, 2) DEFAULT 0.00,
    ADD COLUMN IF NOT EXISTS valor_aditivo DECIMAL(15, 2) DEFAULT 0.00,
    ADD COLUMN IF NOT EXISTS custo_mao_obra DECIMAL(15, 2) DEFAULT 0.00,
    ADD COLUMN IF NOT EXISTS despesas_gerais DECIMAL(15, 2) DEFAULT 0.00,
    ADD COLUMN IF NOT EXISTS lucro DECIMAL(15, 2) DEFAULT 0.00,
    ADD COLUMN IF NOT EXISTS status VARCHAR(30) DEFAULT 'PLANEJADA' CHECK (status IN (
        'PLANEJADA',
        'EM_ANDAMENTO',
        'PAUSADA',
        'CONCLUIDA',
        'CANCELADA'
    )),
    ADD COLUMN IF NOT EXISTS endereco_rua VARCHAR(255),
    ADD COLUMN IF NOT EXISTS endereco_numero VARCHAR(20),
    ADD COLUMN IF NOT EXISTS endereco_complemento VARCHAR(100),
    ADD COLUMN IF NOT EXISTS endereco_bairro VARCHAR(100),
    ADD COLUMN IF NOT EXISTS endereco_cidade VARCHAR(100),
    ADD COLUMN IF NOT EXISTS endereco_estado VARCHAR(2),
    ADD COLUMN IF NOT EXISTS endereco_cep VARCHAR(10);

CREATE INDEX IF NOT EXISTS idx_obra_cliente ON obra(cliente_id);
CREATE INDEX IF NOT EXISTS idx_obra_responsavel ON obra(responsavel_id);
CREATE INDEX IF NOT EXISTS idx_obra_status ON obra(status);
CREATE INDEX IF NOT EXISTS idx_obra_data_inicio ON obra(data_inicio);
