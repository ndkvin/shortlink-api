package visit_link

import (
	"shortlink/pkg/common/resources/visit_link"

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

func (v *Validation) VisitLinkValidation(req *visit_link.VisitLinkRequest) (err error) {
	if err = v.Validator.Struct(req); err!=nil {
		err = fiber.NewError(400, err.Error())
		return
	}

	return
}