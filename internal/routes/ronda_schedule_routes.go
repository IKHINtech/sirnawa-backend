package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RondaScheduleRoutes(route fiber.Router) {
	repository := repository.NewRondaScheduleRepository(database.DB)
	services := services.NewRondaScheduleServices(repository, database.DB)
	handlers := handlers.NewRondaScheduleHandler(services)

	route.Get("/", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Delete("/:id", handlers.Delete)
}
