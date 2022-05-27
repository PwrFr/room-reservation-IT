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
		AllowedOrigins:   []string{Env("FRONTEND_IP")},
		AllowedMethods:   []string{"GET, POST, PUT, DELETE"},
		AllowedHeaders:   []string{"Origin, Content-Type, Accept, Authorization"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
}
