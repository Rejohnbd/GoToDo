package cmd

import (
	"fmt"
	"log"
	"todo/config"
	"todo/internal/app"
	"todo/migrations"
)

func StartServer() {
	// Load configuration
	config.LoadConfig()

	// Connect to the database
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Database connection failed: %s", err)
		return
	}

	// Run database migrations
	migrations.RunMigrations()

	// Initialize the application
	app := app.NewFiberApp()

	// Start the server
	fmt.Println("Server is starting on port:", config.AppConfig.Port)
	fmt.Println(app.Listen(":" + config.AppConfig.Port))
}
