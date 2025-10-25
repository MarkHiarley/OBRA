# 📋 Documentação Completa do Sistema de Obras

## 🎯 Visão Geral

Sistema completo de gestão de obras com controle financeiro, diário de obra, fornecedores e relatórios detalhados.

---

## 📊 Funcionalidades Principais

### 1. 👥 Gestão de Pessoas/Empresas

**Campos disponíveis:**
- ✅ Nome Completo
- ✅ Nome Fantasia (apenas para CNPJ)
- ✅ CPF/CNPJ (validação automática)
- ✅ Função (desabilitado para CNPJ)
- ✅ Endereço Completo (Rua, Número, Complemento, Bairro, Cidade, Estado, CEP)
- ✅ Contato (telefone, e-mail)
- ✅ Observações
- ✅ Status (Ativo/Inativo)

**Tipos:**
- **Pessoa Física (CPF)**: Profissionais (Engenheiro, Arquiteto, Mestre de Obra, etc.)
- **Pessoa Jurídica (CNPJ)**: Empresas (Construtoras, Fornecedores, etc.)

---

### 2. 🏗️ Gestão de Obras

**Campos disponíveis:**
- ✅ Nome da Obra
- ✅ ART (Anotação de Responsabilidade Técnica)
- ✅ Cliente (busca de pessoa cadastrada)
- ✅ Responsável Técnico (busca de pessoa)
- ✅ Parceiro (busca de pessoa)
- ✅ Endereço da Obra (preenchimento automático do cliente)
- ✅ Data de Início
- ✅ Data de Previsão de Término
- ✅ Data de Término Real
- ✅ Status (Planejada, Em Andamento, Pausada, Concluída, Cancelada)

**Controle Financeiro:**
- ✅ Orçamento Inicial
- ✅ Valor Total da Obra
- ✅ Valor Aditivo
- ✅ Custo de Mão de Obra
- ✅ Despesas Gerais
- ✅ Lucro

---

### 3. 💰 Gestão de Despesas

**Campos disponíveis:**
- ✅ Data da Despesa
- ✅ Descrição Detalhada
- ✅ Categoria:
  - Material
  - Mão de Obra
  - Combustível
  - Alimentação
  - Material Elétrico
  - Aluguel de Equipamento
  - Transporte
  - Outros
- ✅ Fornecedor (busca de fornecedor cadastrado)
- ✅ Valor
- ✅ Forma de Pagamento:
  - PIX
  - Boleto
  - Cartão de Crédito
  - Cartão de Débito
  - Transferência
  - Espécie
  - Cheque
- ✅ Status do Pagamento (Pendente, Pago, Cancelado)
- ✅ Data de Pagamento
- ✅ Responsável pelo Pagamento
- ✅ Observações

---

### 4. 📖 Diário de Obra

**Campos disponíveis:**
- ✅ Data
- ✅ Descrição do que foi feito
- ✅ Ferramentas Utilizadas
- ✅ Quantidade de Pessoas que Trabalharam
- ✅ Responsável pela Execução
- ✅ Clima (Ensolarado, Nublado, Chuvoso, Ventoso, Outros)
- ✅ Progresso Percentual
- ✅ Problemas Encontrados
- ✅ Fotos Anexadas (URLs)
- ✅ Observações Adicionais

---

### 5. 📦 Gestão de Fornecedores

**Campos disponíveis:**
- ✅ Nome/Razão Social
- ✅ Tipo (CPF/CNPJ)
- ✅ Documento
- ✅ E-mail
- ✅ Telefone
- ✅ Endereço Completo
- ✅ Cidade/Estado
- ✅ Status (Ativo/Inativo)

---

## 📊 Relatórios Disponíveis

### 1. Relatório de Obra
- Orçamento previsto vs. Gasto realizado
- Status da obra
- Cronograma (início, previsão, conclusão)
- Resumo financeiro completo

### 2. Relatório de Despesas
- Despesas por categoria
- Despesas por fornecedor
- Despesas por período
- Status de pagamentos

### 3. Relatório de Pagamentos
- Pagamentos realizados
- Pagamentos pendentes
- Por forma de pagamento
- Por período

