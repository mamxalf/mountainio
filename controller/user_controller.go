package controller

import (
	"github.com/gofiber/fiber/v2"
	"mountainio/app/exception"
	"mountainio/domain/model"
	"mountainio/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{*userService}
}

func (controller *UserController) Route(app *fiber.App) {
	v1 := app.Group("/v1/users")
	v1.Get("/")
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var request model.RegisterUser
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, _ := controller.UserService.RegisterUser(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
