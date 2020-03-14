package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jopoleon/rustamViewer/models"
	"github.com/k0kubun/pp"
)

func (a *Controllers) CreateNewUserTmpl(w http.ResponseWriter, r *http.Request) {

	IsAdmin := r.Context().Value("isAdmin").(bool)
	if !IsAdmin {
		http.Error(w, "not admin", http.StatusUnauthorized)
		return
	}
	user, ok := r.Context().Value("userData").(models.User)
	if !ok || user.Login == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	err := Templates.ExecuteTemplate(w, "createUser", IndexData{User: user})
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) CreateNewUser(w http.ResponseWriter, r *http.Request) {
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
	cred := models.CredStruct{}
	defer r.Body.Close()
	err = json.Unmarshal(b, &cred)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := cred.Validate(); err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	user := models.User{
		ProfileName: cred.Profilename,
		Login:       cred.Login,
		FirstName:   cred.FirstName,
		SecondName:  cred.SecondName,
		CompanyName: cred.CompanyName,
		Email:       cred.Email,
		IsAdmin:     false,
		Password:    cred.Password,
	}
	pp.Println(user)
	err = a.Repository.DB.CreateUser(user)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("User %s created.", user.Login)))
}
