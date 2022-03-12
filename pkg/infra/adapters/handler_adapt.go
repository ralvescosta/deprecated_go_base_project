package adapters

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"markets/pkg/app/interfaces"
	httpServer "markets/pkg/infra/http_server"

	"github.com/gin-gonic/gin"
)

var readAllBody = ioutil.ReadAll

func HandlerAdapt(handler func(httpRequest httpServer.HttpRequest) httpServer.HttpResponse, logger interfaces.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, err := readAllBody(ctx.Request.Body)
		if err != nil {
			logger.Error("[HandlerAdapt] error while read request bytes")
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		params := make(map[string]string)
		for _, param := range ctx.Params {
			params[param.Key] = param.Value
		}

		request := httpServer.HttpRequest{
			Body:    body,
			Headers: ctx.Request.Header,
			Params:  params,
			Query:   ctx.Request.URL.Query(),
			Ctx:     ctx.Request.Context(),
		}

		result := handler(request)

		ctx.JSON(result.StatusCode, result.Body)
	}
}
