# Configuração do Ambiente com Docker e Docker Compose

## **Visão Geral**

Este projeto será executado inteiramente dentro de contêineres Docker, eliminando a necessidade de instalar **Go** localmente. Para isso, utilizaremos **Docker Compose** para orquestrar o serviço da aplicação.

## **Tecnologias Utilizadas**

- **Docker** para contêinerização do ambiente.
- **Docker Compose** para gerenciar o serviço.
- **Imagem Base da Aplicação:** `golang:1.24-alpine` para otimização e leveza.

## **Estrutura dos Serviços no Docker Compose**

- **`app`**: Serviço principal onde a aplicação Go será executada dentro do contêiner.

## **Execução Totalmente Dentro do Container**

- O contêiner da aplicação terá o ambiente Go configurado para rodar diretamente dentro dele.
- O código será montado como um volume para desenvolvimento dinâmico sem necessidade de reconstrução manual do container.
- O serviço será executado com **TTY habilitado** (`tty: true`) para manter o processo ativo.

## **Bind Mounts**

- O código da aplicação será montado como um bind mount para desenvolvimento, logo, qualquer alteração no código será refletida automaticamente no container.

## **Arquivo `docker-compose.yaml` (Descrição Geral)**

- Define o serviço **`app`** para a API de consulta de temperatura.
- Configura bind mount para o código-fonte.
- Expõe a porta da aplicação (8080).
- Carrega variáveis de ambiente do arquivo `.env`.
- Exibe logs da aplicação em tempo real.

## **Ambiente de Desenvolvimento**

- Execute `docker compose up -d` para iniciar o ambiente de desenvolvimento.
- Execute comandos diretamente no container com `docker compose exec app <comando>`.
- Para testes de desenvolvimento, use `docker compose exec app go test ./...`.
- Para executar a aplicação em modo de desenvolvimento, use `docker compose exec app go run cmd/app/main.go`.

## **Ambiente de Produção (Cloud Run)**

- O projeto será containerizado para deploy no Google Cloud Run.
- O Dockerfile de produção otimizará a imagem usando multi-stage build:
  - Stage de build: compila o código Go.
  - Stage final: imagem mínima contendo apenas o binário compilado.
- As variáveis de ambiente serão configuradas diretamente no serviço Cloud Run.
