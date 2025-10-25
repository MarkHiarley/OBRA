# 🗄️ Diagrama do Banco de Dados

## Estrutura Completa das Tabelas

```
┌─────────────────────────────────────────────────────────────┐
│                        PESSOA                                │
├─────────────────────────────────────────────────────────────┤
│ PK  id                                                       │
│     nome                                                     │
│     tipo (CPF/CNPJ)                                         │
│     documento                                               │
│     email                                                   │
│     telefone                                                │
│     cargo                                                   │
│     ativo                                                   │
│ ✨  nome_fantasia                                           │
│ ✨  eh_pessoa_juridica                                      │
│ ✨  funcao                                                  │
│ ✨  endereco_rua                                            │
│ ✨  endereco_numero                                         │
│ ✨  endereco_complemento                                    │
│ ✨  endereco_bairro                                         │
│ ✨  endereco_cidade                                         │
│ ✨  endereco_estado                                         │
│ ✨  endereco_cep                                            │
│ ✨  observacao                                              │
│     created_at                                              │
│     updated_at                                              │
└─────────────────────────────────────────────────────────────┘
                           │
                           │ (cliente_id, responsavel_id, parceiro_id)
                           ▼
┌─────────────────────────────────────────────────────────────┐
│                         OBRA                                 │
├─────────────────────────────────────────────────────────────┤
│ PK  id                                                       │
│     nome                                                     │
│ ✨  art                                                      │
│ FK  cliente_id → pessoa                                     │
│ FK  responsavel_id → pessoa                                 │
│ FK  parceiro_id → pessoa                                    │
│ ✨  data_inicio                                             │
│ ✨  data_previsao_termino                                   │
│ ✨  data_termino_real                                       │
│ ✨  orcamento_inicial                                       │
│ ✨  valor_total                                             │
│ ✨  valor_aditivo                                           │
│ ✨  custo_mao_obra                                          │
│ ✨  despesas_gerais                                         │
│ ✨  lucro                                                   │
│ ✨  status (PLANEJADA|EM_ANDAMENTO|PAUSADA|CONCLUIDA)       │
│ ✨  endereco_rua                                            │
│ ✨  endereco_numero                                         │
│ ✨  endereco_complemento                                    │
│ ✨  endereco_bairro                                         │
│ ✨  endereco_cidade                                         │
│ ✨  endereco_estado                                         │
│ ✨  endereco_cep                                            │
│     created_at                                              │
│     updated_at                                              │
└─────────────────────────────────────────────────────────────┘
          │                           │
          │ (obra_id)                 │ (obra_id)
          ▼                           ▼
┌─────────────────────────┐    ┌──────────────────────────────┐
│       DIARIO            │    │        DESPESA                │
├─────────────────────────┤    ├──────────────────────────────┤
│ PK  id                  │    │ PK  id                        │
│ FK  obra_id             │    │ FK  obra_id                   │
│     data                │    │ FK  fornecedor_id             │
│     descricao           │    │     data_despesa              │
│     observacoes         │    │     descricao                 │
│ ✨  ferramentas_        │    │     categoria                 │
│     utilizadas          │    │     valor                     │
│ ✨  quantidade_pessoas  │    │     forma_pagamento           │
│ ✨  responsavel_execucao│    │     status_pagamento          │
│ ✨  clima               │    │     data_pagamento            │
│ ✨  progresso_percentual│    │     responsavel_pagamento     │
│ ✨  problemas_          │    │     observacao                │
│     encontrados         │    │     created_at                │
│ ✨  fotos_anexadas      │    │     updated_at                │
│     created_at          │    └──────────────────────────────┘
│     updated_at          │                   │
└─────────────────────────┘                   │ (fornecedor_id)
                                              ▼
                                   ┌──────────────────────────┐
                                   │      FORNECEDOR          │
                                   ├──────────────────────────┤
                                   │ PK  id                   │
                                   │     nome                 │
                                   │     tipo_documento       │
                                   │     documento            │
                                   │     email                │
                                   │     telefone             │
                                   │     endereco             │
                                   │     cidade               │
                                   │     estado               │
                                   │     ativo                │
                                   │     created_at           │
                                   │     updated_at           │
                                   └──────────────────────────┘

┌─────────────────────────┐
│       USUARIO           │
├─────────────────────────┤
│ PK  id                  │
│     email               │
│     nome                │
│     senha_hash          │
│     tipo_documento      │
│     documento           │
│     telefone            │
│     perfil_acesso       │
│     ativo               │
│     created_at          │
│     updated_at          │
└─────────────────────────┘
```

