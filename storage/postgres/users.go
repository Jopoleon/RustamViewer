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
	q := `INSERT INTO users 
    (login, email, password, is_admin, first_name, second_name, company_name, company_id) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err = db.DB.Exec(q,
		user.Login, user.Email, string(bpas), false, user.FirstName,
		user.SecondName, user.CompanyName, user.CompanyID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) DeleteUser(userID int) error {

	q := `DELETE FROM users WHERE id = $1;`
	_, err := db.DB.Exec(q, userID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}

	return nil
}

func (db *DB) UpdateUser(user *models.User) error {

	bpas, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	_, err = db.DB.Exec("UPDATE users SET password = $1, first_name = $2, second_name=$3 WHERE id = $4;",
		bpas, user.FirstName, user.SecondName, user.ID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}

func (db *DB) GetUserByID(id int) (*models.User, error) {
	user := models.User{}
	err := db.DB.Get(&user, "SELECT * FROM users WHERE id=$1 LIMIT 1;", id)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	var apps []models.Application
	q := `SELECT projects.id, project_name, description
    			FROM projects,users_projects
    			WHERE projects.id = users_projects.project_id AND
           			users_projects.user_id = $1;`
	err = db.DB.Select(&apps, q, user.ID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	user.Apps = apps
	for _, a := range apps {
		user.AppNames = append(user.AppNames, a.ProjectName)
	}
	return &user, nil
}

func (db *DB) GetAllUsers() ([]models.User, error) {
	res := []models.User{}
	err := db.DB.Select(&res, "SELECT * FROM users;")
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (db *DB) GetUserByEmailOrLogin(email string) (*models.User, error) {
	user := models.User{}
	err := db.DB.Get(&user, "SELECT * FROM users WHERE email=$1 OR login=$1;", email)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}

	var apps []models.Application
	q := `SELECT projects.id, project_name, description
    			FROM projects,users_projects
    			WHERE projects.id = users_projects.project_id AND
           			users_projects.user_id = $1;`
	err = db.DB.Select(&apps, q, user.ID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	user.Apps = apps
	for _, a := range apps {
		user.AppNames = append(user.AppNames, a.ProjectName)
	}

	return &user, nil
}

func (db *DB) CheckStatus() {
	//query := `DROP TABLE IF EXISTS test22221;`
	query := `DROP TABLE IF EXISTS asrresults;
	DROP TABLE IF EXISTS calls_all;
	DROP TABLE IF EXISTS calls_outbound;
	DROP TABLE IF EXISTS companies;
	DROP TABLE IF EXISTS project_companies;
	DROP TABLE IF EXISTS projects;
	DROP TABLE IF EXISTS users;
	DROP TABLE IF EXISTS users_projects;
	DROP TABLE IF EXISTS users_sessions;
	DROP TABLE IF EXISTS var;`
	db.DB.Exec(query)
}

func (db *DB) AddUserToApp(projectID int, userID int) error {
	query := `INSERT INTO users_projects (user_id, project_id)
	SELECT $1, $2
	WHERE NOT EXISTS (
		SELECT (user_id, project_id) FROM users_projects WHERE 
		user_id = $1 AND 
		project_id = $2);`

	_, err := db.DB.Exec(query, userID, projectID)
	if err != nil {
		db.Logger.Error(errors.WithStack(err))
		return errors.WithStack(err)
	}
	return nil
}
