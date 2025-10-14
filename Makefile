# Carregar variáveis do .env
include .env
export

# URL de conexão do Postgres
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

.PHONY: help migrate-up migrate-down migrate-create docker-up docker-down run install-migrate

help: ## Mostra comandos disponíveis
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

install-migrate: ## Instala golang-migrate
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

docker-up: ## Sobe o Postgres
	docker-compose up -d db_obras

docker-down: ## Para o Docker
	docker-compose down

migrate-up: ## Roda todas as migrations
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down: ## Reverte a última migration
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

migrate-create: ## Cria nova migration (uso: make migrate-create NAME=create_pessoa)
	migrate create -ext sql -dir migrations -seq $(NAME)

run: ## Roda a API
	go run cmd/main.go