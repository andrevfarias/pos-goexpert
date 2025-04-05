# Estado Atual do Projeto

Este arquivo fornece uma referência rápida do contexto atual do projeto,
detalhando features ativas, tarefas em andamento, tarefas concluídas e
prioridades futuras.

## Componentes Implementados

### Core

- ✅ Configuração centralizada
- ✅ Entidades de domínio
- ✅ Interfaces de serviço
- ✅ Use cases
  - ✅ GetTemperatureByZipCode

### Infraestrutura

- ✅ Clientes HTTP
  - ✅ ViaCEP
  - ✅ WeatherAPI
- ✅ Adaptadores de serviço
- ✅ Handlers HTTP
- ✅ Roteamento

### Testes

- ✅ Testes de entidades
  - ✅ Address (100% cobertura)
  - ✅ Weather (100% cobertura)
- ✅ Testes de clientes HTTP
  - ✅ ViaCEP (94.1% cobertura)
  - ✅ WeatherAPI (90.5% cobertura)
- ✅ Testes de use cases
  - ✅ GetTemperatureByZipCode (100% cobertura)
- ✅ Testes de handlers HTTP
  - ✅ TemperatureHandler (81.8% cobertura)
- ✅ Testes de middlewares
  - ✅ Logger (94.1% cobertura)
  - ✅ Recoverer (94.1% cobertura)
- [ ] Testes de integração

### Documentação

- ✅ README principal
- ✅ Documentação de features
- ✅ Documentação da API (Swagger)
- [ ] Guia de contribuição

### DevOps

- ✅ Configuração Docker
- ✅ Pipeline CI/CD
  - ✅ Testes automatizados
  - ✅ Build automatizado
  - ✅ Deploy automatizado
- [ ] Configuração Cloud Run

## Próximas Tarefas

1. [ ] Configurar deploy no Cloud Run
2. [ ] Implementar logging estruturado
3. [ ] Adicionar métricas e monitoramento
4. [ ] Implementar testes de integração

## Observações

- Todos os testes implementados estão passando
- Cobertura média de testes acima de 90%
- Projeto seguindo Clean Architecture e DDD
- Containers Docker configurados e funcionando
- Pipeline CI/CD configurado com GitHub Actions

## Features em Andamento

### Testing Implementation

- ✅ Implementação de testes para entidades de domínio
  - ✅ Testes unitários para Weather (100% cobertura)
  - ✅ Testes unitários para Address (100% cobertura)
- 🔄 Testes unitários para casos de uso
  - 🔄 GetTemperatureByZipCode em implementação
  - Mocks dos serviços em preparação
- ⏳ Testes de integração para adaptadores (Planejado)

## Features Concluídas

- ✅ Testes das entidades de domínio
- ✅ Testes dos clientes HTTP
- ✅ Testes dos use cases
- ✅ Testes do handler de temperatura

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
