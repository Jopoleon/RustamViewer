package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/Jopoleon/rustamViewer/models"
)

func (a *Controllers) CreateNewApplication(w http.ResponseWriter, r *http.Request) {
	actor := a.UserFromContext(w, r)

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
	err = a.Repository.DB.CreateNewApp(&app)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.Logger.Info("application ", app, " CREATED by: ", actor)
	w.Write([]byte(fmt.Sprintf("Проект %s добавлен.", app.ProjectName)))
}

func (a *Controllers) UpdateApplication(w http.ResponseWriter, r *http.Request) {
	actor := a.UserFromContext(w, r)

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
	err = a.Repository.DB.UpdateApp(&app)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.Logger.Info("application ", app, " UPDATED by: ", actor)
	w.Write([]byte(fmt.Sprintf("Проект %s изменен.", app.ProjectName)))
}

func (a *Controllers) DeleteApplication(w http.ResponseWriter, r *http.Request) {
	actor := a.UserFromContext(w, r)
	appID := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(appID)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = a.Repository.DB.DeleteApplication(id)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.Logger.Info("application with ID ", id, " DELETED by: ", actor)
	w.Write([]byte(fmt.Sprintf("Проект удален.")))
}
