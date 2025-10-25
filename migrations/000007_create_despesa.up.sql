-- Tabela de Despesas por Obra
CREATE TABLE IF NOT EXISTS despesa (
    id SERIAL PRIMARY KEY,
    obra_id INTEGER NOT NULL REFERENCES obra(id) ON DELETE CASCADE,
    fornecedor_id INTEGER REFERENCES fornecedor(id) ON DELETE SET NULL,
    data_despesa DATE NOT NULL,
    descricao TEXT NOT NULL,
    categoria VARCHAR(50) CHECK (categoria IN (
        'MATERIAL', 
        'MAO_DE_OBRA', 
        'COMBUSTIVEL', 
        'ALIMENTACAO',
        'MATERIAL_ELETRICO', 
        'ALUGUEL_EQUIPAMENTO',
        'TRANSPORTE',
        'OUTROS'
    )),
    valor DECIMAL(15, 2) NOT NULL CHECK (valor >= 0),
    forma_pagamento VARCHAR(30) CHECK (forma_pagamento IN (
        'PIX', 
        'BOLETO', 
        'CARTAO_CREDITO', 
        'CARTAO_DEBITO',
        'TRANSFERENCIA',
        'ESPECIE',
        'CHEQUE'
    )),
    status_pagamento VARCHAR(20) DEFAULT 'PENDENTE' CHECK (status_pagamento IN (
        'PENDENTE', 
        'PAGO', 
        'CANCELADO'
    )),
    data_pagamento DATE,
    responsavel_pagamento VARCHAR(255),
    observacao TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_despesa_obra ON despesa(obra_id);
CREATE INDEX idx_despesa_fornecedor ON despesa(fornecedor_id);
CREATE INDEX idx_despesa_data ON despesa(data_despesa);
CREATE INDEX idx_despesa_categoria ON despesa(categoria);
CREATE INDEX idx_despesa_status ON despesa(status_pagamento);
