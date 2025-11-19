# 📅 EXPLICAÇÃO: AGRUPAMENTO POR SEMANA NO DIÁRIO DE OBRAS

## 🎯 Como Funciona

### 1️⃣ **Entrada**
O usuário fornece:
- `obra_id`: ID da obra
- `data_inicio`: Ex: "2024-11-01"
- `data_fim`: Ex: "2024-11-30"

### 2️⃣ **Processamento**

#### Passo 1: Buscar Diários do Período
```sql
SELECT data, atividades_realizadas, observacoes
FROM diario_obra
WHERE obra_id = $1 
  AND data >= $2  -- data_inicio
  AND data <= $3  -- data_fim
ORDER BY data ASC
```

#### Passo 2: Dividir em Semanas (7 dias)
```
Exemplo: 2024-11-01 a 2024-11-30

Semana 1: 2024-11-01 → 2024-11-07 (7 dias)
Semana 2: 2024-11-08 → 2024-11-14 (7 dias)
Semana 3: 2024-11-15 → 2024-11-21 (7 dias)
Semana 4: 2024-11-22 → 2024-11-28 (7 dias)
Semana 5: 2024-11-29 → 2024-11-30 (2 dias - última semana pode ter menos)
```

#### Passo 3: Para Cada Semana
1. **Verificar quais diários caem naquela semana**
2. **Coletar os dias que tiveram trabalho**
   - Ex: ["2024-11-05", "2024-11-07"]
3. **Concatenar as atividades realizadas**
   ```
   [2024-11-05] Escavação do terreno para fundação
   Obs: Terreno argiloso

   [2024-11-07] Instalação de formas para sapatas
   ```

### 3️⃣ **Saída**

```json
{
  "dados_obra": {
    "nome_obra": "Casa Residencial",
    "localizacao": "Fortaleza - CE",
    "contrato_numero": "CONTR-2024-001",
    "contratante": "Prefeitura",
    "contratada": "Construtora XYZ"
  },
  "semanas": [
    {
      "numero": 1,
      "data_inicio": "2024-11-01",
      "data_fim": "2024-11-07",
      "descricao": "[2024-11-05] Escavação...\n[2024-11-07] Instalação...",
      "dias_trabalho": ["2024-11-05", "2024-11-07"]
    },
    {
      "numero": 2,
      "data_inicio": "2024-11-08",
      "data_fim": "2024-11-14",
      "descricao": "Nenhuma atividade registrada nesta semana",
      "dias_trabalho": null
    }
  ]
}
```

---

## 🔧 Lógica do Código

### Loop Principal
```go
for inicio.Before(fim) || inicio.Equal(fim) {
    // Calcular fim da semana (7 dias à frente)
    fimSemana := inicio.AddDate(0, 0, 6)
    
    // Se passar do período solicitado, ajustar
    if fimSemana.After(fim) {
        fimSemana = fim
    }
    
    // Buscar diários que caem nesta semana
    for _, diario := range diarios {
        if diarioEstaNaSemana(diario, inicio, fimSemana) {
            // Adicionar à lista de dias trabalhados
            // Concatenar descrição das atividades
        }
    }
    
    // Avançar para próxima semana
    inicio = fimSemana.AddDate(0, 0, 1)
    numeroSemana++
}
```

### Verificação de Diário na Semana
```go
// O diário está dentro desta semana se:
if (dataDiario >= inicio) && (dataDiario <= fimSemana) {
    // Diário pertence a esta semana
}
```

### Concatenação de Descrições
```go
// Para cada diário da semana:
descricaoCompleta += "[2024-11-05] Escavação do terreno\n\n"
descricaoCompleta += "[2024-11-07] Instalação de formas\n"
```

---

## 📊 Exemplo Prático

### Dados de Entrada
```
Período: 01/11/2024 - 15/11/2024

Diários registrados:
- 2024-11-03: "Limpeza do terreno"
- 2024-11-05: "Escavação para fundação"
- 2024-11-12: "Concretagem das sapatas"
```

### Processamento
```
SEMANA 1 (01/11 - 07/11)
├─ Dia 03/11: "Limpeza do terreno"
├─ Dia 05/11: "Escavação para fundação"
└─ Resultado:
   - Dias trabalho: [2024-11-03, 2024-11-05]
   - Descrição: "[2024-11-03] Limpeza do terreno\n\n[2024-11-05] Escavação..."

SEMANA 2 (08/11 - 14/11)
├─ Dia 12/11: "Concretagem das sapatas"
└─ Resultado:
   - Dias trabalho: [2024-11-12]
   - Descrição: "[2024-11-12] Concretagem das sapatas"

SEMANA 3 (15/11 - 15/11)
└─ Resultado:
   - Dias trabalho: []
   - Descrição: "Nenhuma atividade registrada nesta semana"
```

---

## 🎯 Por Que 7 Dias?

Cada semana tem **7 dias** (de segunda a domingo):
- Facilita o planejamento semanal
- Padrão comum em diários de obra
- Permite acompanhamento cronológico claro

Se a última semana tiver menos de 7 dias (como no exemplo da Semana 5: apenas 2 dias), o código ajusta automaticamente.

---

## 💡 Vantagens

✅ **Organização Clara**: Agrupa atividades por semana
✅ **Flexível**: Funciona com qualquer período
✅ **Completo**: Mostra semanas sem atividade também
✅ **Rastreável**: Lista os dias específicos de trabalho
✅ **Descritivo**: Concatena todas as atividades da semana

---

## 🔄 Fluxo Completo

```
1. Usuário envia período (01/11 - 30/11)
   ↓
2. Sistema busca todos os diários do período
   ↓
3. Divide em semanas de 7 dias
   ↓
4. Para cada semana:
   - Filtra diários que caem nela
   - Coleta dias de trabalho
   - Concatena descrições
   ↓
5. Retorna lista de semanas organizadas
```

---

**Implementado em:** `internal/services/diario_semanal.go`
**Método:** `agruparPorSemana()`
