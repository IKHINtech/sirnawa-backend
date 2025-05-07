package routes

import (
	_ "github.com/IKHINtech/sirnawa-backend/docs"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutesApp(app *fiber.App, driveService utils.DriveService) {
	AuthRoutes(app.Group("/auth"))
	DashboardRoutes(app.Group("/dashboard"))
	BlockRoutes(app.Group("/block", middleware.Protected()))
	RtRoutes(app.Group("/rt", middleware.Protected()))
	RwRoutes(app.Group("/rw", middleware.Protected()))
	HouseRoutes(app.Group("/house", middleware.Protected()))
	IplPaymentRoutes(app.Group("/ipl-payment", middleware.Protected()))
	PostCommentRoutes(app.Group("/post-comment", middleware.Protected()))
	PostRoutes(app.Group("/post", middleware.Protected()))
	ResidentRoutes(app.Group("/resident", middleware.Protected()))
	RondaActivityRoutes(app.Group("/ronda-activity", middleware.Protected()))
	RondaAttendanceRoutes(app.Group("/ronda-attendance", middleware.Protected()))
	RondaConstributionRoutes(app.Group("/ronda-constribution", middleware.Protected()))
	RondaGroupRoutes(app.Group("/ronda-constribution", middleware.Protected()))
	RondaScheduleRoutes(app.Group("/ronda-schedule", middleware.Protected()))
	ShopRoutes(app.Group("/shop", middleware.Protected()))
	ShopProductRoutes(app.Group("/shop-product", middleware.Protected()))
	HousingAreaRoutes(app.Group("/housing-area", middleware.Protected()))
	IplRoutes(app.Group("/ipl", middleware.Protected()))
	AnnouncementRoutes(app.Group("/announcement", middleware.Protected()), driveService)
	ResidentHouseRoutes(app.Group("/resident-house", middleware.Protected()))

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
