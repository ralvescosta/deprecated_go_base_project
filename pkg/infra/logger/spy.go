package logger

import (
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type LoggerSpy struct {
	mock.Mock
}

func (pst LoggerSpy) Debug(msg string, fields ...zap.Field) {
	pst.Called(msg, fields)
}
func (pst LoggerSpy) Info(msg string, fields ...zap.Field) {
	pst.Called(msg, fields)
}
func (pst LoggerSpy) Warn(msg string, fields ...zap.Field) {
	pst.Called(msg, fields)
}
func (pst LoggerSpy) Error(msg string, fields ...zap.Field) {
	pst.Called(msg, fields)
}
func NewLoggerSpy() *LoggerSpy {
	return new(LoggerSpy)
}
