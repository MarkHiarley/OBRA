# 🎨 Guia de Integração Front-End - Sistema de Diário de Obras

## 📋 Índice

1. [Introdução](#introdução)
2. [Autenticação](#autenticação)
3. [Fluxo Completo](#fluxo-completo)
4. [Endpoints Principais](#endpoints-principais)
5. [Exemplos de Código](#exemplos-de-código)
6. [Tratamento de Erros](#tratamento-de-erros)
7. [Boas Práticas](#boas-práticas)

---

## 🎯 Introdução

Este guia explica como integrar o front-end com a **NOVA ARQUITETURA** do Sistema de Diário de Obras.

### O que mudou?

**❌ Antiga Arquitetura:**
- Um único endpoint `/diarios` com tudo junto
- Dados em campos TEXT monolíticos

**✅ Nova Arquitetura:**
- Endpoints separados por funcionalidade
- Dados normalizados e estruturados
- Melhor controle e flexibilidade

### Novos Endpoints:

| Recurso | Endpoint Base | Descrição |
|---------|--------------|-----------|
| **Tarefas** | `/tarefas` | Atividades realizadas no dia |
| **Ocorrências** | `/ocorrencias` | Problemas/eventos registrados |
| **Metadados** | `/diarios-consolidado/metadados` | Fotos, observações, aprovação |
| **Relatório** | `/diarios/relatorio-formatado/:obra_id` | Relatório completo para PDF |

---

## 🔐 Autenticação

Todos os endpoints requerem autenticação via **JWT Token**.

### 1. Fazer Login

```javascript
// Requisição
POST http://localhost:9090/login
Content-Type: application/json

{
  "email": "usuario@exemplo.com",
  "senha": "senha123"
}

// Resposta (200 OK)
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 2. Usar o Token

Inclua o token no header `Authorization` de todas as requisições:

```javascript
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Exemplo com Axios:

```javascript
import axios from 'axios';

// Criar instância do Axios com configuração padrão
const api = axios.create({
  baseURL: 'http://localhost:9090',
  headers: {
    'Content-Type': 'application/json'
  }
});

// Interceptor para adicionar token automaticamente
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('access_token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;
```

---

## 🔄 Fluxo Completo

### Cenário: Registrar um dia de trabalho na obra

```
┌─────────────────────────────────────────────────┐
│ 1. Usuário preenche formulário do diário       │
│    - Tarefas realizadas                         │
│    - Ocorrências do dia                         │
│    - Foto da obra                               │
│    - Observações gerais                         │
└─────────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────────┐
│ 2. Front-end faz 3 requisições separadas:      │
│                                                 │
│    POST /tarefas (para cada tarefa)            │
│    POST /ocorrencias (para cada ocorrência)    │
│    POST /diarios-consolidado/metadados         │
│         (foto + observações)                    │
└─────────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────────┐
│ 3. Backend agrupa automaticamente na VIEW      │
│    vw_diario_consolidado                       │
└─────────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────────┐
│ 4. Para visualizar relatório:                  │
│    GET /diarios/relatorio-formatado/:obra_id   │
└─────────────────────────────────────────────────┘
```

---

## 📝 Endpoints Principais

### 1️⃣ TAREFAS (Atividades Realizadas)

#### 📌 Criar Nova Tarefa

```javascript
POST /tarefas
Authorization: Bearer {token}
Content-Type: application/json

// Body
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",                    // "manha", "tarde", "integral", "noite"
  "descricao": "Concretagem das vigas",
  "responsavel_id": 60,                  // ID da pessoa (opcional)
  "status": "em_andamento",              // "planejada", "em_andamento", "concluida", "cancelada"
  "percentual_conclusao": 60,            // 0-100
  "observacao": "Clima favorável"        // opcional
}

// Resposta (201 Created)
{
  "message": "Atividade criada com sucesso",
  "data": {
    "id": 2,
    "obra_id": 5,
    "data": "2024-11-14",
    "periodo": "manha",
    "descricao": "Concretagem das vigas",
    "status": "em_andamento",
    "percentual_conclusao": 60,
    "created_at": "2025-11-14T18:23:53Z"
  }
}
```

#### 📋 Listar Todas as Tarefas

```javascript
GET /tarefas
Authorization: Bearer {token}

// Resposta (200 OK)
{
  "data": [
    {
      "id": 2,
      "obra_id": 5,
      "data": "2024-11-14",
      "periodo": "manha",
      "descricao": "Concretagem das vigas",
      "status": "em_andamento",
      "percentual_conclusao": 60
    },
    // ... mais tarefas
  ]
}
```

#### 🔍 Buscar Tarefas por Obra e Data

```javascript
GET /tarefas/obra/{obra_id}/data/{data}
Authorization: Bearer {token}

// Exemplo: GET /tarefas/obra/5/data/2024-11-14
```

#### ✏️ Atualizar Tarefa

```javascript
PUT /tarefas/{id}
Authorization: Bearer {token}
Content-Type: application/json

// Body (todos os campos obrigatórios)
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",
  "descricao": "Concretagem das vigas",
  "status": "concluida",              // ← Mudou de em_andamento para concluida
  "percentual_conclusao": 100,        // ← Mudou de 60 para 100
  "observacao": "Concluída com sucesso"
}
```

#### 🗑️ Deletar Tarefa

```javascript
DELETE /tarefas/{id}
Authorization: Bearer {token}

// Resposta (200 OK)
{
  "message": "Atividade deletada com sucesso"
}
```

---

### 2️⃣ OCORRÊNCIAS (Problemas/Eventos)

#### 📌 Criar Nova Ocorrência

```javascript
POST /ocorrencias
Authorization: Bearer {token}
Content-Type: application/json

// Body
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",                        // opcional
  "tipo": "seguranca",                       // "seguranca", "qualidade", "prazo", "custo", "clima", "equipamento", "material", "geral"
  "gravidade": "alta",                       // "baixa", "media", "alta", "critica"
  "descricao": "Falta de EPIs na equipe",
  "responsavel_id": 60,                      // opcional
  "status_resolucao": "resolvida",           // "pendente", "em_analise", "resolvida", "nao_aplicavel"
  "acao_tomada": "EPIs fornecidos"          // opcional
}

// Resposta (201 Created)
{
  "message": "Ocorrência criada com sucesso",
  "data": {
    "id": 1,
    "obra_id": 5,
    "tipo": "seguranca",
    "gravidade": "alta",
    "descricao": "Falta de EPIs na equipe",
    "status_resolucao": "resolvida",
    "created_at": "2025-11-14T18:25:57Z"
  }
}
```

#### 📋 Listar Todas as Ocorrências

```javascript
GET /ocorrencias
Authorization: Bearer {token}
```

#### 🔍 Buscar por Obra e Data

```javascript
GET /ocorrencias/obra/{obra_id}/data/{data}
Authorization: Bearer {token}
```

#### 🎯 Buscar por Gravidade

```javascript
GET /ocorrencias/gravidade/{gravidade}
Authorization: Bearer {token}

// Exemplo: GET /ocorrencias/gravidade/alta
// Retorna todas as ocorrências de alta gravidade
```

#### ✏️ Atualizar Ocorrência

```javascript
PUT /ocorrencias/{id}
Authorization: Bearer {token}
Content-Type: application/json

// Body (todos os campos obrigatórios)
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "manha",
  "tipo": "seguranca",
  "gravidade": "alta",
  "descricao": "Falta de EPIs na equipe",
  "status_resolucao": "resolvida",         // ← Atualizado
  "acao_tomada": "EPIs fornecidos e treinamento realizado"  // ← Atualizado
}
```

#### 🗑️ Deletar Ocorrência

```javascript
DELETE /ocorrencias/{id}
Authorization: Bearer {token}
```

---

### 3️⃣ METADADOS DO DIÁRIO (Foto, Observações, Aprovação)

#### 📌 Criar/Atualizar Metadados

```javascript
POST /diarios-consolidado/metadados
Authorization: Bearer {token}
Content-Type: application/json

// Body
{
  "obra_id": 5,
  "data": "2024-11-14",
  "periodo": "integral",                    // "manha", "tarde", "integral", "noite"
  "foto": "data:image/jpeg;base64,/9j/4AAQSkZJRg...",  // Base64 da imagem
  "observacoes": "Dia produtivo. Clima favorável. Equipe trabalhou bem.",
  "responsavel_id": 60,                     // opcional
  "aprovado_por_id": 61,                    // opcional
  "status_aprovacao": "aprovado"            // "pendente", "aprovado", "rejeitado"
}

// Resposta (201 Created)
{
  "message": "Metadados salvos com sucesso",
  "data": {
    "id": 1,
    "obra_id": 5,
    "data": "2024-11-14",
    "periodo": "integral",
    "foto": "data:image/jpeg;base64,...",
    "observacoes": "Dia produtivo...",
    "status_aprovacao": "aprovado",
    "created_at": "2025-11-14T18:26:53Z"
  }
}
```

**⚠️ IMPORTANTE: Foto deve ser Base64!**

```javascript
// Converter imagem para Base64
function imageToBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
}

// Uso:
const file = event.target.files[0];
const base64 = await imageToBase64(file);
// base64 = "data:image/jpeg;base64,/9j/4AAQ..."
```

---

### 4️⃣ DIÁRIO CONSOLIDADO (Visualização Agregada)

#### 📋 Listar Todos os Diários

```javascript
GET /diarios-consolidado
Authorization: Bearer {token}

// Resposta (200 OK)
{
  "data": [
    {
      "diario_id": 1,
      "obra_id": 5,
      "obra_nome": "Casa Residencial - Fortaleza",
      "data": "2024-11-14",
      "periodo": "integral",
      "atividades": "Concretagem (em_andamento - 85%); Armadura (concluida - 100%)",
      "ocorrencias": "[ALTA] Falta de EPIs - resolvida; [MEDIA] Atraso material - pendente",
      "foto": "data:image/jpeg;base64,...",
      "observacoes": "Dia produtivo",
      "responsavel_nome": "João Silva",
      "aprovado_por_nome": "Maria Santos",
      "status_aprovacao": "aprovado",
      "qtd_atividades": 2,
      "qtd_ocorrencias": 2,
      "qtd_equipe": 15,
      "qtd_equipamentos": 8
    }
  ]
}
```

#### 🔍 Buscar por Obra

```javascript
GET /diarios-consolidado/obra/{obra_id}
Authorization: Bearer {token}
```

#### 🔍 Buscar por Data

```javascript
GET /diarios-consolidado/data/{data}
Authorization: Bearer {token}

// Exemplo: GET /diarios-consolidado/data/2024-11-14
```

---

### 5️⃣ RELATÓRIO FORMATADO (Pronto para PDF)

#### 📄 Gerar Relatório Completo

```javascript
GET /diarios/relatorio-formatado/{obra_id}
Authorization: Bearer {token}

// Resposta (200 OK)
{
  "data": {
    "informacoes_obra": {
      "titulo": "Casa Residencial - Fortaleza",
      "numero_contrato": "CONTR-2024-001",
      "contratante": "Prefeitura Municipal",
      "prazo_obra": "180 DIAS",
      "tempo_decorrido": "30 DIAS",
      "contratada": "Construtora ABC LTDA",
      "responsavel_tecnico": "Eng. João Silva",
      "registro_profissional": "CREA-CE 12345"
    },
    "tarefas_realizadas": [
      {
        "descricao": "Concretagem das vigas do 2º pavimento (em_andamento - 85%)",
        "data": "2024-11-14"
      },
      {
        "descricao": "Montagem de armadura da laje (concluida - 100%)",
        "data": "2024-11-14"
      }
    ],
    "ocorrencias": [
      {
        "descricao": "Falta de EPIs na equipe de alvenaria - resolvida",
        "tipo": "CRITICO"
      },
      {
        "descricao": "Atraso na entrega de material - pendente",
        "tipo": "IMPORTANTE"
      },
      {
        "descricao": "Chuva leve durante 30 minutos - nao_aplicavel",
        "tipo": "OBSERVACAO"
      }
    ],
    "equipe_envolvida": [
      {
        "codigo": "PEDREIRO",
        "descricao": "Pedreiro",
        "quantidade_utilizada": 5
      }
    ],
    "equipamentos_utilizados": [
      {
        "codigo": "BETONEIRA",
        "descricao": "Betoneira",
        "quantidade_utilizada": 2
      }
    ],
    "materiais_utilizados": [
      {
        "codigo": "CIMENTO",
        "descricao": "Cimento CP-II",
        "quantidade_utilizada": 50,
        "unidade": "sacos"
      }
    ],
    "fotos": [
      {
        "id": 5,
        "url": "data:image/jpeg;base64,/9j/4AAQSkZJRg...",
        "timestamp": "2024-11-14",
        "categoria": "DIARIO"
      }
    ],
    "responsavel_empresa": {
      "nome": "João Silva",
      "cargo": "Engenheiro Civil",
      "empresa": "Construtora ABC LTDA"
    },
    "responsavel_prefeitura": {
      "nome": "Maria Santos",
      "cargo": "Fiscal da Obra",
      "empresa": "Prefeitura Municipal"
    }
  }
}
```

---

## 💻 Exemplos de Código

### Exemplo Completo: React + Axios

```javascript
import React, { useState } from 'react';
import api from './api'; // Instância do axios configurada

const DiarioForm = () => {
  const [obraId, setObraId] = useState(5);
  const [data, setData] = useState('2024-11-14');
  const [tarefas, setTarefas] = useState([
    { descricao: '', status: 'em_andamento', percentual: 0 }
  ]);
  const [ocorrencias, setOcorrencias] = useState([
    { descricao: '', tipo: 'geral', gravidade: 'baixa' }
  ]);
  const [foto, setFoto] = useState(null);
  const [observacoes, setObservacoes] = useState('');

  // Converter imagem para Base64
  const handleImageChange = async (e) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => setFoto(reader.result);
    }
  };

  // Adicionar nova tarefa
  const addTarefa = () => {
    setTarefas([...tarefas, { descricao: '', status: 'em_andamento', percentual: 0 }]);
  };

  // Adicionar nova ocorrência
  const addOcorrencia = () => {
    setOcorrencias([...ocorrencias, { descricao: '', tipo: 'geral', gravidade: 'baixa' }]);
  };

  // Enviar diário completo
  const handleSubmit = async (e) => {
    e.preventDefault();
    
    try {
      // 1. Criar todas as tarefas
      for (const tarefa of tarefas) {
        if (tarefa.descricao.trim()) {
          await api.post('/tarefas', {
            obra_id: obraId,
            data: data,
            periodo: 'integral',
            descricao: tarefa.descricao,
            status: tarefa.status,
            percentual_conclusao: tarefa.percentual
          });
        }
      }

      // 2. Criar todas as ocorrências
      for (const ocorrencia of ocorrencias) {
        if (ocorrencia.descricao.trim()) {
          await api.post('/ocorrencias', {
            obra_id: obraId,
            data: data,
            periodo: 'integral',
            tipo: ocorrencia.tipo,
            gravidade: ocorrencia.gravidade,
            descricao: ocorrencia.descricao,
            status_resolucao: 'pendente'
          });
        }
      }

      // 3. Criar metadados (foto + observações)
      await api.post('/diarios-consolidado/metadados', {
        obra_id: obraId,
        data: data,
        periodo: 'integral',
        foto: foto,
        observacoes: observacoes,
        status_aprovacao: 'pendente'
      });

      alert('Diário salvo com sucesso!');
      
      // Limpar formulário
      setTarefas([{ descricao: '', status: 'em_andamento', percentual: 0 }]);
      setOcorrencias([{ descricao: '', tipo: 'geral', gravidade: 'baixa' }]);
      setFoto(null);
      setObservacoes('');
      
    } catch (error) {
      console.error('Erro ao salvar diário:', error);
      alert('Erro ao salvar diário: ' + error.response?.data?.error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Novo Diário de Obra</h2>
      
      {/* Obra e Data */}
      <div>
        <label>Obra ID:</label>
        <input 
          type="number" 
          value={obraId} 
          onChange={(e) => setObraId(Number(e.target.value))}
          required
        />
      </div>
      
      <div>
        <label>Data:</label>
        <input 
          type="date" 
          value={data} 
          onChange={(e) => setData(e.target.value)}
          required
        />
      </div>

      {/* Tarefas */}
      <h3>Tarefas Realizadas</h3>
      {tarefas.map((tarefa, index) => (
        <div key={index} style={{ border: '1px solid #ccc', padding: '10px', marginBottom: '10px' }}>
          <input
            type="text"
            placeholder="Descrição da tarefa"
            value={tarefa.descricao}
            onChange={(e) => {
              const newTarefas = [...tarefas];
              newTarefas[index].descricao = e.target.value;
              setTarefas(newTarefas);
            }}
          />
          
          <select
            value={tarefa.status}
            onChange={(e) => {
              const newTarefas = [...tarefas];
              newTarefas[index].status = e.target.value;
              setTarefas(newTarefas);
            }}
          >
            <option value="planejada">Planejada</option>
            <option value="em_andamento">Em Andamento</option>
            <option value="concluida">Concluída</option>
            <option value="cancelada">Cancelada</option>
          </select>
          
          <input
            type="number"
            placeholder="% Conclusão"
            min="0"
            max="100"
            value={tarefa.percentual}
            onChange={(e) => {
              const newTarefas = [...tarefas];
              newTarefas[index].percentual = Number(e.target.value);
              setTarefas(newTarefas);
            }}
          />
        </div>
      ))}
      <button type="button" onClick={addTarefa}>+ Adicionar Tarefa</button>

      {/* Ocorrências */}
      <h3>Ocorrências</h3>
      {ocorrencias.map((ocorrencia, index) => (
        <div key={index} style={{ border: '1px solid #ccc', padding: '10px', marginBottom: '10px' }}>
          <input
            type="text"
            placeholder="Descrição da ocorrência"
            value={ocorrencia.descricao}
            onChange={(e) => {
              const newOcorrencias = [...ocorrencias];
              newOcorrencias[index].descricao = e.target.value;
              setOcorrencias(newOcorrencias);
            }}
          />
          
          <select
            value={ocorrencia.tipo}
            onChange={(e) => {
              const newOcorrencias = [...ocorrencias];
              newOcorrencias[index].tipo = e.target.value;
              setOcorrencias(newOcorrencias);
            }}
          >
            <option value="seguranca">Segurança</option>
            <option value="qualidade">Qualidade</option>
            <option value="prazo">Prazo</option>
            <option value="custo">Custo</option>
            <option value="clima">Clima</option>
            <option value="equipamento">Equipamento</option>
            <option value="material">Material</option>
            <option value="geral">Geral</option>
          </select>
          
          <select
            value={ocorrencia.gravidade}
            onChange={(e) => {
              const newOcorrencias = [...ocorrencias];
              newOcorrencias[index].gravidade = e.target.value;
              setOcorrencias(newOcorrencias);
            }}
          >
            <option value="baixa">Baixa</option>
            <option value="media">Média</option>
            <option value="alta">Alta</option>
            <option value="critica">Crítica</option>
          </select>
        </div>
      ))}
      <button type="button" onClick={addOcorrencia}>+ Adicionar Ocorrência</button>

      {/* Foto */}
      <h3>Foto da Obra</h3>
      <input 
        type="file" 
        accept="image/*"
        onChange={handleImageChange}
      />
      {foto && <img src={foto} alt="Preview" style={{ maxWidth: '300px', marginTop: '10px' }} />}

      {/* Observações */}
      <h3>Observações Gerais</h3>
      <textarea
        value={observacoes}
        onChange={(e) => setObservacoes(e.target.value)}
        placeholder="Observações sobre o dia de trabalho..."
        rows="4"
        style={{ width: '100%' }}
      />

      {/* Botão Submit */}
      <button type="submit" style={{ marginTop: '20px', padding: '10px 20px', fontSize: '16px' }}>
        Salvar Diário
      </button>
    </form>
  );
};

export default DiarioForm;
```

### Exemplo: Visualizar Relatório

```javascript
import React, { useState, useEffect } from 'react';
import api from './api';

const RelatorioView = ({ obraId }) => {
  const [relatorio, setRelatorio] = useState(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    carregarRelatorio();
  }, [obraId]);

  const carregarRelatorio = async () => {
    setLoading(true);
    try {
      const response = await api.get(`/diarios/relatorio-formatado/${obraId}`);
      setRelatorio(response.data.data);
    } catch (error) {
      console.error('Erro ao carregar relatório:', error);
      alert('Erro ao carregar relatório');
    } finally {
      setLoading(false);
    }
  };

  if (loading) return <p>Carregando relatório...</p>;
  if (!relatorio) return <p>Nenhum relatório encontrado</p>;

  return (
    <div className="relatorio">
      <h1>Relatório de Obra</h1>
      
      {/* Informações da Obra */}
      <section>
        <h2>Informações da Obra</h2>
        <p><strong>Título:</strong> {relatorio.informacoes_obra.titulo}</p>
        <p><strong>Contrato:</strong> {relatorio.informacoes_obra.numero_contrato}</p>
        <p><strong>Prazo:</strong> {relatorio.informacoes_obra.prazo_obra}</p>
        <p><strong>Tempo Decorrido:</strong> {relatorio.informacoes_obra.tempo_decorrido}</p>
        <p><strong>Contratada:</strong> {relatorio.informacoes_obra.contratada}</p>
      </section>

      {/* Tarefas Realizadas */}
      <section>
        <h2>Tarefas Realizadas</h2>
        {relatorio.tarefas_realizadas && relatorio.tarefas_realizadas.length > 0 ? (
          <ul>
            {relatorio.tarefas_realizadas.map((tarefa, index) => (
              <li key={index}>
                <strong>{tarefa.data}:</strong> {tarefa.descricao}
              </li>
            ))}
          </ul>
        ) : (
          <p>Nenhuma tarefa registrada</p>
        )}
      </section>

      {/* Ocorrências */}
      <section>
        <h2>Ocorrências</h2>
        {relatorio.ocorrencias && relatorio.ocorrencias.length > 0 ? (
          <ul>
            {relatorio.ocorrencias.map((ocorrencia, index) => (
              <li key={index} className={`ocorrencia-${ocorrencia.tipo.toLowerCase()}`}>
                <span className="badge">{ocorrencia.tipo}</span>
                {ocorrencia.descricao}
              </li>
            ))}
          </ul>
        ) : (
          <p>Nenhuma ocorrência registrada</p>
        )}
      </section>

      {/* Fotos */}
      <section>
        <h2>Fotos</h2>
        {relatorio.fotos && relatorio.fotos.length > 0 ? (
          <div className="fotos-grid">
            {relatorio.fotos.map((foto, index) => (
              <div key={index} className="foto-item">
                <img src={foto.url} alt={`Foto ${index + 1}`} />
                <p>{new Date(foto.timestamp).toLocaleDateString()}</p>
              </div>
            ))}
          </div>
        ) : (
          <p>Nenhuma foto disponível</p>
        )}
      </section>

      {/* Botão para gerar PDF */}
      <button onClick={() => window.print()}>
        Imprimir / Gerar PDF
      </button>
    </div>
  );
};

export default RelatorioView;
```

---

## ⚠️ Tratamento de Erros

### Erros Comuns:

#### 1. **401 Unauthorized**
```json
{
  "error": "token inválido ou expirado"
}
```
**Solução:** Fazer login novamente

#### 2. **404 Not Found**
```json
{
  "error": "obra não encontrada"
}
```
**Solução:** Verificar se o `obra_id` existe

#### 3. **400 Bad Request**
```json
{
  "error": "campo 'descricao' é obrigatório"
}
```
**Solução:** Validar campos obrigatórios antes de enviar

#### 4. **500 Internal Server Error**
```json
{
  "error": "erro ao criar atividade: pq: foreign key constraint..."
}
```
**Solução:** Verificar se `responsavel_id` existe na tabela `pessoa`

### Exemplo de Interceptor de Erros:

```javascript
api.interceptors.response.use(
  (response) => response,
  (error) => {
    const status = error.response?.status;
    const message = error.response?.data?.error || 'Erro desconhecido';

    switch (status) {
      case 401:
        // Token expirado, redirecionar para login
        localStorage.removeItem('access_token');
        window.location.href = '/login';
        break;
      
      case 404:
        alert('Recurso não encontrado: ' + message);
        break;
      
      case 400:
        alert('Dados inválidos: ' + message);
        break;
      
      case 500:
        alert('Erro no servidor: ' + message);
        break;
      
      default:
        alert('Erro: ' + message);
    }

    return Promise.reject(error);
  }
);
```

---

## ✅ Boas Práticas

### 1. **Validação no Front-End**

```javascript
const validarTarefa = (tarefa) => {
  const errors = [];
  
  if (!tarefa.descricao || tarefa.descricao.trim() === '') {
    errors.push('Descrição é obrigatória');
  }
  
  if (tarefa.percentual_conclusao < 0 || tarefa.percentual_conclusao > 100) {
    errors.push('Percentual deve estar entre 0 e 100');
  }
  
  if (!['planejada', 'em_andamento', 'concluida', 'cancelada'].includes(tarefa.status)) {
    errors.push('Status inválido');
  }
  
  return errors;
};
```

### 2. **Loading States**

```javascript
const [loading, setLoading] = useState(false);
const [error, setError] = useState(null);

const salvarDiario = async () => {
  setLoading(true);
  setError(null);
  
  try {
    // ... requisições
    alert('Salvo com sucesso!');
  } catch (err) {
    setError(err.response?.data?.error || 'Erro ao salvar');
  } finally {
    setLoading(false);
  }
};
```

### 3. **Cache de Dados**

```javascript
// Usar React Query ou similar
import { useQuery } from 'react-query';

const useRelatorio = (obraId) => {
  return useQuery(['relatorio', obraId], 
    () => api.get(`/diarios/relatorio-formatado/${obraId}`).then(res => res.data),
    {
      staleTime: 5 * 60 * 1000, // 5 minutos
      cacheTime: 10 * 60 * 1000, // 10 minutos
    }
  );
};
```

### 4. **Debounce em Buscas**

```javascript
import { debounce } from 'lodash';

const buscarOcorrencias = debounce(async (termo) => {
  const response = await api.get(`/ocorrencias?search=${termo}`);
  setResultados(response.data.data);
}, 500);
```

### 5. **Compressão de Imagens**

```javascript
import imageCompression from 'browser-image-compression';

const handleImageUpload = async (file) => {
  try {
    // Comprimir antes de converter para base64
    const options = {
      maxSizeMB: 1,
      maxWidthOrHeight: 1920,
      useWebWorker: true
    };
    
    const compressedFile = await imageCompression(file, options);
    const base64 = await imageToBase64(compressedFile);
    setFoto(base64);
  } catch (error) {
    console.error('Erro ao comprimir imagem:', error);
  }
};
```

---

## 🎯 Checklist de Integração

- [ ] Implementar sistema de autenticação (login/logout)
- [ ] Criar formulário para cadastrar tarefas
- [ ] Criar formulário para cadastrar ocorrências
- [ ] Implementar upload de foto (converter para Base64)
- [ ] Criar campo de observações gerais
- [ ] Implementar envio em lote (tarefas + ocorrências + metadados)
- [ ] Criar tela de listagem de diários
- [ ] Implementar filtros (por obra, por data)
- [ ] Criar visualização de relatório formatado
- [ ] Implementar geração de PDF
- [ ] Adicionar tratamento de erros
- [ ] Implementar loading states
- [ ] Adicionar validações de formulário
- [ ] Testar com diferentes cenários

---

## 📞 Suporte

Em caso de dúvidas ou problemas:

1. Verificar logs do console do navegador
2. Verificar Network tab no DevTools
3. Validar formato dos dados enviados
4. Consultar documentação completa em `TESTES-RESULTADOS.md`

---

**Documentação criada em:** 14 de Novembro de 2025  
**Versão da API:** 2.0 (Nova Arquitetura)  
**Status:** ✅ Todos os endpoints testados e funcionando
