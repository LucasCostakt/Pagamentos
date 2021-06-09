
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

Todos os envios e retornos são no formato de json.

Além das mensagem de retorno é retornado os status code da requisição.

<br>

### Inserir novo usuário 


<br>

Porta:
```http
5050
```

<br>

uri:
```http
  POST /insert
```

Envio

| Parameter   | Type     | Description                |
| :--------   | :------- | :------------------------- |
| `name`      | `string` | Nome do usuário |
| `cpf_cnpj`  | `string` | cpf e cnpj |
| `email`     | `string` | email |
| `password`  | `string` | senha |
| `user_type` | `int`    | tipo de usuário 1- comum e 2- lojista |

<br>

Retorno

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `message` | `string` | mensagem notificando se foi um sucesso ou se ocorreu algum erro  |

<br>

<br>

### Efetuar uma transação

<br>

Porta:
```http
5052
```

<br>

uri:
```http
  POST /transaction
```

Envio

| Parameter   | Type     | Description                |
| :--------   | :------- | :------------------------- |
| `value`     | `float`  | valor a ser transferido |
| `payer`     | `int`    | id/conta do usuário que vai enviar o dinheiro |
| `payee`     | `int`    | id/conta do usuário que irá receber o dinheiro |
| `password`  | `string` | senha |

<br>

Retorno

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `message` | `string` | mensagem notificando se foi um sucesso ou se ocorreu algum erro  |

<br>

<br>

### Efetuar um estorno


<br>

Porta:
```http
5052
```

<br>

uri:
```http
  POST /reversal
```

Envio

| Parameter      | Type     | Description                |
| :--------      | :------- | :------------------------- |
| `user_id`      | `int`    | id/conta do usuário que recebeu a transferência |
| `transfer_id`  | `int`    | id da transferência que irá ser retornada |
| `password`     | `string` | senha do usuário |


<br>

Retorno

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `message` | `string` | mensagem notificando se foi um sucesso ou se ocorreu algum erro  |

<br>

<br>
  
## Variáveis de Environment

To run this project, you will need to add the following environment variables to your .env file

`API_KEY`

`ANOTHER_API_KEY`

  
