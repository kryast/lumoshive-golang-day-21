package service

// import (
// 	"day-21/model"
// 	"day-21/repository"
// )

// type AdminService struct {
// 	RepoAdmin repository.AdminRepositoryDB
// }

// func NewAdminService(repo repository.AdminRepositoryDB) *AdminService {
// 	return &AdminService{RepoAdmin: repo}
// }

// func (as *AdminService) LoginService(user model.Admin) (*model.Response, error) {

// 	admins, err := as.RepoAdmin.GetAdminLogin(user)

// 	if err != nil {
// 		return nil, err
// 	}
// 	response := model.Response{
// 		StatusCode: 200,
// 		Message:    "login success",
// 		Data:       admins,
// 	}
// 	return &response, nil
// }

// func (as *AdminService) GetAdminByID(id int) (*model.Admin, error) {

// 	admin, err := as.RepoAdmin.GetID(id)

// 	if err != nil {
// 		return nil, err

// 	}
// 	return admin, nil
// }
