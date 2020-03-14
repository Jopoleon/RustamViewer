package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Jopoleon/rustamViewer/models"
	"github.com/gorilla/securecookie"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func (a *Controllers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	isLogged := r.Context().Value("isLoggedIn")
	switch isLogged.(type) {
	case bool:
		if isLogged.(bool) == true {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	case nil:
	}

	err := Templates.ExecuteTemplate(w, "login", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Templates.Execute(w, nil)

}

func (a *Controllers) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userID")
	if userID == nil {
		a.Logger.Errorf("Logout handler error: empty userID or isLogged fields")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errA := a.Repository.DB.DeleteUserSession((userID).(int))
	if errA != nil {
		a.Logger.Errorf("%v", errA)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: CookieName, MaxAge: -1})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}

func (a *Controllers) SubmitLogin(w http.ResponseWriter, r *http.Request) {
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
	if cred.Email == "" {
		//a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid email or login"))
		return
	}

	user, err := a.Repository.DB.GetUserByEmailOrLogin(cred.Email)
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
	sessionToken := uuid.NewV4()

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
