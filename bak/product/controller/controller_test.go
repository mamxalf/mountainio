package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"mountainio/app/config"
	"mountainio/bak/product/repository"
	"mountainio/bak/product/service"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	productController.Route(app)
	return app
}

var configuration = config.New("../.env.test")

var database = config.NewMongoDatabase(configuration)
var productRepository = repository.NewProductRepository(database)
var productService = service.NewProductService(&productRepository)

var productController = NewProductController(&productService)

var app = createTestApp()
