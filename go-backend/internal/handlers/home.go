package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/home.html"))
	tmpl.Execute(w, map[string]string{
		"Title": "Welcome",
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}
