# 🏗️ OBRA - Sistema de Gerenciamento de Obras

API RESTful para gerenciamento completo de obras, construída em Go com Gin Framework e PostgreSQL.

## 🚀 Quick Start

```bash
# 1. Clone e configure
git clone https://github.com/MarkHiarley/OBRA.git
cd OBRA

# 2. Configure as variáveis de ambiente
cp .env.example .env

# 3. Inicie os containers
docker compose up -d

# 4. Execute as migrations
chmod +x run-migrations.sh
./run-migrations.sh

# 5. Acesse a API
curl http://localhost:9090/pessoas
```

Pronto! A API está rodando em `http://localhost:9090` 🎉

---

## � Funcionalidades

O sistema OBRA oferece controle completo de obras com:

- **🔐 Autenticação JWT** - Login seguro com tokens de acesso
- **👥 Pessoas** - Cadastro de profissionais e contratantes
- **👤 Usuários** - Gestão de acesso ao sistema
- **🏗️ Obras** - Controle de projetos e contratos
- **📖 Diários de Obra** - Registro diário com suporte a fotos base64
- **🏪 Fornecedores** - Cadastro de empresas e prestadores
- **💰 Despesas** - Controle financeiro por categoria
- **💵 Receitas** - Gestão de entradas e receitas das obras
- **📊 Relatórios** - Dashboards financeiros e operacionais completos

---

## 🛠️ Tecnologias

- **Go 1.25** + **Gin Framework**
- **PostgreSQL 12**
- **Docker & Docker Compose**
- **JWT Authentication**
- **Clean Architecture**

---

## ⚙️ Instalação

### Pré-requisitos
- Docker >= 20.10
- Docker Compose >= 2.0

### Configuração

1. **Clone o repositório:**
```bash
git clone https://github.com/MarkHiarley/OBRA.git
cd OBRA
```

2. **Configure o ambiente:**
```bash
# Crie o arquivo .env com suas configurações
cat > .env << EOF
DB_HOST=localhost
DB_PORT=5432
DB_USER=obras
DB_PASSWORD=7894
DB_NAME=obrasdb
DB_HOST_PORT=5440
API_PORT=9090
SECRET_KEY_JWT=OBRAS
EOF
```

3. **Inicie a aplicação:**
```bash
# Subir containers
docker compose up -d

# Aguardar banco inicializar
sleep 10

# Executar migrations
chmod +x run-migrations.sh
./run-migrations.sh

# Verificar logs
docker logs api_obras
```

### Acessos

- **API**: http://localhost:9090
- **PostgreSQL**: localhost:5440 (user: obras, pass: 7894, db: obrasdb)

---

## 📚 Documentação da API

### 🔐 Autenticação

#### Login
```http
POST /login
Content-Type: application/json

{
  "email": "admin@obras.com",
  "senha": "admin123"
}
```

**Resposta:**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

#### Usar Token
```bash
# Todas as rotas protegidas requerem o header:
Authorization: Bearer <access_token>
```

#### Renovar Token
```http
POST /refresh
Content-Type: application/json

{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

---

### 👥 Pessoas

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/pessoas` | Listar todas as pessoas |
| `GET` | `/pessoas/:id` | Buscar pessoa por ID |
| `POST` | `/pessoas` | Criar nova pessoa |
| `PUT` | `/pessoas/:id` | Atualizar pessoa |
| `DELETE` | `/pessoas/:id` | Deletar pessoa |

**Exemplo - Criar Pessoa:**
```http
POST /pessoas
Authorization: Bearer <token>
Content-Type: application/json

{
  "nome": "João Silva",
  "tipo": "CPF",
  "documento": "123.456.789-00",
  "email": "joao@exemplo.com",
  "telefone": "(11) 98765-4321",
  "cargo": "Engenheiro Civil",
  "endereco_rua": "Av. Principal, 1000",
  "endereco_cidade": "São Paulo",
  "endereco_estado": "SP",
  "endereco_cep": "01000-000",
  "ativo": true
}
```

---

### 👤 Usuários

| Método | Endpoint | Descrição | Autenticação |
|--------|----------|-----------|--------------|
| `POST` | `/usuarios` | Cadastrar usuário | ❌ Público |
| `GET` | `/usuarios` | Listar usuários | ✅ Protegido |
| `GET` | `/usuarios/:id` | Buscar usuário por ID | ✅ Protegido |
| `PUT` | `/usuarios/:id` | Atualizar usuário | ✅ Protegido |
| `DELETE` | `/usuarios/:id` | Deletar usuário | ✅ Protegido |

---

### 🏗️ Obras

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/obras` | Listar todas as obras |
| `GET` | `/obras/:id` | Buscar obra por ID |
| `POST` | `/obras` | Criar nova obra |
| `PUT` | `/obras/:id` | Atualizar obra |
| `DELETE` | `/obras/:id` | Deletar obra |

---

### 📖 Diários de Obra

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/diarios` | Listar todos os diários |
| `GET` | `/diarios/:id` | Buscar diário por ID |
| `GET` | `/diarios/obra/:id` | Buscar diários por obra |
| `POST` | `/diarios` | Criar novo diário |
| `PUT` | `/diarios/:id` | Atualizar diário |
| `DELETE` | `/diarios/:id` | Deletar diário |

**🖼️ Suporte a Fotos:**
```json
{
  "obra_id": 1,
  "data": "2025-11-06",
  "periodo": "manha",
  "atividades_realizadas": "Concretagem da laje",
  "foto": "data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg==",
  "responsavel_id": 4,
  "status_aprovacao": "pendente",
  "clima": "ENSOLARADO"
}
```

**Validações:**
- **Período**: `manha`, `tarde`, `noite`, `integral`
- **Clima**: `ENSOLARADO`, `NUBLADO`, `CHUVOSO`, `VENTOSO`, `OUTROS`
- **Status**: `pendente`, `aprovado`, `rejeitado`

---

### 🏪 Fornecedores

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/fornecedores` | Listar fornecedores |
| `GET` | `/fornecedores/:id` | Buscar fornecedor por ID |
| `POST` | `/fornecedores` | Criar fornecedor |
| `PUT` | `/fornecedores/:id` | Atualizar fornecedor |
| `DELETE` | `/fornecedores/:id` | Deletar fornecedor |

---

### � Despesas

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/despesas` | Listar despesas |
| `GET` | `/despesas/:id` | Buscar despesa por ID |
| `GET` | `/despesas/relatorio/:obra_id` | Relatório de despesas por obra |
| `POST` | `/despesas` | Criar despesa |
| `PUT` | `/despesas/:id` | Atualizar despesa |
| `DELETE` | `/despesas/:id` | Deletar despesa |

