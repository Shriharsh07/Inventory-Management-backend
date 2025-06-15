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
