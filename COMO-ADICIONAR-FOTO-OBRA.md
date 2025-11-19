# 📸 GUIA: Como Adicionar Foto à Obra

**Data:** 19 de novembro de 2025  
**Endpoint:** `PUT /obras/:id`

---

## 🎯 OPÇÕES PARA ADICIONAR FOTO

Existem **3 formas** de adicionar uma foto à obra:

### **Opção 1:** Atualizar SOMENTE o campo foto (PATCH-like)
### **Opção 2:** Atualizar a obra completa (PUT)
### **Opção 3:** Usar SQL direto no banco

---

## 📝 OPÇÃO 1: Atualizar SOMENTE o campo foto

> ⚠️ **Atenção:** O endpoint atual é `PUT`, então requer **TODOS os campos obrigatórios**. 
> Veja a Opção 2 para enviar a obra completa.

---

## ✅ OPÇÃO 2: Atualizar Obra Completa (RECOMENDADO)

### Endpoint
```http
PUT /obras/:id
Authorization: Bearer {token}
Content-Type: application/json
```

### Passo 1: Buscar a obra atual

```bash
curl -X GET "http://localhost:9090/obras/5" \
  -H "Authorization: Bearer {seu_token}" \
  -H "Content-Type: application/json"
```

**Resposta:**
```json
{
  "data": {
    "id": 5,
    "nome": "Casa Residencial - Fortaleza",
    "contrato_numero": "CONTR-2024-001",
    "contratante_id": 2,
    "responsavel_id": 1,
    "data_inicio": "2024-01-01",
    "prazo_dias": 180,
    "data_fim_prevista": "2024-06-29",
    "orcamento": 250000.00,
    "status": "em_andamento",
    "foto": null,
    "endereco_rua": "Rua Principal",
    "endereco_numero": "123",
    "endereco_bairro": "Centro",
    "endereco_cidade": "Fortaleza",
    "endereco_estado": "CE",
    "endereco_cep": "60000-000"
  }
}
```

### Passo 2: Preparar a foto em Base64

```bash
# Converter imagem para Base64
base64 -w 0 sua_foto.jpg > foto_base64.txt

# Ou criar o formato completo
echo "data:image/jpeg;base64,$(base64 -w 0 sua_foto.jpg)" > foto_base64.txt
```

### Passo 3: Atualizar a obra com a foto

```bash
curl -X PUT "http://localhost:9090/obras/5" \
  -H "Authorization: Bearer {seu_token}" \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Casa Residencial - Fortaleza",
    "contrato_numero": "CONTR-2024-001",
    "contratante_id": 2,
    "responsavel_id": 1,
    "data_inicio": "2024-01-01",
    "prazo_dias": 180,
    "data_fim_prevista": "2024-06-29",
    "orcamento": 250000.00,
    "status": "em_andamento",
    "foto": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD...",
    "endereco_rua": "Rua Principal",
    "endereco_numero": "123",
    "endereco_bairro": "Centro",
    "endereco_cidade": "Fortaleza",
    "endereco_estado": "CE",
    "endereco_cep": "60000-000"
  }'
```

---

## 💻 EXEMPLO JAVASCRIPT/TYPESCRIPT

### React - Upload de Foto da Obra

