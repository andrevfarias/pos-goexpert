# Pós-Graduação Go Expert

## Lab 1 - Cloud Run - API de Consulta de Temperatura por CEP

Esta API permite consultar a temperatura atual em uma localidade a partir de um CEP válido. A aplicação consulta o serviço ViaCEP para obter a cidade correspondente ao CEP informado, e então consulta o WeatherAPI para obter a temperatura atual naquela localidade.

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

Copie o arquivo `.env.example` para `.env` e preencha as variáveis necessárias:

```bash
cp .env.example .env
```

Edite o arquivo `.env` e adicione sua chave de API do WeatherAPI:

```bash
WEATHER_API_KEY=sua_chave_da_weather_api
```

Você pode obter uma chave gratuita em: [WeatherAPI](https://www.weatherapi.com/)

### 3. Inicie o ambiente de desenvolvimento

```bash
docker compose up -d
```

### 4. Execute a aplicação

```bash
docker compose exec app go run cmd/app/main.go
```

A aplicação estará disponível em: http://localhost:8080

### 5. Execute os testes

```bash
docker compose exec app go test ./...
```

## Comandos Úteis

### Acessar o terminal do container

```bash
docker compose exec app sh
```

### Verificar logs da aplicação

```bash
docker compose logs -f app
```

### Parar o ambiente

```bash
docker compose down
```

## Estrutura do Projeto

```
.
├── cmd/app/                  # Ponto de entrada da aplicação
├── internal/                 # Código interno da aplicação
│   ├── entity/               # Entidades de domínio
│   ├── handlers/             # Handlers HTTP
│   ├── usecases/             # Casos de uso da aplicação
│   ├── clients/              # Clientes para APIs externas
│   │   ├── viacep/           # Cliente para API ViaCEP
│   │   └── weatherapi/       # Cliente para WeatherAPI
│   └── dto/                  # Data Transfer Objects
├── pkg/                      # Pacotes compartilháveis
│   └── adapter/              # Adaptadores e utilitários
└── Dockerfile                # Configuração do container
```

## Rotas da API

A documentação completa da API está disponível no formato OpenAPI/Swagger em `docs/swagger.yaml`.

### Endpoints Principais

- `GET /temperature?zipcode=12345678` - Consulta a temperatura atual pelo CEP

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

- 200: Sucesso
- 400: CEP não informado
- 422: CEP inválido
- 404: CEP não encontrado
- 500: Erro interno

## Objetivo

Desenvolver um sistema em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

## Requisitos

- [ ] O sistema deve receber um CEP válido de 8 digitos
- [ ] O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- [ ] O sistema deve responder adequadamente nos seguintes cenários:
  - [ ] Em caso de sucesso:
    - Código HTTP: 200
    - Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
  - [ ] Em caso de falha, caso o CEP não seja válido (com formato correto):
    - Código HTTP: 422
    - Mensagem: invalid zipcode
  - ​[ ] ​​Em caso de falha, caso o CEP não seja encontrado:
    - Código HTTP: 404
    - Mensagem: can not find zipcode
- [ ] Deverá ser realizado o deploy no Google Cloud Run.

## Dicas

- Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
- Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
- Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C \* 1,8 + 32
- Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
  - Sendo F = Fahrenheit
  - Sendo C = Celsius
  - Sendo K = Kelvin

## Entrega

- [ ] O código-fonte completo da implementação.
- [ ] Testes automatizados demonstrando o funcionamento.
- [ ] Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- [ ] Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.

ref.:
https://github.com/goexpert/cloud-run

## Integração Contínua e Deploy

O projeto utiliza GitHub Actions para automação de testes, build e deploy:

- **Testes**: Execução automática de testes em cada push/PR
- **Build**: Build automático da imagem Docker
- **Deploy**: Deploy automático no Google Cloud Run (apenas na branch main)

### Status do Build

[![CI/CD Pipeline](https://github.com/andrevfarias/goexpert/lab1-cloudrun/actions/workflows/ci.yaml/badge.svg)](https://github.com/andrevfarias/goexpert/lab1-cloudrun/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/andrevfarias/goexpert/lab1-cloudrun/branch/main/graph/badge.svg)](https://codecov.io/gh/andrevfarias/goexpert/lab1-cloudrun)

### Configuração do Deploy

Para configurar o deploy, você precisa:

1. Criar um projeto no Google Cloud Platform
2. Criar uma conta de serviço com permissões para:
   - Cloud Run Admin
   - Storage Admin
   - Service Account User
3. Baixar a chave JSON da conta de serviço
4. Configurar os seguintes secrets no GitHub:
   - `GCP_PROJECT_ID`: ID do projeto no GCP
   - `GCP_SA_KEY`: Conteúdo do arquivo JSON da conta de serviço
   - `WEATHER_API_KEY`: Chave da API WeatherAPI
