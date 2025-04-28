# Pós-Graduação Go Expert

## Lab 2 - Observability

Este projeto implementa dois microserviços em Go que, em conjunto, permitem consultar a temperatura atual de uma cidade a partir de um CEP. Os serviços são instrumentados com OpenTelemetry e Zipkin para observabilidade.

## Objetivo

Desenvolver um sistema distribuído em Go que receba um CEP, identifique a cidade e retorne o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin) juntamente com o nome da cidade. O sistema implementa rastreamento distribuído com OpenTelemetry e Zipkin.

## Arquitetura

O sistema é composto por dois serviços:

1. **Serviço A (Input)**

   - Recebe o CEP via API REST
   - Valida o formato do CEP
   - Encaminha a requisição para o Serviço B

2. **Serviço B (Orquestração)**
   - Consulta o CEP em uma API externa (ViaCEP)
   - Obtém a temperatura atual da cidade usando uma API de clima (WeatherAPI)
   - Formata e retorna os dados de temperatura

## Requisitos

### Serviço A (responsável pelo input)

- Recebe um CEP de 8 dígitos via POST, através do schema: `{ "cep": "29902555" }`
  **- Valida se o input é válido (contém 8 dígitos) e é uma STRING**
- Caso seja válido, encaminha para o Serviço B via HTTP
- Caso não seja válido, retorna:
  - Código HTTP: 422
  - Mensagem: invalid zipcode

### Serviço B (responsável pela orquestração)

- Recebe um CEP válido de 8 dígitos
- Realiza a pesquisa do CEP e encontra o nome da localização
- Consulta as temperaturas e formata-as em Celsius, Fahrenheit e Kelvin
- Retorna o nome da cidade e as temperaturas
- Respostas HTTP:
  - Em caso de sucesso:
    - Código HTTP: 200
    - Response Body: `{ "city": "São Paulo", "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }`
  - Em caso de CEP com formato inválido:
    - Código HTTP: 422
    - Mensagem: invalid zipcode
  - Em caso de CEP não encontrado:
    - Código HTTP: 404
    - Mensagem: can not find zipcode

### Observabilidade

- Implementação de rastreamento distribuído entre Serviço A e Serviço B com OpenTelemetry
- Utilização de spans para medir o tempo de resposta do serviço de busca de CEP e busca de temperatura
- Visualização dos traces com Zipkin

## Tecnologias Utilizadas

- Go
- Docker/Docker Compose
- OpenTelemetry
- Zipkin
- APIs externas:
  - ViaCEP
  - WeatherAPI

## Como Executar

### Pré-requisitos

- Docker
- Docker Compose
- Make (opcional, para facilitar o uso dos comandos)

### Configuração

1. Clone o repositório:

```bash
git clone https://github.com/andrevfarias/go-expert/lab2-observability.git
cd lab2-observability
```

2. Configure os arquivos de variáveis de ambiente:

```bash
cp .env.example .env
```

3. Edite o arquivo `.env` e adicione sua chave de API do WeatherAPI:

```
WEATHER_API_KEY=sua_chave_api
```

Você pode obter uma chave gratuita em [WeatherAPI](https://www.weatherapi.com/).

### Execução com Make

O projeto inclui um Makefile para facilitar o uso dos comandos mais comuns:

```bash
# Iniciar serviços em modo produção
make start

# Parar todos os serviços
make stop

# Verificar logs dos serviços
make logs

# Verificar status dos contêineres
make status

# Listar todos os comandos disponíveis
make help
```

### Execução Manual (sem Make)

1. Inicie os serviços com Docker Compose:

```bash
docker-compose --profile prod up -d
```

2. Os serviços estarão disponíveis em:
   - Serviço A: http://localhost:8080
   - Zipkin: http://localhost:9411

### API Endpoints

#### Serviço A

**Consulta de temperatura por CEP**

- **URL**: `/cep`
- **Método**: `POST`
- **Payload**: `{ "cep": "01001000" }`
- **Respostas**:
  - **Sucesso (200)**: `{ "city": "São Paulo", "temp_C": 23.5, "temp_F": 74.3, "temp_K": 296.65 }`
  - **CEP Inválido (422)**: `{ "message": "invalid zipcode" }`
  - **CEP Não Encontrado (404)**: `{ "message": "can not find zipcode" }`

## Visualizando Traces no Zipkin

1. Acesse a interface do Zipkin em http://localhost:9411
2. Realize algumas requisições para a API
3. No Zipkin, você poderá visualizar os traces distribuídos entre os serviços
4. Analise os spans para verificar o tempo de resposta de cada operação
