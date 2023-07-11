package main

import (
	"github.com/batistell/catalog-api/config"
	. "github.com/batistell/catalog-api/src/server"
	"github.com/sirupsen/logrus"
	"os"
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

	db, err := NewMongoDB(logger, cfg)
	if err != nil {
		logger.WithError(err).Error("Error NewMongoDB")
	}

	server := NewServer(cfg, *db)
	err = server.Run(logger)
	if err != nil {
		logger.WithError(err).Error("Error Running Server")
	}

	logger.Info("Application initialized")
}
