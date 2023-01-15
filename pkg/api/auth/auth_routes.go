package auth

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

	user := app.Group("/user");

	user.Post("/", h.CreateUser)
	user.Put("/", middleware.Auth(), h.ChangePassword)
	user.Post("/login", h.Login)
}