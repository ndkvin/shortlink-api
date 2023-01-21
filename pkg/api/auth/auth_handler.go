package auth

import (
	"shortlink/pkg/common/resources/auth"
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

func (h *Handler) CreateUser(c *fiber.Ctx) (err error) {
	body := auth.CreateRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = h.Validation.CreateUserValidation(&body); err != nil {
		return
	}

	successResponse, err := h.Repository.CareateUser(&body)
	if err != nil {
		return
	}
	c.Set(fiber.HeaderAccessControlAllowOrigin, "http://localhost:3000")
	return c.Status(fiber.StatusCreated).JSON(successResponse)
}

func (h *Handler) Login(c *fiber.Ctx) (err error) {
	body := auth.LoginRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = h.Validation.LoginValidation(&body); err != nil {
		return
	}

	user, err := h.Repository.Login(&body)

	if err != nil {
		return err
	} 

	jwtToken, _ := tokenize.GenereateToken(user.ID, user.Name)
	response := user.LoginResponse(jwtToken)
	
	return c.Status(200).JSON(response)
}

func (h *Handler) ChangePassword(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	body := auth.ChangePasswordRequest{}

	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = h.Validation.ChangePasswordValidation(&body); err != nil{
		return
	}

	response, err := h.Repository.ChangePassword(&body, userId)

	if err != nil {
		return
	}

	return c.Status(200).JSON(response)
}

func (h *Handler) EditProfile(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	body := auth.EditProfileRequest{}

	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = h.Validation.EditProfileValidation(&body); err != nil{
		return
	}
	
	res, err := h.Repository.EditProfile(&body, userId)

	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *Handler) GetUser(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	res, err := h.Repository.GetUser(userId)

	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}