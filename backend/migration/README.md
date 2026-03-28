# Migration Project

This project follows the `golang-migrate` style described in the article:
https://dev.to/oriiyx/migrations-with-go-postgres-54m9

## Structure

- `cmd/migrate/main.go`: migration runner
- `cmd/migrate/migrations/*.up.sql`: apply schema changes
- `cmd/migrate/migrations/*.down.sql`: rollback schema changes

## Environment

The runner uses:

- `AUTH_DATABASE_URL`

Example:

```bash
export AUTH_DATABASE_URL="postgres://user:password@localhost:5432/shop?sslmode=disable"
```

## Commands

Run migrations up:

```bash
cd backend/migration
go run ./cmd/migrate/main.go up
```

Run migrations down:

```bash
cd backend/migration
go run ./cmd/migrate/main.go down
```

## Notes

- `golang-migrate` creates and manages the `schema_migrations` table itself.
- Migration files are stored in `cmd/migrate/migrations`.
