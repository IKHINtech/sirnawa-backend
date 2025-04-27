package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func ResidentRoutes(route fiber.Router) {
	repo := repository.NewResidentRepository(database.DB)
	userRepository := repository.NewUserRepository(database.DB)
	services := services.NewResidentServices(repo, userRepository, database.DB)
	handlers := handlers.NewResidentHandler(services)

	route.Get("/", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Put("/:id", handlers.Update)
	route.Delete("/:id", handlers.Delete)
}
