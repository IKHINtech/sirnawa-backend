package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func NotificationRoutes(router fiber.Router) {
	repository := repository.NewNotificationRepository(database.DB)
	notificationService := services.NewNotificationService(repository)
	handlers := handlers.NewNotificationHandler(notificationService)

	router.Post("/", handlers.CreateNotification)
	router.Get("/", handlers.GetNotifications)
	router.Get("/unread-count", handlers.GetUnreadCount)
	router.Patch("/:id/read", handlers.MarkAsRead)
	router.Patch("/read-all", handlers.MarkAllAsRead)
	router.Delete("/:id", handlers.DeleteNotification)
}
