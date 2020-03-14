package models

type Company struct {
	ID   int           `json:"id" db:"id"`
	Name string        `json:"name" db:"name"`
	Apps []Application `json:"apps"`
}
