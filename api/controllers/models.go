package controllers

import "github.com/Jopoleon/rustamViewer/models"

type IndexData struct {
	User          *models.User           `json:"user"`
	UserList      []models.User          `json:"userList"`
	Ars           []models.ASR           `json:"Ars"`
	Vars          []models.VAR           `json:"Vars"`
	CallsAll      []models.Calls         `json:"CallsAll"`
	CallsOutbound []models.CallsOutbound `json:"CallsOutbound"`
	Companies     []models.Company       `json:"companies"`
	Apps          []models.Application   `json:"apps"`
}

type File struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Data []byte `json:"data"`
}
