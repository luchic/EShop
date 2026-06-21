package repository

import (
	"database/sql"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {

	return &PostgresProductRepository{db: db}
}
