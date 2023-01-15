package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

    var e *fiber.Error

    if errors.As(err, &e) {
        code = e.Code
    }
		
		if code == 400 {
			err = ctx.Status(code).JSON(fiber.Map{
				"status": "Bad Request",
				"message": err.Error(),
				"code": code,
			})

			return err
		}

		if code == 401 {
			err = ctx.Status(code).JSON(fiber.Map{
				"status": "Unauthorized",
				"message": err.Error(),
				"code": code,
			})

			return err
		}

		if code == 404 {
			err = ctx.Status(code).JSON(fiber.Map{
				"status": "Not Found",
				"message": err.Error(),
				"code": code,
			})

			return err
		}
		
		if code == 500 {
			err = ctx.Status(code).JSON(fiber.Map{
				"status": "Internal Server Error",
				"message": "An intermal server error occur",
				"code": code,
			})

			return err
		}

    return err
}