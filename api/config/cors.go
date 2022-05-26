package config

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsConfig() func(http.Handler) http.Handler {
	return (cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
}
