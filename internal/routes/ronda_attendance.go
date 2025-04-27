package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RondaAttendanceRoutes(route fiber.Router) {
	repository := repository.NewRondaAttendanceRepository(database.DB)
	services := services.NewRondaAttendanceServices(repository, database.DB)
	handlers := handlers.NewRondaAttendanceHandler(services)

	route.Get("/", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Post("/:id", handlers.Update)
	route.Delete("/:id", handlers.Delete)
}
