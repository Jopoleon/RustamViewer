package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jopoleon/rustamViewer/models"
)

func (a *Controllers) CreateNewUserApp(w http.ResponseWriter, r *http.Request) {
	a.UserFromContext(w, r)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app := models.Application{}
	defer r.Body.Close()
	err = json.Unmarshal(b, &app)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := app.Validate(); err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = a.Repository.DB.CreateNewApp(app.CompanyID, app.ProjectName, app.Description)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Проект %s добавлен.", app.ProjectName)))
}
