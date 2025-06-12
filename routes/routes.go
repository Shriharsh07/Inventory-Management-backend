package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Add your API routes here
	r.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello, World!")
	}).Methods("GET", "OPTIONS")
}

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Register your API routes
	RegisterRoutes(r)

	return r
}
