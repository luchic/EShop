package repository

import (
	"database/sql"
	"fmt"
	"shop/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
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
	return db, nil
}
