package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/batistell/catalog-api/config"
	model "github.com/batistell/catalog-api/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

type CatalogRepository interface {
	AddProducts(products []model.Product) error
	GetAllProducts() ([]*model.Product, error)
	GetProductByID(id string) (*model.Product, error)
	UpdateProduct(existingProduct *model.Product) error
	DeleteProduct(id string) error
}

type catalogRepository struct {
	cfg *config.Config
	db  mongo.Database
}

func NewCatalogRepository(cfg *config.Config, db mongo.Database) CatalogRepository {
	return &catalogRepository{cfg: cfg, db: db}
}

func (r *catalogRepository) AddProducts(products []model.Product) error {
	collection := r.db.Collection("products")

	documents := make([]interface{}, 0, len(products))
	documents = append(documents, products)

	_, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		return err
	}

	return nil
}

func (r *catalogRepository) GetAllProducts() ([]*model.Product, error) {
	collection := r.db.Collection("products")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.TODO())

	var products []*model.Product
	for cursor.Next(context.Background()) {
		var product model.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (r *catalogRepository) GetProductByID(id string) (*model.Product, error) {
	collection := r.db.Collection("products")

	var product model.Product
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (r *catalogRepository) UpdateProduct(existingProduct *model.Product) error {
	collection := r.db.Collection("products")

	filter := bson.M{"_id": existingProduct.ID}
	update := bson.M{"$set": existingProduct}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("failed to update product")
	}

	return nil
}

func (r *catalogRepository) DeleteProduct(id string) error {
	collection := r.db.Collection("products")

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("failed to delete product")
	}

	return nil
}
