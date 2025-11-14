-- Tabela para armazenar múltiplas fotos relacionadas a diários, atividades ou ocorrências
CREATE TABLE IF NOT EXISTS foto_diario (
    id SERIAL PRIMARY KEY,
    
    -- Relacionamento polimórfico (pode estar ligada a diferentes entidades)
    entidade_tipo VARCHAR(50) NOT NULL CHECK (entidade_tipo IN ('metadados', 'atividade', 'ocorrencia')),
    entidade_id INTEGER NOT NULL,
    
    -- Dados da foto
    foto TEXT NOT NULL, -- Base64 encoded image
    descricao TEXT,
    ordem INTEGER DEFAULT 0, -- Para ordenar as fotos
    categoria VARCHAR(50) DEFAULT 'DIARIO', -- DIARIO, OBRA, OCORRENCIA, ATIVIDADE, SEGURANCA, etc
    
    -- Metadados
    largura INTEGER,
    altura INTEGER,
    tamanho_bytes INTEGER,
    
    -- Auditoria
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Índices para melhorar performance nas consultas
CREATE INDEX idx_foto_diario_entidade ON foto_diario(entidade_tipo, entidade_id);
CREATE INDEX idx_foto_diario_categoria ON foto_diario(categoria);
CREATE INDEX idx_foto_diario_ordem ON foto_diario(ordem);

-- Comentários
COMMENT ON TABLE foto_diario IS 'Armazena múltiplas fotos em Base64 relacionadas a metadados, atividades ou ocorrências do diário';
COMMENT ON COLUMN foto_diario.entidade_tipo IS 'Tipo da entidade: metadados, atividade, ocorrencia';
COMMENT ON COLUMN foto_diario.entidade_id IS 'ID da entidade relacionada (diario_metadados.id, atividade_diaria.id, ou ocorrencia_diaria.id)';
COMMENT ON COLUMN foto_diario.foto IS 'Imagem em Base64 (data:image/jpeg;base64,...)';
COMMENT ON COLUMN foto_diario.ordem IS 'Ordem de exibição das fotos (0 = primeira)';

-- Trigger para atualizar updated_at
CREATE OR REPLACE FUNCTION update_foto_diario_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_foto_diario_updated_at
    BEFORE UPDATE ON foto_diario
    FOR EACH ROW
    EXECUTE FUNCTION update_foto_diario_updated_at();
