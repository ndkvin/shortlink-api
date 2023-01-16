package middleware

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Register(app *fiber.App) {
	// recover
	app.Use(recover.New())

	// limitter
	app.Use(limiter.New(limiter.Config{
		Max: 20,
	}))

	// logger
	file, err := os.OpenFile("./serverlog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		Output: file,
		Format: "[${time}] ${ip}:${port} ${status} ${latency} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 - 15:04:05",
	}))

	// monitoring
	app.Get("/monitor", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// cors
	app.Use(cors.New(cors.Config{
    AllowOrigins: "http://127.0.0.1",
    AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	//csrf
	// app.Use(csrf.New())
}