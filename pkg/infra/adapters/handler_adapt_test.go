package adapters

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	httpServer "markets/pkg/infra/http_server"
	"markets/pkg/infra/logger"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_HandlerAdapter(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeSut()

		sut.adapt(sut.ctx)

		assert.Equal(t, 1, *sut.handlerCalledTimes)
	})

	t.Run("should return error if some occur when read body", func(t *testing.T) {
		sut := makeSut()

		readAllBody = func(r io.Reader) ([]byte, error) {
			return nil, errors.New("Error")
		}
		sut.logger.On("Error", "[HandlerAdapt] error while read request bytes", []zap.Field(nil))

		sut.adapt(sut.ctx)

		assert.Equal(t, 0, *sut.handlerCalledTimes)
		sut.logger.AssertExpectations(t)
	})
}

type sutReturn struct {
	adapt              gin.HandlerFunc
	logger             *logger.LoggerSpy
	handlerCalledTimes *int
	handlerSpy         func(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	request            *http.Request
	ctx                *gin.Context
}

func makeSut() sutReturn {
	handlerCalledTimes := 0
	handlerSpy := func(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
		handlerCalledTimes++
		return httpServer.HttpResponse{}
	}

	logger := logger.NewLoggerSpy()

	req := &http.Request{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte(nil))),
		Header: http.Header{
			"op": []string{"op"},
		},
		URL: &url.URL{},
	}

	sut := HandlerAdapt(handlerSpy, logger)

	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())
	contextMock.Params = []gin.Param{{Key: "key", Value: "value"}}
	contextMock.Request = req

	return sutReturn{
		handlerSpy:         handlerSpy,
		handlerCalledTimes: &handlerCalledTimes,
		logger:             logger,
		request:            req,
		adapt:              sut,
		ctx:                contextMock,
	}
}
