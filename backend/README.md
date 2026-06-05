# Pakumart Backend

Point-of-Sale and inventory management API built with Go standard library, PostgreSQL, and repository pattern.

## Quick Start

```bash
# 1. Clone and enter the directory
cd pakumart/backend

# 2. Copy and configure environment
cp .env.example .env
# Edit .env with your DATABASE_URL

# 3. Run database migrations
psql "$DATABASE_URL" -f database/schema.sql

# 4. Start the server
go run main.go
```

Server starts on `:8080` (configurable via `PORT` env). Health check at `GET /health`.

## Project Status

Currently in **Phase 3 of 7** (see [PLAN.md](PLAN.md)). Live endpoint:

| Method | Path | Status |
|--------|------|--------|
| GET | `/health` | ✅ Live |

Remaining phases: Service layer, Auth/Product handlers, JWT middleware, Router setup.

## Architecture

```
main.go                     Entry point — inits deps, starts HTTP server
├── database/db.go          pgxpool wrapper, connection config
├── internal/config/        Environment-based configuration
├── internal/model/         Domain types + request/response DTOs
├── internal/repository/    Data access layer (interfaces + pgx impl)
├── internal/handler/       HTTP handlers
├── logger/                 Structured logging (zap)
└── database/schema.sql     PostgreSQL schema + indexes
```

Layers communicate through interfaces — handlers depend on repository interfaces, not concrete types.

### Key Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Router | `net/http` ServeMux | Zero dependencies, sufficient for CRUD |
| DB Driver | pgx v5 | Native PostgreSQL protocol, connection pooling |
| Monetary | `decimal.Decimal` | Avoid float64 precision loss on NUMERIC columns |
| Logging | zap | Structured, performant, leveled logging |
| Auth | JWT (planned) | Stateless, fits POS workflow |

## Configuration

| Env Var | Default | Required | Description |
|---------|---------|----------|-------------|
| `PORT` | `8080` | No | HTTP listen port |
| `DATABASE_URL` | — | Yes | PostgreSQL connection string |
| `ENV` | — | No | Set to `production` for JSON log format |

## Dependencies

```
github.com/jackc/pgx/v5          PostgreSQL driver + pool
github.com/joho/godotenv         .env loader
github.com/shopspring/decimal    Arbitrary-precision decimals
go.uber.org/zap                  Structured logging
```

## Contributing

See [PLAN.md](PLAN.md) for the development roadmap. Phases 4–7 are open for implementation:

- **Phase 4** — Service layer (`internal/service/`)
- **Phase 5** — Auth + Product handlers (`internal/handler/`)
- **Phase 6** — JWT middleware (`internal/middleware/`)
- **Phase 7** — Route setup (`internal/router/`)
