package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type note struct {
	title string
	body  string
	tags  []tag
}

type tag struct {
	title string
	notes []note
}

func main() {
	fmt.Println("Work in progress :)")
	http.HandleFunc("/", notes)
	http.ListenAndServe(":7890", nil)
}

func notes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Panic(err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Panic(err)
	}
}
