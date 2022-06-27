package main

import (
	//"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*"))

func Start(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello Develoteca")
        templates.ExecuteTemplate(w, "start", nil)
}

func main() {
	http.HandleFunc("/", Start)
	log.Println("Server running!!!")
	http.ListenAndServe(":8080", nil)
}
