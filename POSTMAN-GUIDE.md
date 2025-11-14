# Como Importar a Coleção

## Postman

1. Abra o Postman
2. Clique em **Import** (canto superior esquerdo)
3. Selecione o arquivo `postman-collection.json`
4. A coleção "API Diário de Obras" será importada

## Insomnia

1. Abra o Insomnia
2. Clique em **Create** → **Import From** → **File**
3. Selecione o arquivo `postman-collection.json`
4. A coleção será importada automaticamente

---

# Configuração Inicial

## 1. Configurar Variáveis de Ambiente

### No Postman:
1. Clique na coleção "API Diário de Obras"
2. Vá em **Variables**
3. Configure:
   - `base_url`: `http://localhost:9090` (ou URL do servidor)
   - `access_token`: (será preenchido automaticamente após login)

### No Insomnia:
1. Clique no dropdown de ambiente (canto superior esquerdo)
2. **Manage Environments**
3. Crie um novo ambiente:
```json
{
  "base_url": "http://localhost:9090",
  "access_token": ""
}
```

## 2. Fazer Login (Obter Token)

1. Execute a requisição **1. Autenticação → Login**
2. Copie o `access_token` da resposta
3. Cole no campo `access_token` das variáveis de ambiente

### Automatizar no Postman (Script de Teste):
Na requisição de Login, adicione este script na aba **Tests**:
```javascript
// Salvar o token automaticamente
if (pm.response.code === 200) {
    const responseJson = pm.response.json();
    pm.collectionVariables.set("access_token", responseJson.access_token);
    console.log("Token salvo:", responseJson.access_token);
}
```

### Automatizar no Insomnia:
O Insomnia não suporta scripts automáticos, mas você pode usar **Response → Copy Value** para copiar o token rapidamente.

---

# Ordem Recomendada de Uso

## Fluxo Completo de Teste:

### 1. **Autenticação**
- [ ] Fazer login e obter token

### 2. **Criar Tarefas**
- [ ] Criar tarefa da manhã
- [ ] Criar tarefa da tarde
- [ ] Listar todas as tarefas
- [ ] Buscar tarefas por obra e data

### 3. **Criar Ocorrências**
- [ ] Criar ocorrência de gravidade alta
- [ ] Criar ocorrência de gravidade média
- [ ] Criar ocorrência de gravidade baixa
- [ ] Listar todas as ocorrências
- [ ] Buscar ocorrências por gravidade

### 4. **Adicionar Metadados**
- [ ] Criar metadados com foto em Base64
- [ ] Criar metadados para manhã (se necessário)
- [ ] Criar metadados para tarde (se necessário)

### 5. **Consultar View Consolidada**
- [ ] Listar todos os diários
- [ ] Buscar diários por obra
- [ ] Buscar diários por data

### 6. **Gerar Relatório Formatado** ⭐
- [ ] Gerar relatório completo (endpoint principal)
- [ ] Verificar estrutura JSON para PDF

### 7. **Operações de Atualização**
- [ ] Atualizar uma tarefa
- [ ] Atualizar uma ocorrência

### 8. **Limpeza (Opcional)**
- [ ] Deletar tarefas de teste
- [ ] Deletar ocorrências de teste

---

# Dicas de Uso

## 🔐 Autenticação

**Token expira?** Sim, após algumas horas.  
**O que fazer?** Execute novamente a requisição de Login para obter um novo token.

## 📝 Valores Obrigatórios

### Para Tarefas:
- `obra_id`: ID da obra (ex: 5)
- `data`: Formato `YYYY-MM-DD`
- `periodo`: `manha`, `tarde`, ou `integral`
- `descricao`: Texto descritivo
- `status`: `planejada`, `em_andamento`, `concluida`, `cancelada`

### Para Ocorrências:
- `obra_id`: ID da obra
- `data`: Formato `YYYY-MM-DD`
- `periodo`: `manha`, `tarde`, ou `integral`
- `tipo`: `seguranca`, `qualidade`, `prazo`, `custo`, `clima`, ou `outro`
- `gravidade`: `baixa`, `media`, ou `alta`
- `descricao`: Descrição da ocorrência
- `status_resolucao`: `pendente`, `em_analise`, `resolvida`, ou `nao_aplicavel`

### Para Metadados:
- `obra_id`: ID da obra
- `data`: Formato `YYYY-MM-DD`
- `periodo`: `manha`, `tarde`, ou `integral`
- `foto`: Base64 (`data:image/jpeg;base64,...`) ou string vazia
- `status_aprovacao`: `pendente`, `aprovado`, ou `reprovado`

## 🖼️ Como Enviar Fotos

### Formato Base64:
```json
{
  "foto": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD..."
}
```

### Converter Imagem para Base64:

**Online:**
- https://www.base64-image.de/
- https://codebeautify.org/image-to-base64-converter

**JavaScript (Browser):**
```javascript
const fileInput = document.querySelector('input[type="file"]');
fileInput.addEventListener('change', (e) => {
  const file = e.target.files[0];
  const reader = new FileReader();
  reader.onloadend = () => {
    console.log(reader.result); // Base64 string
  };
  reader.readAsDataURL(file);
});
```

