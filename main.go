package main

import (
	"shortlink/pkg/api/auth"
	"shortlink/pkg/common/db"
	"shortlink/pkg/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func main () {
	db:= db.InitDB()

	app := fiber.New()

	middleware.Register(app)

	// register api
	auth.Register(app,db)

	app.Listen(":8080")
}