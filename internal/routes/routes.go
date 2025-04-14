package routes

import (
	_ "github.com/IKHINtech/sirnawa-backend/docs"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutesApp(app *fiber.App) {
	BlockRoutes(app.Group("/block", middleware.Protected()))
	AuthRoutes(app.Group("/auth"))

	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SIRNAWA BACKEND 🚀")
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// 404 Route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	})
}
