package handlers

import (
	"encoding/json"

	"markets/pkg/app/interfaces"
	"markets/pkg/domain/usecases"
	httpServer "markets/pkg/infra/http_server"
	"markets/pkg/interfaces/http/factories"
	viewmodels "markets/pkg/interfaces/http/view_models"
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
	createUseCase  usecases.ICreateMarketUseCase
}

func (pst marketHandlers) Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	vModel := viewmodels.MarketViewModel{}
	if err := json.Unmarshal(httpRequest.Body, &vModel); err != nil {
		return pst.httpResFactory.BadRequest("body is required", nil)
	}

	if validationErrs := pst.validator.ValidateStruct(vModel); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return pst.httpResFactory.BadRequest(validationErrs[0].Message, nil)
	}

	result, alreadyCreated, err := pst.createUseCase.Execute(httpRequest.Ctx, vModel.ToValueObject())
	if err != nil {
		return pst.httpResFactory.ErrorResponseMapper(err, nil)
	}
	if alreadyCreated {
		return pst.httpResFactory.Ok(viewmodels.NewMarketViewModel(result), nil)
	}

	return pst.httpResFactory.Created(viewmodels.NewMarketViewModel(result), nil)
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

func NewMarketHandlers(logger interfaces.ILogger, validator interfaces.IValidator, httpResFactory factories.HttpResponseFactory,
	createUseCase usecases.ICreateMarketUseCase) IMarketHandlers {

	return marketHandlers{
		logger,
		validator,
		httpResFactory,
		createUseCase,
	}
}
