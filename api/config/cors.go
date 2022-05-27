package config

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

var host = Env("FRONTEND_IP")

func CorsConfig() func(http.Handler) http.Handler {

	fmt.Println("fe.host: ", host)
	return (cors.New(cors.Options{
		AllowedOrigins:   []string{host},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
}
