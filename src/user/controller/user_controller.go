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
	response, _ := controller.UserService.FindUserByID(id)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Me(c *fiber.Ctx) error {
	claimToken := middleware.GetClaimToken(c)
	return c.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   claimToken,
	})
}

func (controller *UserController) Index(c *fiber.Ctx) error {
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Users Index PATH!",
	})
}
