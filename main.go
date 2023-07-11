package main

import (
	"context"
	"github.com/batistell/catalog-api/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func main() {
	logger := logrus.New()
	logger.Info("Initializing application...")

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.WithError(err).Error("Error LoadConfig")
	}

	logger.SetLevel(cfg.Logger.Level)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)

	// MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.URI))
	if err != nil {
		logger.WithError(err).Error("Failed to connect to MongoDB")
		return
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			logger.WithError(err).Error("Failed to disconnect from MongoDB")
		}
	}()

	logger.Info("Connected to MongoDB")

	logger.Info("Application initialized")
}
