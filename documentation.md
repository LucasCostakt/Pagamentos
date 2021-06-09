
# Documentação

Essa documentação vai servir para guiar a instalação, como usar a aplicação, revisar os testes e disponibilizar um melhor intendimento de todas as partes do projeto

## Tópicos

- Additional browser support

- Add more integrations

  
## Rodar o Projeto

Clonar

```bash
  git clone https://github.com/LucasCostakt/Pagamentos.git
```

Vá para pasta do projeto

```bash
  cd Pagamentos
```

Tenha instalado

```bash
  docker
  docker-compose
  go
  mysql
```

Na pasta pagamentos rode o comando

```bash
docker-compose up

```

Observação: confira se foi criado corretamente as tabelas e o mock 
de dados no mysql, caso não tenha sido rode 
novamente o container `servicemocks`
## Tests

Para rodar os testes entre na pasta `service-transaction`

```bash
  cd service-transaction
```

Depois rode o comando 

```bash
  go test transaction_test.go
```

Ou o comando 

```bash
  go test transaction_test.go -v
```

  
## API Reference

#### Get all items

```http
  GET /api/items
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### add(num1, num2)

Takes two numbers and returns the sum.

  
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`API_KEY`

`ANOTHER_API_KEY`

  
