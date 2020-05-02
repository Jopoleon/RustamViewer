package models

import "errors"

type Company struct {
	ID    int           `json:"id" db:"id"`
	Name  string        `json:"name" db:"name"`
	Users []User        `json:"users"`
	Apps  []Application `json:"apps"`
}

func (c *Company) Validate() error {
	if c.Name == "" {
		return errors.New("Имя не может быть пустым")
	}
	return nil
}
