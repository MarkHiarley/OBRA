-- Reverter correções de constraints (se necessário)

-- Nota: Esta migration down é complexa porque envolve remover constraints
-- que podem ter sido criados de formas diferentes. Em produção, 
-- geralmente não fazemos rollback de correções de constraints.

-- Se precisar reverter, execute manualmente:
-- ALTER TABLE obra ADD CONSTRAINT obra_status_check CHECK (...);
