package httpServer

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Default(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewHTTPServerSpy()

		sut.Default()
	})
}

func Test_RegisterRoute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewHTTPServerSpy()

		method := "POST"
		path := "/api"
		handlers := []gin.HandlerFunc{func(ctx *gin.Context) {}}

		sut.On("RegisterRoute", method, path).Return(nil)

		error := sut.RegisterRoute(method, path, handlers...)

		assert.NoError(t, error)
	})
}

func Test_Setup(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewHTTPServerSpy()

		sut.Setup()
	})
}

func Test_RunSpy(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewHTTPServerSpy()

		sut.On("Run").Return(nil)

		err := sut.Run()

		assert.NoError(t, err)
	})
}
