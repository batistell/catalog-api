package main

import (
	_ "github.com/batistell/catalog-api/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Catalog API
// @description API for Managing Products from Catalog.
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email batistell.labs@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func sandbox() {
	app := fiber.New()

	// Serve the Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Endpoint to add a product
	app.Post("/products", addProduct)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

// Product represents a product entity
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// AddProduct handles the request to add a new product
// @Summary Add a new product
// @Description Add a new product to the catalog
// @Tags Products
// @Accept json
// @Produce json
// @Param product body Product true "Product object to add"
// @Success 200 {object} ProductResponse
// @Failure 400 {object} ErrorResponse
// @Router /products [post]
func addProduct(c *fiber.Ctx) error {
	// Parse the request body to get the product details
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return err
	}

	// TODO: Add the logic to save the product to the database or perform any other required operations

	// Return a success response
	return c.JSON(ProductResponse{
		Message: "Product added successfully",
		Data:    nil,
	})
}

// ErrorResponse represents the structure for error responses
type ErrorResponse struct {
	Message string `json:"message"`
}

// ProductResponse represents the response structure for the product endpoint
type ProductResponse struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
}
