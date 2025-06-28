package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Shriharsh07/InventoryManagement/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.UserList{})

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
