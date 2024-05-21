package middlewares

import (
	"fmt"
	"net/http"
)

func WithLogging(f http.HandlerFunc) http.HandlerFunc {
	// adapter layer that is used to return the same need signature
	// and this is because i want to di extra stuff so its a way lets say to trick the system
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Methode: ", r.Method)
		fmt.Println("Host: ", r.Host)
		fmt.Println("URL: ", r.URL)
		f.ServeHTTP(w, r)

	}
}
