package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PwrFr/gqlgen/config"
	"github.com/go-chi/chi"
)

const defaultPort = "3001"

func main() {
	router := chi.NewRouter()

	router.Use(config.CorsConfig())

	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", config.ServerConfig())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
