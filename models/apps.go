package models

type Application struct {
	ID          int    `json:"id" db:"id"`
	ProfileName string `json:"profileName" db:"profile_name"`
	UserID      int    `json:"userID" db:"user_id"`
	CompanyID   int    `json:"company_id" db:"company_id"`
}

func (a *Application) Validate() error {
	return nil
}
