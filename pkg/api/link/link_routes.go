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

	link := app.Group("/link");

	link.Post("/", middleware.Auth(), h.CreateLink)
	link.Get("/", middleware.Auth(), h.GetAllLink)
	link.Get("/:id", middleware.Auth(), h.GetLink)
	link.Put("/:id", middleware.Auth(), h.EditLink)
	link.Post("/password/:id", middleware.Auth(), h.AddPassword)
	link.Patch("/password/:id", middleware.Auth(), h.EditPassword)
	link.Delete("/password/:id", middleware.Auth(), h.DeletePassword)
	link.Delete("/:id", middleware.Auth(), h.DeleteLink)
}