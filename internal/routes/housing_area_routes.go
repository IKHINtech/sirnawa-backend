package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func HousingAreaRoutes(route fiber.Router) {
	repository := repository.NewHousingAreaRepository(database.DB)
	services := services.NewHousingAreaServices(repository, database.DB)
	handlers := handlers.NewHousingAreaHandler(services)

	route.Get("/", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Delete("/:id", handlers.Delete)
}
