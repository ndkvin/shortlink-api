package visit_link

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

func Register(app *fiber.App,Db *gorm.DB) {

	repository := NewRepository(Db)
	validator := NewValidation(validator.New())
	h := NewHandler(Db,repository,validator)

	link := app.Group("/visit");

	link.Post("/:slug", h.VisitLink)
	link.Post("/password/:slug", h.VisitLinkPassword)
}