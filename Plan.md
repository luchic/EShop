# Web Shop — Learning Project Plan

**Stack goals:** Go, Redis, Kafka, PostgreSQL, React (or plain HTML/JS)

---

## Architecture Overview

```
Browser
  └── Nginx (reverse proxy)
        ├── Frontend (static files or React)
        └── Backend API (Go)
              ├── PostgreSQL  — persistent data (products, orders, users)
              ├── Redis       — sessions, caching, rate limiting
              └── Kafka       — async events (order placed → inventory update, email, etc.)
```

---

## Services

| Service       | Tech              | Role |
|---------------|-------------------|------|
| `api`         | Go (net/http)     | REST API |
| `db`          | PostgreSQL        | Persistent storage |
| `cache`       | Redis             | Sessions + product cache |
| `broker`      | Kafka + Zookeeper | Async order events |
| `consumer`    | Go                | Kafka consumer (order processor) |
| `frontend`    | React or plain JS | UI |
| `nginx`       | Nginx             | Reverse proxy / static file server |

---

## Phases

### Phase 1 — Foundation
- [ ] Clean up repo: remove old code, reset docker-compose from scratch
- [ ] Set up PostgreSQL with migrations (use `golang-migrate` or plain SQL files)
- [ ] Define DB schema: `users`, `products`, `orders`, `order_items`
- [ ] Go project structure (`cmd/api`, `internal/`, `pkg/`)
- [ ] Basic HTTP server with health check endpoint

### Phase 2 — Auth + Users
- [ ] User registration & login (bcrypt passwords)
- [ ] JWT or session-based auth stored in Redis
- [ ] Middleware for protected routes

### Phase 3 — Products
- [ ] CRUD for products (admin only)
- [ ] List/search products (public)
- [ ] Cache product listings in Redis (TTL-based invalidation)

### Phase 4 — Shopping Cart
- [ ] Cart stored in Redis (keyed by session/user)
- [ ] Add, remove, update cart items
- [ ] Cart expiry via Redis TTL

### Phase 5 — Orders + Kafka
- [ ] Place order endpoint: writes order to PostgreSQL, publishes `order.created` event to Kafka
- [ ] Kafka consumer service: consumes `order.created`, updates inventory, logs fulfillment
- [ ] Order status flow: `pending → processing → completed`

### Phase 6 — Frontend
- [ ] Product listing page
- [ ] Product detail page
- [ ] Cart page
- [ ] Checkout flow
- [ ] Order history (authenticated)

### Phase 7 — Polish
- [ ] Rate limiting via Redis (token bucket or fixed window)
- [ ] Structured logging (slog or zerolog)
- [ ] Docker Compose wiring all services together
- [ ] Basic Nginx config for routing

---

## Project Layout (Go)

```
shop/
├── cmd/
│   ├── api/          # main.go — HTTP server entry point
│   └── consumer/     # main.go — Kafka consumer entry point
├── internal/
│   ├── auth/         # JWT/session logic
│   ├── cart/         # Redis cart
│   ├── order/        # order domain logic + Kafka producer
│   ├── product/      # product domain logic
│   ├── user/         # user domain logic
│   └── db/           # DB connection, migrations
├── frontend/         # React app or plain HTML/JS
├── migrations/       # SQL migration files
├── docker-compose.yml
├── nginx.conf
└── Plan.md
```

---

## Key Learning Points

| Topic    | Where you'll use it |
|----------|---------------------|
| **Go**   | REST API, middleware, consumers, DB queries |
| **Redis**| Sessions, cart state, product cache, rate limiting |
| **Kafka**| Order events, decoupled async processing |
| **PostgreSQL** | Relational data, transactions, migrations |
| **Docker Compose** | Local multi-service orchestration |

---

## Suggested Order of Attack

1. Start with Phase 1 — get Go + Postgres + Docker Compose running
2. Add Redis in Phase 2 (auth) so you feel it early
3. Do Phase 3 (products + cache) before cart — simpler Redis usage first
4. Phase 4 cart is Redis-heavy — good warm-up before Kafka
5. Phase 5 is the Kafka payoff — keep the consumer simple at first
6. Frontend last — use curl/Postman until then
