package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_Logger(t *testing.T) {
	t.Run("should create development logger correctly", func(t *testing.T) {
		os.Setenv("LOG_LEVEL", "debug")
		os.Setenv("GO_ENV", "development")
		logger, err := NewLogger()

		assert.NoError(t, err)
		assert.IsType(t, &zap.Logger{}, logger)
	})

	t.Run("should create production logger correctly", func(t *testing.T) {
		os.Setenv("LOG_LEVEL", "warn")
		os.Setenv("GO_ENV", "production")
		logger, err := NewLogger()

		assert.NoError(t, err)
		assert.IsType(t, &zap.Logger{}, logger)
	})

	t.Run("should create logger with different logger level", func(t *testing.T) {
		os.Setenv("LOG_LEVEL", "info")
		os.Setenv("GO_ENV", "development")

		logger, err := NewLogger()

		assert.NoError(t, err)
		assert.IsType(t, &zap.Logger{}, logger)

		os.Setenv("LOG_LEVEL", "error")

		logger, err = NewLogger()

		assert.NoError(t, err)
		assert.IsType(t, &zap.Logger{}, logger)

		os.Setenv("LOG_LEVEL", "panic")

		logger, err = NewLogger()

		assert.NoError(t, err)
		assert.IsType(t, &zap.Logger{}, logger)

		os.Setenv("LOG_LEVEL", "wrong log level")

		logger, err = NewLogger()

		assert.NoError(t, err)
		assert.IsType(t, &zap.Logger{}, logger)
	})
}
