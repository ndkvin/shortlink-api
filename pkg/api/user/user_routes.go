package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app *fiber.App,DB *gorm.DB) {
	repository := &Repository{
		Db: DB,
	}
	
	h := &Handler{
		Db: DB,
		Repository: repository,
	}

	user := app.Group("/user");

	user.Post("/", h.createUser)
}