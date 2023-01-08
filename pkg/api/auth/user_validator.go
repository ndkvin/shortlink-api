package auth

import (
	"shortlink/pkg/common/resources/auth"

	"gopkg.in/go-playground/validator.v9"
	"github.com/gofiber/fiber/v2"
)

type Validation struct {
	Validator *validator.Validate
}

func NewValidation (Validator *validator.Validate) *Validation {
	return &Validation{
		Validator: Validator,
	}
}

func (v *Validation) CreateUserValidation(c *fiber.Ctx, req *auth.CreateRequest) (*auth.CreateResponseError,error) {
	if err:=v.Validator.Struct(req); err!=nil {
		response := auth.CreateResponseError{
			Status: "Bad Request",
			Message: err.Error(),
		}

		return &response, err;
	}
	response := auth.CreateResponseError{}

	return &response,nil
}