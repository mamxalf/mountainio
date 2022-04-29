package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"mountainio/app/config"
	"mountainio/app/exception"
	"mountainio/repository"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	database := config.ConnectPostgres(configuration)

	// Setup Repository
	_ = repository.NewUserRepository(database)

	// Setup Service
	//productService := service.NewProductService(&productRepository)

	// Setup Controller
	//productController := controller.NewProductController(&productService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	//productController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
