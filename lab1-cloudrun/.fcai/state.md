# Estado Atual do Projeto

Este arquivo fornece uma referência rápida do contexto atual do projeto,
detalhando features ativas, tarefas em andamento, tarefas concluídas e
prioridades futuras.

## Componentes Implementados

### 1. Configuração

- ✅ Configuração centralizada usando Viper
- ✅ Carregamento de variáveis do arquivo `.env`
- ✅ Validação de configurações obrigatórias

### 2. Domínio

- ✅ Entidade `Weather` com comportamentos de conversão de temperatura
  - Cálculo de temperatura em Fahrenheit
  - Cálculo de temperatura em Kelvin
  - Serialização JSON com todas as unidades
  - ✅ Testes unitários com 100% de cobertura
- ✅ Entidade `Address` com validações e formatação
  - Validação de CEP
  - Formatação de CEP (00000-000)
  - Limpeza de CEP (remoção de caracteres especiais)
  - ✅ Testes unitários com 100% de cobertura
- ✅ Interfaces de serviço bem definidas
  - `WeatherService`
  - `ZipCodeFinder`

### 3. Casos de Uso

- ✅ GetTemperatureByZipCode
  - Integração com serviços de CEP e clima
  - Tratamento de erros
  - Retorno padronizado
  - 🔄 Testes unitários em implementação

### 4. Infraestrutura

- ✅ Clientes HTTP
  - Cliente ViaCEP
  - Cliente WeatherAPI
- ✅ Adaptadores de Serviço
  - Implementação do WeatherService
  - Implementação do ZipCodeFinder
- ✅ API REST
  - Handler de temperatura
  - Roteamento com Chi
  - Tratamento de erros HTTP

## Features em Andamento

### Testing Implementation

- ✅ Implementação de testes para entidades de domínio
  - ✅ Testes unitários para Weather (100% cobertura)
  - ✅ Testes unitários para Address (100% cobertura)
- 🔄 Testes unitários para casos de uso
  - 🔄 GetTemperatureByZipCode em implementação
  - Mocks dos serviços em preparação
- ⏳ Testes de integração para adaptadores (Planejado)

## Melhorias Implementadas

1. Remoção de lógica de conversão do cliente WeatherAPI
2. Adição de comportamentos nas entidades de domínio
3. Implementação de validações no domínio
4. Configuração via arquivo `.env`
5. Testes unitários completos para entidades de domínio

## Próximos Passos

1. [ ] Implementar testes automatizados
   - [✅] Testes unitários para entidades
   - [🔄] Testes unitários para casos de uso
   - [ ] Testes de integração para adaptadores
2. [ ] Adicionar documentação Swagger
3. [ ] Configurar CI/CD para Google Cloud Run
4. [ ] Implementar logging estruturado
5. [ ] Adicionar métricas e monitoramento

## Legenda

- ✅ Concluído
- 🔄 Em Andamento
- ⏳ Planejado
- ❌ Bloqueado

## Estrutura básica

- [x] Projeto Go inicializado
- [x] Módulos Go configurados
- [x] Estrutura de diretórios seguindo as convenções

## API

- [x] Implementação dos endpoints REST
- [ ] Documentação da API com Swagger
- [x] Middlewares de log

## Deployment

- [x] Configuração do Docker para desenvolvimento local
- [ ] CI/CD para Cloud Run

## Tarefas Concluídas

- Atualização do arquivo de bibliotecas Go (.fcai/project/architecture/go-libs.md) com as dependências específicas para o projeto de API de Consulta de Temperatura por CEP.
- Atualização do arquivo de variáveis de ambiente (.fcai/project/architecture/env-vars.md) com as configurações necessárias para o projeto.
- Atualização do arquivo de configuração Docker (.fcai/project/architecture/docker.md) com a configuração simplificada para o projeto.
- **Configuração do Docker para desenvolvimento local**: Criação dos arquivos Dockerfile, docker-compose.yaml e configuração das variáveis de ambiente.
- **Implementação da configuração centralizada**: Criação do arquivo `internal/config/config.go` utilizando a biblioteca Viper para gerenciar as configurações da aplicação de forma centralizada.
- **Implementação da estrutura de diretórios em camadas**: Criação das camadas Domain, Application e Infrastructure seguindo os princípios de Clean Architecture.
- **Implementação das entidades de domínio**: Criação das entidades Address e Weather.
- **Implementação dos casos de uso**: Criação do caso de uso GetTemperatureByZipCode para orquestrar a obtenção da temperatura a partir de um CEP.
- **Implementação dos clientes HTTP**: Desenvolvimento dos clientes para as APIs ViaCEP e WeatherAPI.
- **Implementação da API REST**: Desenvolvimento dos handlers e router para expor a funcionalidade via API REST.

## Tarefas Em Andamento

- Nenhuma tarefa em andamento no momento.

## Próximos Passos

- Implementar testes automatizados para as diferentes camadas
- Adicionar documentação Swagger para a API
- Configurar CI/CD para implantação no Google Cloud Run
