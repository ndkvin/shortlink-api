package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	Db         *gorm.DB
	Repository *Repository
	Validation  *Validation
}

func NewHandler(Db *gorm.DB, 	Repository *Repository, Validation  *Validation) *Handler{
	return &Handler{
		Db: Db,
		Repository: Repository,
		Validation: Validation,
	}
}

func (h *Handler) createUser(c *fiber.Ctx) error {
	body := CreateRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if response, err := h.Validation.CreateUserValidation(c, &body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	successResponse, errResponse, err := h.Repository.CreateUser(&body)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(&successResponse)
}
