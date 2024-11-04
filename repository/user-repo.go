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

func (r *UserRepositoryDB) GetUserLogin(user model.User) (*model.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE email=$1 AND password=$2`

	var userResponse model.User

	err := r.DB.QueryRow(query, user.Email, user.Password).Scan(&userResponse.ID, &userResponse.Name, &userResponse.Email, &userResponse.Password)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}
