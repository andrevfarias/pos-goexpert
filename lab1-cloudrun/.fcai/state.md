# Estado Atual do Projeto

Este arquivo fornece uma referência rápida do contexto atual do projeto, detalhando features ativas, tarefas em andamento, tarefas concluídas e prioridades futuras.

## Estrutura básica

- [x] Projeto Go inicializado
- [x] Módulos Go configurados
- [x] Estrutura de diretórios seguindo as convenções

## API

- [ ] Implementação dos endpoints REST
- [ ] Documentação da API com Swagger
- [ ] Middlewares de log

## Deployment

- [x] Configuração do Docker para desenvolvimento local
- [ ] CI/CD para Cloud Run

## Tarefas Concluídas

- Atualização do arquivo de bibliotecas Go (.fcai/project/architecture/go-libs.md) com as dependências específicas para o projeto de API de Consulta de Temperatura por CEP.
- Atualização do arquivo de variáveis de ambiente (.fcai/project/architecture/env-vars.md) com as configurações necessárias para o projeto.
- Atualização do arquivo de configuração Docker (.fcai/project/architecture/docker.md) com a configuração simplificada para o projeto.
- **Configuração do Docker para desenvolvimento local**: Criação dos arquivos Dockerfile, docker-compose.yaml e configuração das variáveis de ambiente.

## Tarefas Em Andamento

- Nenhuma tarefa em andamento no momento.

## Próximos Passos

- Implementar a estrutura de diretórios conforme a arquitetura em camadas definida
- Desenvolver os clientes HTTP para ViaCEP e WeatherAPI
- Implementar os handlers da API REST
