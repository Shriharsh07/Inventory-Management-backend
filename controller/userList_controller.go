package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Shriharsh07/InventoryManagement/config"
	dbservice "github.com/Shriharsh07/InventoryManagement/db_service"
	"github.com/Shriharsh07/InventoryManagement/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func AddUsers(w http.ResponseWriter, r *http.Request) {
	var user models.UserList

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request"})
		return
	}

	_, data := dbservice.CheckEmailAndCreatorID(user.Email, user.CreaterId)
	if data.RowsAffected > 0 {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"message": "Email already exists"})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Something went wrong, Please try again later"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User added successfully!"})
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	var users []models.UserList
	var err error

	vars := mux.Vars(r)
	createrIdStr := vars["userId"] // assuming path param like /userList/{userId}

	if createrIdStr != "" {
		createrId, parseErr := uuid.Parse(createrIdStr)
		if parseErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid userId"})
			return
		}
		users, err = dbservice.GetUserData(&createrId)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Failed to fetch users"})
		return
	}

	// Return only fullname and email
	type userResponse struct {
		FullName string `json:"fullname"`
		Email    string `json:"email"`
	}

	var filteredUsers []userResponse
	for _, user := range users {
		filteredUsers = append(filteredUsers, userResponse{
			FullName: user.FullName,
			Email:    user.Email,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredUsers)
}
