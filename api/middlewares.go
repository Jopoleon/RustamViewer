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
				a.Logger.Errorf("%v", err)
				http.RedirectHandler("/login", http.StatusUnauthorized)
				//http.Error(w, "no cookie found", http.StatusUnauthorized)
				return
			}
			a.Logger.Errorf("%v", err)
			http.Error(w, "can't read cookie", http.StatusBadRequest)
			return
		}
		var value models.Session
		s := securecookie.New([]byte(a.Config.CookieSecret), nil)

		err = s.Decode(CookieAuthName, cookie.Value, &value)
		if err != nil {
			a.Logger.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//ctxUser := a.Controllers.UserFromContext(w, r)
		//if ctxUser == nil || ctxUser.Login == "" {
		//	a.Logger.Errorf("%v", err)
		//	http.Redirect(w, r, "/login", http.StatusSeeOther)
		//	return
		//}
		user, err := a.Repository.DB.GetUserByID(value.UserID)
		if err != nil {
			a.Logger.Errorf("%v", err)
			//This happens when you WriteHeader after a Write, which already flushes the headers.
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		ctx = context.WithValue(ctx, "userData", models.User{
			ID:          user.ID,
			Login:       user.Login,
			FirstName:   user.FirstName,
			SecondName:  user.SecondName,
			CompanyName: user.CompanyName,
			CompanyID:   user.CompanyID,
			AppNames:    user.AppNames,
			Email:       user.Email,
			IsAdmin:     user.IsAdmin,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *API) AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := a.Controllers.UserFromContext(w, r)
		if !user.IsAdmin {
			http.Error(w, "only admin allowed", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
