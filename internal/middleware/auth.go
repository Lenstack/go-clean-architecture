package middleware

import (
	"github.com/Lenstack/clean-architecture/internal/domain"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"os"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET"))

func RouteProtected() fiber.Handler {
	config := jwtware.Config{
		SigningKey:     SecretKey,
		SuccessHandler: Next,
		ErrorHandler:   Unauthorized,
		AuthScheme:     "Bearer",
	}
	return jwtware.New(config)
}

func Unauthorized(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(
		domain.Response{
			Status:  fiber.StatusUnauthorized,
			Message: "Authorization Error",
			Data:    err.Error(),
		},
	)
}

func Next(ctx *fiber.Ctx) error {
	err := ctx.Next()
	if err != nil {
		return err
	}
	return nil
}
