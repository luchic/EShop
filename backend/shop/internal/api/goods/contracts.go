package goods

type Product struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Logo        string `json:"logo"`
	Description string `json:"desc"`
}

type AddProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Price       int    `json:"price"`
	Logo        string `json:"logo"`
}

type AddProductResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PaginatedResponse struct {
	Items []Product `json:"items"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
}
