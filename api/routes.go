package api

import (
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
	r.Use(middleware.Logger)
	r.Use(a.ParseTemplates)
	r.Use(middleware.Timeout(20 * time.Second))
	r.Group(func(r chi.Router) {
		r.MethodFunc("GET", "/login", a.LoginHandler)
		r.MethodFunc("GET", "/templates/*", a.ServeStatic)
		r.MethodFunc("POST", "/login", a.SubmitLogin)

		// Private business logic routes
		r.Group(func(rr chi.Router) {
			rr.Use(a.CheckAuth)
			rr.MethodFunc("GET", "/logout", a.LogoutHandler)
			rr.MethodFunc("GET", "/createNewUser", a.CreateNewUserTmpl)
			rr.MethodFunc("POST", "/createUser", a.CreateNewUser)
			rr.MethodFunc("POST", "/createApp", a.CreateNewUserApp)

			rr.MethodFunc("GET", "/companies", a.ListCompanies) // with apps included
			rr.MethodFunc("POST", "/createCompany", a.CreateNewCompany)
			rr.MethodFunc("GET", "/", a.IndexHandler)
			rr.MethodFunc("GET", "/filterTable", a.GetArsWithFilters)
			rr.MethodFunc("GET", "/waverecord/{ID}", a.GetArs)
		})
	})

	a.Router = r
}
