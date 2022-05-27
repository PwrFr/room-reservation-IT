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

	// x, _ := middleware.GenerateJwtToken("student")
	// fmt.Println(x)

	// x := middleware.RoleValid("student", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM3NTQyNzMsInN1YiI6InN0dWRlbnQifQ.U-TRxnKR_lROvO_1bglpCA4K05NsrDS64sJ8cn0qSd4")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
