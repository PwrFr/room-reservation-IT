package config

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/PwrFr/gqlgen/graph"
	"github.com/PwrFr/gqlgen/graph/generated"
	repo "github.com/PwrFr/gqlgen/repository"
	"github.com/gorilla/websocket"
)

func ServerConfig() *handler.Server {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					RoomRepo: repo.RoomRepo{DB: db},
				}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return r.Host == "example.org"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	return srv
}
