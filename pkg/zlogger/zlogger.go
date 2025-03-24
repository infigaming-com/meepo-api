package zlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZLogger(level string) (*zap.Logger, func(), error) {
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	zapCfg := zap.NewProductionConfig()
	zapCfg.Level = zap.NewAtomicLevelAt(zapLevel)
	zapCfg.EncoderConfig.CallerKey = "ln"
	zapCfg.EncoderConfig.FunctionKey = "fn"
	zapCfg.EncoderConfig.LevelKey = "severity"
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapCfg.OutputPaths = []string{"stdout"}

	lg, err := zapCfg.Build()
	if err != nil {
		return nil, nil, err
	}

	return lg, func() {
		lg.Sync()
	}, nil
}
