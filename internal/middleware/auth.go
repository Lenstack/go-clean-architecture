package middleware

import (
	"github.com/Lenstack/clean-architecture/internal/domain"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"os"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET"))

func RouteProtected() func(*fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey:   SecretKey,
		ErrorHandler: Unauthorized,
	}
	return jwtware.New(config)
}

func Unauthorized(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(
		domain.Response{
			Status:  fiber.StatusUnauthorized,
			Message: "Authentication Error",
			Data:    err.Error(),
		},
	)
}
