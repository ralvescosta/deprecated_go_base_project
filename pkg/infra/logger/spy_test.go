package logger

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func Test_Debug(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewLoggerSpy()

		sut.On("Debug", "[SomeDebug]", []zapcore.Field(nil))

		sut.Debug("[SomeDebug]")

		sut.AssertExpectations(t)
	})
}

func Test_Info(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewLoggerSpy()

		sut.On("Info", "[SomeInfo]", []zapcore.Field(nil))

		sut.Info("[SomeInfo]")

		sut.AssertExpectations(t)
	})
}

func Test_Warn(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewLoggerSpy()

		sut.On("Warn", "[SomeWarn]", []zapcore.Field(nil))

		sut.Warn("[SomeWarn]")

		sut.AssertExpectations(t)
	})
}

func Test_Error(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewLoggerSpy()

		sut.On("Error", "[SomeError]", []zapcore.Field(nil))

		sut.Error("[SomeError]")

		sut.AssertExpectations(t)
	})
}
