package main

import (
	"fmt"
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
