package cmd

import (
	"fmt"
	"todo/config"
	"todo/internal/app"
)

func StartServer() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := config.ConnectDB(cfg)
	if err != nil {
		fmt.Println("Database connection failed: ", err)
		return
	}

	// Initialize the application
	app := app.NewFiberApp(cfg, db)

	// Start the server
	fmt.Println("Server is starting on port ", cfg.Port)
	fmt.Println(app.Listen(":" + cfg.Port))
}
