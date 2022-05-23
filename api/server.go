package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PwrFr/gqlgen/config"
	"github.com/go-chi/chi"
)

const defaultPort = "3001"

func main() {
	router := chi.NewRouter()

	router.Use(config.CorsConfig())

	// c := graph{graph.Resolver: &graph.Resolver{
	// 	RoomRepo: repo.RoomRepo{DB: config.CallDB()},
	// }}

	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	router.Handle("/query", config.ServerConfig())

	err := http.ListenAndServe(":3001", router)
	if err != nil {
		panic(err)
	}
}
