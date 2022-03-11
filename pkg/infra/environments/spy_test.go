package environments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConfigureSpy(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewEnvironmentsSpy()

		sut.On("Configure").Return(nil)

		err := sut.Configure()

		assert.NoError(t, err)
	})
}

func Test_GO_ENVSpy(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewEnvironmentsSpy()

		sut.On("GO_ENV").Return("development")

		result := sut.GO_ENV()

		assert.Equal(t, "development", result)
	})
}

func Test_DEV_ENVSpy(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewEnvironmentsSpy()

		sut.On("DEV_ENV").Return("development")

		result := sut.DEV_ENV()

		assert.Equal(t, "development", result)
	})
}

func Test_STAGING_ENVSpy(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewEnvironmentsSpy()

		sut.On("STAGING_ENV").Return("staging")

		result := sut.STAGING_ENV()

		assert.Equal(t, "staging", result)
	})
}

func Test_PROD_ENVSpy(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewEnvironmentsSpy()

		sut.On("PROD_ENV").Return("production")

		result := sut.PROD_ENV()

		assert.Equal(t, "production", result)
	})
}
