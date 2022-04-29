package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	config2 "mountainio/app/config"
	"mountainio/app/exception"
	"mountainio/controller"
	"mountainio/repository"
	"mountainio/service"
)

func main() {
	// Setup Configuration
	configuration := config2.New()
	database := config2.NewMongoDatabase(configuration)

	// Setup Repository
	productRepository := repository.NewProductRepository(database)

	// Setup Service
	productService := service.NewProductService(&productRepository)

	// Setup Controller
	productController := controller.NewProductController(&productService)

	// Setup Fiber
	app := fiber.New(config2.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	productController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
