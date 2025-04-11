# Payment Gateway with Anti-Fraud System

Este repositório é composto por três aplicações que juntas formam um sistema completo para análise de fraude:

🔌 GO-GATEWAY-API: API Gateway escrita em Go.

🔍 nestjs-anti-fraud: Serviço de antifraude utilizando NestJS + Prisma + PostgreSQL.

🌐 next-frontend: Interface frontend feita em Next.js.

## 📦 Requisitos

Docker + Docker Compose

Node.js 18+ (se quiser rodar o frontend fora do container)

Go 1.20+ (se quiser rodar a gateway fora do container)

## 🚀 Rodando tudo com Docker

1. **Clone o Repositório**
```bash
git clone https://github.com/seu-usuario/fullcycle-fraud-system.git
cd fullcycle-fraud-system
```

2. **Crie os arquivos .env de cada serviço**
🛡️ nestjs-anti-fraud/.env
```bash
DATABASE_URL="postgresql://postgres:root@db:5440/anti-fraud?schema=public"
SUSPICIOUS_VARIATION_PERCENTAGE=50
INVOICES_HISTORY_COUNT=5
SUSPICIOUS_INVOICES_COUNT=3
SUSPICIOUS_TIMEFRAME_HOURS=24
```

3. **Suba os Containers para cada serviço**
```bash
docker compose up --build -d
```

4. **Acesse os Serviços**
Gateway (Go): http://localhost:8080

Backend NestJS: http://localhost:3001

Frontend Next.js: http://localhost:3000

Prisma Studio (opcional): http://localhost:5555

## ⚙️ Primeira execução do Prisma (NestJS)
```bash
docker compose exec nestjs npx prisma migrate deploy
docker compose exec nestjs npx prisma studio
```