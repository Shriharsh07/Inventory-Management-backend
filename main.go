package main

import (
	"fmt"
	"net/http"

	"github.com/Shriharsh07/InventoryManagement/config"
	"github.com/Shriharsh07/InventoryManagement/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("Listening on port: 8080")
	http.ListenAndServe(":8080", r)
}
