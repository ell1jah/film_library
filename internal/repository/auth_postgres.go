package repository

import (
	"github.com/ell1jah/film_library/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateAdmin(admin *models.Admin) (int, error) {
	insertQuery := `INSERT INTO admins (adminname, password) VALUES ($1, $2) RETURNING admin_id;`
	var id int
	res := a.db.QueryRow(insertQuery, admin.Adminname, admin.Password)
	err := res.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetAdmin(adminname, password string) (*models.Admin, error) {
	var admin models.Admin
	selectQuery := `SELECT adminname, password FROM admins WHERE adminname=$1 AND password=$2`
	res := a.db.QueryRow(selectQuery, adminname, password)
	err := res.Scan(&admin.Adminname, &admin.Password)
	if err != nil {
		return &models.Admin{}, err
	}
	return &admin, nil
}
