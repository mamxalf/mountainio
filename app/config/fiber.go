package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mountainio/app/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}

func CorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}
}

func LoggerConfig() logger.Config {
	return logger.Config{
		Format:   "[${time}] - ${method} ${path} ~ [${status}][${latency}]\n",
		TimeZone: "Local",
	}
}
