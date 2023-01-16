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
	user.Get("/", middleware.Auth(), h.GetUser)
	user.Put("/password", middleware.Auth(), h.ChangePassword)
	user.Put("/profile", middleware.Auth(), h.EditProfile)
	user.Post("/login", h.Login)
}