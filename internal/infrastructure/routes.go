package infrastructure

import (
	"github.com/Lenstack/clean-architecture/internal/interfaces"
	"github.com/Lenstack/clean-architecture/internal/middleware"
	"github.com/Lenstack/clean-architecture/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"os"
)

func NewRoutes(app *fiber.App, logger usecases.LoggerRepository, mongo usecases.MongoRepository) {
	UserRoutes(app, logger, mongo)
}

func UserRoutes(app *fiber.App, logger usecases.LoggerRepository, mongo usecases.MongoRepository) {
	userHandler := interfaces.NewUserHandler(logger, mongo)
	version := app.Group(os.Getenv("API_VERSION"))

	//Public Routes
	routes := version.Group("user")

	routes.Post("/", userHandler.Create)
	//Protected Routes
	routes.Get("/", middleware.RouteProtected(), userHandler.Index)
	routes.Get("/:id", middleware.RouteProtected(), userHandler.Show)
	routes.Post("/", middleware.RouteProtected(), userHandler.Create)
	routes.Put("/:id", middleware.RouteProtected(), userHandler.Update)
	routes.Delete("/:id", middleware.RouteProtected(), userHandler.Destroy)
}
