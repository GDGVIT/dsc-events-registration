package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

var whiteList = []string{
	"https://womentechies.dscvit.com",
	"https://solutions.dscvit.com",
	"http://127.0.0.1:8080",
	"http://localhost:8080",
}

func CorsEveryWhere(mux http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:     whiteList,
		AllowCredentials:   false,
		Debug:              false,
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		OptionsPassthrough: true,
		AllowedHeaders:     []string{"*"},
	})
	return c.Handler(mux)
}