**Categorias:** `MATERIAL`, `MAO_DE_OBRA`, `TRANSPORTE`, `EQUIPAMENTO`, `ALIMENTACAO`, `OUTROS`
**Formas de Pagamento:** `PIX`, `BOLETO`, `CARTAO_CREDITO`, `TRANSFERENCIA`, `DINHEIRO`
**Status:** `PENDENTE`, `PAGO`, `VENCIDO`, `CANCELADO`

---

### 💵 Receitas

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/receitas` | Listar todas as receitas |
| `GET` | `/receitas/:id` | Buscar receita por ID |
| `GET` | `/receitas/obra/:obra_id` | Buscar receitas por obra |
| `POST` | `/receitas` | Criar nova receita |
| `PUT` | `/receitas/:id` | Atualizar receita |
| `DELETE` | `/receitas/:id` | Deletar receita |

**Exemplo - Criar Receita:**
```http
POST /receitas
Authorization: Bearer <token>
Content-Type: application/json

{
  "obra_id": 5,
  "fonte_receita": "CONTRATO",
  "descricao": "Pagamento inicial do contrato",
  "valor": 50000.00,
  "data_recebimento": "2025-11-06",
  "responsavel_id": 4,
  "observacoes": "Primeira parcela do contrato"
}
```

**Fontes de Receita:**
- `CONTRATO` - Pagamentos contratuais
- `PAGAMENTO_CLIENTE` - Pagamentos de clientes
- `ADIANTAMENTO` - Adiantamentos recebidos
- `FINANCIAMENTO` - Financiamentos obtidos
- `MEDICAO` - Pagamentos por medição
- `OUTROS` - Outras receitas

---

### � Relatórios

#### Relatório de Obra
```http
GET /relatorios/obra/:obra_id
```
**Retorna:** Orçamento vs Gasto vs Receita, Saldo Atual, Percentual de Lucro

#### Relatório de Despesas por Categoria
```http
GET /relatorios/despesas/:obra_id
```
**Retorna:** Despesas agrupadas por categoria com totais e percentuais

#### Relatório de Pagamentos
```http
GET /relatorios/pagamentos/:obra_id?status=PENDENTE
```
**Retorna:** Status de pagamentos, dias de atraso, formas de pagamento

#### Relatório de Materiais
```http
GET /relatorios/materiais/:obra_id
```
**Retorna:** Total gasto em materiais, quantidade de itens, maior gasto

#### Relatório de Profissionais
```http
GET /relatorios/profissionais/:obra_id
```
**Retorna:** Total mão de obra, quantidade de pagamentos, maior pagamento

**Exemplo de Resposta - Relatório de Obra:**
```json
{
  "data": {
    "obra_id": 5,
    "orcamento_previsto": 0,
    "gasto_realizado": 1750,
    "receita_total": 50000,
    "saldo_atual": 48250,
    "percentual_executado": 3.5,
    "percentual_lucro": 96.5,
    "status_financeiro": "LUCRO"
  }
}
```

---

## � Comandos Úteis

### Docker
```bash
# Iniciar aplicação
docker compose up -d

# Ver logs
docker logs api_obras -f
docker logs db_obras -f

# Rebuild após mudanças
docker compose down
docker compose up -d --build

# Parar aplicação
docker compose down
```

### Banco de Dados
```bash
# Conectar ao PostgreSQL
docker exec -it db_obras psql -U obras -d obrasdb

# Backup
docker exec db_obras pg_dump -U obras obrasdb > backup.sql

# Executar migrations
./run-migrations.sh
```

### Desenvolvimento
```bash
# Rodar localmente (sem Docker)
go run cmd/main.go

# Build da aplicação
go build ./...

# Testes
go test ./...
```

---

## 📁 Estrutura do Projeto

```
OBRA/
├── cmd/main.go                    # Ponto de entrada
├── internal/
│   ├── auth/                      # JWT e middleware
│   ├── controllers/               # Handlers HTTP
│   ├── models/                    # Estruturas de dados
│   ├── services/                  # Acesso ao banco
│   └── usecases/                  # Lógica de negócio
├── migrations/                    # Scripts SQL
├── pkg/postgres/                  # Configuração DB
├── docker-compose.yml             # Orquestração
├── Dockerfile                     # Imagem da API
└── .env                          # Variáveis de ambiente
```

---

## 🤝 Contribuição

1. Fork o projeto
2. Crie sua feature branch (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -am 'Add nova funcionalidade'`)
4. Push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

---

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## 📧 Contato

