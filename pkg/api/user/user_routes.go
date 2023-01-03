package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gopkg.in/go-playground/validator.v9"
)

func Register(app *fiber.App,DB *gorm.DB) {
	repository := &Repository{
		Db: DB,
	}

	validator := &Validator{
		Validator: validator.New(),
	}
	
	h := &Handler{
		Db: DB,
		Repository: repository,
		Validator: validator,
	}

	user := app.Group("/user");

	user.Post("/", h.createUser)
}