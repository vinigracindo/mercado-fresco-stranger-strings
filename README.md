# Mercado Fresco - ITBootcamp Go 

[![build](https://github.com/vinigracindo/mercado-fresco-stranger-strings/actions/workflows/build.yml/badge.svg)](https://github.com/vinigracindo/mercado-fresco-stranger-strings/actions/workflows/build.yml)
[![dependency-review](
https://github.com/vinigracindo/mercado-fresco-stranger-strings/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/vinigracindo/mercado-fresco-stranger-strings/actions/workflows/dependency-review.yml)
[![codeql](
https://github.com/vinigracindo/mercado-fresco-stranger-strings/actions/workflows/codeql.yml/badge.svg)](https://github.com/vinigracindo/mercado-fresco-stranger-strings/actions/workflows/codeql.yml)

Mercado Fresco é uma marketplace de produtos frescos. O objetivo é 
adicionar em sua listagem (oferta) esse tipo de produto. 

## How to Download Dependencies

```shell
go mod tidy
```

## ⚡️ Quick start

1. Clone o projeto
```shell
git clone https://github.com/vinigracindo/mercado-fresco-stranger-strings.git
```

2. Entre na pasta do projeto
```shell
cd mercado-fresco-stranger-strings
```

3. Crie um .env com as variáveis de ambiente do projeto e as edite
```shell
cp .env-example .env
```

4. Crie o database (root/root default username e password - adicionar ao arquivo .env)
```shell
docker-compose up -d
```

5. Rode os testes
```shell
go test ./...
```

6. Rode o projeto
```shell
go run main.go
```

## 📝 Swagger - API Doc

1. Run: go run main.go
2. Open: http://localhost:8080/swagger/index.html

## 📦 Requirements

| Name                                                                  | Version   | Type       |
| --------------------------------------------------------------------- | --------- | ---------- |
| [go](https://go.dev/)                                                 | `v1.18`   | core       |
| [gin-gonic/gin](https://github.com/gin-gonic/gin)                     | `v1.8.0`  | core       |
| [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)         | `v1.6.0`  | database   |
| [joho/godotenv](https://github.com/joho/godotenv)                     | `v1.4.0`  | config     |
| [DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)         | `v1.5.0`  | test       |
| [stretchr/testify](https://github.com/stretchr/testify)               | `v1.7.4`  | test       |
| [swaggo/swag](https://github.com/swaggo/swag)                         | `v1.8.2`  | doc        |
| [swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)           | `v1.4.3`  | doc        |
