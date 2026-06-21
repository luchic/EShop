package repository

import (
	"database/sql"
	"fmt"
	"shop/internal/api"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) PostgresUserRepository {
	return PostgresUserRepository{db: db}
}

func (r PostgresUserRepository) CreateUser(user api.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (first_name, second_name, email, password, role) VALUES ($1, $2, $3, $4, $5)",
		user.FirstName,
		user.SecondName,
		user.Email,
		user.Password,
		user.Role)
	return err
}

func (r PostgresUserRepository) GetUserByEmail(userEmail string) (api.User, error) {
	var user api.User

	rows := r.db.QueryRow(
		"SELECT id, first_name, second_name, email, role, password, created_at FROM users WHERE email = $1",
		userEmail)

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
			return user, fmt.Errorf("GetUserByEmail %s: no such user", userEmail)
		}
		return user, fmt.Errorf("GetUserByEmail %s: %v", userEmail, err)
	}
	return user, nil
}

func (r PostgresUserRepository) GetUserById(userId int64) (api.User, error) {
	var user api.User

	rows := r.db.QueryRow("SELECT id, first_name, second_name, email, role, password, created_at FROM users WHERE id = $1", userId)

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
			return user, fmt.Errorf("GetUserById %d: no such user", userId)
		}
		return user, fmt.Errorf("GetUserById %d: %v", userId, err)
	}
	return user, nil
}
