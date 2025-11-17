-- Migração para refatorar equipe_diario, material_diario e equipamento_diario
-- Remover dependência de diario_obra e relacionar diretamente com obra

-- PRIMEIRO: Dropar a view que depende das colunas diario_id
DROP VIEW IF EXISTS vw_diario_consolidado CASCADE;

-- 1. EQUIPE_DIARIO: Adicionar colunas obra_id e data
ALTER TABLE equipe_diario 
ADD COLUMN IF NOT EXISTS obra_id INTEGER,
ADD COLUMN IF NOT EXISTS data DATE;

UPDATE equipe_diario e
SET obra_id = d.obra_id, data = d.data
FROM diario_obra d
WHERE e.diario_id = d.id AND e.obra_id IS NULL;

-- Tornar as novas colunas NOT NULL após popular
DO $$ 
BEGIN
  IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='equipe_diario' AND column_name='obra_id' AND is_nullable='YES') THEN
    ALTER TABLE equipe_diario ALTER COLUMN obra_id SET NOT NULL;
  END IF;
  IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='equipe_diario' AND column_name='data' AND is_nullable='YES') THEN
    ALTER TABLE equipe_diario ALTER COLUMN data SET NOT NULL;
  END IF;
END $$;

-- Adicionar FK para obra
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_equipe_obra') THEN
    ALTER TABLE equipe_diario
    ADD CONSTRAINT fk_equipe_obra FOREIGN KEY (obra_id) REFERENCES obra(id) ON DELETE CASCADE;
  END IF;
END $$;

-- Remover a FK antiga e o campo diario_id
ALTER TABLE equipe_diario DROP CONSTRAINT IF EXISTS fk_equipe_diario CASCADE;
ALTER TABLE equipe_diario DROP COLUMN IF EXISTS diario_id CASCADE;

-- Criar índices
CREATE INDEX IF NOT EXISTS idx_equipe_obra_id ON equipe_diario(obra_id);
CREATE INDEX IF NOT EXISTS idx_equipe_data ON equipe_diario(data);
CREATE INDEX IF NOT EXISTS idx_equipe_obra_data ON equipe_diario(obra_id, data);

-- 2. MATERIAL_DIARIO: Adicionar colunas obra_id e data
ALTER TABLE material_diario 
ADD COLUMN IF NOT EXISTS obra_id INTEGER,
ADD COLUMN IF NOT EXISTS data DATE;

-- Popular os novos campos
UPDATE material_diario m
SET obra_id = d.obra_id, data = d.data
FROM diario_obra d
WHERE m.diario_id = d.id AND m.obra_id IS NULL;

-- Tornar as novas colunas NOT NULL
DO $$ 
BEGIN
  IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='material_diario' AND column_name='obra_id' AND is_nullable='YES') THEN
    ALTER TABLE material_diario ALTER COLUMN obra_id SET NOT NULL;
  END IF;
  IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='material_diario' AND column_name='data' AND is_nullable='YES') THEN
    ALTER TABLE material_diario ALTER COLUMN data SET NOT NULL;
  END IF;
END $$;

-- Adicionar FK para obra
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_material_obra') THEN
    ALTER TABLE material_diario
    ADD CONSTRAINT fk_material_obra FOREIGN KEY (obra_id) REFERENCES obra(id) ON DELETE CASCADE;
  END IF;
END $$;

-- Remover a FK antiga e o campo diario_id
ALTER TABLE material_diario DROP CONSTRAINT IF EXISTS fk_material_diario CASCADE;
ALTER TABLE material_diario DROP COLUMN IF EXISTS diario_id CASCADE;

-- Criar índices
CREATE INDEX IF NOT EXISTS idx_material_obra_id ON material_diario(obra_id);
CREATE INDEX IF NOT EXISTS idx_material_data ON material_diario(data);
CREATE INDEX IF NOT EXISTS idx_material_obra_data ON material_diario(obra_id, data);

-- 3. EQUIPAMENTO_DIARIO: Adicionar colunas obra_id e data
ALTER TABLE equipamento_diario 
ADD COLUMN IF NOT EXISTS obra_id INTEGER,
ADD COLUMN IF NOT EXISTS data DATE;

