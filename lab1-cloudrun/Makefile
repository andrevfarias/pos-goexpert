.PHONY: up down run exec test sh build run-prod

# Inicia o ambiente de desenvolvimento
up:
	docker compose up -d

# Para o ambiente de desenvolvimento
down:
	docker compose down

# Executa a aplicação no container de desenvolvimento
run:
	docker compose exec app go run cmd/app/main.go

# Executa testes no container de desenvolvimento
test:
	docker compose exec app go test ./...

# Acessa o shell do container de desenvolvimento
sh:
	docker compose exec app sh

# Configura o arquivo de ambiente
env-setup:
	cp .env.example .env
	@echo "Edite o arquivo .env e adicione sua chave de API do WeatherAPI"

# Constrói a imagem de produção usando Dockerfile.prod
build:
	docker build -t lab1-cloudrun-app-prod -f Dockerfile.prod .

# Executa o container de produção
run-prod:
	docker run --rm -p 8080:8080 --env-file .env lab1-cloudrun-app-prod

# Executa todos os testes com cobertura
test-coverage:
	docker compose exec app go test ./... -coverprofile=coverage.out
	docker compose exec app go tool cover -html=coverage.out -o coverage.html

# Comando de ajuda
help:
	@echo "Comandos disponíveis:"
	@echo "  make up             - Inicia o ambiente de desenvolvimento com Docker Compose"
	@echo "  make down           - Para o ambiente de desenvolvimento"
	@echo "  make run            - Executa a aplicação no container de desenvolvimento"
	@echo "  make test           - Executa os testes no container de desenvolvimento"
	@echo "  make sh             - Acessa o shell do container de desenvolvimento"
	@echo "  make env-setup      - Configura o arquivo de ambiente (.env)"
	@echo "  make build          - Constrói a imagem de produção usando Dockerfile.prod"
	@echo "  make run-prod       - Executa o container de produção na porta 8080"
	@echo "  make test-coverage  - Executa testes com cobertura e gera relatório HTML"
	@echo "  make help           - Exibe esta mensagem de ajuda" 