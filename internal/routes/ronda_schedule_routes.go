package routes

import (
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/handlers"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func RondaScheduleRoutes(route fiber.Router) {
	repo := repository.NewRondaScheduleRepository(database.DB)
	rondaGroupMemberRepo := repository.NewRondaGroupMemberRepository(database.DB)
	services := services.NewRondaScheduleServices(repo, rondaGroupMemberRepo, database.DB)
	handlers := handlers.NewRondaScheduleHandler(services)

	route.Get("/", handlers.Paginated)
	route.Get("/:id", handlers.FindByID)
	route.Post("/", handlers.Create)
	route.Put("/:id", handlers.Update)
	route.Delete("/:id", handlers.Delete)
}
