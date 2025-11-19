# 📸 RELATÓRIO FOTOGRÁFICO E DIÁRIO DE OBRAS - NOVA IMPLEMENTAÇÃO

## 🎯 Visão Geral

Implementação de **DOIS** relatórios completamente diferentes conforme solicitação do cliente:

### 1. **RELATÓRIO FOTOGRÁFICO** (Simples)
- Apenas cabeçalho + resumo da obra + fotos
- **SEM** valores financeiros
- **SEM** informações de equipe, materiais, etc.

### 2. **DIÁRIO DE OBRAS** (Por Período Semanal)
- Seleção de período (data início e fim)
- Agrupamento por semana
- Descrição do que foi executado em cada semana
- **SEM** fotos aqui

---

## 📋 1. RELATÓRIO FOTOGRÁFICO

### Endpoint
```http
GET /api/relatorios/fotografico/:obra_id
Authorization: Bearer {token}
```

### Exemplo de Requisição
```bash
curl -X GET "http://localhost:9090/relatorios/fotografico/1" \
  -H "Authorization: Bearer seu_token_aqui"
```

### Estrutura da Resposta
```json
{
  "data": {
    "cabecalho_empresa": {
      "nome_empresa": "EMPRESA CONSTRUTORA",
      "logotipo": null
    },
    "resumo_obra": {
      "nome_obra": "Construção do Prédio Comercial",
      "localizacao": "Rua das Flores, 123 - São Paulo - SP",
      "contrato_numero": "CONT-2024-001",
      "lote": "LOTE-A",
      "descricao_breve": "Construção de edifício comercial de 5 andares",
      "informacoes_gerais": "Relatório fotográfico da execução da obra"
    },
    "fotos": [
      {
        "id": 1,
        "url": "data:image/jpeg;base64,/9j/4AAQSkZJRg...",
        "titulo_legenda": "Foto do período: MANHÃ",
        "data": "2024-11-15",
        "observacao": "Fundação concluída",
        "categoria": "DIARIO"
      },
      {
        "id": 2,
        "url": "data:image/jpeg;base64,/9j/4AAQSkZJRg...",
        "titulo_legenda": "Foto do período: TARDE",
        "data": "2024-11-16",
        "observacao": "Concretagem da laje",
        "categoria": "DIARIO"
      }
    ]
  }
}
```

### O que contém:
✅ Cabeçalho com nome da empresa
✅ Resumo da obra (nome, localização, contrato, descrição)
✅ Lista de TODAS as fotos da obra
✅ Cada foto tem: título/legenda, data, observação

### O que NÃO contém:
❌ Valores financeiros
❌ Informações de equipe
❌ Materiais utilizados
❌ Equipamentos
❌ Atividades detalhadas

---

## 📅 2. DIÁRIO DE OBRAS SEMANAL

### Endpoint
```http
POST /api/diarios/semanal
Authorization: Bearer {token}
Content-Type: application/json
```

### Exemplo de Requisição
```bash
curl -X POST "http://localhost:9090/diarios/semanal" \
  -H "Authorization: Bearer seu_token_aqui" \
  -H "Content-Type: application/json" \
  -d '{
    "obra_id": 1,
    "data_inicio": "2024-11-01",
    "data_fim": "2024-11-30"
  }'
```

### Body da Requisição
```json
{
  "obra_id": 1,
  "data_inicio": "2024-11-01",
  "data_fim": "2024-11-30"
}
```

### Estrutura da Resposta
```json
{
  "data": {
    "dados_obra": {
      "nome_obra": "Construção do Prédio Comercial",
      "localizacao": "Rua das Flores, 123 - São Paulo - SP",
      "contrato_numero": "CONT-2024-001",
      "contratante": "PREFEITURA MUNICIPAL",
      "contratada": "CONSTRUTORA ABC LTDA"
    },
    "semanas": [
      {
        "numero": 1,
        "data_inicio": "2024-11-01",
        "data_fim": "2024-11-07",
        "descricao": "[2024-11-01] Escavação do terreno para fundação\nObs: Terreno argiloso, necessário reforço\n\n[2024-11-03] Instalação de formas para sapatas\n\n[2024-11-05] Concretagem das sapatas\nObs: Concreto FCK 25",
        "dias_trabalho": [
          "2024-11-01",
          "2024-11-03",
          "2024-11-05"
        ]
      },
      {
        "numero": 2,
        "data_inicio": "2024-11-08",
        "data_fim": "2024-11-14",
        "descricao": "[2024-11-10] Levantamento de alvenaria do 1º andar\n\n[2024-11-12] Instalação de vigas baldrame",
        "dias_trabalho": [
          "2024-11-10",
          "2024-11-12"
        ]
      },
      {
        "numero": 3,
        "data_inicio": "2024-11-15",
        "data_fim": "2024-11-21",
        "descricao": "Nenhuma atividade registrada nesta semana",
        "dias_trabalho": []
      }
    ]
  }
}
```

