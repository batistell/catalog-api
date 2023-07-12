package repositories

import (
	"github.com/batistell/catalog-api/config"
	model "github.com/batistell/catalog-api/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CatalogRepository interface {
	AddProduct()
	GetAllProducts() ([]*model.Product, *model.Error)
	GetProductByID(id string) (*model.Product, *model.Error)
	UpdateProduct(existingProduct model.Product) *model.Error
	DeleteProduct(id string) *model.Error
}

type catalogRepository struct {
	cfg *config.Config
	db  mongo.Client
}

func NewCatalogRepository(cfg *config.Config, db mongo.Client) CatalogRepository {
	return &catalogRepository{cfg: cfg, db: db}
}

func (r *catalogRepository) GetAllProducts() ([]*model.Product, *model.Error) {

	panic("implement me")
}

func (r *catalogRepository) GetProductByID(id string) (*model.Product, *model.Error) {
	//TODO implement me
	panic("implement me")
}

func (r *catalogRepository) UpdateProduct(existingProduct model.Product) *model.Error {
	//TODO implement me
	panic("implement me")
}

func (r *catalogRepository) DeleteProduct(id string) *model.Error {
	//TODO implement me
	panic("implement me")
}

func (r *catalogRepository) AddProduct() {

}
