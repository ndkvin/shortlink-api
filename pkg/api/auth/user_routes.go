package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gopkg.in/go-playground/validator.v9"
)

func Register(app *fiber.App,Db *gorm.DB) {
	
	repository := NewRepository(Db)
	validator := NewValidation(validator.New())
	h := NewHandler(Db,repository,validator)

	user := app.Group("/user");

	user.Post("/", h.CreateUser)
}