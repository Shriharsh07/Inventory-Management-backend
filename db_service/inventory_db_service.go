package dbservice

import (
	"github.com/Shriharsh07/InventoryManagement/config"
	"github.com/Shriharsh07/InventoryManagement/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CheckInventoryByUserIDAndSerialNumber(userId uuid.UUID, serialNumber string) *gorm.DB {
	var inventory models.Inventory
	return config.DB.Where("user_id = ? AND serial_number = ?", userId, serialNumber).First(&inventory)
}

func GetAllInventoryByUserID(userId uuid.UUID) ([]models.Inventory, error) {
	var inventories []models.Inventory

	// If userId is the zero UUID, return all inventory
	if userId == uuid.Nil {
		result := config.DB.Find(&inventories)
		return inventories, result.Error
	}

	// Otherwise, filter by userId
	result := config.DB.Where("user_id = ?", userId).Find(&inventories)
	return inventories, result.Error
}
