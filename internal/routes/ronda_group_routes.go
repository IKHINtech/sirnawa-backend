package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RondaGroupRoutes(route fiber.Router) {
	repo := repository.NewRondaGroupRepository(database.DB)
	memberRepo := repository.NewRondaGroupMemberRepository(database.DB)
	services := services.NewRondaGroupServices(repo, memberRepo, database.DB)
	handlers := handlers.NewRondaGroupHandler(services)

	route.Get("/", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Put("/:id", handlers.Update)
	route.Delete("/:id", handlers.Delete)
}
