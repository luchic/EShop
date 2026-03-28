package repository

import (
	authapi "backend/shop/internal/api/auth"
	financeapi "backend/shop/internal/api/finance"
	"backend/shop/internal/api/goods"
	"backend/shop/internal/config"
	"context"
	"database/sql"
	"errors"
	"log/slog"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	context context.Context
	db      *sql.DB
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrProductNotFound   = errors.New("product not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrPriceMismatch     = errors.New("price mismatch")
)

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

func (r *PostgresRepository) AddProduct(product goods.AddProductRequest) (goods.AddProductResponse, error) {
	row := r.db.QueryRow(
		`INSERT INTO products (name, logo, price, description)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name`,
		product.Name,
		product.Logo,
		product.Price,
		product.Description,
	)

	var response goods.AddProductResponse
	err := row.Scan(&response.Id, &response.Name)
	if err != nil {
		slog.Error("add product failed", slog.Any("err", err))
		return goods.AddProductResponse{}, err
	}

	return response, nil
}

func (r *PostgresRepository) GetGoodPage(offset int, limit int) []goods.Product {
	rows, err := r.db.Query(
		`SELECT id, name, price, logo, description
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
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Logo, &product.Description)
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
		`SELECT id, name, price, logo, description
		FROM products
		WHERE id = $1`,
		id,
	).Scan(&product.Id, &product.Name, &product.Price, &product.Logo, &product.Description)
	if err != nil {
		if err != sql.ErrNoRows {
			slog.Error("get product by id failed", slog.Any("err", err), slog.Uint64("id", id))
		}
		return goods.Product{}, false
	}

	return product, true
}

func (r *PostgresRepository) GetUserBalance(userID int64) (financeapi.UserBalanceResponse, bool) {
	var response financeapi.UserBalanceResponse
	err := r.db.QueryRow(
		`SELECT id, balance
		FROM users
		WHERE id = $1`,
		userID,
	).Scan(&response.UserID, &response.Balance)
	if err != nil {
		if err != sql.ErrNoRows {
			slog.Error("get user balance failed", slog.Any("err", err), slog.Int64("user_id", userID))
		}
		return financeapi.UserBalanceResponse{}, false
	}

	return response, true
}

func (r *PostgresRepository) RegisterTransaction(request financeapi.RegisterTransactionRequest) (financeapi.Transaction, error) {
	tx, err := r.db.BeginTx(r.context, nil)
	if err != nil {
		slog.Error("begin register transaction failed", slog.Any("err", err))
		return financeapi.Transaction{}, err
	}
	defer tx.Rollback()

	balance, err := r.getUserBalanceForUpdate(tx, request.UserID)
	if err != nil {
		return financeapi.Transaction{}, err
	}

	if balance < request.Price {
		return financeapi.Transaction{}, ErrInsufficientFunds
	}

	err = r.updateUserBalance(tx, request.UserID, balance-request.Price)
	if err != nil {
		return financeapi.Transaction{}, err
	}

	transaction, err := r.createTransaction(tx, request.UserID, request.ProductID, price)
	if err != nil {
		return financeapi.Transaction{}, err
	}

	if err := tx.Commit(); err != nil {
		slog.Error("commit transaction failed", slog.Any("err", err))
		return financeapi.Transaction{}, err
	}

	return transaction, nil
}

func (r *PostgresRepository) GetTransactionByID(id int64) (financeapi.Transaction, bool) {
	var transaction financeapi.Transaction
	err := r.db.QueryRow(
		`SELECT id, product_id, user_id, price, created_at
		FROM transactions
		WHERE id = $1`,
		id,
	).Scan(
		&transaction.ID,
		&transaction.ProductID,
		&transaction.UserID,
		&transaction.Price,
		&transaction.CreatedAt,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			slog.Error("get transaction by id failed", slog.Any("err", err), slog.Int64("id", id))
		}
		return financeapi.Transaction{}, false
	}

	return transaction, true
}

func (r *PostgresRepository) GetTransactionsByUserID(userID int64) ([]financeapi.Transaction, error) {
	if _, ok := r.GetUserBalance(userID); !ok {
		return nil, ErrUserNotFound
	}

	rows, err := r.db.Query(
		`SELECT id, product_id, user_id, price, created_at
		FROM transactions
		WHERE user_id = $1
		ORDER BY id DESC`,
		userID,
	)
	if err != nil {
		slog.Error("get transactions by user id failed", slog.Any("err", err), slog.Int64("user_id", userID))
		return nil, err
	}
	defer rows.Close()

	transactions := make([]financeapi.Transaction, 0)
	for rows.Next() {
		var transaction financeapi.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.ProductID,
			&transaction.UserID,
			&transaction.Price,
			&transaction.CreatedAt,
		)
		if err != nil {
			slog.Error("scan transaction failed", slog.Any("err", err), slog.Int64("user_id", userID))
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		slog.Error("iterate transactions failed", slog.Any("err", err), slog.Int64("user_id", userID))
		return nil, err
	}

	return transactions, nil
}

func (r *PostgresRepository) getUserBalanceForUpdate(tx *sql.Tx, userID int64) (int, error) {
	var balance int
	err := tx.QueryRowContext(
		r.context,
		`SELECT balance
		FROM users
		WHERE id = $1
		FOR UPDATE`,
		userID,
	).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrUserNotFound
		}
		slog.Error("select user balance for transaction failed", slog.Any("err", err), slog.Int64("user_id", userID))
		return 0, err
	}

	return balance, nil
}

func (r *PostgresRepository) updateUserBalance(tx *sql.Tx, userID int64, newBalance int) error {
	_, err := tx.ExecContext(
		r.context,
		`UPDATE users
		SET balance = $2
		WHERE id = $1`,
		userID,
		newBalance,
	)
	if err != nil {
		slog.Error("update user balance failed", slog.Any("err", err), slog.Int64("user_id", userID))
		return err
	}

	return nil
}

func (r *PostgresRepository) createTransaction(tx *sql.Tx, userID int64, productID int64, price int) (financeapi.Transaction, error) {
	var transaction financeapi.Transaction
	err := tx.QueryRowContext(
		r.context,
		`INSERT INTO transactions (product_id, user_id, price)
		VALUES ($1, $2, $3)
		RETURNING id, product_id, user_id, price, created_at`,
		productID,
		userID,
		price,
	).Scan(
		&transaction.ID,
		&transaction.ProductID,
		&transaction.UserID,
		&transaction.Price,
		&transaction.CreatedAt,
	)
	if err != nil {
		slog.Error("insert transaction failed", slog.Any("err", err), slog.Int64("user_id", userID), slog.Int64("product_id", productID))
		return financeapi.Transaction{}, err
	}

	return transaction, nil
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
		`INSERT INTO users (provider_id, login, display_name, email, balance)
		VALUES ($1, $2, $3, $4, 0)
		ON CONFLICT (provider_id) DO UPDATE SET
			login = EXCLUDED.login,
			display_name = EXCLUDED.display_name,
			email = EXCLUDED.email
		RETURNING id, login, display_name, email, balance`,
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
		&appUser.Balance,
	)
	if err != nil {
		slog.Error("upsert oauth user failed", slog.Any("err", err))
		return authapi.AppUser{}, err
	}

	return appUser, nil
}
