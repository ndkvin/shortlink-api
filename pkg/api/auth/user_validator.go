package auth

import (
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

func (v *Validation) CreateUserValidation(c *fiber.Ctx, req *CreateRequest) (*CreateResponseError,error) {
	if err:=v.Validator.Struct(req); err!=nil {
		response := CreateResponseError{
			Status: "Bad Request",
			Message: err.Error(),
		}

		return &response, err;
	}
	response := CreateResponseError{}

	return &response,nil
}