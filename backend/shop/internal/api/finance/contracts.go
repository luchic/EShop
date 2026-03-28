package finance

type UserBalanceResponse struct {
	UserID  int64 `json:"user_id"`
	Balance int   `json:"balance"`
}
