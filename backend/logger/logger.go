package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	env := os.Getenv("ENV")
	var err error
	if env == "production" {
		Logger, err = zap.NewProduction()
	} else {
		config := zap.Config{
			Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
			Development:      false,
			Encoding:         "console",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		Logger, err = config.Build()
	}
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}

func GetLogger() *zap.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}
