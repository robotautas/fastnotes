package main

// middleware for reference
// func logHits(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		message := fmt.Sprintf("Endpoint %s served!", r.URL)
// 		fmt.Println(message)
// 		next.ServeHTTP(w, r)
// 	})
// }
