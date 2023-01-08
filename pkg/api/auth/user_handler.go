package auth

import (
	"shortlink/pkg/common/resources/auth"

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

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	body := auth.CreateRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if response, err := h.Validation.CreateUserValidation(c, &body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	sr, er, code := h.Repository.CareateUser(&body)
	if code != 201 {
		return c.Status(code).JSON(er)
	} 
	return c.Status(code).JSON(sr)
}
