package api

import "time"

type RegisterUser struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type User struct {
	Id         int64     `db:"id"`
	FirstName  string    `db:"first_name"`
	SecondName string    `db:"second_name"`
	Email      string    `db:"email"`
	Password   []byte    `db:"password"`
	CreatedAt  time.Time `db:"created_at"`
}
