package middlewares

import (
	"net/http"
)


func WithLogging(f http.Handler) http.HandlerFunc {
	// adapter layer that is used to return the same need signature 
	// and this is because i want to di extra stuff so its a way lets say to trick the system
	return  func(w http.ResponseWriter , r *http.Request){
		f.ServeHTTP(w , r)
	}
}

