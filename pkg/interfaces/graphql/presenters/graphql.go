package presenters

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/urfave/cli/v2"

	"github.com/ralvescosta/base/pkg/app/interfaces"
	httpServer "github.com/ralvescosta/base/pkg/infra/http_server"
	i "github.com/ralvescosta/base/pkg/interfaces"
	"github.com/ralvescosta/base/pkg/interfaces/graphql/graph/generated"
	"github.com/ralvescosta/base/pkg/interfaces/graphql/resolvers"
)

type graphqlRoutes struct {
	logger interfaces.ILogger
}

func (pst graphqlRoutes) Register(httpServer httpServer.IHTTPServer) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	httpServer.RegisterRoute(http.MethodPost, "/query", func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	})
	httpServer.RegisterRoute(http.MethodGet, "/", func(ctx *gin.Context) {
		h := playground.Handler("GraphQL playground", "/query")
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})
}

func NewGraphQLRoutes(logger interfaces.ILogger) i.IRoutes {
	return graphqlRoutes{logger}
}
