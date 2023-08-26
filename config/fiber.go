package config

import (
	"github.com/gofiber/fiber/v2"

	"github.com/fanchann/excelReader/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		BodyLimit:    5 * 1024 * 1024, // file upload size must < 5 MB
	}
}
