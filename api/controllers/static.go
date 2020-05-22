package controllers

import (
	"net/http"
	"path"
	"text/template"
)

//var Templates = template.Must(template.ParseGlob("api/tmpls/*"))
var Templates *template.Template

//ParseTemplate parse all tempaltes from /tmpls folder before every http request
// so every update in template source code is included in response html
func (a *Controllers) ParseTemplates(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Templates = template.Must(template.ParseGlob("api/tmpls/*"))
		tmpl, err := template.ParseFiles(path.Join("api/templates", "index.html"))
		Templates = template.Must(tmpl, err)
		next.ServeHTTP(w, r)
	})

}
func (a *Controllers) OptionsCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			next.ServeHTTP(w, r)
		}
		next.ServeHTTP(w, r)
	})
}

func (a *Controllers) ServeStatic(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("api/templates"))
	http.StripPrefix("/", fs).ServeHTTP(w, r)
	fs.ServeHTTP(w, r)
}
