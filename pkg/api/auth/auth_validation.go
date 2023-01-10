package auth

import (
	"shortlink/pkg/common/resources/auth"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
)

type Validation struct {
	Validator *validator.Validate
}

func NewValidation (Validator *validator.Validate) (validation *Validation) {
	validation = &Validation{
		Validator: Validator,
	}

	return
}

func (v *Validation) CreateUserValidation(c *fiber.Ctx, req *auth.CreateRequest) (err error) {
	if err = v.Validator.Struct(req); err!=nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}

func (v *Validation) LoginValidation(c *fiber.Ctx, req *auth.LoginRequest) (err error) {
	if err = v.Validator.Struct(req); err!=nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}