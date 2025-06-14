// controller/masterdata.go
package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Shriharsh07/InventoryManagement/constant"
	"github.com/Shriharsh07/InventoryManagement/models" // Import your models package
)

func GetMasterData(w http.ResponseWriter, r *http.Request) {

	response := models.MasterData{
		Countries:        constant.Countries,        // Get values from constants
		DeviceTypes:      constant.DeviceTypes,      // Get values from constants
		MachineModels:    constant.MachineModels,    // Get values from constants
		Status:           constant.Status,           // Get values from constants
		ReturnableStatus: constant.ReturnableStatus, // Get values from constants
		ReturnReason:     constant.ReturnReason,     // Get values from constants
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
