package cmd

import (
	"log"
	"todo/config"
	"todo/internal/app"
)

func StartServer() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the application
	app := app.NewFiberApp(cfg)

	// Start the server
	log.Printf("Server is starting on port %s...\n", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
