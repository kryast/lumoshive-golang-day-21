package service

import (
	"day-21/model"
	"day-21/repository"
	"errors"
	"fmt"
)

type TaskService struct {
	RepoTask repository.TaskRepositoryDB
}

func NewTaskService(repo repository.TaskRepositoryDB) *TaskService {
	return &TaskService{RepoTask: repo}
}

func (ts *TaskService) InputDataTask(description string) error {
	if description == "" {
		return errors.New("username tidak boleh kosong")
	}

	task := model.Task{
		Description: description,
	}

	err := ts.RepoTask.CreateDataTask(task)
	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println("berhasil input data user dengan id ", task.ID)

	return nil
}
