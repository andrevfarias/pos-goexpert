# Pós-Graduação Go Expert

## Lab 3 - Sistema de Leilão com Fechamento Automático

### Objetivo

Desenvolver um sistema de leilão em Go que permite criar leilões, fazer lances e gerenciar usuários. O sistema possui fechamento automático de leilões baseado em tempo configurável e processamento de lances em lote para alta performance.

O projeto implementa:
- Criação de leilões com fechamento automático via goroutines
- Sistema de lances com processamento em lote
- Validação de usuários e leilões ativos
- API REST para gerenciamento

### Requisitos

- Docker
- Docker Compose
- Make (opcional)

### Como Executar

```bash
# Iniciar todos os serviços
make start

# Ambiente de desenvolvimento
make start-dev

# Executar testes
make test

# Parar serviços
make stop
```

### API Endpoints

A aplicação roda na porta **8080**.

**Leilões:**
- `POST /auction` - Criar leilão
- `GET /auction` - Listar leilões
- `GET /auction/{id}` - Buscar leilão por ID
- `GET /auction/winner/{id}` - Buscar lance vencedor

**Lances:**
- `POST /bid` - Criar lance
- `GET /bid/{auctionId}` - Listar lances por leilão

**Usuários:**
- `GET /user/{userId}` - Buscar usuário por ID

### Exemplo de Uso

```bash
# Criar leilão
curl -X POST http://localhost:8080/auction \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "iPhone 13",
    "category": "Electronics", 
    "description": "iPhone 13 em excelente estado",
    "condition": 1
  }'

# Criar lance
curl -X POST http://localhost:8080/bid \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "123e4567-e89b-12d3-a456-426614174000",
    "auction_id": "123e4567-e89b-12d3-a456-426614174001", 
    "amount": 1000.50
  }'
```

### Variáveis de Ambiente

| Variável                | Descrição                        | Valor Padrão |
| ----------------------- | -------------------------------- | ------------ |
| `MONGODB_URL`           | URL de conexão com MongoDB       | -            |
| `MONGODB_DB`            | Nome do banco de dados           | -            |
| `AUCTION_INTERVAL`      | Tempo para fechamento automático | `5m`         |
| `MAX_BATCH_SIZE`        | Tamanho máximo do lote de lances | `5`          |
| `BATCH_INSERT_INTERVAL` | Intervalo para processar lotes   | `3m`         |
