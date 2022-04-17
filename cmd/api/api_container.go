package api

import (
	"github.com/ralvescosta/base/pkg/app/interfaces"
	"github.com/ralvescosta/base/pkg/app/usecases"
	"github.com/ralvescosta/base/pkg/infra/database"
	httpServer "github.com/ralvescosta/base/pkg/infra/http_server"
	"github.com/ralvescosta/base/pkg/infra/logger"
	"github.com/ralvescosta/base/pkg/infra/repositories"
	"github.com/ralvescosta/base/pkg/infra/validator"
	i "github.com/ralvescosta/base/pkg/interfaces"
	gqlPresenters "github.com/ralvescosta/base/pkg/interfaces/graphql/presenters"
	"github.com/ralvescosta/base/pkg/interfaces/graphql/resolvers"
	"github.com/ralvescosta/base/pkg/interfaces/http/factories"
	"github.com/ralvescosta/base/pkg/interfaces/http/handlers"
	"github.com/ralvescosta/base/pkg/interfaces/http/presenters"
)

type HTTPServerContainer struct {
	logger     interfaces.ILogger
	httpServer httpServer.IHTTPServer

	marketsRoutes i.IRoutes
	graphqlRoutes i.IRoutes
}

func NewHTTPContainer(env interfaces.IEnvironments) (HTTPServerContainer, error) {
	logger, err := logger.NewLogger()
	if err != nil {
		return HTTPServerContainer{}, err
	}

	shotdown := make(chan bool)

	db, err := database.Connect(logger, shotdown)
	if err != nil {
		return HTTPServerContainer{}, err
	}

	httpServer := httpServer.NewHTTPServer(env, logger, shotdown)
	vAlidator := validator.NewValidator()
	httpResFactory := factories.NewHttpResponseFactory()
	marketRepository := repositories.NewMarketRepository(logger, db)

	createMarketUseCase := usecases.NewCreateMarketUseCase(marketRepository)
	getByQueryUseCase := usecases.NewGetMarketByQueryUseCase(marketRepository)
	updateMarketUseCase := usecases.NewUpdateMarketUseCase(marketRepository)
	deleteMarketUseCase := usecases.NewDeleteMarketUseCase(marketRepository)
	marketHandlers := handlers.NewMarketHandlers(logger, vAlidator, httpResFactory, createMarketUseCase, getByQueryUseCase, updateMarketUseCase, deleteMarketUseCase)
	marketsRoutes := presenters.NewMarketRoutes(logger, marketHandlers)

	graphqlResolvers := resolvers.NewResolver(createMarketUseCase, getByQueryUseCase, updateMarketUseCase, deleteMarketUseCase)
	graphqlRoutes := gqlPresenters.NewGraphQLRoutes(logger, graphqlResolvers)

	return HTTPServerContainer{
		logger,
		httpServer,

		marketsRoutes,
		graphqlRoutes,
	}, nil
}
