#!/bin/bash

# ====================================
# Script de Verificação de Saúde - API OBRA
# ====================================
# Verifica se tudo está funcionando corretamente
# ====================================

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

log_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

log_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

log_error() {
    echo -e "${RED}❌ $1${NC}"
}

echo ""
echo "======================================"
log_info "🏥 Verificação de Saúde da API OBRA"
echo "======================================"
echo ""

# 1. Verificar containers
log_info "Verificando containers Docker..."
API_STATUS=$(docker inspect -f '{{.State.Status}}' api_obras 2>/dev/null || echo "not_found")
DB_STATUS=$(docker inspect -f '{{.State.Status}}' db_obras 2>/dev/null || echo "not_found")

if [ "$API_STATUS" == "running" ]; then
    log_success "Container API está rodando"
else
    log_error "Container API não está rodando (status: $API_STATUS)"
fi

if [ "$DB_STATUS" == "running" ]; then
    log_success "Container DB está rodando"
else
    log_error "Container DB não está rodando (status: $DB_STATUS)"
fi

# 2. Verificar banco de dados
log_info "Verificando banco de dados..."
if docker exec db_obras pg_isready -U obras -d obrasdb > /dev/null 2>&1; then
    log_success "Banco de dados está acessível"
    
    # Contar tabelas
    TABLES=$(docker exec db_obras psql -U obras -d obrasdb -t -c "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';" | tr -d '[:space:]')
    log_info "Número de tabelas: $TABLES"
    
    # Listar tabelas
    echo ""
    log_info "Tabelas existentes:"
    docker exec db_obras psql -U obras -d obrasdb -c "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE' ORDER BY table_name;"
    
    # Verificar última migration
    echo ""
    log_info "Última migration executada:"
    docker exec db_obras psql -U obras -d obrasdb -c "SELECT version, dirty FROM schema_migrations;" 2>/dev/null || log_warning "Tabela schema_migrations não encontrada"
    
else
    log_error "Banco de dados não está acessível"
fi

# 3. Verificar API
log_info "Verificando API..."
if curl -s http://localhost:9090/ > /dev/null 2>&1; then
    log_success "API está respondendo em http://localhost:9090"
else
    log_warning "API não está respondendo (pode estar em outra porta)"
fi

# 4. Verificar logs recentes
echo ""
log_info "Últimos erros nos logs da API:"
docker logs api_obras --tail 50 2>&1 | grep -i "error\|fatal\|panic" || log_success "Nenhum erro encontrado nos logs recentes"

# 5. Verificar uso de recursos
echo ""
log_info "Uso de recursos dos containers:"
docker stats --no-stream api_obras db_obras 2>/dev/null || log_warning "Não foi possível obter estatísticas"

# 6. Verificar portas
echo ""
log_info "Portas em uso:"
docker compose ps

# 7. Verificar espaço em disco
echo ""
log_info "Espaço em disco:"
df -h | grep -E "Filesystem|/$" 

echo ""
echo "======================================"
log_success "✅ Verificação de saúde concluída"
echo "======================================"
echo ""
