package handlers

import (
	"github.com/batistell/catalog-api/config"
	model "github.com/batistell/catalog-api/src/models"
	"github.com/batistell/catalog-api/src/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

// AddProduct handles the request to add new products
// @Summary Add new products
// @Description Add new products to the catalog
// @Tags Products
// @Accept json
// @Produce json
// @Param products body []model.Product true "Products array to add"
// @Success 200 {object} model.ProductResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /add [post]
func (h catalogHandler) AddProduct(c *fiber.Ctx) error {
	messageID := uuid.New().String()
	logrus.WithField("messageId", messageID).Info("Starting AddProducts request")

	// Parse the request body to get the products details
	var products []model.Product
	if err := c.BodyParser(&products); err != nil {
		logrus.WithField("messageId", messageID).Error(err)
		return model.NewError(fiber.StatusBadRequest, err.Error()).ToResponse(c)
	}

	if err := h.catalogService.AddProducts(messageID, products); err != nil {
		logrus.WithField("messageId", messageID).Error(err)
		return err.ToResponse(c)
	}

	logrus.WithField("messageId", messageID).Info("AddProducts request completed successfully")
	return c.JSON(model.ProductResponse{
		Message: "Products added successfully",
		Data:    nil,
	})
}

// GetAllProducts returns all products in the catalog
// @Summary Get all products
// @Description Retrieve all products from the catalog
// @Tags Products
// @Produce json
// @Success 200 {object} []model.Product
// @Router /products [get]
func (h catalogHandler) GetAllProducts(c *fiber.Ctx) error {
	messageID := uuid.New().String()
	logrus.WithField("messageId", messageID).Info("Starting GetAllProducts request")

	// Get all products from the catalog
	products, err := h.catalogService.GetAllProducts()
	if err != nil {
		logrus.WithField("messageId", messageID).Error(err)
		return err.ToResponse(c)
	}

	logrus.WithField("messageId", messageID).Info("GetAllProducts request completed successfully")
	return c.JSON(products)
}

// GetProductByID retrieves a product by its ID
// @Summary Get product by ID
// @Description Retrieve a product from the catalog by its ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /products/{id} [get]
func (h catalogHandler) GetProductByID(c *fiber.Ctx) error {
	messageID := uuid.New().String()
	logrus.WithField("messageId", messageID).Info("Starting GetProductByID request")

	id := c.Params("id")

	// Retrieve the product by ID from the catalog
	product, err := h.catalogService.GetProductByID(id)
	if err != nil {
		logrus.WithField("messageId", messageID).Error(err)
		return err.ToResponse(c)
	}

	logrus.WithField("messageId", messageID).Info("GetProductByID request completed successfully")
	return c.JSON(product)
}

// UpdateProduct updates a product in the catalog
// @Summary Update a product
// @Description Update a product in the catalog
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.Product true "Updated product object"
// @Success 200 {object} model.ProductResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /products/{id} [put]
func (h catalogHandler) UpdateProduct(c *fiber.Ctx) error {
	messageID := uuid.New().String()
	logrus.WithField("messageId", messageID).Info("Starting UpdateProduct request")

	// Get the product ID from the path parameters
	id := c.Params("id")

	// Parse the request body to get the updated product details
	productToUpdate := new(model.Product)
	if err := c.BodyParser(productToUpdate); err != nil {
		logrus.WithField("messageId", messageID).Error(err)
		return model.NewError(fiber.StatusBadRequest, err.Error()).ToResponse(c)
	}

	// Update the product in the catalog
	product, err := h.catalogService.UpdateProduct(id, productToUpdate)
	if err != nil {
		logrus.WithField("messageId", messageID).Error(err)
		return err.ToResponse(c)
	}

	logrus.WithField("messageId", messageID).Info("UpdateProduct request completed successfully")
	return c.JSON(model.ProductResponse{
		Message: "Product updated successfully",
		Data:    product,
	})
}

// DeleteProduct deletes a product from the catalog
// @Summary Delete a product
// @Description Delete a product from the catalog by its ID
// @Tags Products
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /products/{id} [delete]
func (h catalogHandler) DeleteProduct(c *fiber.Ctx) error {
	messageID := uuid.New().String()
	logrus.WithField("messageId", messageID).Info("Starting DeleteProduct request")

	// Get the product ID from the path parameters
	id := c.Params("id")

	// Delete the product from the catalog
	err := h.catalogService.DeleteProduct(id)
	if err != nil {
		logrus.WithField("messageId", messageID).Error(err)
		return err.ToResponse(c)
	}

	logrus.WithField("messageId", messageID).Info("DeleteProduct request completed successfully")
	return c.SendStatus(fiber.StatusNoContent)
}
