package services

import (
	"github.com/batistell/catalog-api/config"
	"github.com/batistell/catalog-api/src/repositories"
)

type CatalogService interface {
	AddProduct()
}

type catalogService struct {
	cfg               *config.Config
	catalogRepository repositories.CatalogRepository
}

func NewCatalogService(cfg *config.Config, catalogRepository repositories.CatalogRepository) CatalogService {
	return &catalogService{cfg: cfg, catalogRepository: catalogRepository}
}

func (s *catalogService) AddProduct() {

}
