package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func ShopRoutes(route fiber.Router) {
	repository := repository.NewShopRepository(database.DB)
	services := services.NewShopServices(repository, database.DB)
	handlers := handlers.NewShopHandler(services)

	route.Get("/", handlers.FindAll)
	route.Get("/paginated", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Delete("/:id", handlers.Delete)
}
