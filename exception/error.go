package exception

import (
	"github.com/gofiber/fiber/v2"

	"github.com/fanchann/excelReader/app/domain/types"
)

func PanicIfNeeded(err any) {
	if err != nil {
		panic(err)
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.JSON(types.WebResponse{
		Status:  400,
		Message: err.Error(),
		Data:    nil,
	})
}
