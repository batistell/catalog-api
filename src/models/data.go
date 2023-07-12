package model

import (
	"time"
)

type ProductResponse struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
}

type Product struct {
	ID          string      `json:"_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Category    Category    `json:"category"`
	Parents     []Product   `json:"parents"`
	Children    []Product   `json:"children"`
	Discounts   []Discount  `json:"discounts"`
	Attributes  []Attribute `json:"attributes"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type Category struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Discount struct {
	ID         string    `json:"_id"`
	Name       string    `json:"name"`
	Percentage int       `json:"percentage"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Attribute struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
