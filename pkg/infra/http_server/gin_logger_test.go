package httpServer

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"markets/pkg/infra/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_GinLogger(t *testing.T) {
	t.Run("should execute logger correctly", func(t *testing.T) {
		sut := makeGinLoggerSut()

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
					String: headerToString(sut.ginCtx.Request.Header),
				},
				{
					Key:    "request",
					Type:   zapcore.StringType,
					String: string([]byte(nil)),
				},
				{
					Key:    "response",
					Type:   zapcore.StringType,
					String: string([]byte(nil)),
				},
			},
		)

		GinLogger(sut.logger)(sut.ginCtx)
	})

	t.Run("should execute response writer correctly", func(t *testing.T) {
		sut := makeGinLoggerSut()
		resBodyWriter := responseBodyWriter{
			sut.ginCtx.Writer,
			bytes.NewBuffer([]byte("")),
		}
		result := []byte{}
		_, err := resBodyWriter.Write(result)

		assert.NoError(t, err)
	})
}

type ginLoggerSutRtn struct {
	logger      *logger.LoggerSpy
	ginCtx      *gin.Context
	requestBody io.ReadCloser
	time        time.Time
}

func makeGinLoggerSut() ginLoggerSutRtn {
	logger := logger.NewLoggerSpy()
	ginCtx, _ := gin.CreateTestContext(httptest.NewRecorder())
	requestBody := ioutil.NopCloser(bytes.NewBuffer([]byte(nil)))

	ginCtx.Request = &http.Request{
		Body: requestBody,
		Header: http.Header{
			"op": []string{"op"},
		},
		Method: "POST",
	}
	ginCtx.Params = []gin.Param{{Key: "key", Value: "value"}}
	t := time.Now()
	now = func() time.Time {
		return t
	}

	return ginLoggerSutRtn{logger, ginCtx, requestBody, t}
}
