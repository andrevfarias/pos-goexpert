# Deploy no Cloud Run

## Descrição

Configuração do deploy da API no Google Cloud Run, permitindo que a aplicação seja executada em um ambiente serverless e escalável.

## Status

✅ Configurado

## Objetivos

- ✅ Configurar deploy automático via GitHub Actions
- ✅ Configurar variáveis de ambiente
- ✅ Habilitar acesso público à API
- ✅ Documentar processo de deploy

## Configuração Necessária

1. No Google Cloud Platform:

   - Criar um projeto
   - Habilitar a API do Cloud Run
   - Criar uma conta de serviço com permissão de Cloud Run Admin

2. No GitHub:
   - Configurar secret `GCP_SA_KEY` com a chave da conta de serviço
   - Configurar secret `WEATHER_API_KEY` com a chave da API WeatherAPI

## Deploy

O deploy é realizado automaticamente pelo GitHub Actions quando:

- Um push é feito na branch main
- Todos os testes passam

O processo de deploy:

1. Autentica no Google Cloud
2. Faz o build do código fonte
3. Realiza o deploy no Cloud Run
4. Configura as variáveis de ambiente

## Observações

- Deploy simplificado usando `--source` do Cloud Run
- Acesso público habilitado
- Ambiente de produção único
- Variáveis de ambiente configuradas no deploy
