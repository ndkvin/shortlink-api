package user

import (
	"log"
	"shortlink/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) createUser(c *fiber.Ctx) error {
	body := CreateRequest{}

	err := c.BodyParser(&body)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	user.Email = body.Email
	user.Name = body.Name
	user.Password = body.Password

	result := h.db.Create(&user)

	if result.Error != nil {
		log.Fatalln(result.Error)
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	response := CreateResponse{
		Status: "success",
	}

	return c.Status(fiber.StatusCreated).JSON(&response)
}

