package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func DashboardRoutes(route fiber.Router, driveService utils.DriveService) {
	route.Get("/mobile", middleware.Protected(),
		func(c *fiber.Ctx) error {
			return handlers.DashboardMobile(c, driveService)
		})
}
