#!/bin/bash
# Script para adicionar foto a uma obra
# Uso: ./adicionar_foto_obra.sh <obra_id> <caminho_da_foto>

set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configurações
BASE_URL="http://localhost:9090"
EMAIL="teste@teste.com"
SENHA="senha123"

# Verificar argumentos
if [ "$#" -ne 2 ]; then
    echo -e "${RED}❌ Uso incorreto${NC}"
    echo "Uso: $0 <obra_id> <caminho_da_foto>"
    echo "Exemplo: $0 5 /caminho/para/foto.jpg"
    exit 1
fi

OBRA_ID=$1
FOTO_PATH=$2

# Verificar se a foto existe
if [ ! -f "$FOTO_PATH" ]; then
    echo -e "${RED}❌ Arquivo não encontrado: $FOTO_PATH${NC}"
    exit 1
fi

# Verificar tipo do arquivo
MIME_TYPE=$(file --mime-type -b "$FOTO_PATH")
if [[ ! "$MIME_TYPE" =~ ^image/ ]]; then
    echo -e "${YELLOW}⚠️  Aviso: Arquivo não parece ser uma imagem (tipo: $MIME_TYPE)${NC}"
    read -p "Deseja continuar? (s/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Ss]$ ]]; then
        exit 1
    fi
fi

# Verificar tamanho do arquivo (max 5MB)
FILE_SIZE=$(stat -c%s "$FOTO_PATH" 2>/dev/null || stat -f%z "$FOTO_PATH" 2>/dev/null)
MAX_SIZE=$((5 * 1024 * 1024))
if [ "$FILE_SIZE" -gt "$MAX_SIZE" ]; then
    echo -e "${RED}❌ Arquivo muito grande: $(($FILE_SIZE / 1024 / 1024))MB (máximo: 5MB)${NC}"
    exit 1
fi

echo -e "${BLUE}📸 Adicionando foto à obra${NC}"
echo "================================="
echo "Obra ID: $OBRA_ID"
echo "Foto: $FOTO_PATH"
echo "Tamanho: $(($FILE_SIZE / 1024))KB"
echo "================================="
echo ""

# 1. Fazer login
echo -e "${YELLOW}🔐 Fazendo login...${NC}"
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/login" \
    -H "Content-Type: application/json" \
    -d "{\"email\": \"$EMAIL\", \"senha\": \"$SENHA\"}")

TOKEN=$(echo "$LOGIN_RESPONSE" | python3 -c "import sys, json; print(json.load(sys.stdin).get('access_token', ''))" 2>/dev/null)

if [ -z "$TOKEN" ]; then
    echo -e "${RED}❌ Falha no login${NC}"
    echo "$LOGIN_RESPONSE"
    exit 1
fi

echo -e "${GREEN}✅ Login realizado${NC}"
echo ""

