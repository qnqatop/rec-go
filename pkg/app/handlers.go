package app

import (
	"html/template"
	"log"
	"net/http"

	"rec-start/pkg/recomendation"
)

func (a *App) runHTTPServer() error {

	log.Println("Сервер запущен на http://localhost:8080")
	return http.ListenAndServe(":8080", nil)
}

func (a *App) registerHandlers() {
	http.HandleFunc("/", ShowFormHandler)
	http.HandleFunc("/recommendations", recomendation.Rec(a.db))
}

func ShowFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./pkg/template/from.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки формы", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
