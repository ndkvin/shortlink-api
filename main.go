package main

import (
	"shortlink/pkg/api/auth"
	"shortlink/pkg/api/link"
	"shortlink/pkg/api/visit_link"
	"shortlink/pkg/api/middleware"
	"shortlink/pkg/common/db"

	"github.com/gofiber/fiber/v2"
)

func main () {
	db:= db.InitDB()

	app := fiber.New(fiber.Config{
    ErrorHandler: middleware.ErrorHandler,
	})

	middleware.Register(app)

	// register service
	auth.Register(app, db)
	link.Register(app, db)
	visit_link.Register(app, db)
	app.Static("/", "./public")

	app.Listen(":8080")
}