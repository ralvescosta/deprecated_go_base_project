package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"markets/pkg/app/interfaces"
)

func NewLogger() (interfaces.ILogger, error) {
	goEnv := os.Getenv("GO_ENV")

	zapLogLevel := getLogLevel()

	if goEnv == "production" || goEnv == "staging" {
		return zap.NewProduction(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level.Enabled(zapLogLevel)

	return config.Build()
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
