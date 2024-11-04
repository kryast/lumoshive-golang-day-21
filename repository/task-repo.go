package repository

import (
	"database/sql"
	"day-21/model"
)

type TaskRepositoryDB struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepositoryDB {
	return TaskRepositoryDB{DB: db}
}

func (r *TaskRepositoryDB) CreateDataTask(task model.Task) error {
	query := "INSERT INTO tasks (description, status) VALUES ($1, 'incomplete') RETURNING id"

	err := r.DB.QueryRow(query, task.Description).Scan(&task.ID, &task.Description, &task.Status)
	if err != nil {
		return err
	}

	return nil
}
