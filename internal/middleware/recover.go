package middleware

import (
	"fmt"

	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRecovery(app *fiber.App) {
	app.Use(recover.New(
		recover.Config{
			EnableStackTrace:  true,
			StackTraceHandler: defaultStackTraceHandler,
		},
	))
}

func defaultStackTraceHandler(c *fiber.Ctx, e any) {
	h := &utils.ResponseHandler{}
	h.InternalServerError(c, []string{fmt.Sprintf("panic: %v", e)})
}
