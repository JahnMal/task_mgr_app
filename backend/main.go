package main

import (
	"log"
	"net/http"

	"github.com/JahnMal/task_mgr_app/backend/routes"
)

func main() {
	router := routes.SetupRoutes()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
