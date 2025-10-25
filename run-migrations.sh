#!/bin/bash

# Script para rodar migrations no Docker
# Uso: ./run-migrations.sh

echo "🔄 Executando migrations do banco de dados..."

# Configuração
DB_USER="obras"
DB_NAME="obrasdb"
CONTAINER_NAME="db_obras"

# Verifica se o container está rodando
if ! docker ps | grep -q $CONTAINER_NAME; then
    echo "❌ Container $CONTAINER_NAME não está rodando!"
    echo "Execute: docker compose up -d"
    exit 1
fi

# Aguarda o banco estar pronto
echo "⏳ Aguardando banco de dados ficar pronto..."
sleep 3

# Executa as migrations em ordem
for migration in migrations/*.up.sql; do
    echo "📝 Executando: $(basename $migration)"
    docker exec -i $CONTAINER_NAME psql -U $DB_USER -d $DB_NAME < "$migration"
    
    if [ $? -eq 0 ]; then
        echo "✅ $(basename $migration) executado com sucesso"
    else
        echo "❌ Erro ao executar $(basename $migration)"
        exit 1
    fi
done

echo "🎉 Todas as migrations foram executadas com sucesso!"
