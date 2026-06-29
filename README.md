# Shop

A learning project to explore Go and backend development by building a web shop from scratch.

## Tech Stack

| Layer        | Technology                  |
|--------------|-----------------------------|
| Backend      | Go (net/http)               |
| Database     | PostgreSQL 17               |
| Sessions     | Redis                       |
| Reverse Proxy| Nginx                       |
| Frontend     | Plain HTML/CSS/JS           |
| Migrations   | golang-migrate              |
| API Docs     | Swagger (swaggo)            |
| DB Admin     | Adminer                     |
| Containers   | Docker / Docker Compose     |

## Project Structure

```
shop/
├── application/server/
│   ├── cmd/
│   │   ├── shop/          # HTTP server entry point
│   │   └── migrate/       # Database migrations
│   ├── internal/
│   │   ├── api/           # Request/response models
│   │   ├── auth/          # Authentication & session logic
│   │   ├── config/        # App configuration
│   │   ├── handlers/      # HTTP handlers
│   │   ├── product/       # Product domain logic
│   │   ├── repository/    # Database & Redis repositories
│   │   └── services/      # Middleware
│   ├── tests/             # Unit tests
│   └── docs/              # Swagger docs
├── nginx/
│   ├── html/              # Static frontend pages
│   └── nginx.conf
├── docker-compose.yml
└── Makefile
```

## Getting Started

### Prerequisites

- Docker & Docker Compose

### Run

```bash
make up
```

The app will be available at:

- **Frontend:** http://localhost:8082
- **API (via Nginx):** http://localhost:8082/app/
- **API (direct):** http://localhost:8080
- **Adminer (DB UI):** http://localhost:8081

### Other Commands

```bash
make down      # Stop containers
make clean     # Stop and remove images
make fclean    # Stop, remove images and volumes
make re        # Rebuild and restart
make swagger   # Regenerate Swagger docs
```

## Features

- User registration and login with bcrypt password hashing
- Session-based auth stored in Redis
- Product creation
- Reverse-proxied frontend served via Nginx
