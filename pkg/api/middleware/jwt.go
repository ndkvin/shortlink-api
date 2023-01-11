package middleware

import (
	"shortlink/pkg/common/config"

	"github.com/gofiber/fiber/v2"
	jwtfiber "github.com/gofiber/jwt/v2"
)

func Auth() fiber.Handler {
	config, _ := config.InitConfig()

	SigningKey := []byte(config.JWT_TOKEN)

	return jwtfiber.New(jwtfiber.Config{
		SigningKey:   []byte(SigningKey),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing or malformed JWT")
	}
	return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired JWT")
}
