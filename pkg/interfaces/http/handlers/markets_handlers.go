package handlers

import (
	"markets/pkg/app/interfaces"
	httpServer "markets/pkg/infra/http_server"
	"markets/pkg/interfaces/http/factories"
)

type IMarketHandlers interface {
	Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	FindById(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	FindByQuery(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	Update(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	Delete(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
}

type marketHandlers struct {
	logger         interfaces.ILogger
	validator      interfaces.IValidator
	httpResFactory factories.HttpResponseFactory
}

func (pst marketHandlers) Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.httpResFactory.Created(nil, nil)
}

func (pst marketHandlers) FindById(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.httpResFactory.Ok(nil, nil)
}

func (pst marketHandlers) FindByQuery(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.httpResFactory.Ok(nil, nil)
}

func (pst marketHandlers) Update(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.httpResFactory.Ok(nil, nil)
}

func (pst marketHandlers) Delete(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.httpResFactory.Ok(nil, nil)
}

func NewMarketHandlers(logger interfaces.ILogger, validator interfaces.IValidator, httpResFactory factories.HttpResponseFactory) IMarketHandlers {

	return marketHandlers{
		logger,
		validator,
		httpResFactory,
	}
}
