package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"mountainio/app/exception"
	"mountainio/app/response"
	"mountainio/domain/model"
	_authService "mountainio/src/auth/service"
	_userService "mountainio/src/user/service"
)

type AuthController struct {
	UserService _userService.UserService
	AuthService _authService.AuthService
}

func NewAuthController(userService *_userService.UserService, authService *_authService.AuthService) AuthController {
	return AuthController{*userService, *authService}
}

func (controller *AuthController) Route(app fiber.Router) {
	router := app.Group("/auths")
	router.Post("/login", controller.Login)
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	var request model.LoginInput
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	user, err := controller.UserService.FindUserByEmail(request.Email)
	if err != nil {
		return c.JSON(response.ErrorUnprocessableEntity(err))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return c.JSON(response.ErrorBadRequest(err))
	}

	res, err := controller.AuthService.GenerateTokenAuth(user)
	if err != nil {
		return c.JSON(response.ErrorBadRequest(err))
	}
	return c.JSON(response.Success(res))
}
