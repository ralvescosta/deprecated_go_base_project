package httpServer

import (
	"context"
	"net/http"
)

type HttpResponse struct {
	StatusCode int
	Body       interface{}
	Headers    http.Header
}

type HttpRequest struct {
	Body    []byte
	Headers http.Header
	Params  map[string]string
	Auth    interface{}
	Ctx     context.Context
}
