package user

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	Db *gorm.DB
	Repository *Repository
}

func (h *Handler) createUser(c *fiber.Ctx) error {
	body := CreateRequest{}

	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, error := h.Repository.createUser(&body)

	if error != nil {
		log.Fatalln(error.Error())
		response := CreateResponseError{
			Status: "Internal Server error",
			Message: "An intermal server error",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(&response)
	}

	response := CreateResponseSuccess{
		Status: "success",
		Message: "User created",
		UserId: result,
	}

	return c.Status(fiber.StatusCreated).JSON(&response)
}
