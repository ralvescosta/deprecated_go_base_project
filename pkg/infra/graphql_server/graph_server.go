package graphqlserver

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
)

type IGraphqlServer interface {
	Default()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type graphqlServer struct {
	srv *handler.Server
}

func (pst graphqlServer) Default() {
	pst.srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			log.Println("WS::INITFUNC")
			log.Println(initPayload)
			return ctx, nil
		},
	})

	pst.srv.AddTransport(transport.Options{})
	pst.srv.AddTransport(transport.GET{})
	pst.srv.AddTransport(transport.POST{})
	pst.srv.AddTransport(transport.MultipartForm{})

	pst.srv.SetQueryCache(lru.New(100))

	pst.srv.Use(extension.Introspection{})
	pst.srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
}

func (pst graphqlServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pst.srv.ServeHTTP(w, r)
}

func NewGraphQLServer(srv *handler.Server) IGraphqlServer {
	return graphqlServer{srv}
}
