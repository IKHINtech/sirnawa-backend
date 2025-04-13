package main

import (
	"log"

	"github.com/IKHINtech/sirnawa-backend/internal/config"
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/routes"
	"github.com/gofiber/fiber"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config:", err)
	}

	// connect to database
	database.Connect()

	// migrasi database
	database.Migrate()

	// start server

	app := fiber.New()

	middleware.SetupCORS(app)

	routes.SetupRoutesApp(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
