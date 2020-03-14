package api

import (
	"context"
	"net/http"

	"github.com/Jopoleon/rustamViewer/models"

	"github.com/gorilla/securecookie"
)

const (
	CookieAuthName = "AUTH_SESSION"
	CookieName     = "user_session"
)

func (a *API) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cookie, err := r.Cookie(CookieName)
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			a.Logger.Errorf("%v", err)
			w.Write([]byte(err.Error()))
			return
		}
		var value models.Session
		s := securecookie.New([]byte(a.Config.CookieSecret), nil)

		err = s.Decode(CookieAuthName, cookie.Value, &value)
		if err != nil {
			a.Logger.Errorf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		user, err := a.Repository.DB.GetUserByID(value.UserID)
		if err != nil {
			a.Logger.Errorf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		//pp.Println("Middleware user: ", user)
		ctx = context.WithValue(ctx, "userData", models.User{
			ID:          user.ID,
			ProfileName: user.ProfileName,
			Login:       user.Login,
			FirstName:   user.FirstName,
			SecondName:  user.SecondName,
			CompanyName: user.CompanyName,
			AppNames:    user.AppNames,
			Email:       user.Email,
			IsAdmin:     user.IsAdmin,
		})
		ctx = context.WithValue(ctx, "email", user.Email)
		ctx = context.WithValue(ctx, "login", user.Login)
		ctx = context.WithValue(ctx, "userID", user.ID)
		ctx = context.WithValue(ctx, "isAdmin", user.IsAdmin)
		ctx = context.WithValue(ctx, "isLoggedIn", value.LoggedIn)
		ctx = context.WithValue(ctx, "profileName", user.ProfileName)
		//http.SetCookie(w, cookie)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
