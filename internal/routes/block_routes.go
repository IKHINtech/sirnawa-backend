package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func BlockRoutes(route fiber.Router) {
	repository := repository.NewBlockRepository(database.DB)
	services := services.NewBlockServices(repository, database.DB)
	handlers := handlers.NewBlockHandler(services)
	route.Get("/", handlers.FindAll)
}
