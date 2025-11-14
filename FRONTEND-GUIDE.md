# Guia Frontend - API Diário de Obras

**Base URL:** `http://localhost:9090`

---

## 🔐 Autenticação

### Login
```http
POST /login
```

**Payload:**
```json
{
  "email": "admin@teste.com",
  "senha": "123456"
}
```

**Resposta:**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "pessoa_id": 1,
  "tipo_usuario": "administrador"
}
```

**Header para todos os endpoints:**
```http
Authorization: Bearer {access_token}
```

---

## 📝 Tarefas Realizadas

### 1. Criar Tarefa
```http
POST /tarefas
```

**Payload:**
```json
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",
  "descricao": "Concretagem das vigas do 2º pavimento",
  "status": "em_andamento",
  "percentual_conclusao": 85,
  "observacao": "Clima favorável para concretagem"
}
```

**Campos:**
- `periodo`: `"manha"` | `"tarde"` | `"integral"`
- `status`: `"planejada"` | `"em_andamento"` | `"concluida"` | `"cancelada"`
- `percentual_conclusao`: 0-100

**Resposta (201):**
```json
{
  "id": 1,
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",
  "descricao": "Concretagem das vigas do 2º pavimento",
  "status": "em_andamento",
  "percentual_conclusao": 85,
  "observacao": "Clima favorável para concretagem",
  "created_at": "2024-11-14T10:30:00Z",
  "updated_at": "2024-11-14T10:30:00Z"
}
```

---

### 2. Listar Todas as Tarefas
```http
GET /tarefas
```

**Resposta (200):**
```json
[
  {
    "id": 1,
    "obra_id": 5,
    "data": "2024-11-14",
    "periodo": "manha",
    "descricao": "Concretagem das vigas",
    "status": "em_andamento",
    "percentual_conclusao": 85
  },
  {
    "id": 2,
    "obra_id": 5,
    "data": "2024-11-14",
    "periodo": "tarde",
    "descricao": "Montagem de armadura",
    "status": "concluida",
    "percentual_conclusao": 100
  }
]
```

---

### 3. Buscar Tarefas por Obra e Data
```http
GET /tarefas/obra/{obra_id}/data/{data}
```

**Exemplo:**
```http
GET /tarefas/obra/5/data/2024-11-14
```

**Resposta:** Array de tarefas da obra naquela data

---

### 4. Atualizar Tarefa
```http
PUT /tarefas/{id}
```

**Payload:** (mesmo formato do POST)
```json
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",
  "descricao": "Concretagem das vigas do 2º pavimento",
  "status": "concluida",
  "percentual_conclusao": 100,
  "observacao": "Concluída com sucesso"
}
```

---

### 5. Deletar Tarefa
```http
DELETE /tarefas/{id}
```

**Resposta (200):**
```json
{
  "message": "Atividade deletada com sucesso"
}
```

---

## ⚠️ Ocorrências

### 1. Criar Ocorrência
```http
POST /ocorrencias
```

**Payload:**
```json
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",
  "tipo": "seguranca",
  "gravidade": "alta",
  "descricao": "Falta de EPIs na equipe de alvenaria",
  "status_resolucao": "resolvida",
  "acao_tomada": "EPIs fornecidos imediatamente e treinamento realizado"
}
```

**Campos:**
- `periodo`: `"manha"` | `"tarde"` | `"integral"`
- `tipo`: `"seguranca"` | `"qualidade"` | `"prazo"` | `"custo"` | `"clima"` | `"outro"`
- `gravidade`: `"baixa"` | `"media"` | `"alta"` (mapeia para OBSERVACAO/IMPORTANTE/CRITICO)
- `status_resolucao`: `"pendente"` | `"em_analise"` | `"resolvida"` | `"nao_aplicavel"`

**Resposta (201):**
```json
{
  "id": 1,
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",
  "tipo": "seguranca",
  "gravidade": "alta",
  "descricao": "Falta de EPIs na equipe de alvenaria",
  "status_resolucao": "resolvida",
  "acao_tomada": "EPIs fornecidos imediatamente",
  "created_at": "2024-11-14T10:30:00Z"
}
```

---

### 2. Listar Todas as Ocorrências
```http
GET /ocorrencias
```

**Resposta (200):**
```json
[
  {
    "id": 1,
    "obra_id": 5,
    "data": "2024-11-14",
    "periodo": "manha",
    "tipo": "seguranca",
    "gravidade": "alta",
    "descricao": "Falta de EPIs",
    "status_resolucao": "resolvida"
  }
]
```

---

### 3. Buscar Ocorrências por Obra e Data
```http
GET /ocorrencias/obra/{obra_id}/data/{data}
```

**Exemplo:**
```http
GET /ocorrencias/obra/5/data/2024-11-14
```

---

### 4. Buscar por Gravidade
```http
GET /ocorrencias/gravidade/{gravidade}
```

**Exemplo:**
```http
GET /ocorrencias/gravidade/alta
```

**Valores:** `baixa`, `media`, `alta`

---

### 5. Atualizar Ocorrência
```http
PUT /ocorrencias/{id}
```

**Payload:** (mesmo formato do POST)

---

### 6. Deletar Ocorrência
```http
DELETE /ocorrencias/{id}
```

**Resposta (200):**
```json
{
  "message": "Ocorrência deletada com sucesso"
}
```

---

## 📸 Metadados do Diário

### Criar/Atualizar Metadados
```http
POST /diarios-consolidado/metadados
```

**Payload:**
```json
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "integral",
  "foto": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsL...",
  "observacoes": "Dia produtivo. Clima favorável. Equipe trabalhou bem.",
  "status_aprovacao": "aprovado"
}
```

**Campos:**
- `periodo`: `"manha"` | `"tarde"` | `"integral"`
- `foto`: String Base64 no formato `data:image/jpeg;base64,...` ou `""`
- `status_aprovacao`: `"pendente"` | `"aprovado"` | `"reprovado"`

**Resposta (201):**
```json
{
  "id": 1,
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "integral",
  "foto": "data:image/jpeg;base64,...",
  "observacoes": "Dia produtivo",
  "status_aprovacao": "aprovado",
  "created_at": "2024-11-14T18:00:00Z"
}
```

### ⚠️ Importante sobre Fotos:
- Use Base64, **não URL**
- Formato: `data:image/jpeg;base64,{base64_string}`
- Comprima imagens grandes antes de converter

---

## 📊 Diário Consolidado (View)

### 1. Listar Todos os Diários
```http
GET /diarios-consolidado
```

**Resposta (200):**
```json
[
  {
    "obra_id": 5,
    "data": "2024-11-14",
    "periodo": "integral",
    "atividades_agregadas": "Concretagem das vigas (em_andamento - 85%); Montagem de armadura (concluida - 100%)",
    "ocorrencias_agregadas": "[CRITICO] Falta de EPIs - resolvida; [OBSERVACAO] Chuva leve - nao_aplicavel",
    "foto": "data:image/jpeg;base64,...",
    "observacoes": "Dia produtivo",
    "status_aprovacao": "aprovado"
  }
]
```

**Formato dos campos agregados:**
- `atividades_agregadas`: `"desc (status - XX%); desc2 (status - YY%)"`
- `ocorrencias_agregadas`: `"[GRAVIDADE] desc - status; [GRAVIDADE2] desc2 - status"`

---

### 2. Buscar por Obra
```http
GET /diarios-consolidado/obra/{obra_id}
```

**Exemplo:**
```http
GET /diarios-consolidado/obra/5
```

---

### 3. Buscar por Data
```http
GET /diarios-consolidado/data/{data}
```

**Exemplo:**
```http
GET /diarios-consolidado/data/2024-11-14
```

---

## 📄 Relatório Formatado (Endpoint Principal)

### Gerar Relatório Completo
```http
GET /diarios/relatorio-formatado/{obra_id}
```

**Exemplo:**
```http
GET /diarios/relatorio-formatado/5
```

**Resposta (200):**
```json
{
  "informacoes_obra": {
    "titulo": "Casa Residencial - Fortaleza",
    "numero_contrato": "CONTR-2024-001",
    "cliente": "João Silva",
    "responsavel_tecnico": "Eng. Maria Santos",
    "endereco": "Rua das Flores, 123",
    "cidade": "Fortaleza",
    "estado": "CE",
    "data_inicio": "2024-05-01",
    "data_prevista_conclusao": "2024-10-28",
    "prazo_obra": "180 DIAS",
    "tempo_decorrido": "30 DIAS"
  },
  "tarefas_realizadas": [
    {
      "descricao": "Concretagem das vigas do 2º pavimento (em_andamento - 85%)",
      "data": "2024-11-14",
      "periodo": "manha"
    },
    {
      "descricao": "Montagem de armadura da laje (concluida - 100%)",
      "data": "2024-11-14",
      "periodo": "tarde"
    }
  ],
  "ocorrencias": [
    {
      "descricao": "Falta de EPIs na equipe de alvenaria - resolvida",
      "tipo": "CRITICO",
      "data": "2024-11-14"
    },
    {
      "descricao": "Atraso na entrega de material - pendente",
      "tipo": "IMPORTANTE",
      "data": "2024-11-14"
    },
    {
      "descricao": "Chuva leve durante 30 minutos - nao_aplicavel",
      "tipo": "OBSERVACAO",
      "data": "2024-11-14"
    }
  ],
  "fotos": [
    {
      "id": 1,
      "url": "data:image/jpeg;base64,/9j/4AAQSkZJRg...",
      "descricao": "Vista geral da obra",
      "timestamp": "2024-11-14T15:30:00Z",
      "local_foto": null,
      "categoria": "DIARIO"
    }
  ],
  "observacoes_gerais": "Dia produtivo. Clima favorável. Equipe trabalhou bem.",
  "status_aprovacao": "aprovado"
}
```

**Mapeamento de Gravidade:**
| Input (POST) | Output (Relatório) |
|--------------|-------------------|
| `"baixa"`    | `"OBSERVACAO"`    |
| `"media"`    | `"IMPORTANTE"`    |
| `"alta"`     | `"CRITICO"`       |

---

## 🎯 Fluxo Completo

### Salvar um dia de trabalho:

```javascript
// 1. Criar tarefas da manhã
POST /tarefas
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",
  "descricao": "Concretagem das vigas",
  "status": "em_andamento",
  "percentual_conclusao": 85
}

