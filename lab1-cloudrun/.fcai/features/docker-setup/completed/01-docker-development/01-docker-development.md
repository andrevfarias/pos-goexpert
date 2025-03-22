# Tarefa: Configuração do Docker para Desenvolvimento Local

## Descrição

Esta tarefa envolve a criação da configuração Docker para desenvolvimento local da aplicação de consulta de temperatura por CEP. Será necessário criar um Dockerfile e um arquivo docker-compose.yaml que permitam o desenvolvimento da aplicação em um ambiente containerizado.

## Objetivos

- Criar um Dockerfile para o ambiente de desenvolvimento
- Configurar o Docker Compose para orquestrar o serviço
- Configurar volumes para desenvolvimento dinâmico
- Definir variáveis de ambiente para desenvolvimento local
- Documentar o processo de inicialização e uso do ambiente Docker

## Requisitos Técnicos

1. **Dockerfile**:

   - Utilizar imagem base `golang:1.24-alpine`
   - Configurar diretório de trabalho `/app`
   - Instalar ferramentas de desenvolvimento necessárias

2. **Docker Compose**:

   - Definir o serviço `app` para a aplicação
   - Configurar bind mount para o código-fonte
   - Expor a porta 8080
   - Configurar o carregamento de variáveis de ambiente de um arquivo `.env`
   - Configurar TTY para manter o contêiner ativo

3. **Variáveis de Ambiente**:
   - Criar arquivo `.env` com as variáveis necessárias para desenvolvimento
   - Incluir variáveis para as APIs externas (ViaCEP e WeatherAPI)

## Critérios de Aceitação

- Os contêineres devem ser inicializados corretamente com `docker compose up -d`
- As alterações no código devem ser refletidas automaticamente sem reconstrução de imagem
- O ambiente deve permitir a execução de comandos Go (`go run`, `go test`, etc.)
- A aplicação deve ser acessível na porta 8080 da máquina host
- A documentação deve incluir instruções claras para inicialização e uso do ambiente

## Passos Sugeridos

1. Criar arquivo Dockerfile na raiz do projeto
2. Criar arquivo docker-compose.yaml na raiz do projeto
3. Criar arquivo .env.example na raiz do projeto
4. Testar a inicialização dos contêineres
5. Testar a execução de comandos dentro do contêiner
6. Documentar o processo de inicialização e uso

## Status

- [x] Criação do Dockerfile
- [x] Configuração do Docker Compose
- [x] Criação do arquivo .env.example
- [x] Testes de inicialização e execução
- [x] Documentação do ambiente Docker

## Recursos

- Arquivo de configuração Docker em `.fcai/project/architecture/docker.md`
- Arquivo de variáveis de ambiente em `.fcai/project/architecture/env-vars.md`
