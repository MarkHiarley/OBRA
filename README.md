# 🏗️ OBRA - Sistema de Gerenciamento de Obras

API RESTful para gerenciamento de obras, construída em Go com Gin Framework e PostgreSQL.

## � Quick Start

```bash
# 1. Clone e configure
git clone https://github.com/MarkHiarley/OBRA.git
cd OBRA

# 2. Inicie os containers
docker compose up -d

# 3. Execute as migrations
chmod +x run-migrations.sh
./run-migrations.sh

# 4. Acesse a API
curl http://localhost:9090/pessoas
```

Pronto! A API está rodando em `http://localhost:9090` 🎉

---

## �📋 Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Tecnologias](#tecnologias)
- [Arquitetura](#arquitetura)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Configuração](#configuração)
- [Executando o Projeto](#executando-o-projeto)
- [Documentação da API](#documentação-da-api)
  - [Pessoas](#pessoas)
  - [Usuários](#usuários)
  - [Obras](#obras)
  - [Diários de Obra](#diários-de-obra)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Migrations](#migrations)

---

## 🎯 Sobre o Projeto

O sistema OBRA é uma solução completa para gerenciamento de obras, permitindo o controle de:

- **Pessoas**: Cadastro de profissionais envolvidos nas obras
- **Usuários**: Gestão de acesso ao sistema com perfis diferenciados
- **Obras**: Controle completo de projetos, contratos e prazos
- **Diários de Obra**: Registro diário de atividades, ocorrências e aprovações

---

## 🚀 Tecnologias

- **[Go 1.25](https://golang.org/)** - Linguagem de programação
- **[Gin](https://github.com/gin-gonic/gin)** - Framework web HTTP
- **[PostgreSQL 12](https://www.postgresql.org/)** - Banco de dados relacional
- **[Docker](https://www.docker.com/)** - Containerização
- **[Docker Compose](https://docs.docker.com/compose/)** - Orquestração de containers
- **[golang-migrate](https://github.com/golang-migrate/migrate)** - Migrations de banco de dados

---

## 🏛️ Arquitetura

O projeto segue a **Clean Architecture** com separação clara de responsabilidades:

```
┌─────────────┐
│ Controllers │  ← Camada de apresentação (HTTP handlers)
└──────┬──────┘
       │
┌──────▼──────┐
│  Use Cases  │  ← Lógica de negócio
└──────┬──────┘
       │
┌──────▼──────┐
│  Services   │  ← Acesso a dados
└──────┬──────┘
       │
┌──────▼──────┐
│   Models    │  ← Estruturas de dados
└─────────────┘
```

---

## 📦 Pré-requisitos

- Docker >= 20.10
- Docker Compose >= 2.0
- Make (opcional, para comandos facilitados)

---

## 💻 Instalação

### 1. Clone o repositório

```bash
git clone https://github.com/MarkHiarley/OBRA.git
cd OBRA
```

### 2. Configure as variáveis de ambiente

Crie ou edite o arquivo `.env` na raiz do projeto:

```env
# Configuração do Banco de Dados
DB_HOST=db_obras
DB_PORT=5432
DB_USER=obras
DB_PASSWORD=7894
DB_NAME=obrasdb

# Porta de mapeamento do host para o banco
DB_HOST_PORT=5440

# Porta da API
API_PORT=9090
```

> ⚠️ **Importante**: 
> - `DB_PORT=5432` é a porta interna do container (não altere)
> - `DB_HOST_PORT=5440` é a porta exposta no seu computador (pode alterar se necessário)

---

## 🎮 Configuração

### Variáveis de Ambiente

| Variável | Descrição | Padrão |
|----------|-----------|--------|
| `DB_HOST` | Host do banco de dados | `db_obras` |
| `DB_PORT` | Porta interna do PostgreSQL | `5432` |
| `DB_USER` | Usuário do banco | `obras` |
| `DB_PASSWORD` | Senha do banco | `7894` |
| `DB_NAME` | Nome do banco de dados | `obrasdb` |
| `DB_HOST_PORT` | Porta exposta no host | `5440` |
| `API_PORT` | Porta da API | `9090` |

---

## 🚀 Executando o Projeto

### Usando Docker Compose (Recomendado)

```bash
# 1. Iniciar todos os serviços
docker compose up -d

# 2. Aguardar o banco inicializar (cerca de 5-10 segundos)
sleep 10

# 3. Executar as migrations
./run-migrations.sh

# 4. Ver logs da API
docker logs api_obras -f

# 5. Ver logs do banco
docker logs db_obras -f

# Parar os serviços
docker compose down

# Reconstruir e iniciar (após mudanças no código)
docker compose down
docker compose up -d --build
./run-migrations.sh
```

### Fluxo Completo de Inicialização

```bash
# Passo 1: Subir os containers
docker compose up -d

# Passo 2: Executar migrations (script pronto)
chmod +x run-migrations.sh
./run-migrations.sh

# Passo 3: Verificar se a API está rodando
docker logs api_obras

# Passo 4: Testar a API
curl http://localhost:9090/pessoas
```

### Usando Make

```bash
# Iniciar o projeto
make run

# Ver outros comandos disponíveis
make help
```

### Acessar a API

Após iniciar, a API estará disponível em:

```
http://localhost:9090
```

### Acessar o Banco de Dados

Para conectar ao PostgreSQL externamente:

```bash
psql -h localhost -p 5440 -U obras -d obrasdb
```

Ou usando uma ferramenta GUI com as seguintes credenciais:
- **Host**: localhost
- **Port**: 5440
- **Database**: obrasdb
- **User**: obras
- **Password**: 7894

---

## 📚 Documentação da API

Base URL: `http://localhost:9090`

### � Índice de Endpoints

- [👥 Pessoas](#-pessoas) - Gerenciamento de pessoas (contratantes, profissionais)
- [👤 Usuários](#-usuários) - Gerenciamento de usuários do sistema
- [🏗️ Obras](#️-obras) - Gerenciamento de obras e contratos
- [📖 Diários de Obra](#-diários-de-obra) - Registro diário de atividades

### 🔑 Códigos de Status HTTP

| Código | Descrição |
|--------|-----------|
| `200 OK` | Requisição bem-sucedida |
| `201 Created` | Recurso criado com sucesso |
| `204 No Content` | Requisição bem-sucedida sem conteúdo (DELETE) |
| `400 Bad Request` | Dados inválidos ou malformados |
| `404 Not Found` | Recurso não encontrado |
| `500 Internal Server Error` | Erro interno do servidor |

---

### �👥 Pessoas

#### Listar todas as pessoas
```http
GET /pessoas
```

**Resposta (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "nome": "João Silva",
      "tipo": "CPF",
      "documento": "123.456.789-00",
      "email": "joao@exemplo.com",
      "telefone": "(11) 98765-4321",
      "cargo": "Engenheiro Civil",
      "ativo": true,
      "createdAt": "2025-10-16T10:00:00Z",
      "updatedAt": "2025-10-16T10:00:00Z"
    }
  ]
}
```

#### Buscar pessoa por ID
```http
GET /pessoas/:id
```

**Parâmetros:**
- `id` (path): ID da pessoa

**Resposta (200 OK):**
```json
{
  "id": 1,
  "nome": "João Silva",
  "tipo": "CPF",
  "documento": "123.456.789-00",
  "email": "joao@exemplo.com",
  "telefone": "(11) 98765-4321",
  "cargo": "Engenheiro Civil",
  "ativo": true,
  "createdAt": "2025-10-16T10:00:00Z",
  "updatedAt": "2025-10-16T10:00:00Z"
}
```

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "Pessoa não encontrada"
}
```

#### Criar nova pessoa
```http
POST /pessoas
```

**Body:**
```json
{
  "nome": "Maria Santos",
  "tipo": "CPF",
  "documento": "987.654.321-00",
  "email": "maria@exemplo.com",
  "telefone": "(11) 91234-5678",
  "cargo": "Arquiteta",
  "ativo": true
}
```

**Resposta (201 Created):**
```json
{
  "message": "Pessoa criada com sucesso",
  "data": {
    "id": 2,
    "nome": "Maria Santos",
    "tipo": "CPF",
    "documento": "987.654.321-00",
    "email": "maria@exemplo.com",
    "telefone": "(11) 91234-5678",
    "cargo": "Arquiteta",
    "ativo": true,
    "createdAt": "2025-10-16T11:00:00Z",
    "updatedAt": "2025-10-16T11:00:00Z"
  }
}
```

#### Atualizar pessoa
```http
PUT /pessoas/:id
```

**Parâmetros:**
- `id` (path): ID da pessoa

**Body:**
```json
{
  "nome": "Maria Santos Silva",
  "tipo": "CPF",
  "documento": "987.654.321-00",
  "email": "maria.santos@exemplo.com",
  "telefone": "(11) 91234-5678",
  "cargo": "Arquiteta Sênior",
  "ativo": true
}
```

**Resposta (200 OK):**
```json
{
  "id": 2,
  "nome": "Maria Santos Silva",
  "tipo": "CPF",
  "documento": "987.654.321-00",
  "email": "maria.santos@exemplo.com",
  "telefone": "(11) 91234-5678",
  "cargo": "Arquiteta Sênior",
  "ativo": true,
  "createdAt": "2025-10-16T11:00:00Z",
  "updatedAt": "2025-10-16T12:00:00Z"
}
```

#### Deletar pessoa
```http
DELETE /pessoas/:id
```

**Parâmetros:**
- `id` (path): ID da pessoa

**Resposta (204 No Content):**
```
(sem corpo de resposta)
```

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "Pessoa não encontrada"
}
```

**Resposta de Erro (400 Bad Request):**
```json
{
  "error": "ID deve ser um número válido"
}
```

---

### 👤 Usuários

#### Listar todos os usuários
```http
GET /usuarios
```

**Resposta (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "email": "admin@obra.com",
      "nome": "Administrador",
      "tipo_documento": "CPF",
      "documento": "111.222.333-44",
      "telefone": "(11) 99999-9999",
      "perfil_acesso": "admin",
      "ativo": true,
      "createdAt": "2025-10-16T10:00:00Z",
      "updatedAt": "2025-10-16T10:00:00Z"
    }
  ]
}
```

#### Buscar usuário por ID
```http
GET /usuarios/:id
```

**Parâmetros:**
- `id` (path): ID do usuário

**Resposta (200 OK):**
```json
{
  "id": 1,
  "email": "admin@obra.com",
  "nome": "Administrador",
  "tipo_documento": "CPF",
  "documento": "111.222.333-44",
  "telefone": "(11) 99999-9999",
  "perfil_acesso": "admin",
  "ativo": true,
  "createdAt": "2025-10-16T10:00:00Z",
  "updatedAt": "2025-10-16T10:00:00Z"
}
```

#### Criar novo usuário
```http
POST /usuarios
```

**Body:**
```json
{
  "email": "usuario@obra.com",
  "nome": "João Usuário",
  "senha": "senha123",
  "tipo_documento": "CPF",
  "documento": "555.666.777-88",
  "telefone": "(11) 98888-7777",
  "perfil_acesso": "gestor",
  "ativo": true
}
```

**Resposta (201 Created):**
```json
{
  "message": "Usuario criado com sucesso",
  "data": {
    "id": 2,
    "email": "usuario@obra.com",
    "nome": "João Usuário",
    "tipo_documento": "CPF",
    "documento": "555.666.777-88",
    "telefone": "(11) 98888-7777",
    "perfil_acesso": "gestor",
    "ativo": true,
    "createdAt": "2025-10-16T11:00:00Z",
    "updatedAt": "2025-10-16T11:00:00Z"
  }
}
```

> 🔒 **Nota de Segurança**: A senha é automaticamente criptografada usando bcrypt antes de ser armazenada.

#### Atualizar usuário
```http
PUT /usuarios/:id
```

**Parâmetros:**
- `id` (path): ID do usuário

**Body:**
```json
{
  "email": "usuario.atualizado@obra.com",
  "nome": "João Usuário Atualizado",
  "tipo_documento": "CPF",
  "documento": "555.666.777-88",
  "telefone": "(11) 98888-7777",
  "perfil_acesso": "admin",
  "ativo": true
}
```

**Resposta (200 OK):**
```json
{
  "id": 2,
  "email": "usuario.atualizado@obra.com",
  "nome": "João Usuário Atualizado",
  "tipo_documento": "CPF",
  "documento": "555.666.777-88",
  "telefone": "(11) 98888-7777",
  "perfil_acesso": "admin",
  "ativo": true,
  "createdAt": "2025-10-16T11:00:00Z",
  "updatedAt": "2025-10-16T12:00:00Z"
}
```

#### Deletar usuário
```http
DELETE /usuarios/:id
```

**Parâmetros:**
- `id` (path): ID do usuário

**Resposta (204 No Content):**
```
(sem corpo de resposta)
```

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "Usuário não encontrado"
}
```

---

### 🏗️ Obras

#### Listar todas as obras
```http
GET /obras
```

**Resposta (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "nome": "Construção do Edifício Central",
      "contrato_numero": "CNT-2025-001",
      "contratante_id": 1,
      "responsavel_id": 2,
      "data_inicio": "2025-01-15",
      "prazo_dias": 365,
      "data_fim_prevista": "2026-01-15",
      "orcamento": 5000000.00,
      "status": "em_andamento",
      "endereco_rua": "Av. Principal",
      "endereco_numero": "1000",
      "endereco_bairro": "Centro",
      "endereco_cidade": "São Paulo",
      "endereco_estado": "SP",
      "endereco_cep": "01000-000",
      "observacoes": "Projeto prioritário",
      "ativo": true,
      "created_at": "2025-10-16T10:00:00Z",
      "updated_at": "2025-10-16T10:00:00Z"
    }
  ]
}
```

#### Buscar obra por ID
```http
GET /obras/:id
```

**Parâmetros:**
- `id` (path): ID da obra

**Resposta (200 OK):**
```json
{
  "id": 1,
  "nome": "Construção do Edifício Central",
  "contrato_numero": "CNT-2025-001",
  "contratante_id": 1,
  "responsavel_id": 2,
  "data_inicio": "2025-01-15",
  "prazo_dias": 365,
  "data_fim_prevista": "2026-01-15",
  "orcamento": 5000000.00,
  "status": "em_andamento",
  "endereco_rua": "Av. Principal",
  "endereco_numero": "1000",
  "endereco_bairro": "Centro",
  "endereco_cidade": "São Paulo",
  "endereco_estado": "SP",
  "endereco_cep": "01000-000",
  "observacoes": "Projeto prioritário",
  "ativo": true,
  "created_at": "2025-10-16T10:00:00Z",
  "updated_at": "2025-10-16T10:00:00Z"
}
```

#### Criar nova obra
```http
POST /obras
```

**Body:**
```json
{
  "nome": "Reforma do Prédio B",
  "contrato_numero": "CNT-2025-002",
  "contratante_id": 3,
  "responsavel_id": 4,
  "data_inicio": "2025-03-01",
  "prazo_dias": 180,
  "orcamento": 1500000.00,
  "status": "planejamento",
  "endereco_rua": "Rua Secundária",
  "endereco_numero": "500",
  "endereco_bairro": "Jardim",
  "endereco_cidade": "São Paulo",
  "endereco_estado": "SP",
  "endereco_cep": "02000-000",
  "observacoes": "Reforma completa",
  "ativo": true
}
```

**Resposta (201 Created):**
```json
{
  "message": "Obra criada com sucesso",
  "data": {
    "id": 2,
    "nome": "Reforma do Prédio B",
    "contrato_numero": "CNT-2025-002",
    "contratante_id": 3,
    "responsavel_id": 4,
    "data_inicio": "2025-03-01",
    "prazo_dias": 180,
    "data_fim_prevista": "2025-08-28",
    "orcamento": 1500000.00,
    "status": "planejamento",
    "endereco_rua": "Rua Secundária",
    "endereco_numero": "500",
    "endereco_bairro": "Jardim",
    "endereco_cidade": "São Paulo",
    "endereco_estado": "SP",
    "endereco_cep": "02000-000",
    "observacoes": "Reforma completa",
    "ativo": true,
    "created_at": "2025-10-16T11:00:00Z",
    "updated_at": "2025-10-16T11:00:00Z"
  }
}
```

#### Atualizar obra
```http
PUT /obras/:id
```

**Parâmetros:**
- `id` (path): ID da obra

**Body:**
```json
{
  "nome": "Reforma do Prédio B - Atualizado",
  "contrato_numero": "CNT-2025-002",
  "contratante_id": 3,
  "responsavel_id": 4,
  "data_inicio": "2025-03-01",
  "prazo_dias": 200,
  "orcamento": 1600000.00,
  "status": "em_andamento",
  "endereco_rua": "Rua Secundária",
  "endereco_numero": "500",
  "endereco_bairro": "Jardim",
  "endereco_cidade": "São Paulo",
  "endereco_estado": "SP",
  "endereco_cep": "02000-000",
  "observacoes": "Reforma completa com extensão de prazo",
  "ativo": true
}
```

**Resposta (200 OK):**
```json
{
  "id": 2,
  "nome": "Reforma do Prédio B - Atualizado",
  "contrato_numero": "CNT-2025-002",
  "contratante_id": 3,
  "responsavel_id": 4,
  "data_inicio": "2025-03-01",
  "prazo_dias": 200,
  "data_fim_prevista": "2025-09-17",
  "orcamento": 1600000.00,
  "status": "em_andamento",
  "endereco_rua": "Rua Secundária",
  "endereco_numero": "500",
  "endereco_bairro": "Jardim",
  "endereco_cidade": "São Paulo",
  "endereco_estado": "SP",
  "endereco_cep": "02000-000",
  "observacoes": "Reforma completa com extensão de prazo",
  "ativo": true,
  "created_at": "2025-10-16T11:00:00Z",
  "updated_at": "2025-10-16T12:00:00Z"
}
```

#### Deletar obra
```http
DELETE /obras/:id
```

**Parâmetros:**
- `id` (path): ID da obra

**Resposta (204 No Content):**
```
(sem corpo de resposta)
```

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "Obra não encontrada"
}
```

**Resposta de Erro (400 Bad Request):**
```json
{
  "error": "ID deve ser um número válido"
}
```

---

### 📖 Diários de Obra

#### Listar todos os diários
```http
GET /diarios
```

**Resposta (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "obra_id": 1,
      "data": "2025-10-15",
      "periodo": "manhã",
      "atividades_realizadas": "Concretagem da laje do 3º andar",
      "ocorrencias": "Chuva no período da tarde",
      "observacoes": "Trabalho interrompido às 14h",
      "responsavel_id": 2,
      "aprovado_por_id": 1,
      "status_aprovacao": "aprovado",
      "createdAt": "2025-10-15T18:00:00Z",
      "updatedAt": "2025-10-15T19:00:00Z"
    }
  ]
}
```

#### Buscar diário por ID
```http
GET /diarios/:id
```

**Parâmetros:**
- `id` (path): ID do diário

**Resposta (200 OK):**
```json
{
  "id": 1,
  "obra_id": 1,
  "data": "2025-10-15",
  "periodo": "manhã",
  "atividades_realizadas": "Concretagem da laje do 3º andar",
  "ocorrencias": "Chuva no período da tarde",
  "observacoes": "Trabalho interrompido às 14h",
  "responsavel_id": 2,
  "aprovado_por_id": 1,
  "status_aprovacao": "aprovado",
  "createdAt": "2025-10-15T18:00:00Z",
  "updatedAt": "2025-10-15T19:00:00Z"
}
```

#### Buscar diários por obra
```http
GET /diarios/:id/obra
```

**Parâmetros:**
- `id` (path): ID da obra

**Resposta (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "obra_id": 1,
      "data": "2025-10-15",
      "periodo": "manhã",
      "atividades_realizadas": "Concretagem da laje do 3º andar",
      "ocorrencias": "Chuva no período da tarde",
      "observacoes": "Trabalho interrompido às 14h",
      "responsavel_id": 2,
      "aprovado_por_id": 1,
      "status_aprovacao": "aprovado",
      "createdAt": "2025-10-15T18:00:00Z",
      "updatedAt": "2025-10-15T19:00:00Z"
    },
    {
      "id": 2,
      "obra_id": 1,
      "data": "2025-10-16",
      "periodo": "integral",
      "atividades_realizadas": "Instalação de tubulações hidráulicas",
      "ocorrencias": null,
      "observacoes": "Dia produtivo",
      "responsavel_id": 2,
      "aprovado_por_id": null,
      "status_aprovacao": "pendente",
      "createdAt": "2025-10-16T18:00:00Z",
      "updatedAt": "2025-10-16T18:00:00Z"
    }
  ]
}
```

#### Criar novo diário
```http
POST /diarios
```

**Body:**
```json
{
  "obra_id": 1,
  "data": "2025-10-16",
  "periodo": "integral",
  "atividades_realizadas": "Instalação de tubulações hidráulicas e elétricas no 4º andar",
  "ocorrencias": "Entrega de materiais atrasou 2 horas",
  "observacoes": "Equipe trabalhou até às 18h para compensar",
  "responsavel_id": 2,
  "status_aprovacao": "pendente"
}
```

**Resposta (201 Created):**
```json
{
  "message": "Diário criado com sucesso",
  "data": {
    "id": 3,
    "obra_id": 1,
    "data": "2025-10-16",
    "periodo": "integral",
    "atividades_realizadas": "Instalação de tubulações hidráulicas e elétricas no 4º andar",
    "ocorrencias": "Entrega de materiais atrasou 2 horas",
    "observacoes": "Equipe trabalhou até às 18h para compensar",
    "responsavel_id": 2,
    "aprovado_por_id": null,
    "status_aprovacao": "pendente",
    "createdAt": "2025-10-16T19:00:00Z",
    "updatedAt": "2025-10-16T19:00:00Z"
  }
}
```

#### Atualizar diário
```http
PUT /diarios/:id
```

**Parâmetros:**
- `id` (path): ID do diário

**Body:**
```json
{
  "obra_id": 1,
  "data": "2025-10-16",
  "periodo": "integral",
  "atividades_realizadas": "Instalação de tubulações hidráulicas e elétricas no 4º andar - Concluído",
  "ocorrencias": "Entrega de materiais atrasou 2 horas",
  "observacoes": "Equipe trabalhou até às 18h para compensar. Trabalho concluído.",
  "responsavel_id": 2,
  "aprovado_por_id": 1,
  "status_aprovacao": "aprovado"
}
```

**Resposta (200 OK):**
```json
{
  "id": 3,
  "obra_id": 1,
  "data": "2025-10-16",
  "periodo": "integral",
  "atividades_realizadas": "Instalação de tubulações hidráulicas e elétricas no 4º andar - Concluído",
  "ocorrencias": "Entrega de materiais atrasou 2 horas",
  "observacoes": "Equipe trabalhou até às 18h para compensar. Trabalho concluído.",
  "responsavel_id": 2,
  "aprovado_por_id": 1,
  "status_aprovacao": "aprovado",
  "createdAt": "2025-10-16T19:00:00Z",
  "updatedAt": "2025-10-16T20:00:00Z"
}
```

#### Deletar diário
```http
DELETE /diarios/:id
```

**Parâmetros:**
- `id` (path): ID do diário

**Resposta (204 No Content):**
```
(sem corpo de resposta)
```

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "Diário não encontrado"
}
```

**Resposta de Erro (400 Bad Request):**
```json
{
  "error": "ID deve ser um número válido"
}
```

---

## 📂 Estrutura do Projeto

```
OBRA/
├── cmd/
│   └── main.go                    # Ponto de entrada da aplicação
├── internal/
│   ├── controllers/               # Handlers HTTP (Gin)
│   │   ├── diario.go
│   │   ├── obras.go
│   │   ├── pessoa.go
│   │   └── usuario.go
│   ├── models/                    # Estruturas de dados
│   │   ├── diario.go
│   │   ├── obra.go
│   │   ├── pessoa.go
│   │   ├── response.go
│   │   └── usuario.go
│   ├── services/                  # Camada de acesso a dados
│   │   ├── diario.go
│   │   ├── obra.go
│   │   ├── pessoa.go
│   │   └── usuario.go
│   └── usecases/                  # Lógica de negócio
│       ├── diario.go
│       ├── obra.go
│       ├── pessoa.go
│       └── usuario.go
├── migrations/                    # Scripts de migração do banco
│   ├── 000001_create_pessoa.up.sql
│   ├── 000001_create_pessoa.down.sql
│   ├── 000002_create_usuario.up.sql
│   ├── 000002_create_usuario.down.sql
│   ├── 000003_create_obra.up.sql
│   ├── 000003_create_obra.down.sql
│   ├── 000004_create_diario.up.sql
│   ├── 000004_create_diario.down.sql
│   ├── 000005_seed_data.up.sql
│   └── 000005_seed_data.down.sql
├── pkg/
│   └── postgres/                  # Configuração do banco
│       └── postgres.go
├── .env                           # Variáveis de ambiente
├── docker-compose.yml             # Orquestração de containers
├── Dockerfile                     # Imagem da aplicação
├── go.mod                         # Dependências Go
├── go.sum                         # Checksums das dependências
├── Makefile                       # Comandos facilitados
└── README.md                      # Esta documentação
```

---

## 🗄️ Migrations

O projeto usa migrations para versionamento do banco de dados.

### Como Executar as Migrations

#### Opção 1: Usando golang-migrate (Recomendado)

**1. Instalar o golang-migrate:**

```bash
# No Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/migrate

# No macOS
brew install golang-migrate

# Ou usando Go
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

**2. Executar as migrations:**

```bash
# Aplicar todas as migrations (UP)
migrate -path ./migrations -database "postgresql://obras:7894@localhost:5440/obrasdb?sslmode=disable" up

# Reverter última migration (DOWN)
migrate -path ./migrations -database "postgresql://obras:7894@localhost:5440/obrasdb?sslmode=disable" down 1

# Ver status das migrations
migrate -path ./migrations -database "postgresql://obras:7894@localhost:5440/obrasdb?sslmode=disable" version

# Forçar versão específica (use com cuidado)
migrate -path ./migrations -database "postgresql://obras:7894@localhost:5440/obrasdb?sslmode=disable" force 5
```

#### Opção 2: Executar SQL Diretamente no Container

**1. Conectar ao container do PostgreSQL:**

```bash
docker exec -it db_obras psql -U obras -d obrasdb
```

**2. Executar os arquivos SQL manualmente:**

```bash
# Aplicar migration de pessoas
docker exec -i db_obras psql -U obras -d obrasdb < migrations/000001_create_pessoa.up.sql

# Aplicar migration de usuários
docker exec -i db_obras psql -U obras -d obrasdb < migrations/000002_create_usuario.up.sql

# Aplicar migration de obras
docker exec -i db_obras psql -U obras -d obrasdb < migrations/000003_create_obra.up.sql

# Aplicar migration de diários
docker exec -i db_obras psql -U obras -d obrasdb < migrations/000004_create_diario.up.sql

# Aplicar dados de teste (seed)
docker exec -i db_obras psql -U obras -d obrasdb < migrations/000005_seed_data.up.sql
```

**3. Aplicar todas de uma vez:**

```bash
# Aplicar todas as migrations em ordem
for file in migrations/*.up.sql; do
  echo "Aplicando: $file"
  docker exec -i db_obras psql -U obras -d obrasdb < "$file"
done
```

#### Opção 3: Usando Makefile

O projeto já possui um Makefile com comandos prontos:

```bash
# Ver todos os comandos disponíveis
make help

# Instalar golang-migrate
make install-migrate

# Subir apenas o banco de dados
make docker-up

# Executar migrations (requer golang-migrate instalado)
make migrate-up

# Reverter última migration
make migrate-down

# Criar nova migration
make migrate-create NAME=create_nova_tabela

# Rodar a API localmente (sem Docker)
make run
```

#### Opção 4: Script Shell Personalizado

Crie um arquivo `run-migrations.sh`:

```bash
#!/bin/bash

echo "🚀 Iniciando migrations..."

DB_HOST="localhost"
DB_PORT="5440"
DB_USER="obras"
DB_PASSWORD="7894"
DB_NAME="obrasdb"

export PGPASSWORD=$DB_PASSWORD

# Verificar se o banco está acessível
echo "📡 Verificando conexão com o banco..."
until psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c '\q' 2>/dev/null; do
  echo "⏳ Aguardando banco de dados..."
  sleep 2
done

echo "✅ Banco de dados conectado!"

# Aplicar migrations
echo "📦 Aplicando migrations..."

psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f migrations/000001_create_pessoa.up.sql
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f migrations/000002_create_usuario.up.sql
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f migrations/000003_create_obra.up.sql
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f migrations/000004_create_diario.up.sql
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f migrations/000005_seed_data.up.sql

echo "✅ Migrations aplicadas com sucesso!"
```

Depois execute:

```bash
chmod +x run-migrations.sh
./run-migrations.sh
```

### Verificar se as Migrations Foram Aplicadas

```bash
# Conectar ao banco
docker exec -it db_obras psql -U obras -d obrasdb

# Listar todas as tabelas
\dt

# Ver estrutura de uma tabela
\d pessoas
\d usuarios
\d obras
\d diarios_obra

# Sair
\q
```

Você deve ver as seguintes tabelas:
- `pessoas`
- `usuarios`
- `obras`
- `diarios_obra`
- `schema_migrations` (se usar golang-migrate)

### Estrutura das Tabelas

#### Tabela: `pessoas`
```sql
CREATE TABLE pessoas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    tipo_documento VARCHAR(20) NOT NULL,
    documento VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255),
    telefone VARCHAR(20),
    cargo VARCHAR(100),
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Tabela: `usuarios`
```sql
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    nome VARCHAR(255) NOT NULL,
    senha VARCHAR(255) NOT NULL,
    tipo_documento VARCHAR(20) NOT NULL,
    documento VARCHAR(50) NOT NULL UNIQUE,
    telefone VARCHAR(20),
    perfil_acesso VARCHAR(50) NOT NULL,
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Tabela: `obras`
```sql
CREATE TABLE obras (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    contrato_numero VARCHAR(100) NOT NULL UNIQUE,
    contratante_id INTEGER REFERENCES pessoas(id),
    responsavel_id INTEGER REFERENCES usuarios(id),
    data_inicio DATE NOT NULL,
    prazo_dias INTEGER NOT NULL,
    data_fim_prevista DATE,
    orcamento DECIMAL(15, 2),
    status VARCHAR(50) NOT NULL,
    endereco_rua VARCHAR(255),
    endereco_numero VARCHAR(20),
    endereco_bairro VARCHAR(100),
    endereco_cidade VARCHAR(100),
    endereco_estado VARCHAR(2),
    endereco_cep VARCHAR(10),
    observacoes TEXT,
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Tabela: `diarios_obra`
```sql
CREATE TABLE diarios_obra (
    id SERIAL PRIMARY KEY,
    obra_id INTEGER NOT NULL REFERENCES obras(id),
    data DATE NOT NULL,
    periodo VARCHAR(20),
    atividades_realizadas TEXT NOT NULL,
    ocorrencias TEXT,
    observacoes TEXT,
    responsavel_id INTEGER REFERENCES usuarios(id),
    aprovado_por_id INTEGER REFERENCES usuarios(id),
    status_aprovacao VARCHAR(20) DEFAULT 'pendente',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 🔧 Troubleshooting

### Erro: "connection refused"

**Problema**: A API não consegue conectar ao banco de dados.

**Solução**:
1. Verifique se os containers estão rodando: `docker ps`
2. Verifique os logs do banco: `docker logs db_obras`
3. Certifique-se de que `DB_PORT=5432` no `.env`
4. Reinicie os containers: `docker compose down && docker compose up -d`

### Erro: "port is already allocated"

**Problema**: Porta já está em uso.

**Solução**:
1. Mude as portas no `.env`:
   - `DB_HOST_PORT=5441` (ou outra porta disponível)
   - `API_PORT=9091` (ou outra porta disponível)
2. Reinicie: `docker compose down && docker compose up -d`

### Banco de dados não inicia

**Problema**: Container do banco não sobe.

**Solução**:
```bash
# Remover volumes e recomeçar
docker compose down -v
docker compose up -d
```

---

## � Resumo de Rotas da API

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| **Pessoas** |
| GET | `/pessoas` | Listar todas as pessoas |
| GET | `/pessoas/:id` | Buscar pessoa por ID |
| POST | `/pessoas` | Criar nova pessoa |
| PUT | `/pessoas/:id` | Atualizar pessoa |
| DELETE | `/pessoas/:id` | Deletar pessoa |
| **Usuários** |
| GET | `/usuarios` | Listar todos os usuários |
| GET | `/usuarios/:id` | Buscar usuário por ID |
| POST | `/usuarios` | Criar novo usuário |
| PUT | `/usuarios/:id` | Atualizar usuário |
| DELETE | `/usuarios/:id` | Deletar usuário |
| **Obras** |
| GET | `/obras` | Listar todas as obras |
| GET | `/obras/:id` | Buscar obra por ID |
| POST | `/obras` | Criar nova obra |
| PUT | `/obras/:id` | Atualizar obra |
| DELETE | `/obras/:id` | Deletar obra |
| **Diários** |
| GET | `/diarios` | Listar todos os diários |
| GET | `/diarios/:id` | Buscar diário por ID |
| GET | `/diarios/:id/obra` | Buscar diários por obra |
| POST | `/diarios` | Criar novo diário |
| PUT | `/diarios/:id` | Atualizar diário |
| DELETE | `/diarios/:id` | Deletar diário |

---

## 📝 Códigos de Status HTTP

A API utiliza os seguintes códigos de status HTTP:

| Código | Status | Uso |
|--------|--------|-----|
| `200` | OK | Requisição GET ou PUT bem-sucedida com retorno de dados |
| `201` | Created | Recurso criado com sucesso (POST) |
| `204` | No Content | Requisição bem-sucedida sem conteúdo de retorno (DELETE) |
| `400` | Bad Request | Dados inválidos, malformados ou ID inválido |
| `404` | Not Found | Recurso não encontrado |
| `500` | Internal Server Error | Erro interno do servidor |

### Formato de Resposta de Erro

Erros retornam JSON no seguinte formato:

```json
{
  "error": "Descrição do erro"
}
```

**Exemplos:**
- `404 Not Found`: `{"error": "Pessoa não encontrada"}`
- `400 Bad Request`: `{"error": "ID deve ser um número válido"}`
- `500 Internal Server Error`: `{"error": "Erro ao processar requisição"}`

---

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona MinhaFeature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abra um Pull Request

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

---

## 👨‍💻 Autor

**Mark Hiarley**
- GitHub: [@MarkHiarley](https://github.com/MarkHiarley)

---

## 📞 Suporte

Para reportar bugs ou solicitar features, abra uma [issue](https://github.com/MarkHiarley/OBRA/issues) no GitHub.

---

**Última atualização**: 18 de outubro de 2025
