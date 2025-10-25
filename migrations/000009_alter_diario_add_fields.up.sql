-- Adicionar campos detalhados na tabela diario_obra
ALTER TABLE diario_obra
    ADD COLUMN IF NOT EXISTS ferramentas_utilizadas TEXT,
    ADD COLUMN IF NOT EXISTS quantidade_pessoas INTEGER DEFAULT 0,
    ADD COLUMN IF NOT EXISTS responsavel_execucao VARCHAR(255),
    ADD COLUMN IF NOT EXISTS clima VARCHAR(50) CHECK (clima IN (
        'ENSOLARADO',
        'NUBLADO',
        'CHUVOSO',
        'VENTOSO',
        'OUTROS'
    )),
    ADD COLUMN IF NOT EXISTS progresso_percentual DECIMAL(5, 2) DEFAULT 0.00 CHECK (progresso_percentual >= 0 AND progresso_percentual <= 100),
    ADD COLUMN IF NOT EXISTS problemas_encontrados TEXT,
    ADD COLUMN IF NOT EXISTS fotos_anexadas TEXT; -- JSON array de URLs

CREATE INDEX IF NOT EXISTS idx_diario_data ON diario_obra(data);
CREATE INDEX IF NOT EXISTS idx_diario_responsavel ON diario_obra(responsavel_execucao);
