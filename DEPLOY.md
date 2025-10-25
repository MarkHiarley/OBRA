# 🚀 Guia Rápido de Deploy - Novas Funcionalidades

## ✅ O Que Foi Adicionado

### 📦 Novas Tabelas:
1. **fornecedor** - Cadastro de fornecedores
2. **despesa** - Controle de despesas por obra

### 🔄 Tabelas Atualizadas:
1. **obra** - Novos campos financeiros (ART, cliente, responsável, valores, etc.)
2. **diario** - Campos detalhados (ferramentas, pessoas, clima, progresso, etc.)
3. **pessoa** - Campos para empresas (nome fantasia, endereço completo, etc.)

---

## 🗄️ Migrations Criadas

Total: **6 novas migrations**

| # | Nome | Descrição |
|---|------|-----------|
| 000006 | `create_fornecedor` | Cria tabela de fornecedores |
| 000007 | `create_despesa` | Cria tabela de despesas |
| 000008 | `alter_obra_add_fields` | Adiciona campos financeiros em obras |
| 000009 | `alter_diario_add_fields` | Adiciona campos detalhados em diário |
| 000010 | `alter_pessoa_add_fields` | Adiciona campos de empresa em pessoa |
| 000011 | `seed_fornecedores` | Insere fornecedores da planilha |

---

## 🚀 Como Fazer o Deploy

### 1️⃣ **Commit e Push**

```bash
git add .
git commit -m "feat: Adiciona gestão de fornecedores, despesas e relatórios financeiros"
git push
```

### 2️⃣ **No Servidor VPS**

```bash
# Pull das mudanças
cd ~/OBRA
git pull

# Rebuild do container
docker compose down
docker compose build --no-cache

# Sobe os containers
docker compose up -d

# Aguarda banco ficar pronto
sleep 5

# Executa as migrations
chmod +x run-migrations.sh
./run-migrations.sh

# Verifica os logs
docker logs api_obras -f
```

### 3️⃣ **Verificar Tabelas**

```bash
# Conecta no banco
docker exec -it db_obras psql -U obras -d obrasdb

# Lista tabelas
\dt

# Deve mostrar:
#  fornecedor
#  despesa
#  obra (com novos campos)
#  diario (com novos campos)
#  pessoa (com novos campos)

# Sai do psql
\q
```

---

## 🧪 Testar as Migrações

### Verificar Fornecedores:
```bash
docker exec -it db_obras psql -U obras -d obrasdb \
  -c "SELECT COUNT(*) FROM fornecedor;"
```

**Deve retornar:** 10 fornecedores

### Verificar Campos de Obra:
```bash
docker exec -it db_obras psql -U obras -d obrasdb \
  -c "SELECT column_name FROM information_schema.columns WHERE table_name='obra';"
```

**Deve incluir:**
- art
- cliente_id
- responsavel_id
- parceiro_id
- orcamento_inicial
- valor_total
- valor_aditivo
- custo_mao_obra
- despesas_gerais
- lucro
- status

---

## ⚠️ Troubleshooting

### Problema: Migration já executada

```bash
# Erro: "relation already exists"
# Solução: As migrations usam IF NOT EXISTS, pode ignorar
```

### Problema: Constraint violation

```bash
# Se der erro de foreign key, verifique se as tabelas pai existem:
docker exec -it db_obras psql -U obras -d obrasdb -c "\dt"
```

### Problema: Rollback necessário

```bash
# Para reverter UMA migration:
docker exec -it db_obras psql -U obras -d obrasdb < migrations/000011_seed_fornecedores.down.sql

# Para reverter TODAS as novas migrations:
for i in {11..6}; do
  docker exec -it db_obras psql -U obras -d obrasdb < "migrations/0000${i}_*.down.sql"
done
```

---

## 📊 Dados de Exemplo

Após executar as migrations, você terá:

### 10 Fornecedores cadastrados:
- J. GIRÃO (CROATÁ)
- AÇO CEARENSE
- IF3 EMPREENDIMENTOS
- R.N.ALVES ELETRICA
- BETACON
- SERRA VIDROS
- TOP TINTAS
- ANTONIO GOMES SOUSA
- CHINA ELETRICISTA
- DIVERSOS

---

## 🎯 Próximos Passos

Após o deploy das migrations:

1. [ ] Criar models para Fornecedor e Despesa
2. [ ] Criar services para Fornecedor e Despesa
3. [ ] Criar usecases para Fornecedor e Despesa
4. [ ] Criar controllers para Fornecedor e Despesa
5. [ ] Adicionar rotas no main.go
6. [ ] Testar endpoints com Postman/curl
7. [ ] Atualizar documentação da API

---

## 📞 Suporte

Se houver problemas durante o deploy:

1. Verifique os logs: `docker logs api_obras`
2. Verifique o banco: `docker exec -it db_obras psql -U obras -d obrasdb`
3. Revise o arquivo [FEATURES.md](FEATURES.md) para detalhes

---

**Status**: ✅ Pronto para deploy  
**Data**: 24 de outubro de 2025
