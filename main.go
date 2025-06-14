package main

import (
	"fmt"
	"net/http"

	"github.com/Shriharsh07/InventoryManagement/config"
	"github.com/Shriharsh07/InventoryManagement/routes"
	"github.com/rs/cors" // Make sure this is imported
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRouter()

	// Add Swagger route before wrapping
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS
	handler := c.Handler(r)

	fmt.Println("Listening on port: 8080")
	http.ListenAndServe(":8080", handler)
}
