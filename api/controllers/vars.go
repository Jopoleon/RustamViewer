package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (a *Controllers) GetVarsWithFilters(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	user := a.UserFromContext(w, r)

	res, err := a.Repository.DB.GetVarsByFilters(user.AppNames, params)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) GetVarsByCallID(w http.ResponseWriter, r *http.Request) {
	callID := chi.URLParam(r, "callID")
	if callID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := a.UserFromContext(w, r)
	vars, err := a.Repository.DB.GetVarsByCallID(callID, user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(vars)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
