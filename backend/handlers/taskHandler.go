package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JahnMal/task_mgr_app/backend/models"
	"github.com/JahnMal/task_mgr_app/backend/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.Tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = uuid.New().String()
	utils.Tasks = append(utils.Tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for index, task := range utils.Tasks {
		if task.ID == id {
			utils.Tasks = append(utils.Tasks[:index], utils.Tasks[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
