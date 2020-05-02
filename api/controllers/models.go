package controllers

import "github.com/Jopoleon/rustamViewer/models"

type IndexData struct {
	User          *models.User           `json:"user,omitempty"`
	UserList      []models.User          `json:"userList,omitempty"`
	Ars           []models.ASR           `json:"Ars,omitempty"`
	Vars          []models.VAR           `json:"Vars,omitempty"`
	CallsAll      []models.Calls         `json:"CallsAll,omitempty"`
	CallsOutbound []models.CallsOutbound `json:"CallsOutbound,omitempty"`
	Companies     []models.Company       `json:"companies,omitempty"`
	Apps          []models.Application   `json:"apps,omitempty"`
}

type File struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Data []byte `json:"data"`
}
