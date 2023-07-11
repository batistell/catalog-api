package handlers

import (
	"github.com/batistell/catalog-api/config"
	"github.com/batistell/catalog-api/src/services"
	"github.com/gofiber/fiber/v2"
)

type CatalogHandler interface {
	AddProduct(c *fiber.Ctx) error
}

type catalogHandler struct {
	cfg            *config.Config
	catalogService services.CatalogService
}

func NewCatalogHandler(cfg *config.Config, catalogService services.CatalogService) CatalogHandler {
	return &catalogHandler{cfg: cfg, catalogService: catalogService}
}

// Product represents a product entity
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
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

// AddProduct handles the request to add a new product
// @Summary Add a new product
// @Description Add a new product to the catalog
// @Tags Products
// @Accept json
// @Produce json
// @Param product body Product true "Product object to add"
// @Success 200 {object} ProductResponse
// @Failure 400 {object} ErrorResponse
// @Router /add [post]
func (c2 catalogHandler) AddProduct(c *fiber.Ctx) error {
	// Parse the request body to get the product details

	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return err
	}

	// Return a success response
	return c.JSON(ProductResponse{
		Message: "Product added successfully",
		Data:    nil,
	})
}
