package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("❌ DATABASE_DSN not set in environment")
	}

	fmt.Println("📡 Connecting using DSN:", dsn)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ gorm.Open error: %v", err)
	}

	// Test connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

	fmt.Println("Database connected successfully!")
	return nil
}
