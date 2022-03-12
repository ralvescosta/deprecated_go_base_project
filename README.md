# Unico Challenge

## Conteúdo
- [Unico Challenge](#unico-challenge)
  - [Conteúdo](#conteudo)
  - [Estrutura do Projeto](#estrutura-do-projeto)
  - [Recursos/Rotas](#rotas)
  - [Instalacao](#instalacao)

## Estrutura do projeto

O projeto foi estruturado em camadas viabilizando assim a utilização de alguns padrões de projetos e alguns princípios arquiteturais obtendo assim uma aplicação testável e de fácil manutenção. As principais camadas do projeto são:

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
│       │       └── folders
│       │           └── *_usecase.go
│       │           └── *_usecase_test.go
│       │
│       ├── business
│       │   └── value_objects
│       │       └── *_dto.go
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
|       |       └── handlers
│       │           └── *_handler.go
│       │           └── *_handler_test.go
|       |       └── middleware
│       │           └── *_middleware.go
│       │           └── *_middleware_test.go
|       |       └── presenters
│       │           └── *_routes.go
│       │           └── *_routes_test.go
│       │
|       ├── main.go
```

Com intuito de criar uma API REST foi utilizado o pacote Gin Web Framework para facilitar a gestão dos recursos HTTP disponibilizados tão bem como facilitar a configuração de handlers e middlewares. A escolha do Gin passou pelo fato que atualmente este é um projeto mantido pela comunidade com um ótimo resultado de benchmark comparado com os demais como por exemplo o gurila-mux e echo.

Optou-se por utilizar PostgreSQL como mecanismo de persistência.

Optou-se por utilizar a lib Zap para poder ser o gerenciador de logs da aplicação pois esta lib possui um ótimo resultado de benchmark comparado com outras libs como por exemplo o logrus.

Implementou-se uma estratégia de Graceful Shutdown para quando ocorra a perda de conexão com o banco de dados a aplicação não “desligue” cortando todas as conexões TCP ativas. Desta forma perante a ausência do mongo a aplicação para de aceitar pedidos e espera todas as conexões TCP fecharem para poder encerrar a aplicação.

## Rotas

> POST /api/v1/markets

Recurso utilizado registrar novas feiras

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

## Instalacao

### Para executar o projeto com todas as dependências

```bash
make docker-compose
```

**OBS: Na pasta integration contem um par de collection e environment do postman com os endpoints criados para a aplicação.**

### Para executar a aplicação de forma separada

- Obtendo os Pkg's

```bash
go get
```

- Configurando o ambiente

```bash
docker-compose -f docker-compose.environment.yml up -d
```

- Executando a aplicação

```bash
make run
```

- Executando a aplicação em modo Debug: Pressione F5


- Para executar os tests unitários

```bash
make test
```

```bash
make test-cov
```

- Para compilar a aplicação

```bash
make build
```