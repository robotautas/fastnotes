package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	p := &mock_notes
	t, _ := template.ParseFiles("templates/base.html", "templates/index.html")
	t.ExecuteTemplate(w, "index.html", p)
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/base.html", "templates/login.html")
	t.ExecuteTemplate(w, "login.html", nil)
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	details := Login{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	fmt.Println(details)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
