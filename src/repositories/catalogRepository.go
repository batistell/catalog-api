package repositories

import (
	"github.com/batistell/catalog-api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type CatalogRepository interface {
	AddProduct()
}

type catalogRepository struct {
	cfg *config.Config
	db  mongo.Client
}

func NewCatalogRepository(cfg *config.Config, db mongo.Client) CatalogRepository {
	return &catalogRepository{cfg: cfg, db: db}
}

func (r *catalogRepository) AddProduct() {

}
