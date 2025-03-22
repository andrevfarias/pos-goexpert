# Configuração do Docker para Desenvolvimento Local

## Visão Geral

Esta feature implementa a configuração do ambiente de desenvolvimento local utilizando Docker e Docker Compose. O objetivo é criar um ambiente consistente e isolado para o desenvolvimento da aplicação de consulta de temperatura por CEP, garantindo que todos os desenvolvedores trabalhem com as mesmas versões de ferramentas e bibliotecas.

## Objetivos

- Criar um Dockerfile otimizado para desenvolvimento local
- Configurar o Docker Compose para orquestrar o serviço da aplicação
- Configurar volumes para desenvolvimento dinâmico sem necessidade de reconstrução de imagem
- Implementar scripts de inicialização e helper commands para facilitar o desenvolvimento

## Tecnologias Utilizadas

- Docker com a imagem base `golang:1.24-alpine`
- Docker Compose para orquestração do serviço
- Bind mounts para desenvolvimento dinâmico

## Entregáveis

- Arquivo `Dockerfile` para desenvolvimento
- Arquivo `docker-compose.yaml` para orquestração do serviço
- Arquivo `.env` com variáveis de ambiente para desenvolvimento local
- Documentação sobre como utilizar o ambiente Docker

## Responsáveis

- Equipe de Desenvolvimento

## Dependências

- Projeto Go inicializado
- Módulos Go configurados
- Estrutura de diretórios seguindo as convenções
