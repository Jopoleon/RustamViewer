package postgres

import (
	"github.com/Jopoleon/rustamViewer/models"
	"github.com/pkg/errors"
)

func (db *DB) CreateNewCompany(companyName string) error {
	_, err := db.DB.Exec("INSERT INTO companies (name) VALUES ($1);",
		companyName)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

//admin method
func (db *DB) GetCompanies() ([]models.Company, error) {
	var res []models.Company
	err := db.DB.Select(&res, "SELECT * FROM companies")
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	for i, c := range res {
		apps, err := db.GetCompanyApps(c.ID)
		if err != nil {
			db.Logger.Error(errors.WithStack(err))
			return nil, errors.WithStack(err)
		}
		//comp, err := db.GetCompanyByID(c.ID)

		var users []models.User
		err = db.DB.Select(&users, "SELECT * FROM users WHERE company_id = $1;", c.ID)
		if err != nil {
			db.Logger.Error(errors.WithStack(err))
			return nil, errors.WithStack(err)
		}
		res[i].Apps = apps
		res[i].Users = users
	}
	return res, nil
}

func (db *DB) GetCompanyByID(id int) (*models.Company, error) {
	var res models.Company
	err := db.DB.Get(&res, "SELECT * FROM companies WHERE id = $1;", id)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	apps, err := db.GetCompanyApps(res.ID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	var users []models.User
	err = db.DB.Select(&users, "SELECT * FROM users WHERE company_id = $1;", id)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	res.Apps = apps
	res.Users = users

	return &res, nil
}

func (db *DB) DeleteCompany(companyID int) error {
	_, err := db.DB.Exec("DELETE FROM companies WHERE id=$1",
		companyID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	_, err = db.DB.Exec("DELETE FROM project_companies WHERE company_id=$1",
		companyID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}

	return nil
}
