# 🚀 Quick Reference - API Diário de Obras

## 📌 Base URL
```
http://localhost:9090
```

## 🔐 Autenticação
```bash
# Login
POST /login
Body: {"email": "user@email.com", "senha": "123456"}
Response: {"access_token": "...", "refresh_token": "..."}

# Usar token em todas as requisições
Header: Authorization: Bearer {access_token}
```

---

## 📝 TAREFAS

| Ação | Método | Endpoint | Body Exemplo |
|------|--------|----------|--------------|
| **Criar** | POST | `/tarefas` | `{"obra_id":5, "data":"2024-11-14", "descricao":"Concretagem", "status":"em_andamento", "percentual_conclusao":60}` |
| **Listar** | GET | `/tarefas` | - |
| **Buscar** | GET | `/tarefas/obra/5/data/2024-11-14` | - |
| **Atualizar** | PUT | `/tarefas/{id}` | Todos os campos obrigatórios |
| **Deletar** | DELETE | `/tarefas/{id}` | - |

### Status válidos:
`planejada` | `em_andamento` | `concluida` | `cancelada`

### Períodos válidos:
`manha` | `tarde` | `integral` | `noite`

---

## ⚠️ OCORRÊNCIAS

| Ação | Método | Endpoint | Body Exemplo |
|------|--------|----------|--------------|
| **Criar** | POST | `/ocorrencias` | `{"obra_id":5, "data":"2024-11-14", "tipo":"seguranca", "gravidade":"alta", "descricao":"Falta de EPIs"}` |
| **Listar** | GET | `/ocorrencias` | - |
| **Buscar** | GET | `/ocorrencias/obra/5/data/2024-11-14` | - |
| **Por Gravidade** | GET | `/ocorrencias/gravidade/alta` | - |
| **Atualizar** | PUT | `/ocorrencias/{id}` | Todos os campos obrigatórios |
| **Deletar** | DELETE | `/ocorrencias/{id}` | - |

### Tipos válidos:
`seguranca` | `qualidade` | `prazo` | `custo` | `clima` | `equipamento` | `material` | `geral`

### Gravidades válidas:
`baixa` | `media` | `alta` | `critica`

### Status de Resolução:
`pendente` | `em_analise` | `resolvida` | `nao_aplicavel`

### Mapeamento Gravidade → Tipo no Relatório:
- **ALTA/CRÍTICA** → `CRITICO`
- **MÉDIA** → `IMPORTANTE`
- **BAIXA** → `OBSERVACAO`

---

## 📸 METADADOS

| Ação | Método | Endpoint | Body Exemplo |
|------|--------|----------|--------------|
| **Criar/Atualizar** | POST | `/diarios-consolidado/metadados` | `{"obra_id":5, "data":"2024-11-14", "periodo":"integral", "foto":"data:image/jpeg;base64,...", "observacoes":"Dia produtivo"}` |

### ⚠️ IMPORTANTE:
- **Foto DEVE ser Base64**: `data:image/jpeg;base64,/9j/4AAQ...`
- Não use URL de imagem!

### Status de Aprovação:
`pendente` | `aprovado` | `rejeitado`

---

## 📊 DIÁRIO CONSOLIDADO (View)

| Ação | Método | Endpoint | Retorna |
|------|--------|----------|---------|
| **Listar Todos** | GET | `/diarios-consolidado` | Todos os diários agregados |
| **Por Obra** | GET | `/diarios-consolidado/obra/{obra_id}` | Diários de uma obra |
| **Por Data** | GET | `/diarios-consolidado/data/{data}` | Diários de uma data |

### Campos Agregados:
- `atividades`: String com todas atividades separadas por ";"
- `ocorrencias`: String com todas ocorrências com `[GRAVIDADE]`
- `qtd_atividades`, `qtd_ocorrencias`, `qtd_equipe`, etc.

---

## 📄 RELATÓRIO FORMATADO

| Ação | Método | Endpoint | Descrição |
|------|--------|----------|-----------|
| **Gerar Relatório** | GET | `/diarios/relatorio-formatado/{obra_id}` | JSON completo pronto para PDF |

