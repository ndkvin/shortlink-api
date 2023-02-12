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
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func Register(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,X-CSRF-Token,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	// recover
	app.Use(csrf.New(csrf.Config{
		CookieSecure: true,
	}))
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
		Output:     file,
		Format:     "[${time}] ${ip}:${port} ${status} ${latency} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 - 15:04:05",
	}))

	// monitoring
	app.Get("/monitor", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// cors


	//csrf

}
