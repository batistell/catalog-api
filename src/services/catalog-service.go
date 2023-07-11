package services

import "github.com/batistell/catalog-api/config"

type CatalogServiceInterface interface {
	AddProduct()
}

type CatalogService struct {
	cfg config.Config
}
