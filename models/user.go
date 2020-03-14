package models

type User struct {
	ID          int      `db:"id" json:"id"`
	ProfileName string   `db:"profile_name" json:"profileName"`
	FirstName   string   `db:"first_name" json:"firstName"`
	AppNames    []string `json:"appNames"`
	Login       string   `db:"login" json:"login"`
	SecondName  string   `db:"second_name" json:"secondName"`
	CompanyName string   `db:"company_name" json:"companyName"`
	Email       string   `db:"email" json:"email"`
	IsAdmin     bool     `db:"is_admin" json:"isAdmin"`
	Password    string   `db:"password" json:"password"`
}
