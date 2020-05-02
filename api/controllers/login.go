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

func (a *Controllers) AuthHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie(CookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var value models.Session
	s := securecookie.New([]byte(a.Config.CookieSecret), nil)

	err = s.Decode(CookieAuthName, cookie.Value, &value)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("cookies not correct"))
		return
	}

	user, err := a.Repository.DB.GetUserByID(value.UserID)
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("user not found"))
		return
	}
	err = json.NewEncoder(w).Encode(struct {
		FirstName  string `json:"first_name"`
		SecondName string `json:"second_name"`
		IsAdmin    bool   `json:"is_admin"`
	}{
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		IsAdmin:    user.IsAdmin,
	})
	if err != nil {
		a.Logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("userData").(models.User)
	if ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err := Templates.ExecuteTemplate(w, "login", nil)
	if err != nil {
		a.Logger.Errorf("%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Controllers) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	user := a.UserFromContext(w, r)

	errA := a.Repository.DB.DeleteUserSession(user.ID)
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid email or login"))
		return
	}
	//bpas, err := bcrypt.GenerateFromPassword([]byte(cred.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	a.Logger.Error(errors.WithStack(err))
	//	//return errors.WithStack(err)
	//}
	//pp.Println(string(bpas))
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
		Name:   CookieName,
		Value:  encoded,
		Path:   "/",
		MaxAge: CookieMaxAge,
		//HttpOnly: true,
	}

	errA := a.Repository.DB.SetUserSession(user.ID, cookie.Value)
	if errA != nil {
		a.Logger.Errorf("%v", errA)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &cookie)
	err = json.NewEncoder(w).Encode(struct {
		FirstName  string `json:"first_name"`
		SecondName string `json:"second_name"`
		IsAdmin    bool   `json:"is_admin"`
	}{
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		IsAdmin:    user.IsAdmin,
	})
	if err != nil {
		a.Logger.Errorf("%v", errA)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
