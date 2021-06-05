## Regras de negócio obrigatórias: 

1 - Existem dois tipos de usuários. Os Comuns podem enviar e receber dinheiro e Lojistas que não podem enviar mas recebem dinheiro. 

2 - CPF/CNPJ e Email devem ser únicos.

3  - Caso ocorra qualquer problema na transação o pode ser devolvido o dinheiro.

4 - Após efetuar uma transação disparar uma notificação (envio de email, sms) enviada por um serviço de terceiro.

## Extras:
1 - Pode depositar valores 

2 - Cadastro de novos Usuários Lojista e Comuns


# Funcionalidades Obrigatórias:

## 1 - Transação 

### Fluxograma da transação

<br>

![transaction](/images/transaction.png)
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

![reversal](/images/reversal.png)
<br>
<br>

### Payload da solicitação de estorno
```json
{
	"transfer_id":"int",
	"payer":"int",
	"payee":"int"
}
```

## 3 - Envio de notificação de transferência

### Fluxograma do envio das notificação de tranferência

<br>

![serviceNotification](/images/serviceNotification.png)
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

![users](/images/insertUsers.png)

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

![microsservices](/images/microsservices.png)

<br>
<br>
