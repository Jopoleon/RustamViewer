package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jopoleon/rustamViewer/models"
)

func (a *Controllers) CreateNewUserTmpl(w http.ResponseWriter, r *http.Request) {
	user := a.UserFromContext(w, r)
	companies, err := a.Repository.DB.GetCompanies()
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = Templates.ExecuteTemplate(w, "createUser", IndexData{User: user, Companies: companies})
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) GetUser(w http.ResponseWriter, r *http.Request) {
	user := a.UserFromContext(w, r)
	company, err := a.Repository.DB.GetCompanyByID(user.CompanyID)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := IndexData{
		User:      user,
		Companies: []models.Company{*company},
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := models.User{}
	defer r.Body.Close()
	err = json.Unmarshal(b, &user)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := user.Validate(); err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = a.Repository.DB.CreateUser(user)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Пользователь %s %s (%s) создан.",
		user.FirstName, user.SecondName, user.Login)))
}

func (a *Controllers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := a.UserFromContext(w, r)
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
	if cred.Password != "" {
		if len(cred.Password) < 6 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("New password is too short"))
			return
		}
		user.Password = cred.Password
	}
	if cred.FirstName != "" {
		user.FirstName = cred.FirstName
	}
	if cred.SecondName != "" {
		user.SecondName = cred.SecondName
	}
	err = a.Repository.DB.UpdateUser(user)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Пользователь %s обновлен.", user.Login)))
}

func (a *Controllers) AddUserToProjectTmpl(w http.ResponseWriter, r *http.Request) {
	user := a.UserFromContext(w, r)

	apps, err := a.Repository.DB.GetAllApps()
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	users, err := a.Repository.DB.GetAllUsers()
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := IndexData{
		UserList: users,
		User:     user,
		Apps:     apps,
	}
	err = Templates.ExecuteTemplate(w, "addUserToProject", data)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) AddUserToProject(w http.ResponseWriter, r *http.Request) {

	a.UserFromContext(w, r)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cred := models.UsersApp{}
	defer r.Body.Close()
	err = json.Unmarshal(b, &cred)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.Repository.DB.AddUserToApp(cred.ProjectID, cred.UserID)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(
		fmt.Sprintf("Пользователь %s добавлен в %s.",
			cred.UserFullName, cred.ProjectName)))
}
