package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/config"
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
	"github.com/IKHINtech/sirnawa-backend/internal/repository"
	"github.com/IKHINtech/sirnawa-backend/internal/routes"
	"github.com/IKHINtech/sirnawa-backend/internal/services"
	"github.com/IKHINtech/sirnawa-backend/pkg/firebase"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	workers "github.com/IKHINtech/sirnawa-backend/pkg/worker"
	"github.com/gofiber/fiber/v2"
)

// @title My RT
// @version 1.0
// @description This is a API Server for App My RT
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email sarikhin@yahoo.co.id
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes http https
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal load config:", err)
	}

	// connect to database
	database.Connect()

	// migrasi database
	// database.Migrate()

	// Inisialisasi Google Drive Service
	driveService, err := utils.NewDriveService("service-account-my-rt.json")
	if err != nil {
		log.Fatalf("Gagal inisialisasi Drive Service: %v", err)
	}

	// Inisialisasi Firebase Cloud Messaging
	firebase.InitFCM()

	app := fiber.New()

	middleware.SetupCORS(app)
	middleware.SetupRecovery(app)

	tokenRepo := repository.NewFCMTokenRepository(database.DB)
	tokenService := services.NewFCMTokenService(tokenRepo, database.DB)
	routes.SetupRoutesApp(app, driveService, tokenService)
	// Setup worker
	cleanupWorker := workers.NewTokenCleanupWorker(tokenService, 24*time.Hour)
	cleanupWorker.Start()

	// Handle graceful shutdown
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
		<-sigchan

		cleanupWorker.Stop()
		app.Shutdown()
	}()

	if err := app.Listen(":" + config.AppConfig.PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
