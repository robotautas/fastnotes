package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var cfg Config = getConfig("./config.toml")

func main() {
	// session := getSession()
	fmt.Printf("Settings: %v\n", cfg)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(csrfProtect)

	fmt.Printf("Serving Fastnotes on http://%s%s\n", cfg.Server.IP, cfg.Server.Port)

	r.Get("/", index)
	r.Route("/login", func(r chi.Router) {
		r.Get("/", loginGet)
		r.Post("/", loginPost)
	})

	http.ListenAndServe(cfg.Server.IP+cfg.Server.Port, r)
}
