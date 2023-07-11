package models

import "github.com/batistell/catalog-api/src/handlers"

// ErrorResponse represents the structure for error responses
type ErrorResponse struct {
	Message string `json:"message"`
}

// ProductResponse represents the response structure for the product endpoint
type ProductResponse struct {
	Message string            `json:"message"`
	Data    *handlers.Product `json:"data"`
}
