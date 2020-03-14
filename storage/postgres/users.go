package postgres

import (
	"github.com/Jopoleon/rustamViewer/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (db *DB) CreateUser(user models.User) error {

	bpas, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	_, err = db.DB.Exec("INSERT INTO users (login, email, password, profile_name, is_admin, first_name, second_name, company_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);",
		user.Login, user.Email, string(bpas), "", user.IsAdmin, user.FirstName,
		user.SecondName, user.CompanyName)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) GetUserByID(id int) (*models.User, error) {
	res := models.User{}
	err := db.DB.Get(&res, "SELECT * FROM users WHERE id=$1 LIMIT 1;", id)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	var apps []string
	err = db.DB.Select(&apps, "SELECT profile_name FROM users_apps "+
		"WHERE user_id=$1;", res.ID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	res.AppNames = apps
	return &res, nil
}

func (db *DB) GetUserByEmailOrLogin(email string) (*models.User, error) {
	res := models.User{}
	err := db.DB.Get(&res, "SELECT * FROM users WHERE email=$1 OR login=$1;", email)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}

	var apps []string
	err = db.DB.Select(&apps, "SELECT profile_name FROM users_apps "+
		"WHERE user_id=$1;", res.ID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	res.AppNames = apps
	return &res, nil
}
