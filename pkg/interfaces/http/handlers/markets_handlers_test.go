package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"markets/pkg/app/errors"
	"markets/pkg/app/usecases"
	valueObjects "markets/pkg/domain/value_objects"
	httpServer "markets/pkg/infra/http_server"
	"markets/pkg/infra/logger"
	"markets/pkg/infra/validator"
	"markets/pkg/interfaces/http/factories"
	viewmodels "markets/pkg/interfaces/http/view_models"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func Test_Market_Create(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.validator.On("ValidateStruct", sut.marketViewModelMocked).Return([]valueObjects.ValidateResult(nil))
		sut.createUseCase.On("Execute", sut.createMarketHttpRequest.Ctx, sut.marketViewModelMocked.ToValueObject()).Return(valueObjects.MarketValueObjects{}, false, nil)

		res := sut.handler.Create(sut.createMarketHttpRequest)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		sut.validator.AssertExpectations(t)
		sut.createUseCase.AssertExpectations(t)
	})

	t.Run("should return badRequest if body is no present", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		res := sut.handler.Create(httpServer.HttpRequest{Body: []byte("")})

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should return badRequest if body is unformatted", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.logger.On("Error", "[MarketHandler::Create] - Body unformatted - message", []zapcore.Field(nil))
		sut.validator.On("ValidateStruct", sut.marketViewModelMocked).Return([]valueObjects.ValidateResult{{IsValid: true, Message: "message"}})

		res := sut.handler.Create(sut.createMarketHttpRequest)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		sut.validator.AssertExpectations(t)
	})

	t.Run("should return internalServerError if usecase return internalError", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.validator.On("ValidateStruct", sut.marketViewModelMocked).Return([]valueObjects.ValidateResult(nil))
		sut.createUseCase.On("Execute", sut.createMarketHttpRequest.Ctx, sut.marketViewModelMocked.ToValueObject()).Return(valueObjects.MarketValueObjects{}, false, errors.NewInternalError("some error"))

		res := sut.handler.Create(sut.createMarketHttpRequest)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		sut.validator.AssertExpectations(t)
		sut.createUseCase.AssertExpectations(t)
	})

	t.Run("should return Ok if the market was already created", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.validator.On("ValidateStruct", sut.marketViewModelMocked).Return([]valueObjects.ValidateResult(nil))
		sut.createUseCase.On("Execute", sut.createMarketHttpRequest.Ctx, sut.marketViewModelMocked.ToValueObject()).Return(valueObjects.MarketValueObjects{}, true, nil)

		res := sut.handler.Create(sut.createMarketHttpRequest)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		sut.validator.AssertExpectations(t)
		sut.createUseCase.AssertExpectations(t)
	})
}

func Test_Market_GetByQuey(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.getByQueyUseCase.On(
			"Execute",
			sut.getByQueryHTTPRequest.Ctx,
			viewmodels.MarketViewModel{Bairro: "bairro", NomeFeira: "nomeFeira", Coddist: 10}.ToValueObject(),
		).Return([]valueObjects.MarketValueObjects{{}}, nil)

		res := sut.handler.GetByQuery(sut.getByQueryHTTPRequest)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		sut.getByQueyUseCase.AssertExpectations(t)
	})

	t.Run("should return badRequest if received a invalid query parameter", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.getByQueryHTTPRequest.Query = map[string][]string{"wrong": {"wrong"}}

		res := sut.handler.GetByQuery(sut.getByQueryHTTPRequest)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should internalServerError if usecase return internalError", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.getByQueyUseCase.On(
			"Execute",
			sut.getByQueryHTTPRequest.Ctx,
			viewmodels.MarketViewModel{Bairro: "bairro", NomeFeira: "nomeFeira", Coddist: 10}.ToValueObject(),
		).Return([]valueObjects.MarketValueObjects(nil), errors.NewInternalError(""))

		res := sut.handler.GetByQuery(sut.getByQueryHTTPRequest)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		sut.getByQueyUseCase.AssertExpectations(t)
	})
}

func Test_Market_Update(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.marketViewModelMocked.Registro = ""
		sut.updateUseCase.On("Execute", sut.updateHTTPRequest.Ctx, "registro", sut.marketViewModelMocked.ToValueObject()).Return(valueObjects.MarketValueObjects{}, nil)

		res := sut.handler.Update(sut.updateHTTPRequest)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		sut.updateUseCase.AssertExpectations(t)
	})

	t.Run("should return badRequest if body is unformatted", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.updateHTTPRequest.Body = []byte("")

		res := sut.handler.Update(sut.updateHTTPRequest)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should return badRequest if body contains the filed 'registro'", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		body, _ := json.Marshal(sut.marketViewModelMocked)
		sut.updateHTTPRequest.Body = body

		res := sut.handler.Update(sut.updateHTTPRequest)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should return badRequest if not receive the registerCode parameter", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.updateHTTPRequest.Params = make(map[string]string)

		res := sut.handler.Update(sut.updateHTTPRequest)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should return internalServerError if usecase return internalError", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.marketViewModelMocked.Registro = ""
		sut.updateUseCase.On("Execute", sut.updateHTTPRequest.Ctx, "registro", sut.marketViewModelMocked.ToValueObject()).Return(valueObjects.MarketValueObjects{}, errors.NewInternalError(""))

		res := sut.handler.Update(sut.updateHTTPRequest)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		sut.updateUseCase.AssertExpectations(t)
	})
}

