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

func GetDashboardInventoryByUserID(userId uuid.UUID) ([]models.DashboardInventoryDetails, error) {
	inventories := []models.DashboardInventoryDetails{}
	// Otherwise, filter by userId
	query := `SELECT inv.serial_number, inv.end_of_warranty, inv.last_physical_inventory, inv.model, inv.status, inv.location
			FROM inventories inv
			WHERE inv.user_id = ?;`
	result := config.DB.Raw(query, userId).Scan(&inventories)
	return inventories, result.Error
}
