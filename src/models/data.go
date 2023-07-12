package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type ProductResponse struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
}

type Product struct {
	ID          string      `json:"_id" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Price       float64     `json:"price" validate:"required,gt=0"`
	Category    Category    `json:"category" validate:"required"`
	Parents     []Product   `json:"parents" validate:"customParentsValidation"`
	Children    []Product   `json:"children" validate:"customChildrenValidation"`
	Discounts   []Discount  `json:"discounts" validate:"customDiscountsValidation"`
	Attributes  []Attribute `json:"attributes" validate:"customAttributesValidation"`
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

func (p *Product) ValidateProductData() error {
	// Create a new instance of the validator
	validate := validator.New()

	// Register custom validation functions
	err := validate.RegisterValidation("customParentsValidation", customParentsValidation)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("customChildrenValidation", customChildrenValidation)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("customDiscountsValidation", customDiscountsValidation)
	if err != nil {
		return err
	}
	err = validate.RegisterValidation("customAttributesValidation", customAttributesValidation)
	if err != nil {
		return err
	}

	// Validate the fields of the Product struct
	if err := validate.Struct(p); err != nil {
		return err
	}

	return nil
}

// Custom validation function for the Parents field
func customParentsValidation(fl validator.FieldLevel) bool {
	parents := fl.Field().Interface().([]Product)
	for _, parent := range parents {
		if parent.ID == "" && (parent.Name == "" || parent.Description == "" || parent.Price == 0) {
			return false
		}
	}
	return true
}

// Custom validation function for the Children field
func customChildrenValidation(fl validator.FieldLevel) bool {
	children := fl.Field().Interface().([]Product)
	for _, child := range children {
		if child.ID == "" && (child.Name == "" || child.Description == "" || child.Price == 0) {
			return false
		}
	}
	return true
}

// Custom validation function for the Discounts field
func customDiscountsValidation(fl validator.FieldLevel) bool {
	discounts := fl.Field().Interface().([]Discount)
	for _, discount := range discounts {
		if discount.ID == "" && (discount.Name == "" || discount.Percentage == 0) {
			return false
		}
	}
	return true
}

// Custom validation function for the Attributes field
func customAttributesValidation(fl validator.FieldLevel) bool {
	attributes := fl.Field().Interface().([]Attribute)
	for _, attribute := range attributes {
		if attribute.ID == "" && attribute.Name == "" {
			return false
		}
	}
	return true
}
