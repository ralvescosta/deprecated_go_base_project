package interfaces

import "go.uber.org/zap"

type LogField struct {
	Key   string
	Value interface{}
}

type ILogger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}
