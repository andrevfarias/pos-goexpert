# Pós-Graduação Go Expert - Módulo 1

## Desafio 3 - Clean Architecture

Olá devs!
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:

- [x] Endpoint REST (GET /order)
- [x] Service ListOrders com GRPC
- [x] Query ListOrders GraphQL
- [x] Criar as migrações necessárias
- [x] Criar o arquivo api.http com a request para criar e listar as orders.
- [x] Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml),
      com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
- [x] Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

### Requisitos

- Docker e Docker Compose

### Utilização

Na raiz do projeto executar o comando:

```bash
make docker-up
```

Este comando irá subir o banco de dados mysql, executar as migrations, subir o rabitmq e a aplicação.

A partir desta estapa já será possível testar os serviços:

1.  REST API

    A api rest roda no endereço http://localhost:8000

    Os endpoints disponíveis são:

    - GET /orders
    - POST /orders

    Podem ser testados utilizando o arquivo [`./api/orders.http`](./api/orders.http) juntamente com a extensão

    [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) do VS Code.

2.  GraphQL

    Para utilizar o GraphQL basta acessar http://localhost:8080 e utilizar os scripts abaixo:

    > Criando uma order:

    ```graphql
    mutation createOrder {
      createOrder(input: { Id: "b", Price: 10, Tax: 0.5 }) {
        Id
        Price
        Tax
        FinalPrice
      }
    }
    ```

    > Listando todas as orders

    ```graphql
    query queryOrders {
      getOrders {
        Id
        Price
        Tax
        FinalPrice
      }
    }
    ```

3.  GRPC

    Para testar os serviços gRPC será necessária a utilização de um client gRPC como o [Postman](https://www.postman.com) ou o [Evans](https://github.com/ktr0731/evans).

    Utilizando o Evans:

    ```
    evans -r repl
    ```

    Dentro do coonsole do Evans execute os seguintes comandos:

    ```
    package pb
    service OrdersService
    ```

    Para exibir as chamadas gRPC disponíveis:

    ```
    show rpc
    ```

    Para executar uma chamada em CreateOrder:

    ```
    call CreateOrder
    ```

    Para executar uma chamada em ListOrders:

    ```proto
    call ListOrders
    ```

---
