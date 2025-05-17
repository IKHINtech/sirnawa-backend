package routes

import (
	_ "github.com/IKHINtech/sirnawa-backend/docs"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutesApp(app *fiber.App, driveService utils.DriveService, tokenService *services.FCMTokenService) {
	AuthRoutes(app.Group("/auth"))
	DashboardRoutes(app.Group("/dashboard"), driveService)
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
	RondaGroupRoutes(app.Group("/ronda-group", middleware.Protected()))
	RondaGroupMemberRoutes(app.Group("/ronda-group-member", middleware.Protected()))
	RondaScheduleRoutes(app.Group("/ronda-schedule", middleware.Protected()))
	ShopRoutes(app.Group("/shop", middleware.Protected()))
	ItemRoutes(app.Group("/item", middleware.Protected()))
	IplBillRoutes(app.Group("/ipl-bill", middleware.Protected()))
	IplBillDetailRoutes(app.Group("/ipl-bill-detail", middleware.Protected()))
	IplRateRoutes(app.Group("/ipl-rate", middleware.Protected()))
	IplRateDetailRoutes(app.Group("/ipl-rate-detail", middleware.Protected()))
	ShopProductRoutes(app.Group("/shop-product", middleware.Protected()))
	HousingAreaRoutes(app.Group("/housing-area", middleware.Protected()))
	AnnouncementRoutes(app.Group("/announcement", middleware.Protected()), driveService)
	ResidentHouseRoutes(app.Group("/resident-house", middleware.Protected()))
	UserFcmTokenRoutes(app.Group("/fcm", middleware.Protected()), tokenService)
	NotificationRoutes(app.Group("/notification", middleware.Protected()))

	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(" My RT 🚀")
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// 404 Route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	})
}
