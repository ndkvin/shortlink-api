package link

import (
	"shortlink/pkg/api/middleware"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

func Register(app *fiber.App,Db *gorm.DB) {

	repository := NewRepository(Db)
	validator := NewValidation(validator.New())
	h := NewHandler(Db,repository,validator)

	user := app.Group("/link");

	user.Post("/", middleware.Auth(), h.CreateLink)
	user.Get("/", middleware.Auth(), h.GetAllLink)
	user.Get("/:id", middleware.Auth(), h.GetLink)
	user.Put("/:id", middleware.Auth(), h.EditLink)
}