package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

var whiteList = []string{
	"https://womentechies.dscvit.com",
	"https://solutions.dscvit.com",
}

func CorsEveryWhere(mux http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   whiteList,
		AllowCredentials: true,
		Debug:            false,
	})
	return c.Handler(mux)
}
