package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PwrFr/gqlgen/config"
	"github.com/go-chi/chi"
)

const defaultPort = "3001"

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
	   port = "3001"
	   fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port 
}


func main() {
	router := chi.NewRouter()

	router.Use(config.CorsConfig())

	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", config.ServerConfig())

	err := http.ListenAndServe(getPort(), router)
	if err != nil {
		panic(err)
	}
}
