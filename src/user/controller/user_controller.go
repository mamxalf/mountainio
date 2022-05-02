package controller

import (
	"github.com/gofiber/fiber/v2"
	"mountainio/app/exception"
	"mountainio/app/middleware"
	"mountainio/app/response"
	"mountainio/domain/model"
	"mountainio/src/user/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{*userService}
}

func (controller *UserController) Route(app fiber.Router) {
	router := app.Group("/users")
	router.Post("/", controller.Register)
	router.Get("/me", middleware.AuthProtected(), controller.Me)
	router.Get("/:id", controller.FindByID)
	router.Get("/", controller.Index)
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var request model.RegisterUser
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	res, err := controller.UserService.RegisterUser(request)
	if err != nil {
		return c.JSON(response.ErrorUnprocessableEntity(err))
	}
	return c.JSON(response.Success(res))
}

func (controller *UserController) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := controller.UserService.FindUserByID(id)
	if err != nil {
		return c.JSON(response.ErrorUnprocessableEntity(err))
	}
	return c.JSON(response.Success(res))
}

func (controller *UserController) Me(c *fiber.Ctx) error {
	claimToken := middleware.GetClaimToken(c)
	return c.JSON(response.Success(claimToken))
}

func (controller *UserController) Index(c *fiber.Ctx) error {
	return c.JSON(response.Success("Users Index PATH!"))
}
