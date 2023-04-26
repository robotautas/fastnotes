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
	http.HandleFunc("/", index)
	http.ListenAndServe(":7890", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")
	if err != nil {
		log.Panic(err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Panic(err)
	}
}