func Test_Market_Delete(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.deleteUseCase.On("Execute", sut.deleteMarketHTTPRequest.Ctx, "registro").Return(nil)

		res := sut.handler.Delete(sut.deleteMarketHTTPRequest)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("should return badRequest if registerCode params is empty", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.deleteMarketHTTPRequest.Params = make(map[string]string)
		sut.deleteUseCase.On("Execute", sut.deleteMarketHTTPRequest.Ctx, "registro").Return(nil)

		res := sut.handler.Delete(sut.deleteMarketHTTPRequest)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should return internalServerError if usecase return internalError", func(t *testing.T) {
		sut := makeMarketHandlersSut()

		sut.deleteUseCase.On("Execute", sut.deleteMarketHTTPRequest.Ctx, "registro").Return(errors.NewInternalError(""))

		res := sut.handler.Delete(sut.deleteMarketHTTPRequest)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}

type marketHandlersSutRtn struct {
	logger                  *logger.LoggerSpy
	validator               *validator.ValidatorSpy
	httpResFactory          factories.HttpResponseFactory
	createUseCase           *usecases.CreateMarketUseCaseSpy
	getByQueyUseCase        *usecases.GetMarketByQueryUseCaseSpy
	updateUseCase           *usecases.UpdateMarketUseCaseSpy
	deleteUseCase           *usecases.DeleteMarketUseCaseSpy
	handler                 IMarketHandlers
	marketViewModelMocked   viewmodels.MarketViewModel
	createMarketHttpRequest httpServer.HttpRequest
	getByQueryHTTPRequest   httpServer.HttpRequest
	updateHTTPRequest       httpServer.HttpRequest
	deleteMarketHTTPRequest httpServer.HttpRequest
}

func makeMarketHandlersSut() marketHandlersSutRtn {
	logger := logger.NewLoggerSpy()
	validator := validator.NewValidatorSpy()
	httpResFactor := factories.NewHttpResponseFactory()
	createUseCase := usecases.NewCreateMarketUseCaseSpy()
	getByQueryUseCase := usecases.NewGetMarketByQueryUseCaseSpy()
	updateUseCase := usecases.NewUpdateMarketUseCaseSpy()
	deleteUseCase := usecases.NewDeleteMarketUseCaseSpy()

	handler := NewMarketHandlers(logger, validator, httpResFactor, createUseCase, getByQueryUseCase, updateUseCase, deleteUseCase)

	marketViewModelMocked := viewmodels.MarketViewModel{
		Long:       -100,
		Lat:        -200,
		Setcens:    "setcens",
		Areap:      "areap",
		Coddist:    10,
		Distrito:   "distrito",
		Codsubpref: 10,
		Subpref:    "subpref",
		Regiao5:    "regiao5",
		Regiao8:    "regiao8",
		NomeFeira:  "nomeFeira",
		Registro:   "registro",
		Logradouro: "logradouro",
		Numero:     "numero",
		Bairro:     "bairro",
		Referencia: "referencia",
	}

	createMarketBody, _ := json.Marshal(marketViewModelMocked)
	createMarketHTTPRequest := httpServer.HttpRequest{
		Ctx:  context.Background(),
		Body: createMarketBody,
	}

	getByQueryHTTPRequest := httpServer.HttpRequest{
		Ctx:   context.Background(),
		Query: map[string][]string{"bairro": {"bairro"}, "nome_feira": {"nomeFeira"}, "coddist": {"10"}},
	}

	c := marketViewModelMocked
	c.Registro = ""
	updateMarketBody, _ := json.Marshal(c)
	updateHTTPRequest := httpServer.HttpRequest{
		Ctx:    context.Background(),
		Body:   updateMarketBody,
		Params: map[string]string{"registerCode": "registro"},
	}

	deleteMarketHTTPRequest := httpServer.HttpRequest{
		Ctx:    context.Background(),
		Params: map[string]string{"registerCode": "registro"},
	}

	return marketHandlersSutRtn{
		logger,
		validator,
		httpResFactor,
		createUseCase,
		getByQueryUseCase,
		updateUseCase,
		deleteUseCase,
		handler,
		marketViewModelMocked,
		createMarketHTTPRequest,
		getByQueryHTTPRequest,
		updateHTTPRequest,
		deleteMarketHTTPRequest,
	}
}
