package config

import (
	"net/http"
	"os"

	"github.com/rs/cors"
)

var local = Env("FRONTEND_IP")

func CorsConfig() func(http.Handler) http.Handler {
	host := os.Getenv("FRONTEND_HOST")
	if host == "" {
		host = local
	}

	return (cors.New(cors.Options{
		AllowedOrigins:   []string{host},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
}
