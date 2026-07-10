package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	if err := indexTmpl.Execute(w, nil); err != nil {
		log.Printf("template error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
