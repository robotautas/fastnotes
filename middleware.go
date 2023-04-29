package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// minimal example FOR REFERENCE
// func logHits(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		message := fmt.Sprintf("Endpoint %s served!", r.URL)
// 		fmt.Println(message)
// 		next.ServeHTTP(w, r)
// 	})
// }

func csrfProtect(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
