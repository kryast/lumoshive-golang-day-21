package repository

import (
	"database/sql"
	"day-21/model"
)

type UserRepositoryDB struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryDB {
	return UserRepositoryDB{DB: db}
}

func (r *UserRepositoryDB) CreateDataUser(user model.User) error {
	query := "INSERT INTO users (name , email, password) VALUES ($1 , $2 , $3) RETURNING id"

	err := r.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return err
	}

	return nil
}