// 2. Criar tarefas da tarde
POST /tarefas
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "tarde",
  "descricao": "Montagem de armadura",
  "status": "concluida",
  "percentual_conclusao": 100
}

// 3. Registrar ocorrências (se houver)
POST /ocorrencias
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",
  "tipo": "seguranca",
  "gravidade": "alta",
  "descricao": "Falta de EPIs",
  "status_resolucao": "resolvida",
  "acao_tomada": "EPIs fornecidos"
}

// 4. Adicionar foto e observações
POST /diarios-consolidado/metadados
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "integral",
  "foto": "data:image/jpeg;base64,...",
  "observacoes": "Dia produtivo",
  "status_aprovacao": "aprovado"
}

// 5. Gerar relatório completo
GET /diarios/relatorio-formatado/5
```

---

## 🖼️ Converter Imagem para Base64

### JavaScript (Browser):
```javascript
function imageToBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onloadend = () => resolve(reader.result);
    reader.onerror = reject;
    reader.readAsDataURL(file);
  });
}

// Uso:
const file = event.target.files[0];
const base64 = await imageToBase64(file);
console.log(base64); // "data:image/jpeg;base64,..."
```

### JavaScript (Compressão):
```javascript
async function compressAndConvert(file, maxWidth = 1024) {
  return new Promise((resolve) => {
    const reader = new FileReader();
    reader.onload = (e) => {
      const img = new Image();
      img.onload = () => {
        const canvas = document.createElement('canvas');
        const scale = maxWidth / img.width;
        canvas.width = maxWidth;
        canvas.height = img.height * scale;
        
        const ctx = canvas.getContext('2d');
        ctx.drawImage(img, 0, 0, canvas.width, canvas.height);
        
        resolve(canvas.toDataURL('image/jpeg', 0.8));
      };
      img.src = e.target.result;
    };
    reader.readAsDataURL(file);
  });
}
```

### Python:
```python
import base64

