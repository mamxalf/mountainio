package router

import (
	"github.com/gofiber/fiber/v2"
	"mountainio/app/config"
	"mountainio/controller"
	"mountainio/src/user/repository"
	"mountainio/src/user/service"
)

func SetupRoutesV1(app fiber.Router) {
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
	authController := controller.NewAuthController(&userService)

	//productController.Route(app)
	userController.Route(app)
	authController.Route(app)
}
