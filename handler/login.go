package handler

// import (
// 	"day-21/database"
// 	"day-21/model"
// 	"day-21/repository"
// 	"day-21/service"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// )

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	admin := model.Admin{}
// 	err := json.NewDecoder(r.Body).Decode(&admin)

// 	badResponse := model.Response{
// 		StatusCode: http.StatusBadRequest,
// 		Message:    "Error Server",
// 		Data:       nil,
// 	}
// 	if err != nil {
// 		json.NewEncoder(w).Encode(badResponse)
// 		return
// 	}

// 	db, err := database.InitDB()
// 	if err != nil {
// 		badResponse := model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "Error Server",
// 			Data:       nil,
// 		}
// 		json.NewEncoder(w).Encode(badResponse)
// 		return
// 	}
// 	defer db.Close()

// 	repo := repository.NewAdminRepository(db)
// 	adminService := service.NewAdminService(repo)

// 	adminData, err := adminService.LoginService(admin)
// 	if err != nil {
// 		badResponse := model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "Login failed",
// 			Data:       nil,
// 		}
// 		json.NewEncoder(w).Encode(badResponse)
// 		return
// 	}

// 	Response := model.Response{
// 		StatusCode: http.StatusOK,
// 		Message:    "Success",
// 		Data:       adminData.Data,
// 	}
// 	json.NewEncoder(w).Encode(Response)

// }

// func GetAdminByIDHandler(w http.ResponseWriter, r *http.Request) {

// 	query := r.URL.Query()
// 	idString := query.Get("id")
// 	id, err := strconv.Atoi(idString)

// 	badResponse := model.Response{
// 		StatusCode: http.StatusBadRequest,
// 		Message:    "Error Server",
// 		Data:       nil,
// 	}

// 	if err != nil {
// 		json.NewEncoder(w).Encode(badResponse)
// 		return
// 	}

// 	db, err := database.InitDB()
// 	if err != nil {
// 		badResponse := model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "Error Server",
// 			Data:       nil,
// 		}
// 		json.NewEncoder(w).Encode(badResponse)
// 		return
// 	}
// 	defer db.Close()

// 	repo := repository.NewAdminRepository(db)
// 	adminService := service.NewAdminService(repo)

// 	adminData, err := adminService.GetAdminByID(id)
// 	if err != nil {
// 		badResponse := model.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "Login failed",
// 			Data:       nil,
// 		}
// 		json.NewEncoder(w).Encode(badResponse)
// 		return
// 	}
// 	response := model.Response{
// 		StatusCode: http.StatusOK,
// 		Message:    "Success",
// 		Data:       adminData,
// 	}
// 	json.NewEncoder(w).Encode(response)
// }
