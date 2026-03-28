package repository

import (
	authapi "backend/shop/internal/api/auth"
	"backend/shop/internal/api/goods"
	"backend/shop/internal/config"
	"context"
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	context context.Context
	db      *sql.DB
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
		db:      db,
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

func (r *PostgresRepository) GetGoodByID(id uint64) (goods.Product, bool) {
	var product goods.Product
	err := r.db.QueryRow(
		`SELECT id, name, description
		FROM products
		WHERE id = $1`,
		id,
	).Scan(&product.Id, &product.Name, &product.Description)
	if err != nil {
		if err != sql.ErrNoRows {
			slog.Error("get product by id failed", slog.Any("err", err), slog.Uint64("id", id))
		}
		return goods.Product{}, false
	}

	return product, true
}

func (r *PostgresRepository) SaveOAuthState(state string) error {
	_, err := r.db.Exec(
		`INSERT INTO oauth_states (state)
		VALUES ($1)`,
		state,
	)
	if err != nil {
		slog.Error("save oauth state failed", slog.Any("err", err))
		return err
	}

	return nil
}

func (r *PostgresRepository) HasOAuthState(state string) bool {
	var exists bool
	err := r.db.QueryRow(
		`SELECT EXISTS (
			SELECT 1
			FROM oauth_states
			WHERE state = $1
		)`,
		state,
	).Scan(&exists)
	if err != nil {
		slog.Error("check oauth state failed", slog.Any("err", err))
		return false
	}

	return exists
}

func (r *PostgresRepository) DeleteOAuthState(state string) error {
	_, err := r.db.Exec(
		`DELETE FROM oauth_states
		WHERE state = $1`,
		state,
	)
	if err != nil {
		slog.Error("delete oauth state failed", slog.Any("err", err))
		return err
	}

	return nil
}

func (r *PostgresRepository) UpsertOAuthUser(user authapi.OAuthUser) (authapi.AppUser, error) {
	row := r.db.QueryRow(
		`INSERT INTO users (provider_id, login, display_name, email)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (provider_id) DO UPDATE SET
			login = EXCLUDED.login,
			display_name = EXCLUDED.display_name,
			email = EXCLUDED.email
		RETURNING id, login, display_name, email`,
		user.ProviderID,
		user.Login,
		user.DisplayName,
		user.Email,
	)

	var appUser authapi.AppUser
	err := row.Scan(
		&appUser.ID,
		&appUser.Login,
		&appUser.DisplayName,
		&appUser.Email,
	)
	if err != nil {
		slog.Error("upsert oauth user failed", slog.Any("err", err))
		return authapi.AppUser{}, err
	}

	return appUser, nil
}
