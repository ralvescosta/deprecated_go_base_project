package graphqlserver

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
)

type IGraphqlServer interface {
	Default()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type IGqlHandlerServer interface {
	AddTransport(graphql.Transport)
	SetQueryCache(graphql.Cache)
	Use(graphql.HandlerExtension)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type graphqlServer struct {
	srv IGqlHandlerServer
}

var (
	ws = &transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
	}
)

func (pst graphqlServer) Default() {
	ws.InitFunc = pst.initFunc
	pst.srv.AddTransport(ws)
	pst.srv.AddTransport(transport.Options{})
	pst.srv.AddTransport(transport.GET{})
	pst.srv.AddTransport(transport.POST{})
	pst.srv.AddTransport(transport.MultipartForm{})

	pst.srv.SetQueryCache(lru.New(100))
	pst.srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	pst.srv.Use(extension.Introspection{})
}

func (pst graphqlServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pst.srv.ServeHTTP(w, r)
}

func (pst graphqlServer) initFunc(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
	log.Println("WS::INITFUNC")
	log.Println(initPayload)
	return ctx, nil
}

func NewGraphQLServer(srv IGqlHandlerServer) IGraphqlServer {
	return graphqlServer{srv}
}
