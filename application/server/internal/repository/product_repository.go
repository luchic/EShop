package repository

import (
	"database/sql"
	"shop/internal/api"
)

// PostgresProductRepository
type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) PostgresProductRepository {
	return PostgresProductRepository{db: db}
}

func (r PostgresProductRepository) CreateProduct(product api.Product) error {
	_, err := r.db.Exec(
		"INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4)",
		product.Name,
		product.Description,
		product.Price,
		product.Stock)
	return err
}

func (r PostgresProductRepository) GetProductById(id int64) (api.Product, error) {
	row := r.db.QueryRow(
		"SELECT id, name, description, price, stock, image_url, created_at FROM products WHERE id = $1",
		id)

	var product api.Product
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.ImageUrl,
		&product.CreatedAt)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r PostgresProductRepository) GetProductsByName(name string) ([]api.Product, error) {
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
			&product.Description,
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

// func (r *PostgresProductRepository) UpdateProduct() {

// }

// func (r *PostgresProductRepository) DeleteProduct() {

// }
