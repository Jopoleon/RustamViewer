package controllers

import (
	"net/http"
)

func (a *Controllers) IndexHandler(w http.ResponseWriter, r *http.Request) {

	user := a.UserFromContext(w, r)

	asrs, err := a.Repository.DB.GetWaveRecordByProfileNames(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	vars, err := a.Repository.DB.GetVarsByProjectIDs(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	callsAll, err := a.Repository.DB.GetCallsAllProjectIDs(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	callOutbound, err := a.Repository.DB.GetCallsOutboundProjectIDs(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := IndexData{
		Vars:          vars,
		CallsOutbound: callOutbound,
		CallsAll:      callsAll,
		Ars:           asrs,
		User:          user,
	}
	err = Templates.ExecuteTemplate(w, "indexPage", data)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
