package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func UserFcmTokenRoutes(route fiber.Router, tokenService *services.FCMTokenService) {
	tokenHandler := handlers.NewFCMTokenHandler(tokenService)
	route.Get("/register", tokenHandler.RegisterToken)
	route.Get("/remove", tokenHandler.RemoveToken)
}
