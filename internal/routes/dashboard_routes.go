package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func DashboardRoutes(route fiber.Router) {
	route.Get("/mobile", middleware.Protected(), handlers.DashboardMobile)
}
