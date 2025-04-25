package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func ResidentHouseRoutes(route fiber.Router) {
	repository := repository.NewResidentHouseRepository(database.DB)
	services := services.NewResidentHouseServices(repository, database.DB)
	handlers := handlers.NewResidentHouseHandler(services)

	route.Post("/", handlers.AssignResidentToHouse)
	route.Get("/:id", handlers.ChangeToPrimary)
	route.Delete("/:id", handlers.Delete)
}
