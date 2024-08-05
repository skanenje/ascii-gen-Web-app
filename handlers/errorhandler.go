package handlers

import (
	"html/template"
	"net/http"
)

func Error404Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/error404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func Error405Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/error405.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func Error500Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/error500.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func Error400Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/error400.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
