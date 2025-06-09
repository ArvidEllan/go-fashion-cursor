package config

import (
	"fmt"
	"log"
	"os"

	"virtual-tryon/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitializeDatabase sets up the database connection and migrations
func InitializeDatabase() {
	// Get database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Create database connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Set global database instance
	DB = db

	// Run migrations
	err = DB.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
		&models.Product{},
		&models.Size{},
		&models.TryOn{},
		&models.TryOnHistory{},
	)
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Database connection established and migrations completed")
}
