package controllers

import (
	"net/http"

	"github.com/Jopoleon/rustamViewer/models"
)

func (a *Controllers) IndexHandler(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("userData").(models.User)
	if !ok || user.Login == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	asrs, err := a.Repository.DB.GetWaveRecordByProfileNames(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = Templates.ExecuteTemplate(w, "indexPage", IndexData{Ars: asrs, User: user})
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
