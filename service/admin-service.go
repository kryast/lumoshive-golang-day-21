package service

import (
	"day-21/model"
	"day-21/repository"
)

type AdminService struct {
	RepoAdmin repository.AdminRepositoryDB
}

func NewAdminService(repo repository.AdminRepositoryDB) *AdminService {
	return &AdminService{RepoAdmin: repo}
}

func (cs *AdminService) LoginService(user model.Admin) (*model.Response, error) {

	admins, err := cs.RepoAdmin.GetAdminLogin(user)

	if err != nil {
		return nil, err
	}
	response := model.Response{
		StatusCode: 200,
		Message:    "login success",
		Data:       admins,
	}
	return &response, nil
}
