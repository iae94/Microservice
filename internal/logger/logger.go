package logger

import (
	serviceConfig "github.com/iae94/Microservice/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func CreateLogger(config *serviceConfig.Config) (logger *zap.Logger, err error) {

	var level zap.AtomicLevel
	switch config.Logger.Level {
	case "debug":
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warning":
		level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	zapConfig := zap.Config{
		Encoding:         config.Logger.Encoding,
		Level:            level,
		OutputPaths:      config.Logger.OutputPaths,
		ErrorOutputPaths: config.Logger.ErrorOutputPaths,
		EncoderConfig: zapcore.EncoderConfig{
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			LevelKey:     "level",
			MessageKey:   "message",
		},
	}

	logger, err = zapConfig.Build()
	if err != nil {
		log.Printf("Logger config build error: %v \n", err)
		return nil, err
	}
	return logger, nil

}