- **GitHub**: [@MarkHiarley](https://github.com/MarkHiarley)
- **Email**: markhiarley@exemplo.com

---

**🏗️ OBRA - Construindo o futuro da gestão de obras! 🚀**

---

## 📚 Documentação da API

Base URL: `http://localhost:9090`

### � Índice de Endpoints

- [� Autenticação](#-autenticação) - Login e renovação de tokens JWT
- [�👥 Pessoas](#-pessoas) - Gerenciamento de pessoas (contratantes, profissionais)
- [👤 Usuários](#-usuários) - Gerenciamento de usuários do sistema
- [🏗️ Obras](#️-obras) - Gerenciamento de obras e contratos
- [📖 Diários de Obra](#-diários-de-obra) - Registro diário de atividades (com suporte a fotos base64)
- [🏪 Fornecedores](#-fornecedores) - Gerenciamento de fornecedores e prestadores
- [💰 Despesas](#-despesas) - Controle financeiro e relatórios
- [💵 Receitas](#-receitas) - Gerenciamento de receitas e entradas financeiras das obras
- [📊 Relatórios](#-relatórios) - Sistema completo de relatórios financeiros e operacionais

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

## 🔐 Autenticação

A API utiliza **JWT (JSON Web Tokens)** para autenticação. Existem dois tipos de tokens:

- **Access Token**: Válido por 15 minutos, usado em todas as requisições protegidas
- **Refresh Token**: Válido por 7 dias, usado para renovar o access token

### Fluxo de Autenticação

```
1. Login → Recebe access_token + refresh_token
2. Usa access_token em requisições (header Authorization: Bearer <token>)
3. Quando access_token expirar (401) → Usa refresh_token para renovar
4. Recebe novos tokens → Continua usando a API
```

---

### Login

```http
POST /login
```

**Body:**
```json
{
  "email": "usuario@exemplo.com",
  "senha": "senha123"
}
```

**Resposta (200 OK):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Resposta de Erro (401 Unauthorized):**
```json
{
  "error": "credenciais inválidas"
}
```

---

### Renovar Token

```http
POST /refresh
```

**Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Resposta (200 OK):**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Resposta de Erro (401 Unauthorized):**
```json
{
  "error": "Refresh token inválido ou expirado"
}
```

---

### Como Usar os Tokens

Todas as rotas protegidas requerem o access token no header `Authorization`:

```bash
curl -X GET http://localhost:9090/pessoas \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

**Exemplo de Fluxo Completo:**

```bash
# 1. Fazer login
curl -X POST http://localhost:9090/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@email.com", "senha": "senha123"}'

# Resposta: { "access_token": "...", "refresh_token": "..." }

# 2. Usar o access_token nas requisições
curl -X GET http://localhost:9090/pessoas \
  -H "Authorization: Bearer <access_token>"

# 3. Quando o access_token expirar (após 15 min), renovar:
curl -X POST http://localhost:9090/refresh \
  -H "Content-Type: application/json" \
  -d '{"refresh_token": "<refresh_token>"}'

# 4. Usar o novo access_token
curl -X GET http://localhost:9090/pessoas \
  -H "Authorization: Bearer <novo_access_token>"
```

---

### Rotas Públicas vs Protegidas

| Tipo | Rotas | Autenticação |
|------|-------|--------------|
| **Públicas** | `POST /login`, `POST /usuarios`, `POST /refresh` | ❌ Não requer token |
| **Protegidas** | Todas as outras rotas (GET, PUT, DELETE) | ✅ Requer access token |

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
      "endereco_rua": "Av. Principal",
      "endereco_numero": "1000",
      "endereco_complemento": null,
      "endereco_bairro": "Centro",
      "endereco_cidade": "São Paulo",
      "endereco_estado": "SP",
      "endereco_cep": "01000-000",
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
  "endereco_rua": "Av. Principal",
  "endereco_numero": "1000",
  "endereco_complemento": null,
  "endereco_bairro": "Centro",
  "endereco_cidade": "São Paulo",
  "endereco_estado": "SP",
  "endereco_cep": "01000-000",
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
  "endereco_rua": "Rua das Flores",
  "endereco_numero": "123",
  "endereco_complemento": "Apto 12",
  "endereco_bairro": "Jardim",
  "endereco_cidade": "São Paulo",
  "endereco_estado": "SP",
  "endereco_cep": "02000-000",
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
  "endereco_rua": "Rua das Flores",
  "endereco_numero": "123",
  "endereco_complemento": "Apto 12",
  "endereco_bairro": "Jardim",
  "endereco_cidade": "São Paulo",
  "endereco_estado": "SP",
  "endereco_cep": "02000-000",
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

> 📌 **Nota de Autenticação:**  
> - `POST /usuarios` (cadastro) é **público** - não requer token
> - Todas as outras operações (GET, PUT, DELETE) são **protegidas** - requerem token JWT

#### Cadastrar novo usuário

```http
POST /usuarios
```

**Autenticação:** ❌ Pública (não requer token)

**Body:**
```json
{
  "email": "novo@obra.com",
  "nome": "Novo Usuário",
  "senha": "senha123",
  "tipo_documento": "CPF",
  "documento": "123.456.789-00",
  "telefone": "(11) 98765-4321",
  "perfil_acesso": "usuario",
  "ativo": true
}
```

**Resposta (201 Created):**
```json
{
  "id": 1,
  "email": "novo@obra.com",
  "nome": "Novo Usuário",
  "tipo_documento": "CPF",
  "documento": "123.456.789-00",
  "telefone": "(11) 98765-4321",
  "perfil_acesso": "usuario",
  "ativo": true,
  "createdAt": "2025-10-19T10:00:00Z",
  "updatedAt": "2025-10-19T10:00:00Z"
}
```

> 💡 **Dica:** Após cadastrar, use `POST /login` com o email e senha para obter os tokens JWT.

---

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
      "art": null,
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
  "art": null,
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
      "status_aprovacao": "APROVADO",
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
  "status_aprovacao": "APROVADO",
  "createdAt": "2025-10-15T18:00:00Z",
  "updatedAt": "2025-10-15T19:00:00Z"
}
```

#### Buscar diários por obra
```http
GET /diarios/obra/:id
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
  "status_aprovacao": "PENDENTE",
  "clima": "ENSOLARADO",
  "progresso_percentual": 10.5,
  "foto": "data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg=="
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
    "foto": "data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg==",
    "responsavel_id": 2,
    "aprovado_por_id": null,
    "status_aprovacao": "PENDENTE",
    "clima": "ENSOLARADO",
    "progresso_percentual": 10.5,
    "createdAt": "2025-10-16T19:00:00Z",
    "updatedAt": "2025-10-16T19:00:00Z"
  }
}
```

**Validações e Enums:**

**Período:**
- `manha` - Período da manhã
- `tarde` - Período da tarde  
- `noite` - Período noturno
- `integral` - Dia integral

**Clima:**
- `ENSOLARADO` - Dia ensolarado
- `NUBLADO` - Dia nublado
- `CHUVOSO` - Dia chuvoso
- `VENTOSO` - Dia ventoso
- `OUTROS` - Outras condições

**Status de Aprovação:**
- `pendente` - Aguardando aprovação
- `aprovado` - Aprovado
- `rejeitado` - Rejeitado

**Campo Foto:**
- Suporte a imagens em formato base64
- Formato aceito: `data:image/[tipo];base64,[dados]`
- Exemplo: `data:image/jpeg;base64,/9j/4AAQSkZJRgABA...`
- Campo opcional (nullable)

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
  "status_aprovacao": "APROVADO"
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

### 🏪 Fornecedores

#### Listar todos os fornecedores
```http
GET /fornecedores
```

**Resposta (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "nome": "Materiais Silva LTDA",
      "tipo_documento": "CNPJ",
      "documento": "12.345.678/0001-90",
      "email": "contato@materiaissilva.com.br",
      "telefone": "(11) 98765-4321",
      "endereco": "Av. das Construções, 1000",
      "cidade": "São Paulo",
      "estado": "SP",
      "ativo": true,
      "created_at": "2025-10-16T10:00:00Z",
      "updated_at": "2025-10-16T10:00:00Z"
    }
  ]
}
```

#### Buscar fornecedor por ID
```http
GET /fornecedores/:id
```

**Parâmetros:**
- `id` (path): ID do fornecedor

**Resposta (200 OK):**
```json
{
  "id": 1,
  "nome": "Materiais Silva LTDA",
  "tipo_documento": "CNPJ",
  "documento": "12.345.678/0001-90",
  "email": "contato@materiaissilva.com.br",
  "telefone": "(11) 98765-4321",
  "endereco": "Av. das Construções, 1000",
  "cidade": "São Paulo",
  "estado": "SP",
  "ativo": true,
  "created_at": "2025-10-16T10:00:00Z",
  "updated_at": "2025-10-16T10:00:00Z"
}
```

#### Criar novo fornecedor
```http
POST /fornecedores
```

**Body:**
```json
{
  "nome": "Ferragens Moderna",
  "tipo_documento": "CNPJ",
  "documento": "98.765.432/0001-10",
  "email": "vendas@ferragensmoderna.com",
  "telefone": "(11) 91234-5678",
  "endereco": "Rua dos Materiais, 500",
  "cidade": "São Paulo",
  "estado": "SP",
  "contato_nome": "Rafael Souza",
  "contato_telefone": "(11) 91234-0000",
  "contato_email": "rafael@ferragensmoderna.com",
  "ativo": true
}
```

**Resposta (201 Created):**
```json
{
  "message": "Fornecedor criado com sucesso",
  "data": {
    "id": 2,
    "nome": "Ferragens Moderna",
    "tipo_documento": "CNPJ",
    "documento": "98.765.432/0001-10",
    "email": "vendas@ferragensmoderna.com",
    "telefone": "(11) 91234-5678",
    "endereco": "Rua dos Materiais, 500",
    "cidade": "São Paulo",
    "estado": "SP",
    "ativo": true,
    "created_at": "2025-10-16T11:00:00Z",
    "updated_at": "2025-10-16T11:00:00Z"
  }
}
```

#### Atualizar fornecedor
```http
PUT /fornecedores/:id
```

**Parâmetros:**
- `id` (path): ID do fornecedor

**Body:**
```json
{
  "nome": "Ferragens Moderna LTDA",
  "tipo_documento": "CNPJ",
  "documento": "98.765.432/0001-10",
  "email": "comercial@ferragensmoderna.com",
  "telefone": "(11) 91234-5678",
  "endereco": "Rua dos Materiais, 500 - Sala 2",
  "cidade": "São Paulo",
  "estado": "SP",
  "ativo": true
}
```

**Resposta (200 OK):**
```json
{
  "id": 2,
  "nome": "Ferragens Moderna LTDA",
  "tipo_documento": "CNPJ",
  "documento": "98.765.432/0001-10",
  "email": "comercial@ferragensmoderna.com",
  "telefone": "(11) 91234-5678",
  "endereco": "Rua dos Materiais, 500 - Sala 2",
  "cidade": "São Paulo",
  "estado": "SP",
  "ativo": true,
  "created_at": "2025-10-16T11:00:00Z",
  "updated_at": "2025-10-16T12:00:00Z"
}
```

#### Deletar fornecedor
```http
DELETE /fornecedores/:id
```

**Parâmetros:**
- `id` (path): ID do fornecedor

**Resposta (204 No Content):**
```
(sem corpo de resposta)
```

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "Fornecedor não encontrado"
}
```

---

### 💰 Despesas

#### Listar todas as despesas
```http
GET /despesas
```

**Resposta (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "obra_id": 1,
      "fornecedor_id": 2,
      "descricao": "Compra de cimento Portland",
      "categoria": "MATERIAL",
      "valor": 1500.00,
      "data_vencimento": "2025-11-15",
      "data_pagamento": null,
      "forma_pagamento": "BOLETO",
      "status_pagamento": "PENDENTE",
      "observacoes": "Entrega prevista para 10/11",
      "created_at": "2025-10-16T14:00:00Z",
      "updated_at": "2025-10-16T14:00:00Z"
    }
  ]
}
```

#### Buscar despesa por ID
```http
GET /despesas/:id
```

**Parâmetros:**
- `id` (path): ID da despesa

**Resposta (200 OK):**
```json
{
  "id": 1,
  "obra_id": 1,
  "fornecedor_id": 2,
  "descricao": "Compra de cimento Portland",
  "categoria": "MATERIAL",
  "valor": 1500.00,
  "data_vencimento": "2025-11-15",
  "data_pagamento": null,
  "forma_pagamento": "BOLETO",
  "status_pagamento": "PENDENTE",
  "observacoes": "Entrega prevista para 10/11",
  "created_at": "2025-10-16T14:00:00Z",
  "updated_at": "2025-10-16T14:00:00Z"
}
```

#### Criar nova despesa
```http
POST /despesas
```

**Body:**
Note: the API accepts either `data` (date of the expense) or `data_vencimento` (due date). If `data` is omitted and `data_vencimento` is provided, the server will fallback to using `data_vencimento` as `data`.

```json
{
  "obra_id": 1,
  "fornecedor_id": 3,
  "descricao": "Pagamento de pedreiros - semana 42",
  "categoria": "MAO_DE_OBRA",
  "valor": 2800.00,
  "data_vencimento": "2025-10-25T00:00:00Z",
  "forma_pagamento": "PIX",
  "status_pagamento": "PENDENTE",
  "observacoes": "Pagamento semanal da equipe"
}
```

**Resposta (201 Created):**
```json
{
  "message": "Despesa criada com sucesso",
  "data": {
    "id": 2,
    "obra_id": 1,
    "fornecedor_id": 3,
    "descricao": "Pagamento de pedreiros - semana 42",
    "categoria": "MAO_DE_OBRA",
    "valor": 2800.00,
    "data_vencimento": "2025-10-25",
    "data_pagamento": null,
    "forma_pagamento": "PIX",
    "status_pagamento": "PENDENTE",
    "observacoes": "Pagamento semanal da equipe",
    "created_at": "2025-10-16T15:00:00Z",
    "updated_at": "2025-10-16T15:00:00Z"
  }
}
```

#### Atualizar despesa
```http
PUT /despesas/:id
```

**Parâmetros:**
- `id` (path): ID da despesa

**Body:**
```json
{
  "obra_id": 1,
  "fornecedor_id": 3,
  "descricao": "Pagamento de pedreiros - semana 42",
  "categoria": "MAO_DE_OBRA",
  "valor": 2800.00,
  "data_vencimento": "2025-10-25",
  "data_pagamento": "2025-10-24",
  "forma_pagamento": "PIX",
  "status_pagamento": "PAGO",
  "observacoes": "Pagamento realizado via PIX"
}
```

**Resposta (200 OK):**
```json
{
  "id": 2,
  "obra_id": 1,
  "fornecedor_id": 3,
  "descricao": "Pagamento de pedreiros - semana 42",
  "categoria": "MAO_DE_OBRA",
  "valor": 2800.00,
  "data_vencimento": "2025-10-25",
  "data_pagamento": "2025-10-24",
  "forma_pagamento": "PIX",
  "status_pagamento": "PAGO",
  "observacoes": "Pagamento realizado via PIX",
  "created_at": "2025-10-16T15:00:00Z",
  "updated_at": "2025-10-16T16:00:00Z"
}
```

#### Deletar despesa
```http
DELETE /despesas/:id
```

**Parâmetros:**
- `id` (path): ID da despesa

**Resposta (204 No Content):**
```
(sem corpo de resposta)
```

#### Relatório de despesas por obra
```http
GET /despesas/relatorio/:obra_id
```

**Parâmetros:**
- `obra_id` (path): ID da obra

**Resposta (200 OK):**
```json
{
  "obra_id": 1,
  "totais_por_categoria": {
    "MATERIAL": 15750.00,
    "MAO_DE_OBRA": 8400.00,
    "TRANSPORTE": 1200.00,
    "EQUIPAMENTO": 3500.00
  },
  "total_geral": 28850.00,
  "despesas": [
    {
      "id": 1,
      "descricao": "Compra de cimento Portland",
      "categoria": "MATERIAL",
      "valor": 1500.00,
      "fornecedor": "Materiais Silva LTDA",
      "status_pagamento": "PENDENTE",
      "data_vencimento": "2025-11-15"
    },
    {
      "id": 2,
      "descricao": "Pagamento de pedreiros - semana 42",
      "categoria": "MAO_DE_OBRA",
      "valor": 2800.00,
      "fornecedor": "João Pedreiro",
      "status_pagamento": "PAGO",
      "data_pagamento": "2025-10-24"
    }
  ]
}
```

#### Enums e Validações

**Categorias de Despesa:**
- `MATERIAL` - Materiais de construção
- `MAO_DE_OBRA` - Mão de obra e serviços
- `TRANSPORTE` - Fretes e transportes
- `EQUIPAMENTO` - Aluguel de equipamentos
- `ALIMENTACAO` - Alimentação da equipe
- `OUTROS` - Outras despesas

**Formas de Pagamento:**
- `PIX` - Transferência PIX
- `BOLETO` - Boleto bancário
- `CARTAO_CREDITO` - Cartão de crédito
- `CARTAO_DEBITO` - Cartão de débito
- `TRANSFERENCIA` - Transferência bancária
- `DINHEIRO` - Dinheiro
- `CHEQUE` - Cheque

**Status de Pagamento:**
- `PENDENTE` - Aguardando pagamento
- `PAGO` - Pagamento realizado
- `VENCIDO` - Pagamento em atraso
- `CANCELADO` - Despesa cancelada

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "Despesa não encontrada"
}
```

---

### 💵 Receitas

#### Listar todas as receitas
```http
GET /receitas
```

**Resposta (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "obra_id": 5,
      "descricao": "Pagamento inicial da obra",
      "valor": 50000.00,
      "data": "2025-11-06",
      "fonte_receita": "CONTRATO",
      "numero_documento": "CONTR-2024-001",
      "responsavel_id": 4,
      "observacao": "Primeira parcela do contrato",
      "created_at": "2025-11-06T16:28:24Z",
      "updated_at": "2025-11-06T16:28:24Z",
      "obra_nome": "Casa Residencial - Fortaleza",
      "responsavel_nome": "João Silva"
    }
  ]
}
```

#### Buscar receita por ID
```http
GET /receitas/:id
```

**Parâmetros:**
- `id` (path): ID da receita

**Resposta (200 OK):**
```json
{
  "id": 1,
  "obra_id": 5,
  "descricao": "Pagamento inicial da obra",
  "valor": 50000.00,
  "data": "2025-11-06",
  "fonte_receita": "CONTRATO",
  "numero_documento": "CONTR-2024-001",
  "responsavel_id": 4,
  "observacao": "Primeira parcela do contrato",
  "created_at": "2025-11-06T16:28:24Z",
  "updated_at": "2025-11-06T16:28:24Z",
  "obra_nome": "Casa Residencial - Fortaleza",
  "responsavel_nome": "João Silva"
}
```

#### Buscar receitas por obra
```http
GET /receitas/obra/:obra_id
```

**Parâmetros:**
- `obra_id` (path): ID da obra

**Resposta (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "obra_id": 5,
      "descricao": "Pagamento inicial da obra",
      "valor": 50000.00,
      "data": "2025-11-06",
      "fonte_receita": "CONTRATO",
      "numero_documento": "CONTR-2024-001",
      "responsavel_id": 4,
      "observacao": "Primeira parcela do contrato",
      "created_at": "2025-11-06T16:28:24Z",
      "updated_at": "2025-11-06T16:28:24Z",
      "obra_nome": "Casa Residencial - Fortaleza",
      "responsavel_nome": "João Silva"
    }
  ]
}

#### Criar nova receita
```http
POST /receitas
```

**Body:**
```json
{
  "obra_id": 5,
  "fonte_receita": "CONTRATO",
  "descricao": "Pagamento inicial do contrato",
  "valor": 50000.00,
  "data_recebimento": "2025-11-06",
  "numero_documento": "CONTR-2024-001",
  "responsavel_id": 4,
  "observacoes": "Primeira parcela do contrato"
}
```

**Resposta (201 Created):**
```json
{
  "message": "Receita criada com sucesso",
  "data": {
    "id": 1,
    "obra_id": 5,
    "fonte_receita": "CONTRATO",
    "descricao": "Pagamento inicial do contrato",
    "valor": 50000.00,
    "data": "2025-11-06",
    "numero_documento": "CONTR-2024-001",
    "responsavel_id": 4,
    "observacao": "Primeira parcela do contrato",
    "created_at": "2025-11-06T16:28:24Z",
    "updated_at": "2025-11-06T16:28:24Z"
  }
}
```

#### Atualizar receita
```http
PUT /receitas/:id
```

**Parâmetros:**
- `id` (path): ID da receita

**Body:**
```json
{
  "obra_id": 5,
  "fonte_receita": "CONTRATO",
  "descricao": "Pagamento inicial do contrato - Atualizado",
  "valor": 55000.00,
  "data_recebimento": "2025-11-06",
  "numero_documento": "CONTR-2024-001-UPDATED",
  "responsavel_id": 4,
  "observacoes": "Primeira parcela do contrato com ajuste"
}
```

**Resposta (200 OK):**
```json
{
  "id": 1,
  "obra_id": 5,
  "fonte_receita": "CONTRATO",
  "descricao": "Pagamento inicial do contrato - Atualizado",
  "valor": 55000.00,
  "data": "2025-11-06",
  "numero_documento": "CONTR-2024-001-UPDATED",
  "responsavel_id": 4,
  "observacao": "Primeira parcela do contrato com ajuste",
  "created_at": "2025-11-06T16:28:24Z",
  "updated_at": "2025-11-06T17:30:15Z"
}
```

#### Deletar receita
```http
DELETE /receitas/:id
```

**Parâmetros:**
- `id` (path): ID da receita

**Resposta (204 No Content):**
```
(sem corpo de resposta)
```

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "Receita não encontrada"
}
```

**Fontes de Receita:**
- `CONTRATO` - Pagamentos contratuais
- `PAGAMENTO_CLIENTE` - Pagamentos de clientes
- `ADIANTAMENTO` - Adiantamentos recebidos
- `FINANCIAMENTO` - Financiamentos obtidos
- `MEDICAO` - Pagamentos por medição
- `OUTROS` - Outras receitas

---

### 📊 Relatórios

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/relatorios/obra/:obra_id` | Relatório financeiro completo da obra |
| `GET` | `/relatorios/despesas/:obra_id` | Despesas agrupadas por categoria |
| `GET` | `/relatorios/pagamentos/:obra_id` | Status de pagamentos e atrasos |
| `GET` | `/relatorios/materiais/:obra_id` | Relatório de materiais consumidos |
| `GET` | `/relatorios/profissionais/:obra_id` | Relatório de mão de obra |

#### Relatório de Obra
```http
GET /relatorios/obra/:obra_id
```

**Descrição:** Relatório financeiro completo da obra com orçamento vs gasto vs receita.

**Resposta (200 OK):**
```json
{
  "data": {
    "obra_id": 5,
    "obra_nome": "Casa Residencial - Fortaleza",
    "orcamento_previsto": 0,
    "gasto_realizado": 1750,
    "receita_total": 50000,
    "saldo_atual": 48250,
    "pagamento_pendente": 2700.5,
    "status": "EM_ANDAMENTO",
    "percentual_executado": 3.5,
    "percentual_lucro": 96.5,
    "total_despesas": 5,
    "total_receitas": 1
  }
}
```

#### Relatório de Despesas por Categoria
```http
GET /relatorios/despesas/:obra_id
```

**Descrição:** Despesas agrupadas por categoria com totais e percentuais.

**Resposta (200 OK):**
```json
{
  "data": {
    "obra_id": 5,
    "obra_nome": "Casa Residencial - Fortaleza",
    "total_geral": 4450.5,
    "categorias": [
      {
        "categoria": "MATERIAL",
        "total": 3500.5,
        "percentual": 78.6,
        "quantidade_itens": 2
      },
      {
        "categoria": "MAO_DE_OBRA", 
        "total": 750,
        "percentual": 16.9,
        "quantidade_itens": 1
      },
      {
        "categoria": "OUTROS",
        "total": 200,
        "percentual": 4.5,
        "quantidade_itens": 1
      }
    ]
  }
}
```

#### Relatório de Pagamentos
```http
GET /relatorios/pagamentos/:obra_id?status=PENDENTE
```

**Parâmetros Query (opcionais):**
- `status` - Filtrar por status: `PENDENTE`, `PAGO`, `VENCIDO`

**Descrição:** Status de pagamentos com dias de atraso e formas de pagamento.

**Resposta (200 OK):**
```json
{
  "data": {
    "obra_id": 5,
    "total_pendente": 2700.5,
    "total_pago": 1750,
    "pagamentos_em_atraso": 1,
    "pagamentos": [
      {
        "id": 3,
        "descricao": "Ferro 10mm - 50 barras",
        "valor": 2500.5,
        "status_pagamento": "PENDENTE",
        "forma_pagamento": "BOLETO",
        "data_vencimento": "2025-11-03",
        "dias_atraso": 3,
        "fornecedor_nome": "Distribuidora ABC"
      }
    ]
  }
}
```

#### Relatório de Materiais
```http
GET /relatorios/materiais/:obra_id
```

**Descrição:** Total gasto em materiais, quantidade de itens e maior gasto.

**Resposta (200 OK):**
```json
{
  "data": {
    "total_materiais": 3500.5,
    "quantidade_itens": 2,
    "maior_gasto_valor": 2500.5,
    "maior_gasto_descricao": "Ferro 10mm - 50 barras",
    "materiais": [
      {
        "descricao": "Ferro 10mm - 50 barras",
        "valor": 2500.5,
        "data": "2025-11-06",
        "fornecedor": "Distribuidora ABC"
      },
      {
        "descricao": "Cimento CP-II 50kg - 20 sacos",
        "valor": 1000,
        "data": "2025-11-06", 
        "fornecedor": "Materiais Silva"
      }
    ]
  }
}
```

#### Relatório de Profissionais
```http
GET /relatorios/profissionais/:obra_id
```

**Descrição:** Total de mão de obra, quantidade de pagamentos e maior pagamento.

**Resposta (200 OK):**
```json
{
  "data": {
    "total_mao_obra": 750,
    "quantidade_pagamentos": 1,
    "maior_pagamento_valor": 750,
    "maior_pagamento_descricao": "Pedreiro 5 dias",
    "profissionais": [
      {
        "descricao": "Pedreiro 5 dias",
        "valor": 750,
        "data": "2025-11-06",
        "responsavel": "João da Silva"
      }
    ]
  }
}
```

---

## 📂 Estrutura do Projeto

```
OBRA/
├── cmd/
│   └── main.go                    # Ponto de entrada da aplicação
├── internal/
│   ├── auth/                      # 🔐 Autenticação e Autorização (NOVO)
│   │   ├── jwt.go                 # Geração e validação de tokens JWT
│   │   └── middleware.go          # Middleware de autenticação
│   ├── controllers/               # Handlers HTTP (Gin)
│   │   ├── diario.go
│   │   ├── login.go               # 🆕 Controller de login
│   │   ├── obras.go
│   │   ├── pessoa.go
│   │   └── usuario.go
│   ├── models/                    # Estruturas de dados
│   │   ├── Claims.go              # 🆕 JWT Claims
│   │   ├── diario.go
│   │   ├── login.go               # 🆕 Model de login
│   │   ├── obra.go
│   │   ├── pessoa.go
│   │   ├── response.go
│   │   └── usuario.go
│   ├── services/                  # Camada de acesso a dados
│   │   ├── diario.go
│   │   ├── login.go               # 🆕 Service de autenticação
│   │   ├── obra.go
│   │   ├── pessoa.go
│   │   └── usuario.go
│   └── usecases/                  # Lógica de negócio
│       ├── diario.go
│       ├── login.go               # 🆕 UseCase de login
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
├── .env                           # Variáveis de ambiente (SECRET_KEY_JWT)
├── .env.example                   # 🆕 Exemplo de variáveis de ambiente
├── docker-compose.yml             # Orquestração de containers
├── Dockerfile                     # Imagem da aplicação
├── go.mod                         # Dependências Go
├── go.sum                         # Checksums das dependências
├── Makefile                       # Comandos facilitados
└── README.md                      # Esta documentação
```

### 🔐 Novos Componentes de Autenticação

| Arquivo | Responsabilidade |
|---------|------------------|
| `internal/auth/jwt.go` | Geração de access_token e refresh_token, validação de tokens JWT |
| `internal/auth/middleware.go` | Middleware que protege rotas, valida tokens e injeta claims no contexto |
| `internal/controllers/login.go` | Handler HTTP para `/login` e `/refresh` |
| `internal/usecases/login.go` | Lógica de validação de credenciais e geração de tokens |
| `internal/services/login.go` | Busca usuário no banco de dados por email |
| `internal/models/login.go` | Estrutura de request de login (email + senha) |
| `internal/models/Claims.go` | Estrutura de claims JWT (email, expiração, etc.) |
| `.env.example` | Template de variáveis de ambiente |

### 📊 Fluxo de Autenticação

```
┌─────────────────────────────────────────────────────────────┐
│                   Fluxo de Autenticação                     │
└─────────────────────────────────────────────────────────────┘

1. POST /login
   ├─> LoginController.CreateLogin
   ├─> LoginUseCase.LoginUseCase
   │   ├─> LoginService.CheckUser (busca hash no banco)
   │   ├─> bcrypt.CompareHashAndPassword (valida senha)
   │   └─> auth.GenerateAccessToken + auth.GenerateRefreshToken
   └─> Retorna: { access_token, refresh_token }

2. POST /refresh
   ├─> LoginController.RefreshToken
   ├─> auth.ValidateToken (valida refresh_token)
   └─> Retorna: { novo_access_token, novo_refresh_token }

3. Rotas Protegidas
   ├─> auth.AuthMiddleware (intercepta requisição)
   ├─> Extrai token do header Authorization
   ├─> auth.ValidateToken (valida access_token)
   ├─> Injeta email no contexto (ctx.Set)
   └─> Chama handler da rota
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

### Notas sobre as migrations recentes

Foram adicionados os seguintes arquivos de migration (UP) ao diretório `migrations/` para corrigir inconsistências detectadas entre frontend e API:

- `000017_fix_diario_aprovador.up.sql` — altera `diario_obra.aprovado_por_id` para permitir NULL e adiciona a constraint `ck_diario_aprovador_status` para validar a relação entre `status_aprovacao` e `aprovado_por_id`.
- `000018_rename_data_despesa_to_data_vencimento.up.sql` — renomeia `despesa.data_despesa` para `despesa.data_vencimento` quando aplicável.
- `000019_add_endereco_pessoa.up.sql` — adiciona colunas de endereço na tabela `pessoa`.
- `000020_add_art_obra.up.sql` — adiciona coluna `art` na tabela `obra`.

Importante: revise os dados existentes antes de aplicar constraints mais restritivas (ex.: cheque por diários com `status_aprovacao = 'PENDENTE'` mas `aprovado_por_id IS NOT NULL`).


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

## 📋 Resumo Completo de Endpoints

### 🔐 Autenticação (Públicas)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `POST` | `/login` | Login e geração de tokens JWT |
| `POST` | `/refresh` | Renovar access token |
| `POST` | `/usuarios` | Cadastrar novo usuário |

### 👥 Pessoas (Protegidas)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/pessoas` | Listar todas as pessoas |
| `GET` | `/pessoas/:id` | Buscar pessoa por ID |
| `POST` | `/pessoas` | Criar nova pessoa |
| `PUT` | `/pessoas/:id` | Atualizar pessoa |
| `DELETE` | `/pessoas/:id` | Deletar pessoa |

### 👤 Usuários (Protegidas)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/usuarios` | Listar todos os usuários |
| `GET` | `/usuarios/:id` | Buscar usuário por ID |
| `PUT` | `/usuarios/:id` | Atualizar usuário |
| `DELETE` | `/usuarios/:id` | Deletar usuário |

### 🏗️ Obras (Protegidas)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/obras` | Listar todas as obras |
| `GET` | `/obras/:id` | Buscar obra por ID |
| `POST` | `/obras` | Criar nova obra |
| `PUT` | `/obras/:id` | Atualizar obra |
| `DELETE` | `/obras/:id` | Deletar obra |

### 📖 Diários de Obra (Protegidas)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/diarios` | Listar todos os diários |
| `GET` | `/diarios/:id` | Buscar diário por ID |
| `GET` | `/diarios/obra/:id` | Buscar diários por obra |
| `POST` | `/diarios` | Criar novo diário (com suporte a foto base64) |
| `PUT` | `/diarios/:id` | Atualizar diário |
| `DELETE` | `/diarios/:id` | Deletar diário |

### 🏪 Fornecedores (Protegidas)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/fornecedores` | Listar todos os fornecedores |
| `GET` | `/fornecedores/:id` | Buscar fornecedor por ID |
| `POST` | `/fornecedores` | Criar novo fornecedor |
| `PUT` | `/fornecedores/:id` | Atualizar fornecedor |
| `DELETE` | `/fornecedores/:id` | Deletar fornecedor |

### 💰 Despesas (Protegidas)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/despesas` | Listar todas as despesas |
| `GET` | `/despesas/:id` | Buscar despesa por ID |
| `GET` | `/despesas/relatorio/:obra_id` | Relatório de despesas por obra |
| `POST` | `/despesas` | Criar nova despesa |
| `PUT` | `/despesas/:id` | Atualizar despesa |
| `DELETE` | `/despesas/:id` | Deletar despesa |

### 💵 Receitas (Protegidas)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/receitas` | Listar todas as receitas |
| `GET` | `/receitas/:id` | Buscar receita por ID |
| `GET` | `/receitas/obra/:obra_id` | Buscar receitas por obra |
| `POST` | `/receitas` | Criar nova receita |
| `PUT` | `/receitas/:id` | Atualizar receita |
| `DELETE` | `/receitas/:id` | Deletar receita |

### 📊 Relatórios (Protegidas)
| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/relatorios/obra/:obra_id` | Relatório financeiro completo da obra |
| `GET` | `/relatorios/despesas/:obra_id` | Despesas agrupadas por categoria |
| `GET` | `/relatorios/pagamentos/:obra_id` | Status de pagamentos e atrasos |
| `GET` | `/relatorios/materiais/:obra_id` | Relatório de materiais consumidos |
| `GET` | `/relatorios/profissionais/:obra_id` | Relatório de mão de obra |

**Total de Endpoints:** 52 endpoints (3 públicos + 49 protegidos)

---

| Método | Endpoint | Autenticação | Descrição |
|--------|----------|--------------|-----------|
| **Autenticação** |
| POST | `/login` | ❌ Pública | Login e geração de tokens JWT |
| POST | `/refresh` | ❌ Pública | Renovar access token |
| **Usuários** |
| POST | `/usuarios` | ❌ Pública | Criar novo usuário (cadastro) |
| GET | `/usuarios` | ✅ Protegida | Listar todos os usuários |
| GET | `/usuarios/:id` | ✅ Protegida | Buscar usuário por ID |
| PUT | `/usuarios/:id` | ✅ Protegida | Atualizar usuário |
| DELETE | `/usuarios/:id` | ✅ Protegida | Deletar usuário |
| **Pessoas** |
| GET | `/pessoas` | ✅ Protegida | Listar todas as pessoas |
| GET | `/pessoas/:id` | ✅ Protegida | Buscar pessoa por ID |
| POST | `/pessoas` | ✅ Protegida | Criar nova pessoa |
| PUT | `/pessoas/:id` | ✅ Protegida | Atualizar pessoa |
| DELETE | `/pessoas/:id` | ✅ Protegida | Deletar pessoa |
| **Obras** |
| GET | `/obras` | ✅ Protegida | Listar todas as obras |
| GET | `/obras/:id` | ✅ Protegida | Buscar obra por ID |
| POST | `/obras` | ✅ Protegida | Criar nova obra |
| PUT | `/obras/:id` | ✅ Protegida | Atualizar obra |
| DELETE | `/obras/:id` | ✅ Protegida | Deletar obra |
| **Diários** |
| GET | `/diarios` | ✅ Protegida | Listar todos os diários |
| GET | `/diarios/:id` | ✅ Protegida | Buscar diário por ID |
| GET | `/diarios/obra/:id` | ✅ Protegida | Buscar diários por obra |
| POST | `/diarios` | ✅ Protegida | Criar novo diário |
| PUT | `/diarios/:id` | ✅ Protegida | Atualizar diário |
| DELETE | `/diarios/:id` | ✅ Protegida | Deletar diário |
| **Fornecedores** |
| GET | `/fornecedores` | ✅ Protegida | Listar todos os fornecedores |
| GET | `/fornecedores/:id` | ✅ Protegida | Buscar fornecedor por ID |
| POST | `/fornecedores` | ✅ Protegida | Criar novo fornecedor |
| PUT | `/fornecedores/:id` | ✅ Protegida | Atualizar fornecedor |
| DELETE | `/fornecedores/:id` | ✅ Protegida | Deletar fornecedor |
| **Despesas** |
| GET | `/despesas` | ✅ Protegida | Listar todas as despesas |
| GET | `/despesas/:id` | ✅ Protegida | Buscar despesa por ID |
| POST | `/despesas` | ✅ Protegida | Criar nova despesa |
| PUT | `/despesas/:id` | ✅ Protegida | Atualizar despesa |
| DELETE | `/despesas/:id` | ✅ Protegida | Deletar despesa |
| GET | `/despesas/relatorio/:obra_id` | ✅ Protegida | Relatório de despesas por obra |

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

## � Exemplos de Uso Completo

### Fluxo Completo: Do Cadastro ao Acesso Protegido

```bash
# 1. Cadastrar novo usuário (PÚBLICO - sem token)
curl -X POST http://localhost:9090/usuarios \
  -H "Content-Type: application/json" \
  -d '{
    "email": "joao@obra.com",
    "nome": "João Silva",
    "senha": "senha123",
    "tipo_documento": "CPF",
    "documento": "123.456.789-00",
    "telefone": "(11) 98765-4321",
    "perfil_acesso": "usuario",
    "ativo": true
  }'

# Resposta: {"id": 1, "email": "joao@obra.com", ...}

# 2. Fazer login para obter tokens
curl -X POST http://localhost:9090/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "joao@obra.com",
    "senha": "senha123"
  }'

# Resposta:
# {
#   "access_token": "eyJhbGc...",
#   "refresh_token": "eyJhbGc..."
# }

# 3. Usar o access_token para criar uma pessoa (PROTEGIDO)
curl -X POST http://localhost:9090/pessoas \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGc..." \
  -d '{
    "nome": "Maria Santos",
    "tipo": "CPF",
    "documento": "987.654.321-00",
    "email": "maria@exemplo.com",
    "telefone": "(11) 91234-5678",
    "cargo": "Arquiteta",
    "ativo": true
  }'

# 4. Listar pessoas (PROTEGIDO)
curl -X GET http://localhost:9090/pessoas \
  -H "Authorization: Bearer eyJhbGc..."

# 5. Se o access_token expirar (após 15 min), renovar:
curl -X POST http://localhost:9090/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "eyJhbGc..."
  }'

# Resposta: novos tokens
# {
#   "access_token": "eyJhbGc...",
#   "refresh_token": "eyJhbGc..."
# }

# 6. Continuar usando a API com o novo access_token
curl -X GET http://localhost:9090/obras \
  -H "Authorization: Bearer <novo_access_token>"
```

### Testando Sem Autenticação (Deve Falhar)

```bash
# Tentar acessar rota protegida sem token
curl -X GET http://localhost:9090/pessoas

# Resposta: 401 Unauthorized
# {"error": "Token não fornecido"}

# Tentar acessar com token inválido
curl -X GET http://localhost:9090/pessoas \
  -H "Authorization: Bearer token_invalido"

# Resposta: 401 Unauthorized
# {"error": "Token inválido ou expirado"}
```

---

## �🤝 Contribuindo

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

**Última atualização**: 19 de outubro de 2025
