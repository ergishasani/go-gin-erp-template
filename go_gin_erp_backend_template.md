# Go + Gin ERP Backend Template

Production-grade modular monolith for business ERPs (JWT, RBAC, PostgreSQL, Docker, CI).

---

## Folder Structure

```text
.
├─ cmd/
│  └─ api/
│     └─ main.go
├─ internal/
│  ├─ auth/
│  │  ├─ handler.go
│  │  ├─ service.go
│  │  ├─ middleware.go
│  │  └─ routes.go
│  ├─ users/
│  │  ├─ handler.go
│  │  ├─ service.go
│  │  ├─ repository.go
│  │  ├─ model.go
│  │  └─ routes.go
│  ├─ clients/
│  │  └─ (empty scaffold)
│  ├─ config/
│  │  └─ config.go
│  ├─ db/
│  │  ├─ db.go
│  │  └─ migrate/
│  │     └─ 001_init.sql
│  ├─ errors/
│  │  └─ http.go
│  ├─ logger/
│  │  └─ logger.go
│  ├─ middleware/
│  │  ├─ cors.go
│  │  ├─ rate_limit.go
│  │  └─ recovery.go
│  ├─ jobs/
│  │  └─ worker.go
│  └─ health/
│     └─ handler.go
├─ .env.example
├─ .gitignore
├─ Dockerfile
├─ docker-compose.yml
├─ go.mod
├─ README.md
└─ .github/
   └─ workflows/
      └─ ci.yml
```

---

## cmd/api/main.go

```go
package main

import (
  "context"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/gin-gonic/gin"

  "app/internal/auth"
  "app/internal/config"
  "app/internal/db"
  "app/internal/health"
  "app/internal/logger"
  mw "app/internal/middleware"
  "app/internal/users"
)

func main() {
  cfg := config.Load()
  log := logger.New(cfg.Env)

  database := db.Connect(cfg, log)

  r := gin.New()
  r.Use(mw.Recovery(log))
  r.Use(mw.CORS())
  r.Use(mw.RateLimit())

  api := r.Group("/api/v1")

  health.Register(api)
  auth.Register(api, database, cfg)
  users.Register(api, database)

  srv := &http.Server{
    Addr:    ":" + cfg.Port,
    Handler: r,
  }

  go func() {
    log.Info("server started", "port", cfg.Port)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatal("listen", err)
    }
  }()

  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  _ = srv.Shutdown(ctx)
}
```

---

## internal/config/config.go

```go
package config

import (
  "log"
  "os"

  "github.com/joho/godotenv"
)

type Config struct {
  Env       string
  Port      string
  DBUrl     string
  JWTSecret string
}

func Load() Config {
  _ = godotenv.Load()

  cfg := Config{
    Env:       get("APP_ENV", "dev"),
    Port:      get("PORT", "8080"),
    DBUrl:     must("DATABASE_URL"),
    JWTSecret: must("JWT_SECRET"),
  }

  return cfg
}

func get(k, d string) string {
  if v := os.Getenv(k); v != "" {
    return v
  }
  return d
}

func must(k string) string {
  if v := os.Getenv(k); v != "" {
    return v
  }
  log.Fatalf("missing env %s", k)
  return ""
}
```

---

## Dockerfile

```dockerfile
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/api /app/api
EXPOSE 8080
CMD ["/app/api"]
```

## docker-compose.yml

```yaml
version: '3.9'
services:
  api:
    build: .
    env_file: .env
    ports:
      - "8080:8080"
    depends_on:
      - db
  db:
    image: postgres:16
    environment:
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: app
    ports:
      - "5432:5432"
```

---

## .env.example

```env
APP_ENV=dev
PORT=8080
DATABASE_URL=postgres://postgres:postgres@db:5432/app?sslmode=disable
JWT_SECRET=change-me
```

---

## README.md (excerpt)

```md
# Business ERP Backend – Go + Gin

Production-ready backend template for modular ERP systems.

Features:
- Gin HTTP API
- PostgreSQL (pgx)
- JWT + Refresh Tokens
- Role-based access
- Modular monolith
- Background workers
- Docker + Compose
- CI pipeline

## Quick Start

```bash
cp .env.example .env
docker compose up --build
```

API available at `http://localhost:8080/api/v1`
```

---

This template is intentionally minimal in surface area and maximal in architectural correctness. Add new modules by copying the `users` pattern.

