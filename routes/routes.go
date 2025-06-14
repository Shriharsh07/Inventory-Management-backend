package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Shriharsh07/InventoryManagement/Auth"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods("GET")

	r.HandleFunc("/signup", Auth.Signup).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", Auth.Login).Methods("POST", "OPTIONS")
}

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Register your API routes
	RegisterRoutes(r)

	return r
}
