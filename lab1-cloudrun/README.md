# Pós-Graduação Go Expert

## Lab 1 - Cloud Run - API de Consulta de Temperatura por CEP

## Objetivo

Desenvolver um sistema em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

## Requisitos

- [x] O sistema deve receber um CEP válido de 8 digitos
- [x] O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- [x] O sistema deve responder adequadamente nos seguintes cenários:
  - [x] Em caso de sucesso:
    - Código HTTP: 200
    - Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
  - [x] Em caso de falha, caso o CEP não seja válido (com formato correto):
    - Código HTTP: 422
    - Mensagem: invalid zipcode
  - ​[X] ​​Em caso de falha, caso o CEP não seja encontrado:
    - Código HTTP: 404
    - Mensagem: can not find zipcode
- [x] Deverá ser realizado o deploy no Google Cloud Run.

## Dicas

- Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
- Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
- Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C \* 1,8 + 32
- Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
  - Sendo F = Fahrenheit
  - Sendo C = Celsius
  - Sendo K = Kelvin

## Entrega

- [x] O código-fonte completo da implementação.
- [x] Testes automatizados demonstrando o funcionamento.
- [x] Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- [x] Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.

## Solução proposta

A API permite consultar a temperatura atual em uma localidade a partir de um CEP válido. A aplicação consulta o serviço ViaCEP para obter a cidade correspondente ao CEP informado, e então consulta o WeatherAPI para obter a temperatura atual naquela localidade.

## Requisitos

- Docker
- Docker Compose

## Configuração do Ambiente de Desenvolvimento

### 1. Clone o repositório

```bash
git clone https://github.com/andrevfarias/goexpert/lab1-cloudrun.git
cd lab1-cloudrun
```

### 2. Configure as variáveis de ambiente

```bash
make env-setup
```

Edite o arquivo `.env` e adicione sua chave de API do WeatherAPI:

```bash
WEATHER_API_KEY=sua_chave_da_weather_api
```

Você pode obter uma chave gratuita em: [WeatherAPI](https://www.weatherapi.com/)

### 3. Inicie o ambiente de desenvolvimento

```bash
make up
```

### 4. Execute a aplicação

```bash
make run
```

A aplicação estará disponível em: http://localhost:8080

### 5. Execute os testes

```bash
make test
```

## Comandos Úteis

O projeto inclui um Makefile com os seguintes comandos:

```bash
# Ver todos os comandos disponíveis
make help

# Iniciar o ambiente de desenvolvimento
make up

# Parar o ambiente de desenvolvimento
make down

# Executar a aplicação
make run

# Executar os testes
make test

# Acessar o shell do container
make sh

# Configurar o arquivo de ambiente
make env-setup

# Executar testes com cobertura
make test-coverage

# Construir a imagem de produção
make build

# Executar o container de produção
make run-prod
```

## Estrutura do Projeto

.
├── cmd/app/ # Ponto de entrada da aplicação
├── internal/ # Código interno da aplicação
│ ├── entity/ # Entidades de domínio
│ ├── infra/ # Infraestrutura da aplicação
│ │ ├── api/ # APIs da aplicação
│ │ │ └── handler/ # Handlers HTTP
│ │ └── service/ # Serviços para APIs externas
│ ├── usecase/ # Casos de uso da aplicação
│ └── dto/ # Data Transfer Objects
├── Dockerfile.prod # Configuração do container para produção
└── Dockerfile # Configuração do container para desenvolvimento

## Rotas da API

### Endpoints Principais

- `GET /cep/{cep}` - Consulta a temperatura atual pelo CEP

  - Retorna a temperatura em Celsius, Fahrenheit e Kelvin
  - CEP deve ter 8 dígitos numéricos

- `GET /health` - Verifica a saúde da API

### Exemplo de Resposta

```json
{
  "temp_c": 25.5,
  "temp_f": 77.9,
  "temp_k": 298.65
}
```

### Códigos de Status

- 200: success
- 404: zipcode not found
- 422: invalid zipcode