-- Popular os novos campos
UPDATE equipamento_diario e
SET obra_id = d.obra_id, data = d.data
FROM diario_obra d
WHERE e.diario_id = d.id AND e.obra_id IS NULL;

-- Tornar as novas colunas NOT NULL
DO $$ 
BEGIN
  IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='equipamento_diario' AND column_name='obra_id' AND is_nullable='YES') THEN
    ALTER TABLE equipamento_diario ALTER COLUMN obra_id SET NOT NULL;
  END IF;
  IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='equipamento_diario' AND column_name='data' AND is_nullable='YES') THEN
    ALTER TABLE equipamento_diario ALTER COLUMN data SET NOT NULL;
  END IF;
END $$;

-- Adicionar FK para obra
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_equipamento_obra') THEN
    ALTER TABLE equipamento_diario
    ADD CONSTRAINT fk_equipamento_obra FOREIGN KEY (obra_id) REFERENCES obra(id) ON DELETE CASCADE;
  END IF;
END $$;

-- Remover a FK antiga e o campo diario_id
ALTER TABLE equipamento_diario DROP CONSTRAINT IF EXISTS fk_equipamento_diario CASCADE;
ALTER TABLE equipamento_diario DROP COLUMN IF EXISTS diario_id CASCADE;

-- Criar índices
CREATE INDEX IF NOT EXISTS idx_equipamento_obra_id ON equipamento_diario(obra_id);
CREATE INDEX IF NOT EXISTS idx_equipamento_data ON equipamento_diario(data);
CREATE INDEX IF NOT EXISTS idx_equipamento_obra_data ON equipamento_diario(obra_id, data);

-- RECRIAR a view vw_diario_consolidado sem depender de diario_id
CREATE OR REPLACE VIEW vw_diario_consolidado AS
SELECT 
    NULL::INTEGER AS diario_id,  -- Campo deprecated, sempre NULL
    dm.obra_id,
    o.nome AS obra_nome,
    dm.data,
    dm.periodo,
    STRING_AGG(DISTINCT a.descricao, '; ') AS atividades,
    STRING_AGG(DISTINCT oc.descricao, '; ') AS ocorrencias,
    dm.foto,
    dm.observacoes,
    dm.responsavel_id,
    p1.nome AS responsavel_nome,
    dm.aprovado_por_id,
    p2.nome AS aprovado_por_nome,
    dm.status_aprovacao,
    COUNT(DISTINCT a.id) AS qtd_atividades,
    COUNT(DISTINCT oc.id) AS qtd_ocorrencias,
    COUNT(DISTINCT eq.id) AS qtd_equipe,
    COUNT(DISTINCT equip.id) AS qtd_equipamentos,
    COUNT(DISTINCT m.id) AS qtd_materiais,
    dm.created_at,
    dm.updated_at
FROM diario_metadados dm
LEFT JOIN obra o ON o.id = dm.obra_id
LEFT JOIN pessoa p1 ON p1.id = dm.responsavel_id
LEFT JOIN pessoa p2 ON p2.id = dm.aprovado_por_id
LEFT JOIN atividade_diaria a ON a.obra_id = dm.obra_id AND a.data = dm.data
LEFT JOIN ocorrencia_diaria oc ON oc.obra_id = dm.obra_id AND oc.data = dm.data
LEFT JOIN equipe_diario eq ON eq.obra_id = dm.obra_id AND eq.data = dm.data
LEFT JOIN equipamento_diario equip ON equip.obra_id = dm.obra_id AND equip.data = dm.data
LEFT JOIN material_diario m ON m.obra_id = dm.obra_id AND m.data = dm.data
GROUP BY dm.id, dm.obra_id, o.nome, dm.data, dm.periodo, dm.foto, dm.observacoes, 
         dm.responsavel_id, p1.nome, dm.aprovado_por_id, p2.nome, dm.status_aprovacao, 
         dm.created_at, dm.updated_at;

COMMENT ON VIEW vw_diario_consolidado IS 'View consolidada do diário de obra - agregando atividades, ocorrências, equipe, equipamentos e materiais por obra/data';

