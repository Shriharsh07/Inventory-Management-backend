package dbservice

import (
	"github.com/Shriharsh07/InventoryManagement/Auth"
	"github.com/Shriharsh07/InventoryManagement/config"
	"gorm.io/gorm"
)

func DBServiceGetUserByEmail(email string) (*gorm.DB, Auth.User) {
	var user Auth.User
	query := `SELECT users.id, users.name, users.email FROM users WHERE users.email = ?`
	data := config.DB.Raw(query, email).Scan(&user)
	return data, user
}
