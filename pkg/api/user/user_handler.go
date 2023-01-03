package user

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	Db         *gorm.DB
	Repository *Repository
	Validator  *Validator
}

func (h *Handler) createUser(c *fiber.Ctx) error {
	body := CreateRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if response, err := h.Validator.CreateUserValidator(c, &body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	result, err := h.Repository.CreateUser(&body)

	if err != nil {
		log.Fatalln(err.Error())
		response := CreateResponseError{
			Status:  "Internal Server error",
			Message: "An intermal server error",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(&response)
	}

	response := CreateResponseSuccess{
		Status:  "success",
		Message: "User created",
		UserId:  result,
	}

	return c.Status(fiber.StatusCreated).JSON(&response)
}
