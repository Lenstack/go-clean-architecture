package infrastructure

import (
	"github.com/Lenstack/clean-architecture/internal/middleware"
	"github.com/Lenstack/clean-architecture/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"os"
)

func Dispatch(logger usecases.LoggerRepository, mongo usecases.MongoRepository) {
	app := fiber.New()
	middleware.NewMiddleware(app)
	NewRoutes(app, logger, mongo)

	if err := app.Listen(":" + os.Getenv("SERVER_PORT")); err != nil {
		logger.LogError("%s", err)
	}
}
