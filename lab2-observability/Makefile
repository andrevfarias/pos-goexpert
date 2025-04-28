.PHONY: prod dev stop logs help build status dev-only

# Cores para melhorar a visualização
GREEN = \x1b[32m
YELLOW = \x1b[33m
BLUE = \x1b[34m
NC = \x1b[0m # Sem cor

help:
	@echo "$(YELLOW)Uso: make [comando]$(NC)"
	@echo ""
	@echo "Comandos disponíveis:"
	@echo "  $(YELLOW)start$(NC)        - Inicia os serviços (modo padrão/produção) na porta 8080"
	@echo "  $(YELLOW)start-dev$(NC)    - Inicia os serviços em ambiente de desenvolvimento na porta 8080"
	@echo "  $(YELLOW)start-infra$(NC)  - Inicia apenas a infra"
	@echo "  $(YELLOW)logs$(NC)         - Mostra os logs de todos os serviços"
	@echo "  $(YELLOW)stop$(NC)         - Para todos os serviços"
	@echo "  $(YELLOW)build-prod$(NC)   - Constrói a imagem para produção"
	@echo "  $(YELLOW)build-dev$(NC)    - Constrói a imagem para desenvolvimento"
	@echo "  $(YELLOW)status$(NC)       - Mostra o status dos contêineres"
	@echo "  $(YELLOW)help$(NC)         - Mostra esta ajuda"

stop:
	@echo "$(YELLOW)Parando todos os contêineres...$(NC)"
	docker compose --profile prod --profile dev down --remove-orphans

start: stop
	@echo "$(GREEN)Iniciando serviços em modo produção...$(NC)"
	docker compose --profile prod up -d

# Iniciar o ambiente de desenvolvimento, parando apenas o serviço-a 
start-dev: stop
	@echo "$(GREEN)Iniciando serviços em modo desenvolvimento...$(NC)"
	docker compose --profile dev up -d

start-infra: stop
	@echo "$(GREEN)Iniciando infra...$(NC)"
	docker compose up -d

logs:
	@echo "$(YELLOW)Mostrando logs...$(NC)"
	docker compose --profile prod --profile dev logs -f

build-prod:
	@echo "$(BLUE)Construindo a imagem para produção...$(NC)"
	docker compose build service-a
	docker compose build service-b

build-dev:
	@echo "$(BLUE)Construindo a imagem para desenvolvimento...$(NC)"
	docker compose build service-a-dev
	docker compose build service-b-dev
	
status:
	@echo "$(YELLOW)Status dos contêineres:$(NC)"
	docker compose ps

# Define help como o target padrão
.DEFAULT_GOAL := help 