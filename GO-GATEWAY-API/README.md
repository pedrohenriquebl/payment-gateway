# 🚀 Gateway API em Go

Este projeto é uma API gateway desenvolvida em Go, utilizando PostgreSQL como banco de dados e o framework Chi para rotas.

---

## 📦 Requisitos

- Go 1.20+
- PostgreSQL
- Docker (opcional, para ambiente de desenvolvimento)

---

## ⚙️ Configuração do Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
# Ambiente local
HTTP_PORT=8080
DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gateway
DB_SSLMODE=disable

# Ambiente Docker (exemplo)
# DB_HOST=db
# DB_PORT=5432
```
## ▶️ Rodando o Projeto

```bash
    go run cmd/app/main.go
```

## Baixar as dependências do projeto

```bash
    go mod tidy
```
## 🛠️ Migrations

```bash
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

⚠️ Certifique-se de que o diretório $GOPATH/bin ou $HOME/go/bin esteja no seu PATH.

Rodar migrations
```bash
    migrate -database "postgres://postgres:postgres@localhost:5433/gateway?sslmode=disable" -path migrations up

```

## 🐳 Docker (opcional)

```bash
    docker-compose up --build && docker-compose up -d
```

## 📁 Estrutura do Projeto

```bash
.
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── repository/
│   ├── service/
│   └── web/
│       ├── handlers/
│       ├── middleware/
│       └── server/
├── migrations/
├── go.mod
├── go.sum
└── .env
```