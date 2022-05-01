package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"mountainio/app/config"
	"mountainio/app/exception"
	"mountainio/app/router"
)

func main() {
	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig()))
	app.Use(logger.New(config.LoggerConfig()))

	// Setup Routing API
	api := app.Group("/api")

	// Setup V1
	v1 := api.Group("/v1")
	router.SetupRoutesV1(v1)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
