package main

import (
	"log"
	"net/http"

	"SecureFileshare/service/backend/controllers"
	"SecureFileshare/service/backend/routes"
)

func main() {

	controller := controllers.NewControllers()

	routes.RegisterRoutes(controller)

	// Start the server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
