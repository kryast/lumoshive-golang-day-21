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

	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(repo)

	tasks, err := taskService.UpdateTaskById(task.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Task status updated successfully",
		Data:       tasks,
	}
	json.NewEncoder(w).Encode(response)
}
