package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"markets/pkg/app/errors"
	"markets/pkg/app/interfaces"
)

func NewLogger() (interfaces.ILogger, error) {
	goEnv := os.Getenv("GO_ENV")

	zapLogLevel := getLogLevel()

	logFile := os.Getenv("LOG_FILE")
	if logFile == "" {
		return nil, errors.NewInternalError("missing log file configuration")
	}
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	if goEnv == "production" || goEnv == "staging" {
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		fileEncoder := zapcore.NewJSONEncoder(config)

		core := zapcore.NewCore(fileEncoder, zapcore.AddSync(f), zapLogLevel)
		return zap.New(core), nil
	}

	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), zapLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapLogLevel),
	)
	return zap.New(core), nil
}

func getLogLevel() zapcore.Level {
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}
