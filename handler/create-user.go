package handler

import (
	"day-21/database"
	"day-21/model"
	"day-21/repository"
	"day-21/service"
	"encoding/json"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

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

	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		badResponse.Message = "Database connection error"
		json.NewEncoder(w).Encode(badResponse)
		return
	}
	defer db.Close()

	// Create user repository and service
	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)

	// Call the service to create a user
	err = userService.InputDataUser(user.Name, user.Email, user.Password)
	if err != nil {
		badResponse.Message = err.Error()
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	// Success response
	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "User created successfully",
		Data:       user,
	}
	json.NewEncoder(w).Encode(response)
}
