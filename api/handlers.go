package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/securecookie"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/chi"

	"github.com/Jopoleon/rustamViewer/models"
)

const (
	CookieMaxAge   = 259200 //cookie expiration time in seconds (3 days)
	CookieAuthName = "AUTH_SESSION"
	CookieName     = "user_session"
)

func (a *API) LoginHandler(w http.ResponseWriter, r *http.Request) {
	isLogged := r.Context().Value("isLoggedIn")
	switch isLogged.(type) {
	case bool:
		if isLogged.(bool) == true {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	case nil:
	}

	tmpl := template.Must(template.ParseFiles("api/templates/login.html"))

	tmpl.Execute(w, nil)

}

type CredStruct struct {
	Login    string `json:"email"`
	Password string `json:"password"`
}

func (a *API) SubmitLogin(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cred := CredStruct{}
	defer r.Body.Close()
	err = json.Unmarshal(b, &cred)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := a.Repository.DB.GetUserByEmailOrLogin(cred.Login)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User with such email or login not found"))
			return
		}
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong password"))
		return
	}
	var s = securecookie.New([]byte(a.Config.CookieSecret), nil)
	sessionToken:= uuid.NewV4()

	sessionStruct := models.Session{
		LoggedIn:     true,
		UserID:       user.ID,
		SessionToken: sessionToken.String(),
	}

	encoded, errS := s.Encode(CookieAuthName, sessionStruct)
	if errS != nil {
		a.Logger.Errorf("%v", errS)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     CookieName,
		Value:    encoded,
		Path:     "/",
		MaxAge:   CookieMaxAge,
		HttpOnly: true,
	}

	errA := a.Repository.DB.SetUserSession(user.ID, cookie.Value)
	if errA != nil {
		a.Logger.Errorf("%v", errA)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &cookie)
	return
}

type ArsData struct {
	Ars []models.ASR
}

func (a *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("api/templates/index.html"))
	p := r.Context().Value("profileName")
	if p == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	profile := p.(string)
	asrs, err := a.Repository.DB.GetWaveRecordByProfileName(profile)
	if err != nil {
		a.Logger.Errorf("%v", err)
	}

	err = tmpl.Execute(w, ArsData{Ars: asrs})
	if err != nil {
		a.Logger.Errorf("%v", err)
	}

}

func (a *API) GetArs(w http.ResponseWriter, r *http.Request) {
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
	http.ServeContent(w, r, wav.Interpritation, time.Now(), bytes.NewReader(wav.WAVRecord))
}

func (a *API) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	secretToken := chi.URLParam(r, "secretToken")
	if secretToken != a.Config.CreateUserSecret {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := a.Repository.DB.CreateUser()
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.Write([]byte("Try again, internal error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("{'login':%s, 'password':%s", user.Login, user.Password)))

}
