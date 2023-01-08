package main

import (
	"shortlink/pkg/api/auth"
	"shortlink/pkg/common/db"

	"github.com/gofiber/fiber/v2"
)

func main () {
	db:= db.InitDB()

	app := fiber.New()

	auth.Register(app,db)

	app.Listen(":8080")
}