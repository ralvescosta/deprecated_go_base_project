package presenters

import (
	"markets/pkg/app/interfaces"
	"markets/pkg/infra/adapters"
	httpServer "markets/pkg/infra/http_server"
	"markets/pkg/interfaces/http/handlers"
)

type marketRoutes struct {
	logger   interfaces.ILogger
	handlers handlers.IMarketHandlers
}

func (pst marketRoutes) Register(httpServer httpServer.IHTTPServer) {
	httpServer.RegisterRoute("POST", "/api/v1/markets", adapters.HandlerAdapt(pst.handlers.Create, pst.logger))
	//@TODO: verify if we need something in path to allow the query params
	httpServer.RegisterRoute("GET", "/api/v1/markets", adapters.HandlerAdapt(pst.handlers.GetByQuery, pst.logger))
	//@TODO: verify the vest approach to update (HTTP Verb and the id parameter)
	httpServer.RegisterRoute("PATCH", "/api/v1/markets/:id", adapters.HandlerAdapt(pst.handlers.Update, pst.logger))
	httpServer.RegisterRoute("DELETE", "/api/v1/markets/:registerCode", adapters.HandlerAdapt(pst.handlers.Delete, pst.logger))
}

func NewMarketRoutes(logger interfaces.ILogger, handlers handlers.IMarketHandlers) IRoutes {
	return marketRoutes{
		logger,
		handlers,
	}
}
