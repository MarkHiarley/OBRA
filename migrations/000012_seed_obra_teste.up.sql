-- Inserir obra de teste para desenvolvimento
INSERT INTO obra (nome, contrato_numero, data_inicio, prazo_dias, status, ativo, endereco_cidade, endereco_estado)
VALUES 
    ('Casa Residencial - Fortaleza', 'CONTR-2024-001', '2024-10-01', 180, 'EM_ANDAMENTO', true, 'Fortaleza', 'CE')
ON CONFLICT (contrato_numero) DO NOTHING;
