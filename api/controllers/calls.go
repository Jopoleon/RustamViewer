package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (a *Controllers) GetCallsAll(w http.ResponseWriter, r *http.Request) {

	user := a.UserFromContext(w, r)
	if user.IsAdmin {
		callsAll, err := a.Repository.DB.GetCallsAllAdmin()
		if err != nil {
			a.Logger.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(callsAll)
		if err != nil {
			a.Logger.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	callsAll, err := a.Repository.DB.GetCallsAllProjectIDs(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(callsAll)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) GetCallsOutAll(w http.ResponseWriter, r *http.Request) {

	user := a.UserFromContext(w, r)
	if user.IsAdmin {
		callOutbound, err := a.Repository.DB.GetCallsOutboundAdmin()
		if err != nil {
			a.Logger.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(callOutbound)
		if err != nil {
			a.Logger.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	callOutbound, err := a.Repository.DB.GetCallsOutboundProjectIDs(user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(callOutbound)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) GetCallsAllByCallID(w http.ResponseWriter, r *http.Request) {

	callID := chi.URLParam(r, "callID")
	if callID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := a.UserFromContext(w, r)

	if user.IsAdmin {
		callsAll, err := a.Repository.DB.GetCallsAllByCallIDAdmin(callID)
		if err != nil {
			a.Logger.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(callsAll)
		if err != nil {
			a.Logger.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	callsAll, err := a.Repository.DB.GetCallsAllByCallID(callID, user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(callsAll)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) GetCallsOutByCallID(w http.ResponseWriter, r *http.Request) {
	callID := chi.URLParam(r, "callID")
	if callID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := a.UserFromContext(w, r)
	if user.IsAdmin {
		callOutbound, err := a.Repository.DB.GetCallsOutboundByCallIDAdmin(callID)
		if err != nil {
			a.Logger.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(callOutbound)
		if err != nil {
			a.Logger.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	callOutbound, err := a.Repository.DB.GetCallsOutboundByCallID(callID, user.AppNames)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(callOutbound)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
