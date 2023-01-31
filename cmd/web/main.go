package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

var port string = "80"

type Note struct {
	ID              int       `json:"note_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	TextColor       string    `json:"text_color"`
	BackgroundColor string    `json:"background_color"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "body.page.gohtml")
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		render(w, "notes.page.gohtml")
	})

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, layout string) {
	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
	}

	var templateSlice []string

	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", layout))
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// nil - place for data - any
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
