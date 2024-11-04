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

func (ts *TaskService) GetAllDataTask() ([]model.Task, error) {

	tasks, err := ts.RepoTask.GetAll()

	if err != nil {
		return nil, err
	}

	if tasks == nil || len(*tasks) == 0 {
		fmt.Println("No tasks found")
		return nil, err
	}

	return *tasks, nil

}

func (ts *TaskService) UpdateTaskById(id int) (*model.Task, error) {

	task, err := ts.RepoTask.UpdateStatusTask(id)

	if err != nil {
		return nil, err

	}
	return task, nil
}
