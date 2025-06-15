package controller

import (
	"encoding/json"
	"net/http"

	dbservice "github.com/Shriharsh07/InventoryManagement/db_service"
	"github.com/gorilla/mux"
)

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	// Retrieve the user by email from your database
	data, user := dbservice.DBServiceGetUserByEmail(email)
	if data.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	}

	json.NewEncoder(w).Encode(user)
}
