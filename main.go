package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	http.HandleFunc("/", helloworld)
	http.HandleFunc("/add-film/", addfilm)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
}

func addfilm(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}