```javascript
import React, { useState } from 'react';

function AtualizarFotoObra({ obra, token, onSuccess }) {
  const [loading, setLoading] = useState(false);
  const [preview, setPreview] = useState(null);

  const handleFileChange = async (event) => {
    const file = event.target.files[0];
    if (!file) return;

    // Validar tipo de arquivo
    if (!file.type.startsWith('image/')) {
      alert('Apenas imagens são permitidas');
      return;
    }

    // Validar tamanho (max 5MB)
    if (file.size > 5 * 1024 * 1024) {
      alert('Imagem muito grande. Máximo 5MB');
      return;
    }

    // Preview
    const reader = new FileReader();
    reader.onload = (e) => setPreview(e.target.result);
    reader.readAsDataURL(file);
  };

  const handleUpload = async () => {
    if (!preview) {
      alert('Selecione uma foto primeiro');
      return;
    }

    setLoading(true);

    try {
      // Montar payload com TODOS os campos da obra
      const payload = {
        nome: obra.nome,
        contrato_numero: obra.contrato_numero,
        contratante_id: obra.contratante_id,
        responsavel_id: obra.responsavel_id,
        data_inicio: obra.data_inicio,
        prazo_dias: obra.prazo_dias,
        data_fim_prevista: obra.data_fim_prevista,
        orcamento: obra.orcamento,
        status: obra.status,
        foto: preview,  // Base64 da imagem
        endereco_rua: obra.endereco_rua,
        endereco_numero: obra.endereco_numero,
        endereco_bairro: obra.endereco_bairro,
        endereco_cidade: obra.endereco_cidade,
        endereco_estado: obra.endereco_estado,
        endereco_cep: obra.endereco_cep,
        observacoes: obra.observacoes,
        contratada: obra.contratada,
        art: obra.art
      };

      const response = await fetch(
        `http://localhost:9090/obras/${obra.id}`,
        {
          method: 'PUT',
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(payload)
        }
      );

      if (response.ok) {
        alert('Foto atualizada com sucesso!');
        if (onSuccess) onSuccess();
      } else {
        const error = await response.json();
        alert(`Erro: ${error.error || 'Falha ao atualizar'}`);
      }
    } catch (error) {
      console.error('Erro:', error);
      alert('Erro ao enviar foto');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <h3>Adicionar Foto da Obra</h3>
      
      <input 
        type="file" 
        accept="image/*"
        onChange={handleFileChange}
        disabled={loading}
      />
      
      {preview && (
        <div style={{ margin: '20px 0' }}>
          <img 
            src={preview} 
            alt="Preview" 
            style={{ maxWidth: '300px', maxHeight: '300px' }}
          />
        </div>
      )}
      
      <button 
        onClick={handleUpload}
        disabled={loading || !preview}
      >
        {loading ? 'Enviando...' : 'Salvar Foto'}
      </button>
    </div>
  );
}

export default AtualizarFotoObra;
```

---

## 🗄️ OPÇÃO 3: SQL Direto (Para Desenvolvimento/Teste)

### Atualizar foto diretamente no banco

```sql
-- Conectar ao PostgreSQL
docker exec -it db_obras psql -U obra_user -d obras_db

-- Atualizar foto da obra ID 5
UPDATE obra 
SET foto = 'data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD...',
    updated_at = NOW()
WHERE id = 5;

-- Verificar
SELECT id, nome, 
       CASE 
         WHEN foto IS NULL THEN 'SEM FOTO'
         WHEN foto = '' THEN 'VAZIO'
         ELSE CONCAT('COM FOTO (', LENGTH(foto), ' bytes)')
       END as status_foto
FROM obra 
WHERE id = 5;
```

---

## 📋 CAMPOS OBRIGATÓRIOS DO PUT

O endpoint `PUT /obras/:id` requer os seguintes campos:

| Campo | Tipo | Obrigatório | Exemplo |
|-------|------|-------------|---------|
| `nome` | string | ✅ Sim | "Casa Residencial" |
| `contrato_numero` | string | ✅ Sim | "CONTR-2024-001" |
| `contratante_id` | integer | ✅ Sim | 2 |
| `responsavel_id` | integer | ✅ Sim | 1 |
| `data_inicio` | string | ✅ Sim | "2024-01-01" |
| `prazo_dias` | integer | ✅ Sim | 180 |
| `data_fim_prevista` | string | ✅ Sim | "2024-06-29" |
| `orcamento` | number | ✅ Sim | 250000.00 |
| `status` | string | ✅ Sim | "em_andamento" |
| `endereco_rua` | string | ✅ Sim | "Rua Principal" |
| `endereco_numero` | string | ✅ Sim | "123" |
| `endereco_bairro` | string | ✅ Sim | "Centro" |
| `endereco_cidade` | string | ✅ Sim | "Fortaleza" |
| `endereco_estado` | string | ✅ Sim | "CE" |
| `foto` | string | ❌ Opcional | "data:image/jpeg;base64,..." |
| `endereco_cep` | string | ❌ Opcional | "60000-000" |
| `observacoes` | string | ❌ Opcional | "..." |
| `contratada` | string | ❌ Opcional | "Empresa XYZ" |
| `art` | string | ❌ Opcional | "ART-123456" |

---

## 🧪 TESTE COMPLETO

### 1. Fazer login
```bash
TOKEN=$(curl -s -X POST "http://localhost:9090/login" \
  -H "Content-Type: application/json" \
  -d '{"email": "teste@teste.com", "senha": "senha123"}' \
  | python3 -c "import sys, json; print(json.load(sys.stdin)['access_token'])")

