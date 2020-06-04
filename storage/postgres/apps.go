package postgres

import (
	"fmt"
	"strings"

	"github.com/Jopoleon/rustamViewer/models"
	"github.com/pkg/errors"
)

func (db *DB) CreateNewApp(app *models.Application) error {

	rows, err := db.DB.Query("INSERT INTO projects (project_name, description, is_recording, transcription, language) VALUES ($1,$2,$3,$4,$5) RETURNING id;",
		app.ProjectName, app.Description, app.IsRecording, app.Transcription, app.Language)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return errors.WithStack(errors.New(
				fmt.Sprintf(ERROR_NON_UNIQUE, app.ProjectName)))
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
		projectID, app.CompanyID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) DeleteUserFromApplication(userID, projectID int) error {

	q := `DELETE FROM users_projects WHERE user_id = $1 AND project_id = $2;`
	_, err := db.DB.Exec(q, userID, projectID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}

	return nil
}
func (db *DB) DeleteUserFromAllApps(user_id int) error {
	_, err := db.DB.Exec("DELETE FROM users_projects WHERE user_id=$1",
		user_id)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) DeleteApplication(appID int) error {
	_, err := db.DB.Exec("DELETE FROM projects WHERE id=$1",
		appID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	_, err = db.DB.Exec("DELETE FROM project_companies WHERE project_id=$1",
		appID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	_, err = db.DB.Exec("DELETE FROM users_projects WHERE project_id=$1",
		appID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}

	return nil
}

func (db *DB) GetCompanyApps(companyID int) ([]models.Application, error) {
	var res []models.Application
	q := `SELECT projects.id, projects.project_name,projects.description FROM projects,project_companies 
				WHERE projects.id = project_companies.project_id 
				AND project_companies.company_id = $1;`
	err := db.DB.Select(&res, q, companyID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	for i, a := range res {
		var users []models.User
		appUserQ :=
			`SELECT users.id, login, email, password, is_admin, first_name, 
				second_name, company_name, company_id
    			FROM users,users_projects
    			WHERE users.id = users_projects.user_id AND
           			users_projects.project_id= $1;`
		err = db.DB.Select(&users, appUserQ, a.ID)
		if err != nil {
			db.Logger.Error(errors.WithStack(err))
			return nil, errors.WithStack(err)
		}
		res[i].AppUsers = users
	}
	return res, nil
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
