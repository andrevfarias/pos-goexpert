# Estado Atual do Projeto

Este arquivo fornece uma referência rápida do contexto atual do projeto, detalhando features ativas, tarefas em andamento, tarefas concluídas e prioridades futuras.

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
