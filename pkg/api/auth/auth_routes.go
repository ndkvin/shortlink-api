package auth

import (
	"fmt"
	"shortlink/pkg/api/middleware"

	"shortlink/pkg/common/tokenize"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
	
	user.Post("/test", middleware.Auth(), func(c *fiber.Ctx) error {
		
		a:=c.Locals("user").(*jwt.Token)

		test :=tokenize.GetUserId(a.Raw)
		fmt.Printf("test: %v\n", test)
		return c.SendStatus(200)
	})
}