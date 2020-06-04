package api

import (
	chiLog "github.com/766b/chi-logger"
	"github.com/go-chi/chi"
)

func (a *API) InitRouter() {

	r := chi.NewRouter()
	r.Use(chiLog.NewLogrusMiddleware("router", a.Logger.Log))
	r.Use(a.ParseTemplates)

	//Enabling Cross Origin Resource Sharing
	r.Use(a.OptionsCors)
	//r.Use(middleware.Timeout(20 * time.Second))

	r.Group(func(r chi.Router) {
		r.MethodFunc("GET", "/login", a.IndexHandler)

		r.MethodFunc("GET", "/auth", a.AuthHandler)

		r.MethodFunc("GET", "/*", a.ServeStatic)

		r.MethodFunc("POST", "/login", a.SubmitLogin)

		r.MethodFunc("GET", "/", a.IndexHandler)
		r.MethodFunc("GET", "/profile", a.IndexHandler)
		r.MethodFunc("GET", "/profile/edit", a.IndexHandler)
		r.MethodFunc("GET", "/user/add", a.IndexHandler)
		r.MethodFunc("GET", "/project/add", a.IndexHandler)
		r.MethodFunc("GET", "/company/add", a.IndexHandler)
		r.MethodFunc("GET", "/tables", a.IndexHandler)
		r.MethodFunc("GET", "/companieslist", a.IndexHandler) // with apps and users included
		// Private business logic routes
		r.Group(func(rr chi.Router) {
			rr.Use(a.CheckAuth)

			rr.MethodFunc("GET", "/logout", a.LogoutHandler)
			rr.MethodFunc("GET", "/user", a.GetUser)
			rr.MethodFunc("PUT", "/user", a.UpdateUser)

			rr.MethodFunc("GET", "/waverecord/{ID}", a.GetArs)
			rr.MethodFunc("GET", "/file", a.GetFile)
			rr.MethodFunc("GET", "/file", a.ListFiles)

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
				admin.MethodFunc("GET", "/companies", a.ListCompanies)
				//admin.MethodFunc("GET", "/companieslist", a.IndexHandler)        // with apps and users included
				admin.MethodFunc("GET", "/companies/{companyID}", a.CompanyByID) // with apps and users included
				admin.MethodFunc("POST", "/company", a.CreateNewCompany)
				admin.MethodFunc("DELETE", "/company/{ID}", a.DeleteCompany)

				admin.MethodFunc("POST", "/user", a.CreateNewUser)
				admin.MethodFunc("DELETE", "/user/{ID}", a.DeleteUser)
				admin.MethodFunc("DELETE", "/user/{ID}/project/{ProjectID}", a.DeleteUseFromProject)

				admin.MethodFunc("POST", "/projects/{userID}", a.AddUserToProject)
				admin.MethodFunc("POST", "/project", a.CreateNewApplication)
				admin.MethodFunc("DELETE", "/project/{ID}", a.DeleteApplication)
			})
		})
	})

	a.Router = r
}
