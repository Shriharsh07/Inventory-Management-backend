package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Shriharsh07/InventoryManagement/config"
	dbservice "github.com/Shriharsh07/InventoryManagement/db_service"
	"github.com/Shriharsh07/InventoryManagement/models"
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
