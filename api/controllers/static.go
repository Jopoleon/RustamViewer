package controllers

import (
	"net/http"
	"text/template"
)

//var Templates = template.Must(template.ParseGlob("api/tmpls/*"))
var Templates *template.Template

//ParseTemplate parse all tempaltes from /tmpls folder before every http request
// so every update in template source code is included in response html
func (a *Controllers) ParseTemplates(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Templates = template.Must(template.ParseGlob("api/tmpls/*"))
		next.ServeHTTP(w, r)
	})

}

func (a *Controllers) ServeStatic(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("api/templates"))
	http.StripPrefix("/templates", fs).ServeHTTP(w, r)
	fs.ServeHTTP(w, r)
}
