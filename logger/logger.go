package logger

import (
	"go.uber.org/zap"
)

// InitLogger initializes the uber-go zap production sugary zap.S().
func InitLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	config.EncoderConfig.StacktraceKey = "stacktrace"
	zapLogger, _ := config.Build()
	//errorLogger = zapzap.S().Sugar()
	return zapLogger
}
