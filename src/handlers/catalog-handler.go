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
func (h *catalogHandler) AddProduct(c *fiber.Ctx) error {
	return c.SendString("OK")
}
