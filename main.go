package main

import (
	"shortlink/pkg/api/auth"
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

	// register api
	auth.Register(app,db)

	app.Listen(":8080")
}