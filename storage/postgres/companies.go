package postgres

import (
	"database/sql"

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
		var apps []models.Application
		query := `SELECT projects.id, project_name, description 
FROM projects,project_companies
WHERE projects.id = project_companies.project_id AND 
project_companies.company_id = $1;
`
		err = db.DB.Select(&apps, query, c.ID)
		if err != nil && err != sql.ErrNoRows {
			db.Logger.Error(errors.WithStack(err))
			return nil, errors.WithStack(err)
		}
		res[i].Apps = apps
	}
	return res, nil
}
