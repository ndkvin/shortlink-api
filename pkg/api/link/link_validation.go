package link

import (
	"shortlink/pkg/common/resources/link"

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

func (v *Validation) CreateLinkValidation(req *link.CreateRequest) (err error) {
	if err = v.Validator.Struct(req); err!=nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}

func (v *Validation) AddPasswordValidation(req *link.AddPasswordRequest) (err error) {
	if err = v.Validator.Struct(req); err!=nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}

func (v *Validation) EditPasswordValidation(req *link.EditPasswordRequest) (err error) {
	if err = v.Validator.Struct(req); err!=nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}