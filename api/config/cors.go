package config

import (
	"fmt"
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
	fmt.Println("fe.host: ", host)
	return (cors.New(cors.Options{
		AllowedOrigins:   []string{host},
		AllowedMethods:   []string{"GET, POST, PUT, DELETE"},
		AllowedHeaders:   []string{"Origin, Content-Type, Accept, Authorization"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
}
