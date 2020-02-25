package models

type User struct {
	ID          int    `json:"id" db:"id"`
	ProfileName string `json:"profileName" db:"profile_name"`
	Login       string `json:"login" db:"login"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
}
