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

func (r *TaskRepositoryDB) GetAll() (*[]model.Task, error) {
	query := `SELECT id, description, status FROM tasks`
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []model.Task{}

	for rows.Next() {
		var task model.Task
		rows.Scan(&task.ID, &task.Description, &task.Status)

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (r *TaskRepositoryDB) UpdateStatusTask(id int) (*model.Task, error) {
	var task model.Task
	query := "UPDATE tasks SET status = 'complete' WHERE id = $1 "
	err := r.DB.QueryRow(query, id).Scan(&task.ID, &task.Description, &task.Status)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
