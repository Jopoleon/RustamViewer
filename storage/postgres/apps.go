package postgres

import (
	"fmt"
	"strings"

	"github.com/Jopoleon/rustamViewer/models"
	"github.com/pkg/errors"
)

func (db *DB) CreateNewApp(companyID int, projectName, description string) error {

	rows, err := db.DB.Query("INSERT INTO projects (project_name, description) VALUES ($1,$2) RETURNING id;",
		projectName, description)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return errors.WithStack(errors.New(
				fmt.Sprintf(ERROR_NON_UNIQUE, projectName)))
		}
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}

	var projectID int
	if rows.Next() {
		err := rows.Scan(&projectID)
		if err != nil {
			db.Logger.Error(errors.WithStack(err))
			return errors.WithStack(err)
		}
	}

	_, err = db.DB.Exec("INSERT INTO project_companies (project_id, company_id) VALUES ($1,$2);",
		projectID, companyID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) DeleteUserFromApp(user_id int) error {
	_, err := db.DB.Exec("DELETE FROM users_projects WHERE user_id=$1",
		user_id)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) GetUserApps(user_id int) ([]models.Application, error) {
	var res []models.Application
	q := `SELECT * FROM projects,users_projects 
				WHERE projects.id = users_projects.project_id 
				AND users_projects.user_id = $1;`
	err := db.DB.Select(&res, q, user_id)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetAllApps() ([]models.Application, error) {

	var apps []models.Application

	appsWithCompaniesQuery := `select projects.id as "id" ,projects.project_name, 
		companies.name as "company_name", companies.id as "company_id"
		from projects
         join project_companies
              on projects.id = project_companies.project_id
         join companies
              on project_companies.company_id = companies.id`

	err := db.DB.Select(&apps, appsWithCompaniesQuery)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	//pp.Println("GetAllApps: ", apps)

	for i, app := range apps {
		var users []models.User
		appUserQ :=
			`SELECT users.id, login, email, password, is_admin, first_name, second_name, company_name, company_id
    			FROM users,users_projects
    			WHERE users.id = users_projects.user_id AND
           			users_projects.project_id= $1;`
		err = db.DB.Select(&users, appUserQ, app.ID)
		if err != nil {
			db.Logger.Error(errors.WithStack(err))
			return nil, errors.WithStack(err)
		}
		apps[i].AppUsers = users
	}
	return apps, nil
}