## 📊 Legenda

- `PK` - Primary Key (Chave Primária)
- `FK` - Foreign Key (Chave Estrangeira)
- `✨` - Campos novos adicionados
- `→` - Referência para outra tabela

## 🔗 Relacionamentos

### OBRA
- **cliente_id** → PESSOA (Cliente da obra)
- **responsavel_id** → PESSOA (Responsável técnico)
- **parceiro_id** → PESSOA (Parceiro/Colaborador)

### DESPESA
- **obra_id** → OBRA (Obra relacionada)
- **fornecedor_id** → FORNECEDOR (Fornecedor da despesa)

### DIARIO
- **obra_id** → OBRA (Obra do diário)

## 📈 Índices Criados

### Tabela OBRA
- `idx_obra_cliente` - Busca por cliente
- `idx_obra_responsavel` - Busca por responsável
- `idx_obra_status` - Filtro por status
- `idx_obra_data_inicio` - Ordenação por data

### Tabela DESPESA
- `idx_despesa_obra` - Busca por obra
- `idx_despesa_fornecedor` - Busca por fornecedor
- `idx_despesa_data` - Ordenação por data
- `idx_despesa_categoria` - Filtro por categoria
- `idx_despesa_status` - Filtro por status de pagamento

### Tabela DIARIO
- `idx_diario_data` - Ordenação por data
- `idx_diario_responsavel` - Busca por responsável

### Tabela FORNECEDOR
- `idx_fornecedor_documento` - Busca por CPF/CNPJ
- `idx_fornecedor_ativo` - Filtro por ativos

### Tabela PESSOA
- `idx_pessoa_juridica` - Filtro por tipo de pessoa
- `idx_pessoa_funcao` - Busca por função

## 🎯 Constraints

### CHECK Constraints

**DESPESA:**
- `categoria IN ('MATERIAL', 'MAO_DE_OBRA', 'COMBUSTIVEL', ...)`
- `forma_pagamento IN ('PIX', 'BOLETO', 'CARTAO_CREDITO', ...)`
- `status_pagamento IN ('PENDENTE', 'PAGO', 'CANCELADO')`
- `valor >= 0`

**OBRA:**
- `status IN ('PLANEJADA', 'EM_ANDAMENTO', 'PAUSADA', 'CONCLUIDA', 'CANCELADA')`

**DIARIO:**
- `clima IN ('ENSOLARADO', 'NUBLADO', 'CHUVOSO', 'VENTOSO', 'OUTROS')`
- `progresso_percentual >= 0 AND progresso_percentual <= 100`

**FORNECEDOR:**
- `tipo_documento IN ('CPF', 'CNPJ')`

## 📝 Campos Únicos

- `pessoa.documento` - CPF/CNPJ único
- `usuario.email` - Email único
- `fornecedor.documento` - CPF/CNPJ único

## 🔄 Cascade Actions

- **DESPESA.obra_id** - ON DELETE CASCADE (se a obra for deletada, suas despesas também)
- **DESPESA.fornecedor_id** - ON DELETE SET NULL (se o fornecedor for deletado, o campo fica NULL)
- **DIARIO.obra_id** - ON DELETE CASCADE (se a obra for deletada, seus diários também)

---

**Total de Tabelas**: 6  
**Total de Relacionamentos**: 5  
**Total de Índices**: 14  
**Campos com ✨ (novos)**: 34
