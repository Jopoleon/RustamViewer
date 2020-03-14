package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/Jopoleon/rustamViewer/models"
)

func (a *Controllers) CreateNewCompany(w http.ResponseWriter, r *http.Request) {
	IsAdmin := r.Context().Value("isAdmin").(bool)
	if !IsAdmin {
		http.Error(w, "not admin", http.StatusUnauthorized)
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = b
}

func (a *Controllers) ListCompanies(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("userData").(models.User)
	if !ok || user.Login == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !user.IsAdmin {
		http.Error(w, "not admin", http.StatusUnauthorized)
		return
	}
	comps, err := a.Repository.DB.GetCompanies()
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := IndexData{Companies: comps, User: user}
	err = Templates.ExecuteTemplate(w, "companiesList", data)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
