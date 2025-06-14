package models

import (
	"time"

	"github.com/google/uuid"
)

type Inventory struct {
	SerialNumber          int       `json:"serialNumber"`
	PurchaseDate          time.Time `json:"purchaseDate"`
	DeliveryDate          time.Time `json:"deliveryDate" gorm:"default:NULL"`
	EndOfWarranty         time.Time `json:"endOfWarranty" gorm:"default:NULL"`
	LastPhysicalInventory time.Time `json:"lastPhysicalInventory" gorm:"default:NULL"`
	OrderNumber           int       `json:"orderNumber"`
	ManufacturerID        int       `json:"manufacturerID"`
	Model                 string    `json:"model"`
	StatusID              int       `json:"statusID"`
	DeviceType            string    `json:"deviceType"`
	Location              string    `json:"location"`
	Description           string    `json:"description"`
	UserID                uuid.UUID `json:"userID"`
}
