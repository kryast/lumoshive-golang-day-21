package service

import (
	"day-21/model"
	"day-21/repository"
	"errors"
	"fmt"
)

type UserService struct {
	RepoUser repository.UserRepositoryDB
}

func NewUserService(repo repository.UserRepositoryDB) *UserService {
	return &UserService{RepoUser: repo}
}

func (us *UserService) InputDataUser(name string, email string, password string) error {
	if name == "" {
		return errors.New("username tidak boleh kosong")
	}
	if email == "" {
		return errors.New("email tidak boleh kosong")
	}
	if password == "" {
		return errors.New("password tidak boleh kosong")
	}

	user := model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := us.RepoUser.CreateDataUser(user)
	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println("berhasil input data user dengan id ", user.ID)

	return nil
}

func (us *UserService) UserLoginService(user model.User) (*model.Response, error) {

	users, err := us.RepoUser.GetUserLogin(user)

	if err != nil {
		return nil, err
	}
	response := model.Response{
		StatusCode: 200,
		Message:    "login success",
		Data:       users,
	}
	return &response, nil
}
