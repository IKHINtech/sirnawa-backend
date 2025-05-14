package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func IplBillRoutes(route fiber.Router) {
	repo := repository.NewIplBillRepository(database.DB)
	iplRateRepo := repository.NewIplRateRepository(database.DB)
	services := services.NewIplBillServices(repo, iplRateRepo, database.DB)
	handlers := handlers.NewIplBillHandler(services)

	route.Get("/", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Put("/:id", handlers.Update)
	route.Delete("/:id", handlers.Delete)
}
