# Desafio - Clean Architecture
 
Olá devs!
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

Google Wire
```
go generate
```

GraphQL
```
go run github.com/99designs/gqlgen generate.
```

# Como rodar a aplicação

Clone o repositório
```
git clone https://github.com/nicolastanski/go-clean-architecture
```

Execute o comando para executar o Server
```
cd cmd/ordersystem/
go run main.go wire_gen.go
```

```
http://localhost:8000/order 
```

Testar com gRPC
```
evans -r repl

call CreateOrder
```

Testar com GraphQL
```
http://localhost:8080

mutation createOrder {
    createOrder(input: {id: "ccc", Price: 10.0}) {
        id
        Price
        Tax
        FinalPrice
    }
}
```