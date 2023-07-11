package server

import (
	"context"
	"errors"
	"github.com/batistell/catalog-api/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoDB(logger *logrus.Logger, cfg *config.Config) (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.URI))
	if err != nil {
		logger.WithError(err).Error("Failed to connect to MongoDB")
		return
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			logger.WithError(err).Error("Failed to disconnect from MongoDB")
		}
	}()

	if client == nil {
		err = errors.New("MongoDB client is nil")
		logger.WithError(err).Error("Failed to connect to MongoDB")
		return
	}

	logger.Info("Connected to MongoDB")
	return
}
