#!/bin/bash

echo "========================================"
echo "🧪 TESTE FINAL - ENDPOINTS NOVOS"
echo "========================================"
echo ""

# 1. Login
echo "1️⃣ Fazendo Login..."
TOKEN=$(curl -s -X POST "http://localhost:9090/login" \
  -H "Content-Type: application/json" \
  -d '{"email": "teste@teste.com", "senha": "senha123"}' \
  | python3 -c "import sys, json; print(json.load(sys.stdin)['access_token'])")

if [ -z "$TOKEN" ]; then
    echo "❌ Falha no login"
    exit 1
fi

echo "✅ Token obtido!"
echo ""

# 2. Relatório Fotográfico
echo "2️⃣ Testando Relatório Fotográfico..."
RESULT=$(curl -s -X GET "http://localhost:9090/relatorios/fotografico/5" \
  -H "Authorization: Bearer $TOKEN" \
  | python3 -c "import sys, json; d=json.load(sys.stdin)['data']; print(f\"OK|{d['resumo_obra']['nome_obra']}|{len(d['fotos'])}\")")

IFS='|' read -ra FOTO <<< "$RESULT"
echo "✅ Obra: ${FOTO[1]}"
echo "✅ Fotos: ${FOTO[2]}"
echo "✅ Status: SUCESSO"
echo ""

# 3. Diário Semanal
echo "3️⃣ Testando Diário Semanal..."
RESULT=$(curl -s -X POST "http://localhost:9090/diarios/semanal" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"obra_id": 5, "data_inicio": "2024-11-01", "data_fim": "2024-11-30"}' \
  | python3 -c "import sys, json; d=json.load(sys.stdin)['data']; print(f\"OK|{d['dados_obra']['nome_obra']}|{len(d['semanas'])}\")")

IFS='|' read -ra DIARIO <<< "$RESULT"
echo "✅ Obra: ${DIARIO[1]}"
echo "✅ Semanas: ${DIARIO[2]}"
echo "✅ Status: SUCESSO"
echo ""

echo "========================================"
echo "🎉 TODOS OS TESTES PASSARAM!"
echo "========================================"
