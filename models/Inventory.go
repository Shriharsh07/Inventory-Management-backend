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
