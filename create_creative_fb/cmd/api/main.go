package main

import (
	"fmt"
	"log"

	"creative_fb/internal/config"
	"creative_fb/internal/facebook"
	"creative_fb/internal/handlers"
	"creative_fb/internal/repositories"
	"creative_fb/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Initialize config
	cfg := config.LoadConfig()
	if cfg.FacebookAccessToken == "" {
		log.Fatal("FACEBOOK_ACCESS_TOKEN environment variable is required")
	}

	// Initialize layers
	handler := handlers.NewCreativeHandler()
	client := facebook.NewClient(cfg.FacebookAccessToken)
	repo := repositories.NewCreativeRepository(client)
	service := services.NewCreativeService(repo)

	// Parse command line flags
	input := handler.ParseFlags()

	result, err := service.CreateCreative(input)
	if err != nil {
		log.Fatalf("Failed to create creative: %v", err)
	}

	fmt.Printf("✓ Creative created successfully!\n")
	fmt.Printf("  Creative ID: %s\n", result.ID)
}
