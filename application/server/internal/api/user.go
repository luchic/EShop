package api

import "time"

type User struct {
	Id         int64     `db:"id"`
	FirstName  string    `db:"first_name"`
	SecondName string    `db:"second_name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	CreatedAt  time.Time `db:"created_at"`
}