**Python:**
```python
import base64

with open("imagem.jpg", "rb") as image_file:
    encoded = base64.b64encode(image_file.read()).decode()
    base64_string = f"data:image/jpeg;base64,{encoded}"
    print(base64_string)
```

**Bash:**
```bash
echo "data:image/jpeg;base64,$(base64 -w 0 imagem.jpg)"
```

## 📊 Mapeamento de Gravidade

Quando você criar uma ocorrência com gravidade específica, ela será mapeada no relatório formatado:

| Gravidade (Input) | Tipo (Output no Relatório) |
|------------------|----------------------------|
| `alta`           | `CRITICO`                  |
| `media`          | `IMPORTANTE`               |
| `baixa`          | `OBSERVACAO`               |

## 🔄 IDs Dinâmicos

Os IDs nas URLs (`/tarefas/2`, `/ocorrencias/1`) são dinâmicos.  
Após criar um registro, use o `id` retornado na resposta para atualizar ou deletar.

**Exemplo:**
```json
// Resposta do POST /tarefas
{
  "id": 42,
  "obra_id": 5,
  "data": "2024-11-14",
  ...
}

// Use o ID 42 para atualizar:
PUT /tarefas/42
```

## 🎯 Endpoint Principal

**Relatório Formatado** (`GET /diarios/relatorio-formatado/{obra_id}`) é o endpoint mais importante.  
Ele retorna o JSON completo pronto para gerar PDF com:
- Informações da obra
- Tarefas realizadas (formatadas com status e %)
- Ocorrências (com gravidade mapeada)
- Fotos (em Base64)

---

# Troubleshooting

## Erro 401 Unauthorized
**Causa:** Token expirado ou não configurado  
**Solução:** Faça login novamente e atualize a variável `access_token`

## Erro 404 Not Found
**Causa:** ID não existe ou endpoint incorreto  
**Solução:** Verifique se o ID existe listando os registros primeiro

## Erro 400 Bad Request
**Causa:** JSON malformado ou campos obrigatórios faltando  
**Solução:** Valide o JSON e confira os campos obrigatórios acima

## Erro 500 Internal Server Error
**Causa:** Erro no servidor (constraints, validações, etc)  
**Solução:** Verifique os logs da API (`docker logs <container_id>`)

## Base64 muito grande
**Causa:** Imagem muito pesada (> 5MB)  
**Solução:** Comprima a imagem antes de converter para Base64

---

# Scripts Úteis

## Limpar Todos os Dados de Teste (cURL)

```bash
#!/bin/bash

# Configurar variáveis
BASE_URL="http://localhost:9090"
TOKEN="seu_token_aqui"

# Listar IDs para deletar
TAREFA_IDS=(2 3 4)
OCORRENCIA_IDS=(1 2 3)

# Deletar tarefas
for id in "${TAREFA_IDS[@]}"; do
  curl -X DELETE "$BASE_URL/tarefas/$id" \
    -H "Authorization: Bearer $TOKEN"
done

# Deletar ocorrências
for id in "${OCORRENCIA_IDS[@]}"; do
  curl -X DELETE "$BASE_URL/ocorrencias/$id" \
    -H "Authorization: Bearer $TOKEN"
done

echo "Limpeza concluída!"
```

## Criar Diário Completo (Bash)

```bash
#!/bin/bash

BASE_URL="http://localhost:9090"
TOKEN="seu_token_aqui"
OBRA_ID=5
DATA="2024-11-14"

# 1. Criar tarefa
curl -X POST "$BASE_URL/tarefas" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "obra_id": '$OBRA_ID',
    "data": "'$DATA'",
    "periodo": "manha",
    "descricao": "Concretagem das vigas",
    "status": "em_andamento",
    "percentual_conclusao": 70
  }'

# 2. Criar ocorrência
curl -X POST "$BASE_URL/ocorrencias" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "obra_id": '$OBRA_ID',
    "data": "'$DATA'",
    "periodo": "manha",
    "tipo": "clima",
    "gravidade": "baixa",
    "descricao": "Chuva leve",
    "status_resolucao": "nao_aplicavel"
  }'

# 3. Adicionar metadados com foto
curl -X POST "$BASE_URL/diarios-consolidado/metadados" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "obra_id": '$OBRA_ID',
    "data": "'$DATA'",
    "periodo": "integral",
    "foto": "data:image/jpeg;base64,/9j/4AAQSkZJRg...",
    "observacoes": "Dia produtivo",
    "status_aprovacao": "aprovado"
  }'

# 4. Gerar relatório
curl -X GET "$BASE_URL/diarios/relatorio-formatado/$OBRA_ID" \
  -H "Authorization: Bearer $TOKEN"

echo "Diário criado e relatório gerado!"
```

---

# Documentação Adicional

- **Guia Completo**: `FRONTEND-GUIDE.md`
- **Referência Rápida**: `QUICK-REFERENCE.md`
- **Resultados de Testes**: `TESTES-RESULTADOS.md`
- **README Principal**: `README.md`

---

# Suporte

Dúvidas sobre:
- **Endpoints**: Consulte `QUICK-REFERENCE.md`
- **Integração Frontend**: Consulte `FRONTEND-GUIDE.md`
- **Validação de Testes**: Consulte `TESTES-RESULTADOS.md`

**Contato:** [Adicione informações de contato]
