package cmd

import (
	"log"
	"todo/config"
	"todo/internal/app"
)

func StartServer() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %s", err)
		return
	}

	// Initialize the application
	app := app.NewFiberApp(cfg, db)

	// Start the server
	log.Printf("Server is starting on port %s...\n", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
