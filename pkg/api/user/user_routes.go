package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func Register(app *fiber.App,DB *gorm.DB) {
	h := &Handler{
		db: DB,
	}

	user := app.Group("/user");

	user.Post("/", h.createUser)
}