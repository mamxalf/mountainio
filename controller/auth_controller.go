package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"mountainio/app/exception"
	"mountainio/app/middleware"
	"mountainio/domain/model"
	"mountainio/src/user/service"
	"net/http"
	"time"
)

type AuthController struct {
	UserService service.UserService
}

func NewAuthController(userService *service.UserService) AuthController {
	return AuthController{*userService}
}

func (controller *AuthController) Route(app fiber.Router) {
	router := app.Group("/auths")
	router.Post("/login", controller.Login)
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	var request model.LoginInput
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	user, _ := controller.UserService.FindUserByEmail(request.Email)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		})
	}

	auth := model.AuthClaim{
		UserID:  user.ID,
		Role:    user.Role,
		Expired: time.Now().Add(time.Hour * 72).Unix(),
	}

	token, err := middleware.GenerateToken(auth)
	if err != nil {
		return c.JSON(model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "ERROR",
			Data:   err,
		})
	}

	response := model.LoginSuccess{
		UserID: user.ID,
		Email:  user.Email,
		Token:  token,
	}

	return c.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
