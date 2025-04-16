package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RondaGroupRoutes(route fiber.Router) {
	repository := repository.NewRondaGroupRepository(database.DB)
	services := services.NewRondaGroupServices(repository, database.DB)
	handlers := handlers.NewRondaGroupHandler(services)

	route.Get("/", handlers.FindAll)
	route.Get("/paginated", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Delete("/:id", handlers.Delete)
}
