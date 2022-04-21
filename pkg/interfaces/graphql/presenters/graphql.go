package presenters

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/urfave/cli/v2"

	"github.com/ralvescosta/base/pkg/app/interfaces"
	graphqlserver "github.com/ralvescosta/base/pkg/infra/graphql_server"
	httpServer "github.com/ralvescosta/base/pkg/infra/http_server"
)

type GraphqlRoutes struct {
	logger interfaces.ILogger
}

func (pst GraphqlRoutes) Register(httpServer httpServer.IHTTPServer, graphqlServer graphqlserver.IGraphqlServer) {
	httpServer.RegisterRoute(http.MethodPost, "/api/gql/query", func(ctx *gin.Context) {
		graphqlServer.ServeHTTP(ctx.Writer, ctx.Request)
	})
	httpServer.RegisterRoute(http.MethodGet, "/api/gql/subscriptions", func(ctx *gin.Context) {
		graphqlServer.ServeHTTP(ctx.Writer, ctx.Request)
	})
	httpServer.RegisterRoute(http.MethodGet, "/api/gql/playground", func(ctx *gin.Context) {
		h := playground.Handler("GraphQL playground", "/api/gql/query")
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})
}

func NewGraphQLRoutes(logger interfaces.ILogger) GraphqlRoutes {
	return GraphqlRoutes{logger}
}
