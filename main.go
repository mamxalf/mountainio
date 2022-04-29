package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"mountainio/app/config"
	"mountainio/app/exception"
	"mountainio/controller"
	"mountainio/repository"
	"mountainio/service"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	database := config.ConnectPostgres(configuration)

	// Setup Repository
	//productRepository := repository.NewProductRepository(database)
	userRepository := repository.NewUserRepository(database)

	// Setup Service
	//productService := service.NewProductService(&productRepository)
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	//productController := controller.NewProductController(&productService)
	userController := controller.NewUserController(&userService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	//productController.Route(app)
	userController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
