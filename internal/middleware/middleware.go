package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewMiddleware(app *fiber.App) {
	app.Use(cors.New())
	//app.Use(csrf.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] - [${ip}:${port}] ${status} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
	}))
}
