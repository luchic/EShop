package goods

type Product struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PaginatedResponse struct {
	Items      []Product `json:"items"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
}
