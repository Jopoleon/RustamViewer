package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/Jopoleon/rustamViewer/models"
)

func (a *Controllers) CreateNewCompany(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c := models.Company{}
	defer r.Body.Close()
	err = json.Unmarshal(b, &c)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.Validate()
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = a.Repository.DB.CreateNewCompany(c.Name)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *Controllers) ListCompanies(w http.ResponseWriter, r *http.Request) {

	_ = a.UserFromContext(w, r)

	comps, err := a.Repository.DB.GetCompanies()
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := IndexData{Companies: comps}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) CompanyByID(w http.ResponseWriter, r *http.Request) {

	_ = a.UserFromContext(w, r)

	companyID := chi.URLParam(r, "companyID")
	if companyID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cID, err := strconv.Atoi(companyID)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comps, err := a.Repository.DB.GetCompanyByID(cID)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := IndexData{Companies: []models.Company{*comps}}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