with open("foto.jpg", "rb") as f:
    base64_string = f"data:image/jpeg;base64,{base64.b64encode(f.read()).decode()}"
```

### Bash:
```bash
echo "data:image/jpeg;base64,$(base64 -w 0 foto.jpg)"
```

---

## ⚠️ Erros Comuns

| Código | Erro | Solução |
|--------|------|---------|
| 401 | Unauthorized | Token expirado ou inválido - faça login novamente |
| 400 | Bad Request | Campos obrigatórios faltando ou formato inválido |
| 404 | Not Found | ID não existe - verifique se o registro foi criado |
| 500 | Internal Server Error | Erro no servidor - verifique os logs |

---

## 📚 Valores Válidos

### Período:
- `"manha"`
- `"tarde"`
- `"integral"`

### Status (Tarefas):
- `"planejada"`
- `"em_andamento"`
- `"concluida"`
- `"cancelada"`

### Tipo (Ocorrências):
- `"seguranca"`
- `"qualidade"`
- `"prazo"`
- `"custo"`
- `"clima"`
- `"outro"`

### Gravidade:
- `"baixa"` → Aparece como **OBSERVACAO** no relatório
- `"media"` → Aparece como **IMPORTANTE** no relatório
- `"alta"` → Aparece como **CRITICO** no relatório

### Status Resolução:
- `"pendente"`
- `"em_analise"`
- `"resolvida"`
- `"nao_aplicavel"`

### Status Aprovação:
- `"pendente"`
- `"aprovado"`
- `"reprovado"`

---

## 📖 Documentação Adicional

- **Guia Postman:** `POSTMAN-GUIDE.md` - Como importar e usar a coleção
- **Referência Rápida:** `QUICK-REFERENCE.md` - Cheat sheet da API
- **Testes:** `TESTES-RESULTADOS.md` - Resultados dos testes dos endpoints
- **README:** `README.md` - Documentação geral do projeto

---

## 💡 Dicas

1. **Sempre salve o `access_token`** após o login
2. **Fotos usam Base64**, não URLs
3. **Comprima imagens grandes** antes de enviar (máx. 5MB recomendado)
4. **Use o endpoint de relatório formatado** para gerar PDFs
5. **A gravidade é mapeada automaticamente** (baixa→OBSERVACAO, media→IMPORTANTE, alta→CRITICO)
6. **O período pode ser diferente** entre tarefas e metadados (ex: tarefa=manha, metadados=integral)
7. **Metadados são opcionais** - você pode criar tarefas/ocorrências sem foto
