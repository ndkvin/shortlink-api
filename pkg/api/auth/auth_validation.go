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

func (v *Validation) CreateUserValidation(req *auth.CreateRequest) (err error) {
	if err = v.Validator.Struct(req); err!=nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}

func (v *Validation) LoginValidation(req *auth.LoginRequest) (err error) {
	if err = v.Validator.Struct(req); err!=nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}

func (v *Validation) ChangePasswordValidation(req *auth.ChangePasswordRequest) (err error) {
	if err = v.Validator.Struct(req); err != nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}

func (v *Validation) EditProfileValidation(req *auth.EditProfileReques) (err error) {
	if err = v.Validator.Struct(req); err != nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}