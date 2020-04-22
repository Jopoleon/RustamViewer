package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (a *Controllers) GetArsWithFilters(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	user := a.UserFromContext(w, r)

	res, err := a.Repository.DB.GetAllAsrWithFilters(user.AppNames, params)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = Templates.ExecuteTemplate(w, "tableASR", IndexData{Ars: res})
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (a *Controllers) GetAllArsWithFilters(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	user := a.UserFromContext(w, r)

	res, err := a.Repository.DB.GetAllAsrWithFilters(user.AppNames, params)
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
func (a *Controllers) GetArs(w http.ResponseWriter, r *http.Request) {
	ids := chi.URLParam(r, "ID")
	if ids == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(ids)
	if err != nil {
		a.Logger.Errorf("%v", err)
		return
	}
	user := a.UserFromContext(w, r)
	wav, err := a.Repository.DB.GetWaveRecordByID(id, user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		return
	}

	http.ServeContent(w, r, *wav.Interpretation, time.Now(), bytes.NewReader(wav.Waverecord))
}

func (a *Controllers) GetArsByCallID(w http.ResponseWriter, r *http.Request) {
	callID := chi.URLParam(r, "callID")
	if callID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := a.UserFromContext(w, r)
	ars, err := a.Repository.DB.GetWaveRecordByCallID(callID, user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(ars)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
