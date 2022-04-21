package presenters

import (
	"github.com/ralvescosta/base/pkg/app/interfaces"
	"github.com/ralvescosta/base/pkg/infra/adapters"
	httpServer "github.com/ralvescosta/base/pkg/infra/http_server"
	"github.com/ralvescosta/base/pkg/interfaces/http/handlers"
)

type marketRoutes struct {
	logger   interfaces.ILogger
	handlers handlers.IMarketHandlers
}

func (pst marketRoutes) Register(httpServer httpServer.IHTTPServer) {
	httpServer.RegisterRoute("POST", "/api/v1/markets", adapters.HandlerAdapt(pst.handlers.Create, pst.logger))
	httpServer.RegisterRoute("GET", "/api/v1/markets", adapters.HandlerAdapt(pst.handlers.GetByQuery, pst.logger))
	httpServer.RegisterRoute("PATCH", "/api/v1/markets/:registerCode", adapters.HandlerAdapt(pst.handlers.Update, pst.logger))
	httpServer.RegisterRoute("DELETE", "/api/v1/markets/:registerCode", adapters.HandlerAdapt(pst.handlers.Delete, pst.logger))
}

func NewMarketRoutes(logger interfaces.ILogger, handlers handlers.IMarketHandlers) IRoutes {
	return marketRoutes{
		logger,
		handlers,
	}
}
