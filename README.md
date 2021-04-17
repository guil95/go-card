# API para simular transferencias de contas

# Stack
    - Go
    - MySql

# Estruturação do projeto
    - api (Configuração de rotas)
    - app (Regras de negócio)
        |
        |domains(Domínios da aplicação)
    - infra (Dependencias externas, como por exemplo banco de dados)

# Como subir a aplicação?
```shell
$ docker-compose up
```
A apliacação ficará na porta 8000 de sua máquina, acesse localhost:8000 a aplicação deverá retornar `"Go card!"`
