package api

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *API) InitRouter() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(a.ParseTemplates)
	//Enabling Cross Origin Resource Sharing
	r.Use(a.OptionsCors)
	r.Use(middleware.Timeout(20 * time.Second))

	r.Group(func(r chi.Router) {
		r.MethodFunc("GET", "/login", a.LoginHandler)
		r.MethodFunc("GET", "/auth", a.AuthHandler)
		r.MethodFunc("GET", "/templates/*", a.ServeStatic)
		r.MethodFunc("POST", "/login", a.SubmitLogin)

		// Private business logic routes
		r.Group(func(rr chi.Router) {
			rr.Use(a.CheckAuth)
			rr.MethodFunc("GET", "/", a.IndexHandler)

			rr.MethodFunc("GET", "/logout", a.LogoutHandler)
			rr.MethodFunc("GET", "/user", a.GetUser)
			rr.MethodFunc("PUT", "/user", a.UpdateUser)

			rr.MethodFunc("GET", "/waverecord/{ID}", a.GetArs)
			rr.MethodFunc("GET", "/file", a.GetFile)

			rr.MethodFunc("GET", "/export/calls", a.ExportCallsAll)
			rr.MethodFunc("GET", "/export/callsOut", a.ExportCallsOutBound)
			rr.MethodFunc("GET", "/export/vars", a.ExportVars)
			rr.MethodFunc("GET", "/export/ars", a.ExportArsresults)

			rr.MethodFunc("GET", "/calls/{callID}", a.GetCallsAllByCallID)
			rr.MethodFunc("GET", "/callsout/{callID}", a.GetCallsOutByCallID)
			rr.MethodFunc("GET", "/vars/{callID}", a.GetVarsByCallID)
			rr.MethodFunc("GET", "/ars/{callID}", a.GetArsByCallID)

			rr.MethodFunc("GET", "/calls/{projectID}", a.GetCallsAll)
			rr.MethodFunc("GET", "/callsout/{projectID}", a.GetCallsOutAll)
			rr.MethodFunc("GET", "/vars/{projectID}", a.GetVarsWithFilters)
			rr.MethodFunc("GET", "/ars/{projectID}", a.GetArsByProjectID)

			rr.MethodFunc("GET", "/callsout", a.GetCallsOutAll)
			rr.MethodFunc("GET", "/calls", a.GetCallsAll)
			rr.MethodFunc("GET", "/vars", a.GetVarsWithFilters)
			rr.MethodFunc("GET", "/ars", a.GetAllArsWithFilters)

			rr.Group(func(admin chi.Router) {
				admin.Use(a.AdminOnly)
				admin.MethodFunc("GET", "/companies", a.ListCompanies)           // with apps and users included
				admin.MethodFunc("GET", "/companies/{companyID}", a.CompanyByID) // with apps and users included
				admin.MethodFunc("POST", "/company", a.CreateNewCompany)
				admin.MethodFunc("GET", "/createUser", a.CreateNewUserTmpl)
				admin.MethodFunc("POST", "/user", a.CreateNewUser)
				admin.MethodFunc("GET", "/projects", a.AddUserToProjectTmpl)
				admin.MethodFunc("POST", "/projects/{userID}", a.AddUserToProject)
				admin.MethodFunc("POST", "/project", a.CreateNewUserApp)
			})
		})
	})

	a.Router = r
}
