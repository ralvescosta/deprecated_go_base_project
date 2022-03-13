package presenters

import (
	httpServer "markets/pkg/infra/http_server"
	"markets/pkg/infra/logger"
	"markets/pkg/interfaces/http/handlers"
	"testing"
)

func Test_Market_Register(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeMarketsPresentersSut()

		sut.handlers.On("Create").Return(httpServer.HttpResponse{})
		sut.handlers.On("GetByQuery").Return(httpServer.HttpResponse{})
		sut.handlers.On("Update").Return(httpServer.HttpResponse{})
		sut.handlers.On("Delete").Return(httpServer.HttpResponse{})
		sut.server.On("RegisterRoute", "POST", "/api/v1/markets").Return(nil)
		sut.server.On("RegisterRoute", "GET", "/api/v1/markets").Return(nil)
		sut.server.On("RegisterRoute", "PATCH", "/api/v1/markets/:registerCode").Return(nil)
		sut.server.On("RegisterRoute", "DELETE", "/api/v1/markets/:registerCode").Return(nil)

		sut.routes.Register(sut.server)

		sut.server.AssertExpectations(t)
	})
}

type marketsPresentersSutRtn struct {
	logger   *logger.LoggerSpy
	handlers *handlers.MarketsHandlersSpy
	server   *httpServer.HTTPServerSpy
	routes   IRoutes
}

func makeMarketsPresentersSut() marketsPresentersSutRtn {
	logger := logger.NewLoggerSpy()
	handlers := handlers.NewMarketsHandlersSpy()
	server := httpServer.NewHTTPServerSpy()

	routes := NewMarketRoutes(logger, handlers)

	return marketsPresentersSutRtn{logger, handlers, server, routes}
}
