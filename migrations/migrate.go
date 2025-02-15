package migrations

import (
	"fmt"
	"log"
	"todo/config"
	"todo/internal/models"
)

// RunMigrations executes database migrations
func RunMigrations() {
	db := config.DB // Get the database connection

	fmt.Println("Running migrations...")

	// AutoMigrate will create or update tables based on model structures
	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migrations completed successfully!")
}
