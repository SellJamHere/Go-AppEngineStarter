package handlers

import (
	"net/http"

	"github.com/codegangsta/martini"
)

func Origin(allowedOrigins []string) martini.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigins != nil && len(allowedOrigins) != 0 {
			if currentOrigin := r.Header.Get("Origin"); currentOrigin != "" && originAllowed(allowedOrigins, currentOrigin) {
				w.Header().Add("Access-Control-Allow-Credentials", "true")
				w.Header().Add("Access-Control-Allow-Origin", currentOrigin)
			}
		}
	}
}

func originAllowed(allowedOrigins []string, currentOrigin string) bool {
	if allowedOrigins[0] == "*" {
		return true
	} else {
		for _, origin := range allowedOrigins {
			if origin == currentOrigin {
				return true
			}
		}
	}

	return false
}
