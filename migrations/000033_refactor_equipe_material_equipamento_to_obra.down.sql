-- Rollback: Reverter equipe_diario, material_diario e equipamento_diario para usar diario_id

-- 1. EQUIPE_DIARIO: Recriar coluna diario_id
ALTER TABLE equipe_diario ADD COLUMN diario_id INTEGER;

-- Remover índices novos
DROP INDEX IF EXISTS idx_equipe_obra_data;
DROP INDEX IF EXISTS idx_equipe_data;
DROP INDEX IF EXISTS idx_equipe_obra_id;

-- Remover FK e colunas novas
ALTER TABLE equipe_diario DROP CONSTRAINT IF EXISTS fk_equipe_obra;
ALTER TABLE equipe_diario DROP COLUMN obra_id;
ALTER TABLE equipe_diario DROP COLUMN data;

-- Recriar FK antiga
ALTER TABLE equipe_diario
ADD CONSTRAINT fk_equipe_diario FOREIGN KEY (diario_id) REFERENCES diario_obra(id) ON DELETE CASCADE;

CREATE INDEX idx_equipe_diario_id ON equipe_diario(diario_id);

-- 2. MATERIAL_DIARIO: Recriar coluna diario_id
ALTER TABLE material_diario ADD COLUMN diario_id INTEGER;

-- Remover índices novos
DROP INDEX IF EXISTS idx_material_obra_data;
DROP INDEX IF EXISTS idx_material_data;
DROP INDEX IF EXISTS idx_material_obra_id;

-- Remover FK e colunas novas
ALTER TABLE material_diario DROP CONSTRAINT IF EXISTS fk_material_obra;
ALTER TABLE material_diario DROP COLUMN obra_id;
ALTER TABLE material_diario DROP COLUMN data;

-- Recriar FK antiga
ALTER TABLE material_diario
ADD CONSTRAINT fk_material_diario FOREIGN KEY (diario_id) REFERENCES diario_obra(id) ON DELETE CASCADE;

CREATE INDEX idx_material_diario_id ON material_diario(diario_id);

-- 3. EQUIPAMENTO_DIARIO: Recriar coluna diario_id
ALTER TABLE equipamento_diario ADD COLUMN diario_id INTEGER;

-- Remover índices novos
DROP INDEX IF EXISTS idx_equipamento_obra_data;
DROP INDEX IF EXISTS idx_equipamento_data;
DROP INDEX IF EXISTS idx_equipamento_obra_id;

-- Remover FK e colunas novas
ALTER TABLE equipamento_diario DROP CONSTRAINT IF EXISTS fk_equipamento_obra;
ALTER TABLE equipamento_diario DROP COLUMN obra_id;
ALTER TABLE equipamento_diario DROP COLUMN data;

-- Recriar FK antiga
ALTER TABLE equipamento_diario
ADD CONSTRAINT fk_equipamento_diario FOREIGN KEY (diario_id) REFERENCES diario_obra(id) ON DELETE CASCADE;

CREATE INDEX idx_equipamento_diario_id ON equipamento_diario(diario_id);
