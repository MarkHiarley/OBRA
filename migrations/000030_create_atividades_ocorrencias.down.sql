-- Remover view consolidada
DROP VIEW IF EXISTS vw_diario_consolidado;

-- Remover índices
DROP INDEX IF EXISTS idx_diario_meta_aprovacao;
DROP INDEX IF EXISTS idx_diario_meta_obra_data;

DROP INDEX IF EXISTS idx_ocorrencia_status;
DROP INDEX IF EXISTS idx_ocorrencia_responsavel;
DROP INDEX IF EXISTS idx_ocorrencia_gravidade;
DROP INDEX IF EXISTS idx_ocorrencia_tipo;
DROP INDEX IF EXISTS idx_ocorrencia_periodo;
DROP INDEX IF EXISTS idx_ocorrencia_obra_data;

DROP INDEX IF EXISTS idx_atividade_status;
DROP INDEX IF EXISTS idx_atividade_responsavel;
DROP INDEX IF EXISTS idx_atividade_periodo;
DROP INDEX IF EXISTS idx_atividade_obra_data;

-- Remover tabelas
DROP TABLE IF EXISTS diario_metadados;
DROP TABLE IF EXISTS ocorrencia_diaria;
DROP TABLE IF EXISTS atividade_diaria;
