package httpServer

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type HTTPServerSpy struct {
	mock.Mock
}

func (pst HTTPServerSpy) Default() {

}

func (pst HTTPServerSpy) RegisterRoute(method string, path string, handlers ...gin.HandlerFunc) error {
	args := pst.Called(method, path)

	return args.Error(0)
}

func (pst HTTPServerSpy) Setup() {

}

func (pst HTTPServerSpy) Run() error {
	args := pst.Called()

	return args.Error(0)
}

func NewHTTPServerSpy() *HTTPServerSpy {
	return new(HTTPServerSpy)
}
