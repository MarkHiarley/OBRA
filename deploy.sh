    #!/bin/bash

    # ====================================
    # Script de Deploy Automático - API OBRA
    # ====================================
    # Este script atualiza a aplicação na VPS
    # Puxa código do GitHub, roda migrations e reinicia os containers
    # ====================================

    set -e  # Parar em caso de erro

    # Cores para output
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[1;33m'
    BLUE='\033[0;34m'
    NC='\033[0m' # Sem cor

    # Função para log colorido
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

    # Início do deploy
    echo ""
    echo "======================================"
    log_info "🚀 Iniciando Deploy da API OBRA"
    echo "======================================"
    echo ""

    # 1. Verificar se está no diretório correto
    log_info "Verificando diretório do projeto..."
    if [ ! -f "go.mod" ]; then
        log_error "Arquivo go.mod não encontrado! Execute este script na raiz do projeto."
        exit 1
    fi
    log_success "Diretório correto confirmado"

    # 2. Fazer backup da versão atual
    log_info "Criando backup da versão atual..."
    BACKUP_DIR="backup_$(date +%Y%m%d_%H%M%S)"
    if [ -f "main" ]; then
        mkdir -p backups
        cp main "backups/main_$BACKUP_DIR"
        log_success "Backup criado: backups/main_$BACKUP_DIR"
    else
        log_warning "Nenhum binário anterior encontrado para backup"
    fi

    # 3. Verificar status do Git
    log_info "Verificando status do Git..."
    if [ -n "$(git status --porcelain)" ]; then
        log_warning "Há alterações não commitadas no repositório"
        read -p "Deseja continuar mesmo assim? (s/n): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Ss]$ ]]; then
            log_error "Deploy cancelado pelo usuário"
            exit 1
        fi
    fi

    # 4. Puxar código do GitHub
    log_info "Puxando código do GitHub..."
    CURRENT_BRANCH=$(git branch --show-current)
    log_info "Branch atual: $CURRENT_BRANCH"

    git fetch origin
    git pull origin $CURRENT_BRANCH

    if [ $? -eq 0 ]; then
        log_success "Código atualizado com sucesso do GitHub"
    else
        log_error "Erro ao puxar código do GitHub"
        exit 1
    fi

    # 5. Verificar se há novas dependências Go
    log_info "Verificando dependências Go..."
    go mod tidy
    if [ $? -eq 0 ]; then
        log_success "Dependências verificadas"
    else
        log_error "Erro ao verificar dependências"
        exit 1
    fi

    # 6. Compilar o projeto
    log_info "Compilando o projeto..."
    go build -o main cmd/main.go

    if [ $? -eq 0 ]; then
        log_success "Projeto compilado com sucesso"
    else
        log_error "Erro na compilação do projeto"
        exit 1
    fi

    # 7. Parar containers Docker (se estiverem rodando)
log_info "Parando containers Docker..."
docker compose down || true
log_success "Containers parados"

# 🧹 Limpeza preventiva de containers antigos
log_info "Removendo containers antigos, se existirem..."
docker rm -f api_obras db_obras 2>/dev/null || true
log_info "Removendo rede antiga, se existir..."
docker network rm obra_obras_network 2>/dev/null || true
log_success "Ambiente limpo com sucesso"

# 8. Reconstruir e iniciar containers
log_info "Iniciando containers Docker..."
docker compose up -d --build

if [ $? -eq 0 ]; then
    log_success "Containers iniciados com sucesso"
else
    log_error "Erro ao iniciar containers"
    exit 1
fi


    # 9. Aguardar banco de dados ficar pronto
    log_info "Aguardando banco de dados ficar pronto..."
    sleep 10

    # Verificar se o banco está acessível
    MAX_ATTEMPTS=30
    ATTEMPT=0
    while [ $ATTEMPT -lt $MAX_ATTEMPTS ]; do
        if docker exec db_obras pg_isready -U obras -d obrasdb > /dev/null 2>&1; then
            log_success "Banco de dados está pronto!"
            break
        fi
        ATTEMPT=$((ATTEMPT + 1))
        echo -n "."
        sleep 2
    done

    if [ $ATTEMPT -eq $MAX_ATTEMPTS ]; then
        log_error "Timeout: Banco de dados não ficou pronto em tempo hábil"
        exit 1
    fi
    echo ""

    # 10. Executar migrations
    log_info "Executando migrations do banco de dados..."
    chmod +x run-migrations.sh
    ./run-migrations.sh

    if [ $? -eq 0 ]; then
        log_success "Migrations executadas com sucesso"
    else
        log_error "Erro ao executar migrations"
        log_warning "Verifique os logs acima para mais detalhes"
        exit 1
    fi

    # 11. Testar conexão com o banco de dados
    log_info "Testando estrutura do banco de dados..."

    # Verificar tabelas criadas
    TABLES=$(docker exec db_obras psql -U obras -d obrasdb -t -c "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';")
    TABLES=$(echo $TABLES | tr -d '[:space:]')

    if [ "$TABLES" -ge "7" ]; then
        log_success "Banco de dados possui $TABLES tabelas"
    else
        log_error "Banco de dados possui apenas $TABLES tabelas (esperado pelo menos 7)"
        exit 1
    fi

    # Listar tabelas
    log_info "Tabelas no banco de dados:"
    docker exec db_obras psql -U obras -d obrasdb -c "\dt" | grep "public"

    # 12. Verificar se a API está respondendo
    log_info "Verificando se a API está respondendo..."
    sleep 5

    # Tentar acessar a API (assumindo porta 9090)
    if curl -s http://localhost:9090/health > /dev/null 2>&1; then
        log_success "API está respondendo!"
    elif curl -s http://localhost:9090/ > /dev/null 2>&1; then
        log_success "API está respondendo!"
    else
        log_warning "API pode não estar respondendo ainda (isso é normal)"
        log_info "Verifique os logs: docker logs api_obras"
    fi

    # 13. Mostrar status dos containers
    log_info "Status dos containers:"
    docker compose ps

    # 14. Mostrar últimos logs da API
    log_info "Últimos logs da API:"
    docker logs api_obras --tail 20

    echo ""
    echo "======================================"
    log_success "🎉 Deploy concluído com sucesso!"
    echo "======================================"
    echo ""
    log_info "📊 Resumo do Deploy:"
    echo "  • Branch: $CURRENT_BRANCH"
    echo "  • Commit: $(git log -1 --pretty=format:'%h - %s')"
    echo "  • Data: $(date '+%Y-%m-%d %H:%M:%S')"
    echo "  • Tabelas no DB: $TABLES"
    echo ""
    log_info "🔍 Comandos úteis:"
    echo "  • Ver logs da API: docker logs api_obras -f"
    echo "  • Ver logs do DB: docker logs db_obras -f"
    echo "  • Acessar banco: docker exec -it db_obras psql -U obras -d obrasdb"
    echo "  • Reiniciar API: docker compose restart api_obras"
    echo "  • Parar tudo: docker compose down"
    echo ""
    log_success "✨ API está pronta para uso!"
    echo ""
