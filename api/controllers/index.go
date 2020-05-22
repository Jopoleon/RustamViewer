package controllers

import (
	"net/http"
	"text/template"
)

func (a *Controllers) IndexHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("api/templates/index.html")
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
