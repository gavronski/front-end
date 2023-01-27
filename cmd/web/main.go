package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var port string = "80"

type Note struct {
	Name            string
	Description     string
	TextColor       string
	BackgroundColor string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		notes := []Note{
			{
				Name:            "Test1",
				Description:     "Description1",
				TextColor:       "#FDF6F2",
				BackgroundColor: "#4466EE",
			},
			{
				Name:            "Test1",
				Description:     "Description1",
				TextColor:       "#FDF6F2",
				BackgroundColor: "#44EE47",
			},
			{
				Name:            "Test1",
				Description:     "Description1",
				TextColor:       "#FDF6F2",
				BackgroundColor: "#EE6644",
			},
			{
				Name:            "Test1",
				Description:     "Description1",
				TextColor:       "#FDF6F2",
				BackgroundColor: "#E83092",
			},
			{
				Name:            "Test1",
				Description:     "Description1",
				TextColor:       "#FDF6F2",
				BackgroundColor: "#4466EE",
			},
			{
				Name:            "Test1",
				Description:     "Description1",
				TextColor:       "#FDF6F2",
				BackgroundColor: "#9A961D",
			},
		}
		render(w, "body.page.gohtml", notes)
	})

	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {

		var note Note = Note{
			Name:            "test",
			Description:     "test descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest descriptiontest description",
			TextColor:       "#FDF6F2",
			BackgroundColor: "#EE8044",
		}

		render(w, "notes.page.gohtml", note)
	})

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, layout string, td any) {
	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	var templateSlice []string

	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", layout))
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, td); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
