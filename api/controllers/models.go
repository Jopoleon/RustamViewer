package controllers

import "github.com/Jopoleon/rustamViewer/models"

type IndexData struct {
	User      models.User          `json:"user"`
	Ars       []models.ASR         `json:"Ars"`
	Companies []models.Company     `json:"companies"`
	Apps      []models.Application `json:"apps"`
}
