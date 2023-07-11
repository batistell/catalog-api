package server

import (
	"encoding/json"
	"github.com/batistell/catalog-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server struct
type Server struct {
	config *config.Config
	db     mongo.Client
}

func healthz(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func addProduct(c *fiber.Ctx) error {
	return c.SendString("OK")
}

// NewServer New Server constructor
func NewServer(config *config.Config, db mongo.Client) *Server {
	return &Server{config: config, db: db}
}

func (s *Server) Run(logger *logrus.Logger) error {
	app := fiber.New(fiber.Config{
		AppName:     s.config.Server.ServiceName,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(recover.New())

	app.Get("/health", healthz)

	v1 := app.Group("")
	v1.Post("/add", addProduct)

	logger.WithField("port", s.config.Server.Port).Infof("Server is running")
	err := app.Listen(s.config.Server.Port)
	if err != nil {
		logger.WithError(err).Error("Error running server")
		return err
	}
	return nil
}
