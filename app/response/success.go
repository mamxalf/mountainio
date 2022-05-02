package response

import (
	"github.com/gofiber/fiber/v2"
	"mountainio/domain/model"
)

func Success(data interface{}) model.WebResponse {
	return model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	}
}