# 2. Buscar dados atuais da obra
echo -e "${YELLOW}📋 Buscando dados da obra...${NC}"
OBRA_RESPONSE=$(curl -s -X GET "$BASE_URL/obras/$OBRA_ID" \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json")

# Verificar se a obra existe
if echo "$OBRA_RESPONSE" | grep -q "error"; then
    echo -e "${RED}❌ Erro ao buscar obra${NC}"
    echo "$OBRA_RESPONSE"
    exit 1
fi

# Extrair dados da obra
OBRA_DATA=$(echo "$OBRA_RESPONSE" | python3 -c "import sys, json; print(json.dumps(json.load(sys.stdin).get('data', {})))")

if [ "$OBRA_DATA" = "{}" ]; then
    echo -e "${RED}❌ Obra não encontrada${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Obra encontrada${NC}"
echo ""

# 3. Converter foto para Base64
echo -e "${YELLOW}🖼️  Convertendo foto para Base64...${NC}"
if [[ "$MIME_TYPE" == "image/jpeg" ]] || [[ "$MIME_TYPE" == "image/jpg" ]]; then
    PREFIX="data:image/jpeg;base64,"
elif [[ "$MIME_TYPE" == "image/png" ]]; then
    PREFIX="data:image/png;base64,"
elif [[ "$MIME_TYPE" == "image/gif" ]]; then
    PREFIX="data:image/gif;base64,"
else
    PREFIX="data:image/jpeg;base64,"
fi

FOTO_BASE64="${PREFIX}$(base64 -w 0 "$FOTO_PATH" 2>/dev/null || base64 "$FOTO_PATH" | tr -d '\n')"
echo -e "${GREEN}✅ Foto convertida (${#FOTO_BASE64} caracteres)${NC}"
echo ""

# 4. Preparar payload com todos os campos
echo -e "${YELLOW}📝 Preparando dados para atualização...${NC}"

# Criar JSON com a foto adicionada
UPDATE_PAYLOAD=$(echo "$OBRA_DATA" | python3 -c "
import sys, json

try:
    obra = json.load(sys.stdin)
    
    # Garantir que campos obrigatórios estão presentes
    required_fields = [
        'nome', 'contrato_numero', 'contratante_id', 'responsavel_id',
        'data_inicio', 'prazo_dias', 'data_fim_prevista', 'orcamento',
        'status', 'endereco_rua', 'endereco_numero', 'endereco_bairro',
        'endereco_cidade', 'endereco_estado'
    ]
    
    for field in required_fields:
        if field not in obra or obra[field] is None:
            print(f'Erro: campo obrigatório ausente: {field}', file=sys.stderr)
            sys.exit(1)
    
    # Adicionar a foto
    foto_base64 = sys.argv[1] if len(sys.argv) > 1 else ''
    obra['foto'] = foto_base64
    
    # Remover campos que não devem ser enviados
    for key in ['id', 'created_at', 'updated_at', 'ativo']:
        obra.pop(key, None)
    
    print(json.dumps(obra))
except Exception as e:
    print(f'Erro ao processar dados: {e}', file=sys.stderr)
    sys.exit(1)
" "$FOTO_BASE64")

if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Erro ao preparar payload${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Payload preparado${NC}"
echo ""

# 5. Atualizar obra
echo -e "${YELLOW}🚀 Atualizando obra...${NC}"
UPDATE_RESPONSE=$(curl -s -X PUT "$BASE_URL/obras/$OBRA_ID" \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    -d "$UPDATE_PAYLOAD")

# Verificar resposta
if echo "$UPDATE_RESPONSE" | grep -q "error"; then
    echo -e "${RED}❌ Erro ao atualizar obra${NC}"
    echo "$UPDATE_RESPONSE" | python3 -m json.tool 2>/dev/null || echo "$UPDATE_RESPONSE"
    exit 1
fi

echo -e "${GREEN}✅ Obra atualizada com sucesso!${NC}"
echo ""

# 6. Verificar no relatório fotográfico
echo -e "${YELLOW}🔍 Verificando foto no relatório...${NC}"
RELATORIO_RESPONSE=$(curl -s -X GET "$BASE_URL/relatorios/fotografico/$OBRA_ID" \
    -H "Authorization: Bearer $TOKEN")

FOTO_OBRA=$(echo "$RELATORIO_RESPONSE" | python3 -c "
import sys, json
try:
    data = json.load(sys.stdin)
    foto = data.get('data', {}).get('resumo_obra', {}).get('foto_obra')
    if foto:
        print(f'PRESENTE ({len(foto)} caracteres)')
    else:
        print('AUSENTE')
except:
    print('ERRO')
" 2>/dev/null)

if [ "$FOTO_OBRA" = "AUSENTE" ] || [ "$FOTO_OBRA" = "ERRO" ]; then
    echo -e "${RED}❌ Foto não encontrada no relatório${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Foto presente no relatório: $FOTO_OBRA${NC}"
echo ""

# Resumo final
echo "================================="
echo -e "${GREEN}🎉 SUCESSO!${NC}"
echo "================================="
echo "A foto foi adicionada à obra $OBRA_ID"
echo "Você pode visualizar no relatório fotográfico:"
echo "GET $BASE_URL/relatorios/fotografico/$OBRA_ID"
echo ""
