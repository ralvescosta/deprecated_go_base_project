package cmd

import (
	"markets/pkg/app/interfaces"
	"markets/pkg/app/usecases"
	"markets/pkg/infra/database"
	httpServer "markets/pkg/infra/http_server"
	"markets/pkg/infra/logger"
	"markets/pkg/infra/repositories"
	"markets/pkg/infra/validator"
	"markets/pkg/interfaces/http/factories"
	"markets/pkg/interfaces/http/handlers"
	"markets/pkg/interfaces/http/presenters"
)

type HTTPServerContainer struct {
	logger     interfaces.ILogger
	httpServer httpServer.IHTTPServer

	marketsRoutes presenters.IRoutes
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
	marketHandlers := handlers.NewMarketHandlers(logger, vAlidator, httpResFactory, createMarketUseCase)
	marketsRoutes := presenters.NewMarketRoutes(logger, marketHandlers)

	return HTTPServerContainer{
		logger,
		httpServer,

		marketsRoutes,
	}, nil
}
