# GoLang Base Project

[![CircleCI](https://circleci.com/gh/ralvescosta/go_base_project/tree/main.svg?style=svg)](https://circleci.com/gh/ralvescosta/go_base_project/tree/main)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ralvescosta_unico_idtech_challenge&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ralvescosta_unico_idtech_challenge)
## Conteúdo
- [GO Base Project](#go-base-project)
  - [Conteúdo](#conteudo)
  - [Estrutura do Projeto](#estrutura-do-projeto)
  - [Recursos/Rotas](#rotas)
  - [Instalacao](#instalacao)

## Estrutura do projeto

O projeto foi estruturado em camadas viabilizando a utilização de alguns padrões de projetos e alguns princípios arquiteturais obtendo assim uma aplicação testável e de fácil manutenção. As principais camadas do projeto são:

- **cmd**: Seguindo o padrão da comunidade, esta camada é utilizada para gerenciar configurações iniciais e execução da aplicação.

- **pkg/domain**: Esta camada contém as interfaces das regras de negócios que deverão ser implementadas e suas estruturas de dados.

- **pkg/app (application)**: Esta camada contém a implementação das interfaces criadas na camada domain.

- **pkg/infra (infrastructure)**: Esta camada contém a implementação e adequação de todas as libs/tools que foi utilizado no projeto. Esta camada tem como um dos objetivos isolar as demais camadas de dependências de libs de terceiro. Um outro objetivo desta camada é isolar as camadas mais internas de dependências de ferramentas como por exemplo o banco de dados, viabilizando uma fácil atualização na necessidade de troca de tecnologia de armazenamento.

- **pkg/interfaces**: Esta camada gerencia como os recursos da aplicação são disponibilizados.

Abaixo segue um esquemático simplificado da estrutura da aplicação:

```
|
│   └── cmd
|       └── *_cmd.go  
│       └── *_container.go
|       
│   └── pkg
|       ├── app        
│       │   └── errors
│       │       └── *_error.go
│       │       └── *_error_test.go
│       │   └── interfaces
│       │       └── i_*.go
│       │   └── usecases
│       │       └── *_usecase.go
│       │       └── *_usecase_test.go
│       │
│       ├── business
│       │   └── value_objects
│       │       └── *_.go
│       │   └── usecases
│       │       └──  i_*_usecase.go
│       │
│       ├── infra
│       │   └── adapters
│       │       └── *_adapt.go
│       │       └── *_adapt_test.go
│       │   └── database
│       │       └── *.go
│       │       └── *_test.go
│       │   └── environments
│       │       └── *.go
│       │       └── *_test.go
│       │   └── http_server
│       │       └── *.go
│       │       └── *_test.go
│       │   └── repositories
│       │       └── *_repository.go
│       │   └── folder
│       │       └── *.go
│       │       └── *_test.go
|       |
│       ├── interfaces
│       │   └── http
|       |       └── factories
│       │           └── *_factory.go
│       │           └── *_factory_test.go
|       |       └── handlers
│       │           └── *_handler.go
│       │           └── *_handler_test.go
|       |       └── presenters
│       │           └── *_routes.go
│       │           └── *_routes_test.go
|       |       └── view_models
│       │           └── *_viewmodels.go
│       │           └── *_viewmodels_test.go
│       │
|       ├── main.go
```

## Rotas

### POST /api/v1/markets

Recurso utilizado registrar novas feiras

>REQUEST:
```bash
curl --location --request POST 'https://localhost:3333/api/v1/markets' \
--header 'Content-Type: application/json' \
--data-raw '{
    "long": -46550162,
    "lat": -23558733,
    "setcens": "355030885000091",
    "areap": "3550308005040",
    "coddist": 87,
    "distrito": "VILA FORMOSA",
    "codsubpref": 26,
    "subpref": "ARICANDUVA-FORMOSA-CARRAO",
    "regiao5": "Leste",
    "regiao8": "Leste 1",
    "nome_feira": "VILA FORMOSA",
    "registro": "4041-0",
    "logradouro": "UA MARAGOJIPE",
    "numero": "S/N",
    "bairro": "VL FORMOSA",
    "referencia": "TV RUA PRETORIA"
}'
```
>RESPONSE:
- 201 - Feira criado com sucesso
- 200 - Caso exista uma feira cadastrada com o mesmo 'Registro', retorna a feira ja cadastrada.
- 400 - Erro de contrato - Todos os campos sao obrigatórios para cadastro da feira
- 500 - Error interno

### GET /api/v1/markets?distrito=VILA FORMOSA&regiao5=Leste&nome_feira=VILA FORMOSA&bairro=VL FORMOSA

Recurso utilizado para consultar feiras. Este recurso aceita todos os parâmetros existentes no registro de feiras

>REQUEST:
```bash
curl --location --request GET 'https://localhost:3333/api/v1/markets?distrito=VILA FORMOSA&regiao5=Leste&nome_feira=VILA FORMOSA&bairro=VL FORMOSA'
```
>RESPONSE:
- 200 - Resultado da consulta
- 400 - Caso algum campo nao valido informado na query
- 500 - Error interno

### PATCH /api/v1/markets/:registerCode

Recurso utilizado para atualizar uma feira ja cadastrada. O único campo que nao é possível atualizar é o capo 'registro'

>REQUEST:
```bash
curl --location --request PATCH 'https://localhost:3333/api/v1/markets/4041-0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "long": -46550162,
    "lat": -23558733,
    "setcens": "1234",
    "areap": "3550308005040",
    "coddist": 87,
    "distrito": "VILA FORMOSA",
    "codsubpref": 26,
    "subpref": "ARICANDUVA-FORMOSA-CARRAO",
    "regiao5": "Leste",
    "regiao8": "Leste 1",
    "nome_feira": "VILA FORMOSA",
    "logradouro": "UA MARAGOJIPE",
    "numero": "S/N",
    "bairro": "VL FORMOSA",
    "referencia": "TV RUA PRETORIA"
}'
```

>RESPONSE:
- 200 - Registro atualizado com sucesso
- 400 - Error de contrato
- 404 - Caso o registro solicitado a atualização nao exista na base de dados
- 500 - Erro interno

### DELETE /api/v1/markets/:registerCode

Recurso utilizado para deletar um registro de feira na base de dados.

>REQUEST:
```bash
curl --location --request DELETE 'https://localhost:3333/api/v1/markets/4041-0'
```

>RESPONSE:
- 200 - Registro deletado com sucesso
- 400 - Error de contrato
- 404 - Caso o registro solicitado a atualização nao exista na base de dados
- 500 - Erro interno

## Instalacao

### Para executar o projeto com todas as dependências

```bash
make docker-compose
```

Apos execute o script de carga da base de dados

```bash
make seeder
```

**OBS: Na pasta integration contem um par de collection e environment do postman com os endpoints criados para a aplicação.**

### Para executar a aplicação de forma separada

- Obtendo os Pkg's

```bash
go get
```

- Configurando o ambiente

```bash
docker-compose -f docker-compose.env.yml up -d
```

- Executando o seeder

```bash
make seeder
```
ou

```bash
GO_ENV=development go run ./db/seeder.go
```

- Executando a aplicação

```bash
make run
```

ou

```bash
GO_ENV=development GIN_MODE=debug go run main.go
```

- Executando a aplicação em modo Debug: Pressione F5


- Para executar os tests unitários

```bash
make test
```

ou

```bash
GO_ENV=development GIN_MODE=debug go test ./pkg/... -v
```

```bash
make test-cov
```

ou

```bash
GO_ENV=development go test ./... -cover -v -coverprofile ./coverage/c.out && go tool cover -html=./coverage/c.out -o ./coverage/coverage.html
```

- Para compilar a aplicação

```bash
make build
```

ou

```bash
go build -ldflags "-s -w" main.go
```
