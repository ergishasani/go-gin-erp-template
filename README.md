# Go + Gin ERP Backend Template

Production-ready, modular monolith backend scaffold for business ERP systems using Go, Gin, and PostgreSQL. Includes Docker, JWT authentication, and CI setup.

## Features

- Gin HTTP API
- PostgreSQL (pgx)
- JWT + Refresh Token authentication
- Role-based access control
- Modular monolith architecture (microservice-ready)
- Background workers support (PDF, emails, cleanup)
- Docker + Docker Compose
- CI pipeline ready (GitHub Actions)
- Health check endpoint `/health`
- `.env` configuration support

## Getting Started

### 1. Clone the template

Click **Use this template** on GitHub or clone manually:

```bash
git clone https://github.com/<your-username>/go-gin-erp-template.git
cd go-gin-erp-template
```

### 2. Configure environment variables

Copy the example file and fill in your values:

```bash
cp .env.example .env
# Edit .env to configure DATABASE_URL and JWT_SECRET
```

> **Important:** Do not commit your `.env` file. It contains sensitive credentials.

### 3. Build and run locally

```bash
go mod tidy
go run ./cmd/api
```

Server will start on the port defined in `.env` (default: `8080`).
Visit `http://localhost:8080/api/v1/health` to check status.

### 4. Using Docker

Build and run using Docker Compose:

```bash
docker compose up --build
```

### 5. Project Structure

```text
cmd/api/                # Entry point of the API
internal/                # Application modules and utilities
  auth/                  # Authentication module (JWT + Refresh tokens)
  users/                 # Example module for user CRUD
  clients/               # Example module scaffold
  config/                # Config loader
  db/                    # Database connection & migrations
  logger/                # Logging setup
  middleware/            # Common middleware (CORS, recovery, rate limit)
  jobs/                  # Background workers
  health/                # Health check endpoint
.github/workflows/       # CI pipeline
.env.example             # Example environment variables
Dockerfile               # Docker build file
docker-compose.yml       # Docker Compose for API + DB
```

### 6. CI Pipeline

The repository includes a starter GitHub Actions workflow for:

- Build
- Test
- Lint
- Docker build

### 7. Extending the Template

- Add new modules by copying the `users` pattern: `handler.go`, `service.go`, `repository.go`, `model.go`, `routes.go`
- Implement additional middleware, background jobs, or services as needed
- Maintain `.env.example` for new configuration variables

### 8. Contributing

This template is meant for internal use or reuse across projects. Contributions can improve modularity, logging, and worker implementations.

---

**Note:** This template provides a scaffold. Actual business logic, ERP modules, and production configurations need to be implemented per project.
