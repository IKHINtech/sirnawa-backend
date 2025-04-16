package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RondaConstributionRoutes(route fiber.Router) {
	repository := repository.NewRondaConstributionRepository(database.DB)
	services := services.NewRondaConstributionServices(repository, database.DB)
	handlers := handlers.NewRondaConstributionHandler(services)

	route.Get("/", handlers.FindAll)
	route.Get("/paginated", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Delete("/:id", handlers.Delete)
}