### O que contém:
✅ Dados do cabeçalho da obra (contrato, contratante, etc.)
✅ Semanas organizadas em sequência
✅ Para cada semana:
  - Número da semana
  - Data de início e fim
  - **Descrição do que foi executado** (campo editável)
  - Lista de dias que tiveram trabalho

### O que NÃO contém:
❌ Fotos
❌ Valores financeiros
❌ Relatório fotográfico

---

## 🏗️ Estrutura de Arquivos Criados

### Models
```
internal/models/
├── relatorio_fotografico.go   # Modelo do relatório fotográfico
└── diario_semanal.go          # Modelo do diário semanal
```

### Services
```
internal/services/
├── relatorio_fotografico.go   # Busca dados e fotos da obra
└── diario_semanal.go          # Agrupa diários por semana
```

### Use Cases
```
internal/usecases/
├── relatorio_fotografico.go   # Lógica de negócio do relatório fotográfico
└── diario_semanal.go          # Lógica de negócio do diário semanal
```

### Controllers
```
internal/controllers/
└── relatorio_fotografico.go   # Endpoints HTTP para ambos relatórios
```

---

## 🔧 Como Usar no Frontend

### 1. Buscar Relatório Fotográfico
```javascript
async function buscarRelatorioFotografico(obraId) {
  const response = await fetch(`/api/relatorios/fotografico/${obraId}`, {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  });
  
  const data = await response.json();
  return data.data;
}
```

### 2. Gerar Diário Semanal
```javascript
async function gerarDiarioSemanal(obraId, dataInicio, dataFim) {
  const response = await fetch('/api/diarios/semanal', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    },
    body: JSON.stringify({
      obra_id: obraId,
      data_inicio: dataInicio,  // "2024-11-01"
      data_fim: dataFim          // "2024-11-30"
    })
  });
  
  const data = await response.json();
  return data.data;
}
```

---

## 📝 Regras de Negócio

### Relatório Fotográfico
1. Busca TODAS as fotos cadastradas nos diários da obra
2. Ordena por data (mais recente primeiro)
3. Cada foto mostra: título/legenda, data e observação
4. **Não inclui** informações financeiras ou de recursos

### Diário Semanal
1. Divide o período em semanas (domingo a sábado)
2. Agrupa os diários de cada semana
3. Concatena as atividades realizadas em cada dia
4. Mostra os dias que tiveram trabalho registrado
5. **Não inclui** fotos ou valores

---

## ✅ Diferenças entre os Relatórios

| Característica | Relatório Fotográfico | Diário Semanal |
|---------------|----------------------|----------------|
| **Objetivo** | Mostrar fotos da obra | Descrever atividades executadas |
| **Período** | Todas as fotos | Período selecionado (data início/fim) |
| **Fotos** | ✅ Sim | ❌ Não |
| **Atividades** | ❌ Não | ✅ Sim |
| **Agrupamento** | Nenhum | Por semana |
| **Valores** | ❌ Não | ❌ Não |
| **Equipe/Materiais** | ❌ Não | ❌ Não |

---

## 🧪 Testando os Endpoints

### Teste do Relatório Fotográfico
```bash
# Substitua {token} e {obra_id}
curl -X GET "http://localhost:9090/relatorios/fotografico/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### Teste do Diário Semanal
```bash
# Substitua {token}
curl -X POST "http://localhost:9090/diarios/semanal" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "obra_id": 1,
    "data_inicio": "2024-11-01",
    "data_fim": "2024-11-30"
  }'
```

---

## 📌 Observações Importantes

1. **Relatório Fotográfico**: Busca fotos da tabela `diario_obra` onde `foto IS NOT NULL`
2. **Diário Semanal**: Agrupa diários por semana baseado na data
3. Ambos os relatórios **não duplicam** informações
4. São **independentes** um do outro
5. Seguem as regras específicas solicitadas pelo cliente

---

## 🚀 Próximos Passos

Para implementar no frontend:
1. Criar página de "Relatório Fotográfico" (apenas visualização de fotos)
2. Criar página de "Diário de Obras" com:
   - Seleção de período (date picker)
   - Botão "Gerar Diário"
   - Visualização por semana
   - Campo editável para descrição de cada semana

---

**Implementado em:** 19 de novembro de 2025
**Versão:** 1.0
