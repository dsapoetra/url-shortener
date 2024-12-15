package main

import (
	"log"

	"Backend/config"
	"Backend/db"
	"Backend/handlers"
	"Backend/repositories"
	"Backend/routes"
	"Backend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	database, err := db.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	urlRepo := repositories.NewUrlRepository(database) // Add this

	// Initialize services
	urlService := services.NewUrlService(urlRepo) // Add this

	// Initialize handlers
	urlHandler := handlers.NewUrlHandler(urlService)
	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app, urlHandler) // Update this

	// Setup Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