echo "Token: $TOKEN"
```

### 2. Buscar obra atual
```bash
curl -X GET "http://localhost:9090/obras/5" \
  -H "Authorization: Bearer $TOKEN" \
  -s | python3 -m json.tool > obra_atual.json

cat obra_atual.json
```

### 3. Preparar foto (exemplo com imagem pequena de teste)
```bash
# Criar uma imagem de teste simples (1x1 pixel)
FOTO_BASE64="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg=="

echo "Foto preparada: ${FOTO_BASE64:0:50}..."
```

### 4. Atualizar com foto
```bash
curl -X PUT "http://localhost:9090/obras/5" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Casa Residencial - Fortaleza",
    "contrato_numero": "CONTR-2024-001",
    "contratante_id": 2,
    "responsavel_id": 1,
    "data_inicio": "2024-01-01",
    "prazo_dias": 180,
    "data_fim_prevista": "2024-06-29",
    "orcamento": 250000.00,
    "status": "em_andamento",
    "foto": "'"$FOTO_BASE64"'",
    "endereco_rua": "Rua Principal",
    "endereco_numero": "123",
    "endereco_bairro": "Centro",
    "endereco_cidade": "Fortaleza",
    "endereco_estado": "CE"
  }' \
  -s | python3 -m json.tool
```

### 5. Verificar no relatório fotográfico
```bash
curl -X GET "http://localhost:9090/relatorios/fotografico/5" \
  -H "Authorization: Bearer $TOKEN" \
  -s | python3 -m json.tool | grep -A 2 "foto_obra"
```

**Resultado esperado:**
```json
"foto_obra": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAY..."
```

---

## ⚠️ OBSERVAÇÕES IMPORTANTES

### 1. **Formato da Foto**
- Deve ser uma string Base64
- Prefixo recomendado: `data:image/jpeg;base64,` ou `data:image/png;base64,`
- Sem o prefixo também funciona, mas o prefixo facilita o uso no frontend

### 2. **Tamanho Máximo**
- PostgreSQL suporta TEXT sem limite prático
- **Recomendação:** Limitar a 5-10MB no frontend
- Para fotos maiores, considere usar serviço de storage externo (S3, etc)

### 3. **Compressão**
- Comprimir imagens antes de converter para Base64
- Redimensionar para tamanho adequado (ex: 1200x800)
- Usar qualidade JPEG 80-85%

### 4. **Performance**
- Fotos grandes em Base64 aumentam o payload da API
- Para muitas fotos, considere endpoint separado de upload

### 5. **Diferença entre fotos**
- `obra.foto` = Foto principal/capa da obra (1 foto)
- `diario_obra.foto` = Fotos dos diários (múltiplas fotos)

---

## 🎯 RESUMO RÁPIDO

```bash
# 1. Login
TOKEN="seu_token_aqui"

# 2. Preparar foto
FOTO="data:image/jpeg;base64,/9j/4AAQSkZJRg..."

# 3. Atualizar (com TODOS os campos da obra + foto)
curl -X PUT "http://localhost:9090/obras/5" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "...",
    "contrato_numero": "...",
    ... todos os campos obrigatórios ...
    "foto": "'"$FOTO"'"
  }'

# 4. Verificar
curl -X GET "http://localhost:9090/relatorios/fotografico/5" \
  -H "Authorization: Bearer $TOKEN"
```

---

**Status:** ✅ Pronto para uso  
**Endpoint:** `PUT /obras/:id`  
**Campo:** `foto` (opcional, string Base64)
