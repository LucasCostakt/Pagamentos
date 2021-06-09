
# Documentação

Documentação com os guias de instalação, como executar a aplicação, execução dos testes unitários e disponibilizar um melhor entendimento de todas as partes do projeto

## Tópicos

- [Requisitos necessários](/documentation.md#requisitos-necessários-para-utilizar-o-projeto)

- [Executar o Projeto](/documentation.md#executar-o-projeto)

- [Testes](/documentation.md#testes)

- [Referência para a API](/documentation.md#referência-para-a-api)

- [Variáveis de Environment](/documentation.md#variáveis-de-environment)

## Requisitos necessários para utilizar o projeto

```bash
  docker
  docker-compose
  go
  mysql
```
  
## Executar o Projeto

1- Clonar

```bash
  git clone https://github.com/LucasCostakt/Pagamentos.git
```

2- Vá para pasta do projeto

```bash
  cd Pagamentos
```


3- Na pasta pagamentos rode o comando

```bash
docker-compose up

```

Observação: confira se foi criado corretamente as tabelas e o mock 
de dados no mysql, caso não tenha sido rode 
novamente o container `servicemocks`
## Testes

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

  
## Referência para a API

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

  
## Variáveis de Environment

To run this project, you will need to add the following environment variables to your .env file

`API_KEY`

`ANOTHER_API_KEY`

  
