package handlers

import (
	"github.com/stretchr/testify/mock"

	httpServer "github.com/ralvescosta/base/pkg/infra/http_server"
)

type MarketsHandlersSpy struct {
	mock.Mock
}

func (pst MarketsHandlersSpy) Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	args := pst.Called(httpRequest)

	return args.Get(0).(httpServer.HttpResponse)
}
func (pst MarketsHandlersSpy) GetByQuery(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	args := pst.Called(httpRequest)

	return args.Get(0).(httpServer.HttpResponse)
}
func (pst MarketsHandlersSpy) Update(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	args := pst.Called(httpRequest)

	return args.Get(0).(httpServer.HttpResponse)
}
func (pst MarketsHandlersSpy) Delete(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	args := pst.Called(httpRequest)

	return args.Get(0).(httpServer.HttpResponse)
}

func NewMarketsHandlersSpy() *MarketsHandlersSpy {
	return new(MarketsHandlersSpy)
}
