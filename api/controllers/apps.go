package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Jopoleon/rustamViewer/models"
)

func (a *Controllers) CreateNewUserApp(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("userData").(models.User)
	if !ok || user.Login == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

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
		w.Write([]byte(err.Error()))
		return
	}
	err = a.Repository.DB.CreateNewApp(user.ID, app.CompanyID, app.ProfileName)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
