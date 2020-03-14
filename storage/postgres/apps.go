package postgres

import (
	"github.com/Jopoleon/rustamViewer/models"
	"github.com/pkg/errors"
)

func (db *DB) CreateNewApp(user_id, companyID int, profileName string) error {
	_, err := db.DB.Exec("INSERT INTO users_apps (user_id,profile_name, company_id) VALUES ($1,$2, $3);",
		user_id, profileName, companyID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) DeleteUsersApp(user_id int, profileName string) error {
	_, err := db.DB.Exec("DELETE FROM users_apps WHERE user_id=$1 AND profile_name=$2;",
		user_id, profileName)

	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) GetUserApps(user_id int) ([]models.Application, error) {
	var res []models.Application
	err := db.DB.Select(&res, "SELECT * FROM users_apps WHERE user_id = $1;", user_id)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}
