package repository

import (
	"backend/shop/internal/api/goods"
	"backend/shop/internal/config"
	"context"
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	context context.Context
	db *sql.DB
}

func NewPostgresRepository(ctx context.Context, cfg config.Config) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}

	return &PostgresRepository{
		context: ctx,
		db: db,
	}, nil
}

func (r *PostgresRepository) AddProduct(product goods.AddProductRequest) {
	_, err := r.db.Exec(
		`INSERT INTO products (name, description) VALUES ($1, $2)`,
		product.Name,
		product.Description,
	)
	if err != nil {
		slog.Error("add product failed", slog.Any("err", err))
	}
}

func (r *PostgresRepository) GetGoodPage(offset int, limit int) []goods.Product {
	rows, err := r.db.Query(
		`SELECT id, name, description
		FROM products
		ORDER BY id
		OFFSET $1
		LIMIT $2`,
		offset,
		limit,
	)
	if err != nil {
		slog.Error("get goods page failed", slog.Any("err", err))
		return []goods.Product{}
	}
	defer rows.Close()

	products := make([]goods.Product, 0, limit)
	for rows.Next() {
		var product goods.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description)
		if err != nil {
			slog.Error("scan product failed", slog.Any("err", err))
			return []goods.Product{}
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		slog.Error("iterate products failed", slog.Any("err", err))
		return []goods.Product{}
	}

	return products
}