### 4. Relatório Financeiro Consolidado
- Valor Total da Obra
- Valor Aditivo
- Custo de Mão de Obra
- Despesas Gerais
- Lucro Real

---

## 🗄️ Estrutura do Banco de Dados

### Tabelas Criadas:

1. **pessoa** - Pessoas físicas e jurídicas
2. **usuario** - Usuários do sistema
3. **obra** - Obras cadastradas
4. **diario** - Diário de obra
5. **fornecedor** - Fornecedores de materiais e serviços
6. **despesa** - Despesas por obra

### Relacionamentos:

```
obra
├── cliente_id → pessoa
├── responsavel_id → pessoa
└── parceiro_id → pessoa

despesa
├── obra_id → obra
└── fornecedor_id → fornecedor

diario
└── obra_id → obra
```

---

## 📥 Exemplo de Dados (Baseado na Planilha)

### Exemplo de Despesa:

```json
{
  "data_despesa": "2024-12-03",
  "descricao": "AREIA (CARRADA)",
  "categoria": "MATERIAL",
  "fornecedor": "J. GIRÃO (CROATÁ)",
  "valor": 1000.00,
  "forma_pagamento": "PIX",
  "status_pagamento": "PAGO",
  "data_pagamento": "2024-12-03"
}
```

### Exemplo de Resumo Financeiro da Obra:

```json
{
  "valor_total": 89673.35,
  "valor_aditivo": 16621.86,
  "custo_mao_obra": 18125.30,
  "despesas_gerais": 45380.21,
  "lucro": 42789.70
}
```

---

## 🚀 Como Executar as Migrations

### 1. Suba o ambiente Docker:
```bash
docker compose up -d
```

### 2. Execute as migrations:
```bash
chmod +x run-migrations.sh
./run-migrations.sh
```

### 3. Verifique as tabelas:
```bash
docker exec -it db_obras psql -U obras -d obrasdb -c "\dt"
```

---

## 📋 Ordem das Migrations

1. `000001_create_pessoa` - Tabela de pessoas
2. `000002_create_usuario` - Tabela de usuários
3. `000003_create_obra` - Tabela de obras
4. `000004_create_diario` - Tabela de diário
5. `000005_seed_data` - Dados iniciais
6. `000006_create_fornecedor` - Tabela de fornecedores ✨ NOVO
7. `000007_create_despesa` - Tabela de despesas ✨ NOVO
8. `000008_alter_obra_add_fields` - Campos financeiros em obra ✨ NOVO
9. `000009_alter_diario_add_fields` - Campos detalhados em diário ✨ NOVO
10. `000010_alter_pessoa_add_fields` - Campos de empresa em pessoa ✨ NOVO
11. `000011_seed_fornecedores` - Dados de fornecedores ✨ NOVO

---

## ✅ Checklist de Implementação

### Banco de Dados
- [x] Tabela de fornecedores
- [x] Tabela de despesas
- [x] Campos financeiros em obras
- [x] Campos detalhados em diário
- [x] Campos de empresa em pessoa
- [x] Seeds com dados da planilha

### Backend (Próximos Passos)
- [ ] Models para Fornecedor e Despesa
- [ ] Services para Fornecedor e Despesa
- [ ] UseCases para Fornecedor e Despesa
- [ ] Controllers para Fornecedor e Despesa
- [ ] Rotas para Fornecedor e Despesa
- [ ] Endpoint de relatórios

### Frontend (Futuro)
- [ ] Tela de cadastro de fornecedores
- [ ] Tela de registro de despesas
- [ ] Dashboard financeiro
- [ ] Relatórios interativos
- [ ] Gráficos de gastos

---

## 🎯 Próximas Funcionalidades

1. **Relatórios em PDF/Excel**
2. **Dashboard com gráficos**
3. **Sistema de notificações**
4. **Controle de estoque de materiais**
5. **Gestão de equipe e ponto**
6. **Upload de fotos no diário**
7. **Integração com sistemas contábeis**

---

**Última atualização**: 24 de outubro de 2025
