package user

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/gofiber/fiber/v2"
)

type Validator struct {
	Validator *validator.Validate
}

func (v *Validator) CreateUserValidator(c *fiber.Ctx, req *CreateRequest) (*CreateResponseError,error) {
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