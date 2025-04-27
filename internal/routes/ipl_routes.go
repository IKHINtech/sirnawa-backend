package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func IplRoutes(route fiber.Router) {
	repo := repository.NewIplRepository(database.DB)
	itemRepository := repository.NewIplDetailRepository(database.DB)
	services := services.NewIplServices(repo, itemRepository, database.DB)
	handlers := handlers.NewIplHandler(services)

	route.Get("/", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Put("/:id", handlers.Update)
	route.Delete("/:id", handlers.Delete)
}
