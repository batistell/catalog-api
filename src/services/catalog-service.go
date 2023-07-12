package services

import (
	"github.com/batistell/catalog-api/config"
	model "github.com/batistell/catalog-api/src/models"
	"github.com/batistell/catalog-api/src/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CatalogService interface {
	AddProducts(messageID string, products []model.Product) *model.Error
	GetAllProducts() ([]*model.Product, *model.Error)
	GetProductByID(id string) (*model.Product, *model.Error)
	UpdateProduct(id string, updatedProduct *model.Product) (*model.Product, *model.Error)
	DeleteProduct(id string) *model.Error
}

type catalogService struct {
	cfg               *config.Config
	catalogRepository repositories.CatalogRepository
}

func NewCatalogService(cfg *config.Config, catalogRepository repositories.CatalogRepository) CatalogService {
	return &catalogService{cfg: cfg, catalogRepository: catalogRepository}
}

func (s *catalogService) AddProducts(messageID string, products []model.Product) *model.Error {
	if products == nil {
		logrus.WithField("messageId", messageID).Error("Invalid product data")
		return model.NewError(fiber.StatusBadRequest, "Invalid product data")
	}

	for _, product := range products {
		if err := product.ValidateProductData(); err != nil {
			logrus.WithField("messageId", messageID).Error(err)
			return model.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	err := s.catalogRepository.AddProducts(products)
	if err != nil {
		logrus.WithField("messageId", messageID).Error(err)
		return model.NewError(fiber.StatusInternalServerError, "Failed to add products")
	}

	return nil
}

func (s *catalogService) GetAllProducts() ([]*model.Product, *model.Error) {
	// Retrieve all products from the catalog repository
	products, err := s.catalogRepository.GetAllProducts()
	if err != nil {
		return nil, model.NewError(fiber.StatusInternalServerError, "Failed to retrieve products")
	}

	return products, nil
}

func (s *catalogService) GetProductByID(id string) (*model.Product, *model.Error) {
	// Validate the product ID
	if id == "" {
		return nil, model.NewError(fiber.StatusBadRequest, "Invalid product ID")
	}

	// Retrieve the product by ID from the catalog repository
	product, err := s.catalogRepository.GetProductByID(id)
	if err != nil {
		return nil, model.NewError(fiber.StatusInternalServerError, "Failed to retrieve product")
	}

	return product, nil
}

func (s *catalogService) UpdateProduct(id string, updatedProduct *model.Product) (*model.Product, *model.Error) {
	// Validate the product ID and updated product data
	if id == "" {
		return nil, model.NewError(fiber.StatusBadRequest, "Invalid product ID")
	}
	if updatedProduct == nil {
		return nil, model.NewError(fiber.StatusBadRequest, "Invalid updated product data")
	}

	// Retrieve the existing product by ID from the catalog repository
	existingProduct, err := s.catalogRepository.GetProductByID(id)
	if err != nil {
		return nil, model.NewError(fiber.StatusInternalServerError, "Failed to retrieve product")
	}

	// Update the product data
	// TODO: Implement the logic to update the product fields based on the updatedProduct data

	// Save the updated product to the catalog repository
	if err := s.catalogRepository.UpdateProduct(existingProduct); err != nil {
		return nil, model.NewError(fiber.StatusInternalServerError, "Failed to update product")
	}

	return existingProduct, nil
}

func (s *catalogService) DeleteProduct(id string) *model.Error {
	// Validate the product ID
	if id == "" {
		return model.NewError(fiber.StatusBadRequest, "Invalid product ID")
	}

	// Delete the product from the catalog repository
	if err := s.catalogRepository.DeleteProduct(id); err != nil {
		return model.NewError(fiber.StatusInternalServerError, "Failed to delete product")
	}

	return nil
}
