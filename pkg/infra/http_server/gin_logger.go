package httpServer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"markets/pkg/app/interfaces"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

var now = time.Now

func GinLogger(logger interfaces.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
		ctx.Writer = w

		startTime := now()
		ctx.Next()
		endTime := now()

		latencyTimeInMileseconds := float64(endTime.Sub(startTime).Nanoseconds() / 1000)

		requestBody, _ := ioutil.ReadAll(ctx.Request.Body)
		responseBody, _ := ioutil.ReadAll(w.body)

		logger.Info("[HTTP Request]",
			zapcore.Field{
				Key:    "method",
				Type:   zapcore.StringType,
				String: ctx.Request.Method,
			},
			zapcore.Field{
				Key:    "uri",
				Type:   zapcore.StringType,
				String: ctx.Request.RequestURI,
			},
			zapcore.Field{
				Key:     "statusCode",
				Type:    zapcore.Int64Type,
				Integer: int64(ctx.Writer.Status()),
			},
			zapcore.Field{
				Key:    "latencyTime",
				Type:   zapcore.StringType,
				String: fmt.Sprintf("%.2f us", latencyTimeInMileseconds),
			},
			zapcore.Field{
				Key:    "headers",
				Type:   zapcore.StringType,
				String: headerToString(ctx.Request.Header),
			},
			zapcore.Field{
				Key:    "request",
				Type:   zapcore.StringType,
				String: string(requestBody),
			},
			zapcore.Field{
				Key:    "response",
				Type:   zapcore.StringType,
				String: string(responseBody),
			},
		)
	}
}

func headerToString(header http.Header) string {
	h := ""
	for k, v := range header {
		h += k + ":" + strings.Join(v, ",") + ";"
	}

	return h
}
