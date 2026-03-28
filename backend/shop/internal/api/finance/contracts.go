package finance

import "time"

type UserBalanceResponse struct {
	UserID  int64 `json:"user_id"`
	Balance int   `json:"balance"`
}

type RegisterTransactionRequest struct {
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	Price     int   `json:"price"`
}

type Transaction struct {
	ID        int64     `json:"id"`
	ProductID int64     `json:"product_id"`
	UserID    int64     `json:"user_id"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
