package models

import (
	"time"

	"github.com/google/uuid"
)

type Inventory struct {
	SerialNumber          string    `json:"serialNumber"`
	PurchaseDate          time.Time `json:"purchaseDate"`
	DeliveryDate          time.Time `json:"deliveryDate" gorm:"default:NULL"`
	EndOfWarranty         time.Time `json:"endOfWarranty" gorm:"default:NULL"`
	LastPhysicalInventory time.Time `json:"lastPhysicalInventory" gorm:"default:NULL"`
	OrderNumber           string    `json:"orderNumber"`
	ManufacturerID        string    `json:"manufacturerID"`
	Model                 string    `json:"model"`
	Status                string    `json:"status"`
	DeviceType            string    `json:"deviceType"`
	Location              string    `json:"location"`
	Description           string    `json:"description"`
	UserID                uuid.UUID `json:"userID"`
}

type StatisticsData struct {
	AllDevices    int `json:"allDevices"`
	InStock       int `json:"inStock"`
	Active        int `json:"active"`
	OutOfWarranty int `json:"outOfWarranty"`
}

type DashboardInventoryDetails struct {
	SerialNumber          string    `json:"serialNumber"`
	WarrantyTill          time.Time `json:"warrantyTill"`
	LastPhysicalInventory time.Time `json:"lastSyncDate"`
	Model                 string    `json:"model"`
	Status                string    `json:"status"`
	TypeAssignedTo        string    `json:"typeAssignedTo"`
	LastLoggedIn          *string   `json:"lastLoggedIn"`   // optional
	LastAssignedTo        *string   `json:"lastAssignedTo"` // optional
	Location              string    `json:"location"`
}
