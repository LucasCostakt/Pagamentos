## Regras de negócio obrigatórias: 

1 - Existem dois tipos de usuários. Os Comuns podem enviar e receber dinheiro e Lojistas que não podem enviar mas recebem dinheiro. 

2 - CPF/CNPJ e Email devem ser únicos.

3  - Caso ocorra qualquer problema na transação o dinheiro pode ser devolvido.

4 - Após efetuar uma transação disparar uma notificação (envio de email, sms) enviada por um serviço de terceiro.

## Extras:
1 - Cadastro de novos Usuários Lojista e Comuns

2 - Poder depositar valores 

# Funcionalidades Obrigatórias:

## 1 - Transação 

### Fluxograma da transação

<br>

![transaction](images/transaction.png)
<br>
<br>

### Payload da transação

```json
{
	"value":"float",
	"payer":"int",
	"payee":"int",
	"password":"string", 
	"transfer_time":"date",
	"end_date_ reversal":"date" 
}
```

## 2 - Solicitação de Estorno 

### Fluxograma de solicitação de estorno

<br>

![reversal](images/reversal.png)
<br>
<br>

### Payload da solicitação de estorno
```json
{
	"user_id":"int",
	"transfer_id":"int",
	"password":"string"
}
```

## 3 - Envio de notificação de transferência

### Fluxograma do envio das notificação de tranferência

<br>

![serviceNotification](images/serviceNotification.png)
<br>
<br>

### Payload do Envio de notificação de transferência

```json
{
	"value":"float",
	"payer_id":"int",
	"payer_name":"string",
	"payee_id":"int",
	"payee_name":"string",
	"transfer_time":"date",
	"email_payer":"string",
	"email_payee":"string"
}
```


# Funcionalidades Extras:

## 1 - Cadastro de Novos Usuários

### Fluxograma do cadastro de novos usuários

<br>

![users](images/insertUsers.png)

<br>
<br>

### Payload do Cadastro de Novos Usuários

```json 
{
	"name":"string",
	"cpf_cnpj":"string",
	"email":"string",
	"password":"string",
	"type":"int"
}
```

## 2 - Depósito (Não implementado)

<br>

### Payload do Depósito
```json
{
"value":"float",
"account":"int",
"password":"string",
"cpf/cnpj":"string"
}
```

## Estrutura dos Microsserviços

<br>

![microsservices](images/microsservices.png)

<br>
<br>

## Comunicação entre os serviços

<br>

![ms_comunication](images/ms_comunication.png)

<br>
<br>

## Modelagem do banco de dados

- ``users``: tabela referente a registros dos usuários;

- ``userstype``: tabela contendo os tipos de usuários, sendo eles, comuns ou lojistas;

- ``transaction``: armazena os registro de todas as transções que foram efetuadas com sucesso;

- ``reversal``: armazena os registro de todos os estornos realizados.

<br>

![mysql](images/model.png)

<br>
<br>

## Futuras melhorias

- Implementar a api getway.

- Implementar Serviço de deposito de valores.

- Implementar os testes para o serviços de: users e sendnotification.

- A chamada da api externa para envio de notificação demora bastante, e isso atrasa um pouco a transação, modificar o fluxo para que essa chamada fique sendo feita separada (através de um goroutine), para que o retorno seja mais rápido.

- Utilização de middlewares
