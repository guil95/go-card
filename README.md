# API para simular transferencias de contas

# Stack
    - Go
    - MySql

# Estruturação do projeto
    └─── api (Configuração de rotas)
    └─── app (Regras de negócio)
         └───domains(Domínios da aplicação)
    └───infra (Dependencias externas, como por exemplo banco de dados)

# Como subir a aplicação?
```shell
$ docker-compose up -d
```
ou
```shell
$ make start
```

A apliacação ficará na porta 8000 de sua máquina, acesse localhost:8000 a aplicação deverá retornar `"Go card!"`

# Recursos da aplicação

## Cadastrar uma conta
`POST` para a rota `http://localhost:8000/accounts` com o seguinte corpo: 
```
    {
	    "document_number": "12345678900"
    }
```

 Caso envie payload inválido a API retornará status code `422` e com o seguinte corpo: 
```
    {
         "message": "Unprocessable entity"
    }
```

 Caso ja exista conta cadastrada com o mesmo documento a API retornará status code `422` e com o seguinte corpo: 
```
    {
         "message": "Account exists"
    }
```

Caso de erro interno a API retornará status code `500` e com o seguinte corpo:
```
    {
         "message": "Internal Error"
    }
```

Caso de sucesso a API retornará status code `201` e com o seguinte corpo:
```
    {
        "id": "fbb8481b-0fa9-4a8a-b85b-c9a8c5e3e753"
        "document_number": "12345678900"
    }
```

## Realizar uma listagem de todas as contas cadastradas
`GET` para a rota `http://localhost:8000/accounts`

Caso não exista contas cadastradas a API retornará status code `404` e com o seguinte corpo: 
```
    {
         "message": "Accounts not found"
    }
```

Caso de erro interno a API retornará status code `500` e com o seguinte corpo:
```
    {
         "message": "Internal Error"
    }
```

Casos de sucesso a API retornará status code `200` e com o seguinte corpo:
```
    [
        {
            "id": "63afaad7-4dd1-452a-867f-ca02dfcbc0ef",
            "document_number": "12345678900"
        },
        {
            "id": "fbb8481b-0fa9-4a8a-b85b-c9a8c5e3e753",
            "document_number": "12345687445"
        }
    ]
```

## Buscar uma conta específica
`GET` para a rota `http://localhost:8000/accounts/{id}`

 Caso envie uuid inválido a API retornará status code `422` e com o seguinte corpo: 
```
    {
         "message": "Unprocessable entity"
    }
```

Caso de erro interno a API retornará status code `500` e com o seguinte corpo:
```
    {
         "message": "Internal Error"
    }
```

Caso não encontre a conta cadastrada a API retornará status code `404` e com o seguinte corpo: 
```
    {
         "message": "Account not found"
    }
```

Caso de sucesso a API retornará status code `200` e com o seguinte corpo:
```
    {
        "id": "43898e63-55a3-4efc-ab81-de3450c5d449",
        "document_number": "12345678907"
    }
```

## Criar Transação
`POST` para a rota `http://localhost:8000/transactions` com o seguinte corpo: 
```
    {
        "account_id": "2d83ef93-c87d-413e-81cf-97241fdc02ba",
        "operation_type_id" : 4,
        "amount": 20
    }
```

 Caso envie payload inválido a API retornará status code `422` e com o seguinte corpo: 
```
    {
         "message": "Unprocessable entity"
    }
```

Caso de erro interno a API retornará status code `500` e com o seguinte corpo:
```
    {
         "message": "Internal Error"
    }
```

Caso de sucesso a API retornará status code `201` e com o seguinte corpo:
```
    {
        "id": "3f9bf1ee-e70f-4257-876e-c6798dad4f9d",
        "account_id": "2d83ef93-c87d-413e-81cf-97241fdc02ba",
        "operation_type_id": 1,
        "amount": 10.50,
        "event_date": "2021-04-18T22:34:16.8053513Z"
    }
```
