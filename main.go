package main

import (
	"github.com/batistell/catalog-api/config"
	. "github.com/batistell/catalog-api/src/server"
	"github.com/sirupsen/logrus"
	"os"
)

// @title Catalog API
// @description API for Managing Products from Catalog.
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email batistell.labs@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	logger := logrus.New()
	logger.Info("Initializing application...")

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.WithError(err).Fatal("Error LoadConfig")
	}

	logger.SetLevel(cfg.Logger.Level)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)

	db, err := NewMongoDB(logger, cfg)
	if err != nil {
		logger.WithError(err).Fatal("Error NewMongoDB")
	}

	server := NewServer(cfg, *db)
	err = server.Run(logger)
	if err != nil {
		logger.WithError(err).Fatal("Error Running Server")
	}

	logger.Info("Application initialized")
}
