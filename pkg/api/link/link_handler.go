package link

import (
	"shortlink/pkg/common/resources/link"
	"shortlink/pkg/common/tokenize"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
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

	qr := uuid.NewString()

	successResponse, err := h.Repository.CreateLink(&body, userId, qr)
	if err != nil {
		return
	}

	if err = CreateQR(qr, successResponse.Data.Slug); err != nil {
		return
	}

	return c.Status(fiber.StatusCreated).JSON(successResponse)
}

func (h *Handler) GetAllLink(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err :=tokenize.GetUserId(h.Db, jwt.Raw)

	if err != nil {
		return
	}

	res, err := h.Repository.GetAllLink(userId)

	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *Handler) GetLink(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err :=tokenize.GetUserId(h.Db, jwt.Raw)

	if err != nil {
		return
	}

	res, err :=h.Repository.GetLink(c.Params("id"), userId)

	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *Handler) EditLink(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
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

	qr := uuid.NewString()

	res, err := h.Repository.EditLink(&body, c.Params("id"), userId, qr)
	if err != nil {
		return
	}

	if err = CreateQR(qr, res.Data.Slug); err != nil {
		return
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *Handler) DeleteLink(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	res, err := h.Repository.DeleteLink(c.Params("id"), userId)

	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *Handler) AddPassword(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	body := link.AddPasswordRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = h.Validation.AddPasswordValidation(&body); err != nil {
		return 
	}

	res, err := h.Repository.AddPassword(c.Params("id"), body.Password, userId)
	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *Handler) EditPassword(c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	body := link.EditPasswordRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = h.Validation.EditPasswordValidation(&body); err != nil {
		return 
	}

	res, err := h.Repository.EditPassword(&body, c.Params("id"), userId)
	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *Handler) DeletePassword (c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	body := link.AddPasswordRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = h.Validation.AddPasswordValidation(&body); err != nil {
		return 
	}

	res, err := h.Repository.DeletePassword(c.Params("id"), body.Password, userId)
	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *Handler) LockLink (c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	res, err := h.Repository.LockLink(c.Params("id"), userId)
	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *Handler) UnlockLink (c *fiber.Ctx) (err error) {
	jwt := c.Locals("user").(*jwt.Token)

	userId, err := tokenize.GetUserId(h.Db, jwt.Raw)
	if err != nil {
		return
	}

	res, err := h.Repository.UnlockLink(c.Params("id"), userId)
	if err != nil {
		return
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
