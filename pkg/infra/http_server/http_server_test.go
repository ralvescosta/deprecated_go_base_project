package httpServer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/ralvescosta/base/pkg/infra/environments"
	"github.com/ralvescosta/base/pkg/infra/logger"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_NewHttpServer(t *testing.T) {
	t.Run("should create a new http server correctly", func(t *testing.T) {
		sut := makeHTTPServerSutRtn("POST")

		server := NewHTTPServer(sut.env, sut.logger, sut.shotdown)

		assert.NotNil(t, server)
	})
}

func Test_RegisterRoutes(t *testing.T) {
	t.Run("should configure a POST route", func(t *testing.T) {
		sut := makeHTTPServerSutRtn("POST")
		sut.httpServer.Default()

		sut.spyLogger()
		sut.httpServer.RegisterRoute("POST", "/api/v1/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})

		response, err := sut.doRequest("POST", "/api/v1/test")

		assert.NoError(t, err)
		assert.Equal(t, response.StatusCode, http.StatusOK)
	})

	t.Run("should configure a GET route", func(t *testing.T) {
		sut := makeHTTPServerSutRtn("GET")
		sut.httpServer.Default()
		sut.spyLogger()
		sut.httpServer.RegisterRoute("GET", "/api/v1/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})

		response, err := sut.doRequest("GET", "/api/v1/test")

		assert.NoError(t, err)
		assert.Equal(t, response.StatusCode, http.StatusOK)
	})

	t.Run("should configure a PUT route", func(t *testing.T) {
		sut := makeHTTPServerSutRtn("PUT")
		sut.httpServer.Default()
		sut.spyLogger()
		sut.httpServer.RegisterRoute("PUT", "/api/v1/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})

		response, err := sut.doRequest("PUT", "/api/v1/test")

		assert.NoError(t, err)
		assert.Equal(t, response.StatusCode, http.StatusOK)
	})

	t.Run("should configure a PATCH route", func(t *testing.T) {
		sut := makeHTTPServerSutRtn("PATCH")
		sut.httpServer.Default()
		sut.spyLogger()
		sut.httpServer.RegisterRoute("PATCH", "/api/v1/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})

		response, err := sut.doRequest("PATCH", "/api/v1/test")

		assert.NoError(t, err)
		assert.Equal(t, response.StatusCode, http.StatusOK)
	})

	t.Run("should configure a DELETE route", func(t *testing.T) {
		sut := makeHTTPServerSutRtn("DELETE")
		sut.httpServer.Default()
		sut.spyLogger()
		sut.httpServer.RegisterRoute("DELETE", "/api/v1/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})

		response, err := sut.doRequest("DELETE", "/api/v1/test")

		assert.NoError(t, err)
		assert.Equal(t, response.StatusCode, http.StatusOK)
	})

	t.Run("should return an error if try to register unsupported method", func(t *testing.T) {
		sut := makeHTTPServerSutRtn("Something")
		sut.httpServer.Default()
		sut.spyLogger()
		err := sut.httpServer.RegisterRoute("Something", "/api/v1/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})

		assert.Error(t, err)
	})
}

func Test_Run(t *testing.T) {
	t.Run("should execute Run with SSL correctly", func(t *testing.T) {
		sut := makeHTTPServerSutRtn("POST")
		sut.env.On("GO_ENV").Return("development")
		sut.env.On("PROD_ENV").Return("production")
		sut.env.On("PROFILING_ENV").Return("disabled")
		os.Setenv("HOST", "localhost")
		os.Setenv("PORT", "1111")
		os.Setenv("TLS_CERT_PATH", "../../../pkg/interfaces/http/certs/cert.pem")
		os.Setenv("TLS_KEY_PATH", "../../../pkg/interfaces/http/certs/key.pem")
		sut.httpServer.Default()
		sut.httpServer.Setup()
		sut.logger.On("Info", "[HttpServer::Run] - Server running at: https://localhost:1111", []zap.Field(nil))

		go func(t *testing.T) {
			err := sut.httpServer.Run()
			assert.NoError(t, err)
		}(t)
		time.Sleep(time.Microsecond * 10)
	})

	t.Run("should execute Run without SSL correctly", func(t *testing.T) {
		sut := makeHTTPServerSutRtn("POST")
		sut.env.On("GO_ENV").Return("production")
		sut.env.On("PROD_ENV").Return("production")
		sut.env.On("PROFILING_ENV").Return("enabled")
		os.Setenv("HOST", "localhost")
		os.Setenv("PORT", "2222")
		sut.httpServer.Default()
		sut.httpServer.Setup()
		sut.logger.On("Info", "[HttpServer::Run] - Server running at: http://localhost:2222", []zap.Field(nil))

		go func(t *testing.T) {
			err := sut.httpServer.Run()
			assert.NoError(t, err)
		}(t)
		time.Sleep(time.Microsecond * 10)
	})
}

type httpServerSutRtn struct {
	httpServer HTTPServer
	logger     *logger.LoggerSpy
	env        *environments.EnvironmentsSpy
	shotdown   chan bool
	ginCtx     *gin.Context
}

func makeHTTPServerSutRtn(httpMethod string) httpServerSutRtn {
	ginCtx, _ := gin.CreateTestContext(httptest.NewRecorder())
	requestBody := ioutil.NopCloser(bytes.NewBuffer([]byte(nil)))

	ginCtx.Request = &http.Request{
		Body:   requestBody,
		Method: httpMethod,
	}
	ginCtx.Params = []gin.Param{{Key: "key", Value: "value"}}
	logger := logger.NewLoggerSpy()
	env := environments.NewEnvironmentsSpy()
	shotdown := make(chan bool)
	httpServer := HTTPServer{
		env:      env,
		logger:   logger,
		shotdown: shotdown,
	}
	t := time.Now()
	now = func() time.Time {
		return t
	}

	return httpServerSutRtn{
		httpServer, logger, env, shotdown, ginCtx,
	}
}

func (sut httpServerSutRtn) doRequest(method, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, path, ioutil.NopCloser(bytes.NewBuffer([]byte(nil))))
	if err != nil {
		return nil, err
	}

	w := httptest.NewRecorder()
	sut.httpServer.router.ServeHTTP(w, req)

	return w.Result(), nil
}

func (sut httpServerSutRtn) spyLogger() {
	sut.logger.On(
		"Info",
		"[HTTP Request]",
		[]zap.Field{
			{
				Key:    "method",
				Type:   zapcore.StringType,
				String: sut.ginCtx.Request.Method,
			},
			{
				Key:    "uri",
				Type:   zapcore.StringType,
				String: sut.ginCtx.Request.RequestURI,
			},
			{
				Key:     "statusCode",
				Type:    zapcore.Int64Type,
				Integer: int64(sut.ginCtx.Writer.Status()),
			},
			{
				Key:    "latencyTime",
				Type:   zapcore.StringType,
				String: fmt.Sprintf("%.2f us", 0.0),
			},
			{
				Key:    "headers",
				Type:   zapcore.StringType,
				String: headerToString(http.Header{}),
			},
			{
				Key:    "request",
				Type:   zapcore.StringType,
				String: string([]byte(nil)),
			},
			{
				Key:    "response",
				Type:   zapcore.StringType,
				String: "null",
			},
		},
	)
}
