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
	for _, product := range products {
		if err := product.ValidateProductData(); err != nil {
			logrus.WithField("messageId", messageID).Error(err)
			return model.NewError(fiber.StatusBadRequest, err.Error())
		}

		s.catalogRepository.AddProduct()
	}

	return nil
}

func (s *catalogService) GetAllProducts() ([]*model.Product, *model.Error) {
	// Retrieve all products from the catalog repository

	// Handle any errors and return the products or an error response

	return nil, nil
}

func (s *catalogService) GetProductByID(id string) (*model.Product, *model.Error) {
	// Validate the product ID

	// Retrieve the product by ID from the catalog repository

	// Handle any errors and return the product or an error response

	return nil, nil
}

func (s *catalogService) UpdateProduct(id string, updatedProduct *model.Product) (*model.Product, *model.Error) {
	// Validate the product ID and updated product data

	// Retrieve the existing product by ID from the catalog repository

	// Update the product data

	// Save the updated product to the catalog repository

	// Handle any errors and return the updated product or an error response

	return nil, nil
}

func (s *catalogService) DeleteProduct(id string) *model.Error {
	// Validate the product ID

	// Delete the product from the catalog repository

	// Handle any errors and return an appropriate response

	return nil
}
