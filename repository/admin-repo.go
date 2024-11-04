package repository

import (
	"database/sql"
	"day-21/model"
)

type AdminRepositoryDB struct {
	DB *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepositoryDB {
	return AdminRepositoryDB{DB: db}
}

func (r *AdminRepositoryDB) GetAdminLogin(admin model.Admin) (*model.Admin, error) {
	query := `SELECT id, name, username, password FROM admin WHERE username=$1 AND password=$2`

	var adminResponse model.Admin

	err := r.DB.QueryRow(query, admin.Username, admin.Password).Scan(&adminResponse.ID, &adminResponse.Name, &adminResponse.Username, &adminResponse.Password)

	if err != nil {
		return nil, err
	}

	return &adminResponse, nil
}