### Estrutura de Resposta:
```json
{
  "data": {
    "informacoes_obra": { ... },
    "tarefas_realizadas": [ ... ],
    "ocorrencias": [ ... ],
    "equipe_envolvida": [ ... ],
    "equipamentos_utilizados": [ ... ],
    "materiais_utilizados": [ ... ],
    "fotos": [ ... ],
    "responsavel_empresa": { ... },
    "responsavel_prefeitura": { ... }
  }
}
```

---

## 💻 Código Rápido

### Axios Config
```javascript
import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:9090',
  headers: { 'Content-Type': 'application/json' }
});

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('access_token');
  if (token) config.headers.Authorization = `Bearer ${token}`;
  return config;
});

export default api;
```

### Converter Imagem para Base64
```javascript
const imageToBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = reject;
  });
};

// Uso:
const base64 = await imageToBase64(file);
```

### Salvar Diário Completo
```javascript
const salvarDiario = async (obraId, data, tarefas, ocorrencias, foto, observacoes) => {
  try {
    // 1. Criar tarefas
    for (const tarefa of tarefas) {
      await api.post('/tarefas', {
        obra_id: obraId,
        data: data,
        periodo: 'integral',
        ...tarefa
      });
    }

    // 2. Criar ocorrências
    for (const ocorrencia of ocorrencias) {
      await api.post('/ocorrencias', {
        obra_id: obraId,
        data: data,
        periodo: 'integral',
        ...ocorrencia
      });
    }

    // 3. Criar metadados
    await api.post('/diarios-consolidado/metadados', {
      obra_id: obraId,
      data: data,
      periodo: 'integral',
      foto: foto,
      observacoes: observacoes
    });

    return { success: true };
  } catch (error) {
    console.error('Erro:', error);
    return { success: false, error: error.response?.data?.error };
  }
};
```

### Buscar Relatório
```javascript
const buscarRelatorio = async (obraId) => {
  try {
    const response = await api.get(`/diarios/relatorio-formatado/${obraId}`);
    return response.data.data;
  } catch (error) {
    console.error('Erro ao buscar relatório:', error);
    throw error;
  }
};
```

---

## ⚠️ Erros Comuns

| Status | Mensagem | Solução |
|--------|----------|---------|
| 401 | Token inválido | Fazer login novamente |
| 404 | Obra não encontrada | Verificar se obra existe |
| 400 | Campo obrigatório | Validar campos antes de enviar |
| 500 | Foreign key constraint | Verificar IDs de relacionamento |

---

## ✅ Fluxo Recomendado

```
1. Login → Obter token ✅
   ↓
2. Criar Tarefas (uma por vez) ✅
   ↓
3. Criar Ocorrências (uma por vez) ✅
   ↓
4. Criar Metadados (foto + observações) ✅
   ↓
5. Visualizar Relatório Formatado ✅
```

---

## 🎯 Validações Importantes

### Tarefas:
- ✅ `obra_id` obrigatório
- ✅ `data` obrigatório (formato: YYYY-MM-DD)
- ✅ `descricao` obrigatório
- ✅ `percentual_conclusao` entre 0-100
- ✅ `status` deve ser válido

### Ocorrências:
- ✅ `obra_id` obrigatório
- ✅ `data` obrigatório
- ✅ `descricao` obrigatório
- ✅ `tipo` e `gravidade` devem ser válidos

### Metadados:
- ✅ `obra_id` obrigatório
- ✅ `data` obrigatório
- ✅ `periodo` obrigatório
- ✅ `foto` deve ser Base64 (não URL!)
- ✅ Formato: `data:image/jpeg;base64,...`

---

## 📞 Links Úteis

- 📄 Documentação Completa: `FRONTEND-GUIDE.md`
- 🧪 Resultados dos Testes: `TESTES-RESULTADOS.md`
- 📚 README: `README.md`

---

**Dica:** Sempre teste os endpoints usando o Postman ou Insomnia antes de integrar no front-end!
