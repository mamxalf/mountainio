package response

import (
	"github.com/gofiber/fiber/v2"
	"mountainio/domain/model"
)

func ErrorUnprocessableEntity(err error) model.WebResponse {
	return model.WebResponse{
		Code:   fiber.StatusUnprocessableEntity,
		Status: "ERROR",
		Data:   err.Error(),
	}
}

func ErrorBadRequest(err error) model.WebResponse {
	return model.WebResponse{
		Code:   fiber.StatusBadRequest,
		Status: "ERROR",
		Data:   err.Error(),
	}
}
