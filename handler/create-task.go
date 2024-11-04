package handler

import (
	"day-21/database"
	"day-21/model"
	"day-21/repository"
	"day-21/service"
	"encoding/json"
	"net/http"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	badResponse := model.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "Error Server",
		Data:       nil,
	}
	if err != nil {
		badResponse.Message = "Invalid input"
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		badResponse.Message = "Database connection error"
		json.NewEncoder(w).Encode(badResponse)
		return
	}
	defer db.Close()

	repo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(repo)

	err = taskService.InputDataTask(task.Description)
	if err != nil {
		badResponse.Message = err.Error()
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "User created successfully",
		Data:       task,
	}
	json.NewEncoder(w).Encode(response)
}
