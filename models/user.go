package models

import (
	"errors"
)

type User struct {
	ID          int           `db:"id" json:"id"`
	ProfileName string        `db:"profile_name" json:"profileName"`
	FirstName   string        `db:"first_name" json:"firstName"`
	AppNames    []string      `json:"appNames"`
	Apps        []Application `json:"apps"`
	Login       string        `db:"login" json:"login"`
	SecondName  string        `db:"second_name" json:"secondName"`
	CompanyName string        `db:"company_name" json:"companyName"`
	CompanyID   int           `db:"company_id" json:"companyID"`
	Email       string        `db:"email" json:"email"`
	IsAdmin     bool          `db:"is_admin" json:"isAdmin"`
	Password    string        `db:"password" json:"password"`
}

func (u *User) Validate() error {
	if u.Login == "" || len(u.Login) < 5 {
		return errors.New("Логин слишком короткий")
	}
	if u.FirstName == "" {
		return errors.New("Пустое имя")
	}
	if u.SecondName == "" {
		return errors.New("Пустая фамилия")
	}
	if u.Password == "" || len(u.Password) < 6 {
		return errors.New("Пароль слишком короткий")
	}
	if u.CompanyName == "" || u.CompanyName == "Выбрать компанию" || u.CompanyID == 0 {
		return errors.New("Пользователь должен быть прикреплен к компании")
	}
	if u.Email == "" {
		return errors.New("Пустой почтовый адрес")
	}
	return nil
}

func ProjectNameAccessRights(projectName string, userApps []string) bool {
	//pp.Println("ProjectNameAccessRights user list::: ", userApps)
	for _, pb := range userApps {
		//pp.Println("ProjectNameAccessRights user project::: ", pb)
		//pp.Println("ProjectNameAccessRights incoming project::: ", projectName)
		if projectName == pb {
			return true
		}
	}
	return false
}

type UsersApp struct {
	ID          int    `db:"id" json:"id"`
	UserID      int    `json:"user_id" db:"user_id"`
	ProjectID   int    `json:"project_id" db:"project_id"`
	ProjectName string `json:"project_name" db:"project_name"`

	UserFullName string `json:"user_full_name"`
}
