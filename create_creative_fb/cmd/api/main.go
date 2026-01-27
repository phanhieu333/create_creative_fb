package main

import (
	"fmt"
	"log"

	"creative_fb/internal/config"
	"creative_fb/internal/facebook"
	"creative_fb/internal/repositories"
	"creative_fb/internal/routes"
	"creative_fb/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Load config
	cfg := config.LoadConfig()

	// Get default account access token from config
	var accessToken string
	if len(cfg.Facebook.Accounts) > 0 {
		// Get first account
		for _, acc := range cfg.Facebook.Accounts {
			accessToken = acc.AccessToken
			break
		}
	}

	if accessToken == "" {
		log.Fatal("No Facebook access token found in config")
	}

	// Initialize layers
	client := facebook.NewClient(accessToken)
	repo := repositories.NewCreativeRepository(client)
	service := services.NewCreativeService(repo)

	// Initialize Fiber router
	app := fiber.New(fiber.Config{
		AppName: "Creative Facebook API",
	})

	// Register routes
	creativeRoutes := routes.NewCreativeRoutes(service)
	creativeRoutes.RegisterRoutes(app)

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Starting server on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
