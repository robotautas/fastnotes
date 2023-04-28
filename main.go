package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const IP = "127.0.0.1"
const port = ":7890"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fmt.Println(r)

	fmt.Printf("Serving Fastnotes on http://%s%s\n", IP, port)

	r.Get("/", index)
	r.Route("/login", func(r chi.Router) {
		r.Get("/", loginGet)
		r.Post("/", loginPost)
	})

	http.ListenAndServe(IP+port, r)
}
