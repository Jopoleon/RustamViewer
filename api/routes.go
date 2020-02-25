package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (a *API) InitRouter() {

	r := chi.NewRouter()
	//Enabling Cross Origin Resource Sharing
	corss := cors.New(cors.Options{
		// Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(corss.Handler)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Group(func(r chi.Router) {
		fs := http.FileServer(http.Dir("api/templates/static"))
		r.Handle("/static/", http.StripPrefix("/static/", fs))
		r.Handle("/static", http.StripPrefix("/static/", fs))
		r.Handle("/static", fs)
		r.Handle("/", http.StripPrefix("/", fs))
		r.Handle("/", http.StripPrefix("/static", fs))
		r.Handle("/", http.StripPrefix("/static/", fs))
		r.MethodFunc("GET", "/login", a.LoginHandler)
		r.MethodFunc("POST", "/login", a.SubmitLogin)
		r.MethodFunc("POST", "/createUser/{secretToken}", a.CreateNewUser)
		// Private business logic routes
		r.Group(func(rr chi.Router) {
			rr.Use(a.CheckAuth)
			rr.MethodFunc("GET", "/", a.IndexHandler)
			rr.MethodFunc("GET", "/waverecord/{ID}", a.GetArs)
		})
	})

	a.Router = r
}
