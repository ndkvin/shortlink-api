package auth

import (
	"shortlink/pkg/common/models"
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

func (v *Validation) CreateUserValidation(c *fiber.Ctx, req *auth.CreateRequest) (response *auth.CreateResponseError,err error) {
	if err=v.Validator.Struct(req); err!=nil {
		response = &auth.CreateResponseError{
			Status: "Bad Request",
			Message: err.Error(),
		}

		return
	}

	return
}

func (v *Validation) LoginValidation(c *fiber.Ctx, req *auth.LoginRequest) (response *auth.CreateResponseError,err error) {
	if err = v.Validator.Struct(req); err!=nil {
		var user models.User

		response = user.CreateResponseFail(err.Error(),"Bad Request")

		return
	}

	return
}