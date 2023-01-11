package auth

import (
	"shortlink/pkg/common/tokenize"
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

	if err := h.Validation.CreateUserValidation(c, &body); err != nil {
		return err
	}

	successResponse, err := h.Repository.CareateUser(&body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(successResponse)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	body := auth.LoginRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.Validation.LoginValidation(c, &body); err != nil {
		return err
	}

	user, err := h.Repository.Login(&body)

	if err != nil {
		return err
	} 

	jwtToken, _ := tokenize.GenereateToken(user.ID.String())
	response := user.CreateLoginResponse(jwtToken)
	
	return c.Status(200).JSON(response)
}