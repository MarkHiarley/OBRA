-- Seed com dados da planilha de gastos
-- Inserir fornecedores
INSERT INTO fornecedor (nome, tipo_documento, documento, telefone, cidade, estado) VALUES
('J. GIRÃO (CROATÁ)', 'CNPJ', '00.000.000/0001-01', '(85) 9999-9999', 'Croatá', 'CE'),
('AÇO CEARENSE', 'CNPJ', '00.000.000/0001-02', '(85) 9999-9998', 'Fortaleza', 'CE'),
('IF3 EMPREENDIMENTOS', 'CNPJ', '00.000.000/0001-03', '(85) 9999-9997', 'Fortaleza', 'CE'),
('R.N.ALVES ELETRICA', 'CNPJ', '00.000.000/0001-04', '(85) 9999-9996', 'Fortaleza', 'CE'),
('BETACON', 'CNPJ', '00.000.000/0001-05', '(85) 9999-9995', 'Fortaleza', 'CE'),
('SERRA VIDROS', 'CNPJ', '00.000.000/0001-06', '(85) 9999-9994', 'Fortaleza', 'CE'),
('TOP TINTAS', 'CNPJ', '00.000.000/0001-07', '(85) 9999-9993', 'Fortaleza', 'CE'),
('ANTONIO GOMES SOUSA', 'CPF', '000.000.000-01', '(85) 9999-9992', 'Fortaleza', 'CE'),
('CHINA ELETRICISTA', 'CPF', '000.000.000-02', '(85) 9999-9991', 'Fortaleza', 'CE'),
('DIVERSOS', 'CNPJ', '00.000.000/0001-99', NULL, NULL, NULL)
ON CONFLICT (documento) DO NOTHING;
