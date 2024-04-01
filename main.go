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
