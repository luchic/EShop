package api

import "time"

type Product struct {
	Id          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float32   `db:"price"`
	Stock       int32     `db:"stock"`
	ImageUrl    string    `db:"image_url"`
	CreatedAt   time.Time `db:"created_at"`
}
