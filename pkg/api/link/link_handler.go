package link

import (
	"shortlink/pkg/common/resources/link"
	"shortlink/pkg/common/tokenize"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

func (h *Handler) CreateLink(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err :=tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	body := link.CreateRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = h.Validation.CreateLinkValidation(&body); err != nil {
		return
	}

	successResponse, err := h.Repository.CreateLink(&body, userId)
	if err != nil {
		return
	}

	return c.Status(fiber.StatusCreated).JSON(successResponse)
}