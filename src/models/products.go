package model

// Product represents a product entity
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ProductResponse represents the response structure for the product endpoint
type ProductResponse struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
}
