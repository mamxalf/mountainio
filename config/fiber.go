package config

import (
	"github.com/gofiber/fiber/v2"
	"mountainio/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
