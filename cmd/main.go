package main

import (
	"SecureFileshare/service/backend/controllers"
	"SecureFileshare/service/backend/routes"
	"log"
	"net/http"
)

func main() {

	controller := controllers.NewControllers()

	routes.RegisterRoutes(controller)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
