package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func ShopProductRoutes(route fiber.Router) {
	repository := repository.NewShopProductRepository(database.DB)
	services := services.NewShopProductServices(repository, database.DB)
	handlers := handlers.NewShopProductHandler(services)

	route.Get("/", handlers.FindAll)
	route.Get("/paginated", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Delete("/:id", handlers.Delete)
}
