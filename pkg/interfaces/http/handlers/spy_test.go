package handlers

import (
	httpServer "markets/pkg/infra/http_server"
	"testing"
)

func Test_MarketHandlerSpy_Create(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewMarketsHandlersSpy()

		req := httpServer.HttpRequest{}

		sut.On("Create", req).Return(httpServer.HttpResponse{})

		sut.Create(req)

		sut.AssertExpectations(t)
	})
}

func Test_MarketHandlerSpy_GetByQuery(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewMarketsHandlersSpy()

		req := httpServer.HttpRequest{}

		sut.On("GetByQuery", req).Return(httpServer.HttpResponse{})

		sut.GetByQuery(req)

		sut.AssertExpectations(t)
	})
}

func Test_MarketHandlerSpy_Update(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewMarketsHandlersSpy()

		req := httpServer.HttpRequest{}

		sut.On("Update", req).Return(httpServer.HttpResponse{})

		sut.Update(req)

		sut.AssertExpectations(t)
	})
}

func Test_MarketHandlerSpy_Delete(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewMarketsHandlersSpy()

		req := httpServer.HttpRequest{}

		sut.On("Delete", req).Return(httpServer.HttpResponse{})

		sut.Delete(req)

		sut.AssertExpectations(t)
	})
}
