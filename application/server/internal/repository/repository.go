package repository

import (
	"database/sql"
	"fmt"
	"shop/internal/api"
	"shop/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Smaller suggestions
// A few things worth considering, none of them bugs:
// Set pool limits. sql.DB is a connection pool (safe for concurrent use — good
// choice for a repository). By default it has unlimited open connections, which can overwhelm Postgres under load. Consider:
// godb.SetMaxOpenConns(25)
// db.SetMaxIdleConns(25)
// db.SetConnMaxLifetime(5 * time.Minute)
// Use context. Accepting a context.Context and using db.PingContext(ctx) lets callers control timeouts and cancellation on startup:
// gofunc NewRepository(ctx context.Context, cfg config.Config) (*Repository, error) {
// 	db, err := sql.Open("pgx", cfg.ConnectionString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := db.PingContext(ctx); err != nil {
// 		db.Close()
// 		return nil, err
// 	}
// 	// ...
// }

// Package name internal. Naming the package itself internal is
// unusual — internal is a special directory name in Go
// (it restricts who can import the directory), but as a package name it's vague.
// Something like package repository or package postgres reads better. This is style, not a problem.
// The good parts: you ping to verify the connection, you close db if the ping
// fails (avoiding a leaked handle), and you propagate errors properly instead
// of calling log.Fatal inside a constructor. That's the right shape for a constructor.
// One thing to decide going forward: importing the stdlib adapter
// means you're using pgx as a driver under database/sql.
// That's a perfectly good choice — you get pgx's quality with the familiar
// standard interface. But if you later want pgx's native features
// (COPY, LISTEN/NOTIFY, richer types), you'd switch to pgxpool and pgx's own
// API instead. Worth being intentional about which world you're in.

type Repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	if cfg == nil {
		return nil, fmt.Errorf("Cfg is null")
	}
	db, err := sql.Open("pgx", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return &Repository{db: db}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) CreateUser(user api.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (first_name, second_name, email, password, role) VALUES ($1, $2, $3, $4, $5)",
		user.FirstName,
		user.SecondName,
		user.Email,
		user.Password,
		user.Role)
	return err
}

func (r *Repository) GetUserByEmail(user_email string) (api.User, error) {
	var user api.User

	rows := r.db.QueryRow("SELECT id, first_name, second_name, email, role, password, created_at FROM users WHERE email = $1", user_email)

	err := rows.Scan(
		&user.Id,
		&user.FirstName,
		&user.SecondName,
		&user.Email,
		&user.Role,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("GetUserByEmail %s: no such user", user_email)
		}
		return user, fmt.Errorf("GetUserByEmail %s: %v", user_email, err)
	}
	return user, nil
}

func (r *Repository) CreateProduct(product api.Product) error {
	_, err := r.db.Exec(
		"INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4)",
		product.Name,
		product.Description,
		product.Price,
		product.Stock)
	return err
}

func (r *Repository) GetProductById(productId int64) (api.Product, error) {
	row := r.db.QueryRow(
		"SELECT id, name, description, price, stock, image_url, created_at FROM products WHERE id = $1",
		productId)

	var product api.Product
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.ImageUrl,
		&product.CreatedAt)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *Repository) GetProductsByName(name string) ([]api.Product, error) {
	rows, err := r.db.Query(
		"SELECT id, name, description, price, stock, image_url, created_at FROM products WHERE name = $1",
		name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []api.Product

	for rows.Next() {
		var product api.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.ImageUrl,
			&product.CreatedAt)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return products, nil
	}

	return products, nil
}

func (r *Repository) UpdateProduct() {

}

func (r *Repository) DeleteProduct() {

}
