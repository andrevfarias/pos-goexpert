# Pós-Graduação Go Expert

## Desafio 4 - Rate Limiter

### Objetivo

Desenvolver um rate limiter em Go que possa ser configurado para limitar o número máximo de requisições por segundo com base em um endereço IP específico ou em um token de acesso.

### Descrição

O objetivo deste desafio é criar um rate limiter em Go que possa ser utilizado para controlar o tráfego de requisições para um serviço web. O rate limiter deve ser capaz de limitar o número de requisições com base em dois critérios:

1. Endereço IP
   O rate limiter deve restringir o número de requisições recebidas de um único endereço IP dentro de um intervalo de tempo definido.
2. Token de Acesso
   O rate limiter deve também limitar as requisições baseadas em um token de acesso único, permitindo diferentes limites de taxa para diferentes tokens. O Token deve ser informado no header no seguinte formato:

   - `API_KEY: <TOKEN>`

   As configurações de limite do token de acesso devem se sobrepor às do IP. Ex: Se o limite por IP é de 10 req/s e a de um determinado token é de 100 req/s, o rate limiter deve utilizar as informações do token.

### Requisitos

- [x] O rate limiter deve poder trabalhar como um middleware que é injetado ao servidor web
- [x] O rate limiter deve permitir a configuração do número máximo de requisições permitidas por segundo
- [x] O rate limiter deve ter a opção de escolher o tempo de bloqueio do IP ou do Token caso a quantidade de requisições tenha sido excedida
- [x] As configurações de limite devem ser realizadas via variáveis de ambiente ou em um arquivo ".env" na pasta raiz
- [x] Deve ser possível configurar o rate limiter tanto para limitação por IP quanto por token de acesso
- [x] O sistema deve responder adequadamente quando o limite é excedido:
  - Código HTTP: 429
  - Mensagem: "You have reached the maximum number of requests or actions allowed within a certain time frame"
- [x] Todas as informações de "limiter" devem ser armazenadas e consultadas de um banco de dados Redis
- [x] Possui uma "strategy" que permite trocar facilmente o Redis por outro mecanismo de persistência
- [x] A lógica do limiter está separada do middleware

### Requisitos técnicos

- Docker e Docker Compose

### Configuração

O Rate Limiter é configurado através do arquivo `.env` na raiz do projeto:

| Variável              | Descrição                                                | Exemplo                                                         |
| --------------------- | -------------------------------------------------------- | --------------------------------------------------------------- |
| `WEB_SERVER_PORT`     | Porta do servidor web                                    | `8080`                                                          |
| `IP_RATE_LIMIT`       | Número máximo de requisições por segundo por IP          | `3`                                                             |
| `API_KEYS_RATE_LIMIT` | Lista de tokens API e seus limites em formato JSON       | `[{"key":"key1","rate_limit":1},{"key":"key2","rate_limit":2}]` |
| `BLOCK_TIME_SECONDS`  | Tempo de bloqueio em segundos quando o limite é excedido | `15`                                                            |
| `STORAGE_TYPE`        | Tipo de armazenamento (redis ou memory)                  | `redis`                                                         |
| `REDIS_HOST`          | Endereço do servidor Redis                               | `localhost:6379`                                                |
| `REDIS_PASSWORD`      | Senha do Redis                                           | `123456`                                                        |
| `REDIS_DB`            | Número do banco de dados Redis                           | `0`                                                             |

### Utilização

Na raiz do projeto, execute os seguintes comandos:

```bash
# Construir a imagem
make build

# Iniciar o serviço
make start

# O servidor estará disponível em http://localhost:8080

# Executar testes de carga
make test

# Parar o serviço
make stop
```

### Exemplos práticos

```bash
# requisição com limiter baseado em ip
curl http://localhost:8080/
# requisição com limiter baseado em api_key
curl -H "API_KEY: key1" http://localhost:8080/
```

### Persistência de dados

O sistema suporta dois mecanismos de armazenamento:

1. **Redis** (padrão): Ideal para ambientes de produção e distribuídos
2. **Memória**: Útil para desenvolvimento e testes locais

A escolha é feita pela variável `STORAGE_TYPE` no arquivo `.env`.

## Status da entrega

- [x] O código-fonte completo da implementação
- [x] Documentação explicando como o rate limiter funciona e como ele pode ser configurado
- [x] Testes automatizados demonstrando a eficácia e a robustez do rate limiter
- [x] Utilização de docker/docker-compose para testes da aplicação
- [x] O servidor web responde na porta 8080
