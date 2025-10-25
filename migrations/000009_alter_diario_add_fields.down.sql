ALTER TABLE diario_obra
    DROP COLUMN IF EXISTS ferramentas_utilizadas,
    DROP COLUMN IF EXISTS quantidade_pessoas,
    DROP COLUMN IF EXISTS responsavel_execucao,
    DROP COLUMN IF EXISTS clima,
    DROP COLUMN IF EXISTS progresso_percentual,
    DROP COLUMN IF EXISTS problemas_encontrados,
    DROP COLUMN IF EXISTS fotos_anexadas;
