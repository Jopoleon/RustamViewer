package controllers

import (
	"net/http"

	"github.com/gocarina/gocsv"
)

const (
	ARS_CSV_NAME            = "ars.csv"
	VARS_CSV_NAME           = "vars.csv"
	CALLS_ALL_CSV_NAME      = "calls_all.csv"
	CALLS_OUTBOUND_CSV_NAME = "calls_outbound.csv"
)

func (a *Controllers) writeCSV(w http.ResponseWriter, content interface{}, filename string) {
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "text/csv")
	err := gocsv.Marshal(content, w)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) ExportCallsAll(w http.ResponseWriter, r *http.Request) {
	user := a.UserFromContext(w, r)

	callsAll, err := a.Repository.DB.GetCallsAllProjectIDs(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.writeCSV(w, callsAll, CALLS_ALL_CSV_NAME)

}
func (a *Controllers) ExportCallsOutBound(w http.ResponseWriter, r *http.Request) {
	user := a.UserFromContext(w, r)

	callsOut, err := a.Repository.DB.GetCallsOutboundProjectIDs(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.writeCSV(w, callsOut, CALLS_OUTBOUND_CSV_NAME)

}
func (a *Controllers) ExportVars(w http.ResponseWriter, r *http.Request) {
	user := a.UserFromContext(w, r)

	vars, err := a.Repository.DB.GetVarsByProjectIDs(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.writeCSV(w, vars, VARS_CSV_NAME)

}
func (a *Controllers) ExportArsresults(w http.ResponseWriter, r *http.Request) {
	user := a.UserFromContext(w, r)

	ars, err := a.Repository.DB.GetWaveRecordByProfileNames(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.writeCSV(w, ars, ARS_CSV_NAME)

}
