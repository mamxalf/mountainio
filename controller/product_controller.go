package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"mountainio/app/exception"
	model2 "mountainio/domain/model"
	"mountainio/service"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{ProductService: *productService}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Post("/api/products", controller.Create)
	app.Get("/api/products", controller.List)
}

func (controller *ProductController) Create(c *fiber.Ctx) error {
	var request model2.CreateProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.ProductService.Create(request)
	return c.JSON(model2.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) List(c *fiber.Ctx) error {
	responses := controller.ProductService.List()
	return c.JSON(model2.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}
