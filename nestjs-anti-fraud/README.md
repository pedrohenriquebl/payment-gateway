# 🛡️ NestJS Anti-Fraud API

Este projeto é uma API de antifraude construída com [NestJS](https://nestjs.com/), [Prisma ORM](https://www.prisma.io/) e banco de dados PostgreSQL, com suporte completo para execução via Docker.

---

## 🚀 Tecnologias

- [Node.js](https://nodejs.org/)
- [NestJS](https://nestjs.com/)
- [Prisma ORM](https://www.prisma.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker + Docker Compose](https://docs.docker.com/compose/)

---

## 📦 Instalação

### 🔧 Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## ▶️ Rodando o projeto

1. **Clone o repositório:**

```bash
git clone https://github.com/seu-usuario/seu-repo.git
cd seu-repo
```

2. **Crie o arquivo .env:**
```
DATABASE_URL="postgresql://postgres:root@db:5440/anti-fraud?schema=public"
SUSPICIOUS_VARIATION_PERCENTAGE=50
INVOICES_HISTORY_COUNT=5
SUSPICIOUS_INVOICES_COUNT=3
SUSPICIOUS_TIMEFRAME_HOURS=24
```

3. **Suba os containers:**
```docker
docker compose up -d
docker compose exec nestjs bash
```

4. **Rodar a Migrate na primeira vez:**
```bash
npx prisma migrate deploy

```

5. **Iniciar a aplicação:**
```bash
npm run start:dev
```

6. **Testes manuais:**
```bash
npm run start:dev -- --entryFile rpl
await get(FraudService).processInvoice({ invoice_id: '6', account_id: '1', amount: 100 })
```
7. **Conferir a criação(Opcional):**
```bash
npx prisma studio
```