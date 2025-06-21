package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Shriharsh07/InventoryManagement/config"
	dbservice "github.com/Shriharsh07/InventoryManagement/db_service"
	"github.com/Shriharsh07/InventoryManagement/models"
	"github.com/google/uuid"
)

func AddInventory(w http.ResponseWriter, r *http.Request) {
	var inventory models.Inventory
	if err := json.NewDecoder(r.Body).Decode(&inventory); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request"})
		return
	}

	data := dbservice.CheckInventoryByUserIDAndSerialNumber(inventory.UserID, inventory.SerialNumber)
	if data.RowsAffected != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Inventory with this serial number already exists"})
		return
	}

	data = config.DB.Create(&inventory)
	if data.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Something went wrong, Please try again later"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Inventory added successfully!"})
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
	userIDParam := r.URL.Query().Get("userId")

	var userID uuid.UUID
	var err error

	if userIDParam != "" {
		userID, err = uuid.Parse(userIDParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid userId"})
			return
		}
	}

	dashboardInventories, err := dbservice.GetDashboardInventoryByUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch inventory"})
		return
	}

	inStock := 0
	allDevices := 0
	active := 0
	outOfWarranty := 0

	for _, dashboardInventory := range dashboardInventories {
		switch dashboardInventory.Status {
		case "InStock":
			inStock++
			allDevices++
		case "Active":
			active++
			allDevices++
		case "OutOfWarranty":
			outOfWarranty++
			allDevices++
		}
	}

	var dashboard models.DashboardResponse
	dashboard.Inventory = dashboardInventories
	dashboard.Statistics.AllDevices = allDevices
	dashboard.Statistics.InStock = inStock
	dashboard.Statistics.Active = active
	dashboard.Statistics.OutOfWarranty = outOfWarranty

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dashboard)
}
