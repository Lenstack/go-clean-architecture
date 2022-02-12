package middleware

import (
	"github.com/Lenstack/clean-architecture/internal/interfaces"
	"github.com/Lenstack/clean-architecture/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"os"
)

func NewRoutes(app *fiber.App, logger usecases.Logger, mongo usecases.Mongo) {
	userHandler := interfaces.NewUserHandler(logger, mongo)
	version := app.Group(os.Getenv("API_VERSION"))

	routes := version.Group("user")
	routes.Get("/", userHandler.Index)
	routes.Get("/:id", userHandler.Show)
	routes.Post("/", userHandler.Create)
	routes.Put("/:id", userHandler.Update)
	routes.Delete("/:id", userHandler.Destroy)
}
