package controllers

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/k0kubun/pp"
)

func (a *Controllers) GetArsWithFilters(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	pp.Println(params)
	res, err := a.Repository.DB.GetAllAsrWithFilters(params)
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
	wav, err := a.Repository.DB.GetWaveRecordByID(id)
	if err != nil {
		a.Logger.Errorf("%v", err)
		return
	}
	http.ServeContent(w, r, wav.Interpretation, time.Now(), bytes.NewReader(wav.Waverecord))
}
