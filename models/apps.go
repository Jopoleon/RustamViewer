package models

import "errors"

type Application struct {
	ID          int    `json:"id" db:"id"`
	ProjectName string `json:"project_name" db:"project_name"`
	CompanyID   int    `json:"company_id" db:"company_id"`
	CompanyName string `json:"company_name" db:"company_name"`
	Description string `json:"description" db:"description"`
	AppUsers    []User
}

func (a *Application) Validate() error {
	if a.ProjectName == "" {
		return errors.New("Имя проекта не может быть пустым")
	}
	return nil
}
