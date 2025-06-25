package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Shriharsh07/InventoryManagement/Auth"
	"github.com/Shriharsh07/InventoryManagement/controller"
	"github.com/Shriharsh07/InventoryManagement/middleware"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods("GET")

	// Auth Routes
	r.HandleFunc("/signup", Auth.Signup).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", Auth.Login).Methods("POST", "OPTIONS")
	r.Handle("/user/{email}", middleware.JWTAuth(http.HandlerFunc(controller.GetUserByEmail))).Methods("GET", "OPTIONS")

	//Inventory
	r.Handle("/inventory", middleware.JWTAuth(http.HandlerFunc(controller.AddInventory))).Methods("POST", "OPTIONS")

	r.Handle("/dashboardInventory", middleware.JWTAuth(http.HandlerFunc(controller.GetInventory))).Methods("Get")
	// MasterData Route
	r.Handle("/masterData", middleware.JWTAuth(http.HandlerFunc(controller.GetMasterData))).Methods("GET")
	r.Handle("/addUser", middleware.JWTAuth(http.HandlerFunc(controller.AddUsers))).Methods("POST")
	r.Handle("/userList/{userId}", middleware.JWTAuth(http.HandlerFunc(controller.GetUserList))).Methods("GET", "OPTIONS")

}

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Register your API routes
	RegisterRoutes(r)

	return r
}
