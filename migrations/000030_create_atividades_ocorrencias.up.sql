-- ========================================
-- NOVA ARQUITETURA: Diário de Obras Refatorado
-- ========================================
-- Modelo normalizado onde atividades e ocorrências são entidades independentes
-- O diário de obras é gerado dinamicamente a partir dessas entidades

-- Criar tabela para atividades diárias (normalizadas)
CREATE TABLE IF NOT EXISTS atividade_diaria (
    id SERIAL PRIMARY KEY,
    obra_id INT NOT NULL REFERENCES obra(id) ON DELETE CASCADE,
    data DATE NOT NULL,
    periodo VARCHAR(10) DEFAULT 'integral' CHECK (periodo IN ('manha', 'tarde', 'integral', 'noite')),
    descricao TEXT NOT NULL,
    responsavel_id INT REFERENCES pessoa(id) ON DELETE SET NULL,
    status VARCHAR(20) DEFAULT 'em_andamento' CHECK (status IN ('planejada', 'em_andamento', 'concluida', 'cancelada')),
    percentual_conclusao INT DEFAULT 0 CHECK (percentual_conclusao >= 0 AND percentual_conclusao <= 100),
    observacao TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Criar tabela para ocorrências diárias (normalizadas)
CREATE TABLE IF NOT EXISTS ocorrencia_diaria (
    id SERIAL PRIMARY KEY,
    obra_id INT NOT NULL REFERENCES obra(id) ON DELETE CASCADE,
    data DATE NOT NULL,
    periodo VARCHAR(10) DEFAULT 'integral' CHECK (periodo IN ('manha', 'tarde', 'integral', 'noite')),
    tipo VARCHAR(30) DEFAULT 'geral' CHECK (tipo IN ('seguranca', 'qualidade', 'prazo', 'custo', 'clima', 'equipamento', 'material', 'geral')),
    gravidade VARCHAR(20) DEFAULT 'baixa' CHECK (gravidade IN ('baixa', 'media', 'alta', 'critica')),
    descricao TEXT NOT NULL,
    responsavel_id INT REFERENCES pessoa(id) ON DELETE SET NULL,
    status_resolucao VARCHAR(20) DEFAULT 'pendente' CHECK (status_resolucao IN ('pendente', 'em_analise', 'resolvida', 'nao_aplicavel')),
    acao_tomada TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Tabela para metadados do diário (foto, aprovação, etc)
CREATE TABLE IF NOT EXISTS diario_metadados (
    id SERIAL PRIMARY KEY,
    obra_id INT NOT NULL REFERENCES obra(id) ON DELETE CASCADE,
    data DATE NOT NULL,
    periodo VARCHAR(10) DEFAULT 'integral' CHECK (periodo IN ('manha', 'tarde', 'integral', 'noite')),
    foto TEXT, -- Base64 encoded image
    observacoes TEXT,
    responsavel_id INT REFERENCES pessoa(id) ON DELETE SET NULL,
    aprovado_por_id INT REFERENCES pessoa(id) ON DELETE SET NULL,
    status_aprovacao VARCHAR(20) DEFAULT 'pendente' CHECK (status_aprovacao IN ('pendente', 'aprovado', 'rejeitado')),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(obra_id, data, periodo)
);

-- Índices para performance
CREATE INDEX idx_atividade_obra_data ON atividade_diaria(obra_id, data);
CREATE INDEX idx_atividade_periodo ON atividade_diaria(periodo);
CREATE INDEX idx_atividade_responsavel ON atividade_diaria(responsavel_id);
CREATE INDEX idx_atividade_status ON atividade_diaria(status);

CREATE INDEX idx_ocorrencia_obra_data ON ocorrencia_diaria(obra_id, data);
CREATE INDEX idx_ocorrencia_periodo ON ocorrencia_diaria(periodo);
CREATE INDEX idx_ocorrencia_tipo ON ocorrencia_diaria(tipo);
CREATE INDEX idx_ocorrencia_gravidade ON ocorrencia_diaria(gravidade);
CREATE INDEX idx_ocorrencia_responsavel ON ocorrencia_diaria(responsavel_id);
CREATE INDEX idx_ocorrencia_status ON ocorrencia_diaria(status_resolucao);

CREATE INDEX idx_diario_meta_obra_data ON diario_metadados(obra_id, data);
CREATE INDEX idx_diario_meta_aprovacao ON diario_metadados(status_aprovacao);

-- View para gerar diário consolidado dinamicamente
-- Esta view agrega atividades, ocorrências, equipe, equipamentos e materiais
CREATE OR REPLACE VIEW vw_diario_consolidado AS
SELECT 
    dm.id as diario_id,
    dm.obra_id,
    o.nome as obra_nome,
    dm.data,
    dm.periodo,
    -- Agregação de atividades (sem ORDER BY no DISTINCT)
    STRING_AGG(DISTINCT a.descricao || ' (' || a.status || ' - ' || a.percentual_conclusao || '%)', '; ') as atividades,
    -- Agregação de ocorrências (sem ORDER BY no DISTINCT)
    STRING_AGG(DISTINCT '[' || UPPER(oc.gravidade) || '] ' || oc.descricao || ' - ' || oc.status_resolucao, '; ') as ocorrencias,
    dm.foto,
    dm.observacoes,
    dm.responsavel_id,
    pr.nome as responsavel_nome,
    dm.aprovado_por_id,
    pa.nome as aprovado_por_nome,
    dm.status_aprovacao,
    -- Contadores
    COUNT(DISTINCT a.id) as qtd_atividades,
    COUNT(DISTINCT oc.id) as qtd_ocorrencias,
    COUNT(DISTINCT e.id) as qtd_equipe,
    COUNT(DISTINCT eq.id) as qtd_equipamentos,
    COUNT(DISTINCT m.id) as qtd_materiais,
    dm.created_at,
    dm.updated_at
FROM diario_metadados dm
INNER JOIN obra o ON dm.obra_id = o.id
LEFT JOIN pessoa pr ON dm.responsavel_id = pr.id
LEFT JOIN pessoa pa ON dm.aprovado_por_id = pa.id
LEFT JOIN atividade_diaria a ON a.obra_id = dm.obra_id AND a.data = dm.data AND a.periodo = dm.periodo
LEFT JOIN ocorrencia_diaria oc ON oc.obra_id = dm.obra_id AND oc.data = dm.data AND oc.periodo = dm.periodo
LEFT JOIN equipe_diario e ON e.diario_id = dm.id
LEFT JOIN equipamento_diario eq ON eq.diario_id = dm.id
LEFT JOIN material_diario m ON m.diario_id = dm.id
GROUP BY 
    dm.id, dm.obra_id, o.nome, dm.data, dm.periodo, 
    dm.observacoes, dm.foto, dm.responsavel_id, pr.nome, 
    dm.aprovado_por_id, pa.nome, dm.status_aprovacao, 
    dm.created_at, dm.updated_at;

-- Comentários para documentação
COMMENT ON TABLE atividade_diaria IS 'Armazena atividades individuais realizadas diariamente na obra';
COMMENT ON TABLE ocorrencia_diaria IS 'Armazena ocorrências/problemas individuais registrados diariamente';
COMMENT ON TABLE diario_metadados IS 'Armazena metadados do diário (foto, observações, aprovação)';
COMMENT ON VIEW vw_diario_consolidado IS 'View que gera diário de obras consolidado agregando atividades, ocorrências e recursos';

COMMENT ON COLUMN atividade_diaria.percentual_conclusao IS 'Percentual de conclusão da atividade (0-100)';
COMMENT ON COLUMN ocorrencia_diaria.gravidade IS 'Nível de gravidade: baixa, media, alta, critica';
COMMENT ON COLUMN ocorrencia_diaria.tipo IS 'Tipo de ocorrência: seguranca, qualidade, prazo, custo, clima, equipamento, material, geral';
