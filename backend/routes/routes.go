package routes

import (
	"net/http"

	"github.com/JahnMal/task_mgr_app/backend/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRoutes() http.Handler {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/api/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// CORS Middleware
	handler := cors.Default().Handler(r)
	return handler
}
