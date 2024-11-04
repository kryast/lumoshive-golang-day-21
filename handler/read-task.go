package handler

import (
	"day-21/database"
	"day-21/model"
	"day-21/repository"
	"day-21/service"
	"encoding/json"
	"net/http"
)

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize the database connection
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Create the repository and service
	repo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(repo)

	// Call the service to get all tasks
	tasks, err := taskService.GetAllDataTask()
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Tasks retrieved successfully",
		Data:       tasks,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
