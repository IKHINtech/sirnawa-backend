package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(route fiber.Router) {
	route.Post("/login", handlers.Login)
	route.Delete("/logout", handlers.Logout)
	route.Get("/me", middleware.Protected(), handlers.Me)
	route.Post("/register", handlers.Register)
	route.Get("/refresh-token", handlers.RefreshToken)
	route.Post("/verify-email-code", handlers.VerifyEmailCode)
	route.Post("/send-email-verification", handlers.SendEmailVerification)
}
