package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
)

type Note struct {
	Title string
	Body  string
	Tags  []Tag
}

type Tag struct {
	Title string
	Notes []Note
}

type Page struct {
	Notes []Note
}

type Login struct {
	Email    string
	Password string
}

const IP = "127.0.0.1"
const port = ":7890"

func main() {
	r := chi.NewRouter()
	fmt.Println(r)

	fmt.Printf("Serving Fastnotes on http://%s%s\n", IP, port)

	r.Get("/", index)
	r.Route("/login", func(r chi.Router) {
		r.Get("/", loginGet)
		r.Post("/", loginPost)
	})

	http.ListenAndServe(IP+port, r)
}

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
