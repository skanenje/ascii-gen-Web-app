package handlers

import (
	"html/template"
	"net/http"
	"os"

	"flags/functions"
)

type PageData struct {
	Result string
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/405", http.StatusSeeOther)
		return
	}

	text := r.FormValue("message")
	banner := r.FormValue("banner-options") + ".txt"

	if text == "" || banner == "" {
		http.Redirect(w, r, "/400", http.StatusSeeOther)
		return
	}

	if err := functions.ContainsNonASCII(text); err != nil {
		http.Redirect(w, r, "/400", http.StatusSeeOther)
		return
	}

	if _, err := os.Stat(banner); os.IsNotExist(err) {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	if err := functions.ValidateFileChecksum(banner); err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	output := functions.FunctionMain(text, banner)
	data := PageData{
		Result: string(output),
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}
