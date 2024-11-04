package handler

import (
	"day-21/database"
	"day-21/model"
	"day-21/repository"
	"day-21/service"
	"encoding/json"
	"net/http"
)

func UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ID int `json:"id"`
	}

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Create task repository and service
	repo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(repo)

	// Call the service to update the task status
	task, err := taskService.UpdateTaskById(request.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Task status updated successfully",
		Data:       task,
	}
	json.NewEncoder(w).Encode(response)
}
