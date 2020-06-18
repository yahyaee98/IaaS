package log

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	zap   *zap.Logger
	sugar *zap.SugaredLogger
}

func newZapLogger() *zapLogger {
	logger, _ := zap.NewProduction(
		zap.AddCallerSkip(1),
	)

	sugar := logger.Sugar()

	return &zapLogger{
		zap:   logger,
		sugar: sugar,
	}
}

func (zl zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	zl.sugar.Infow(msg, keysAndValues...)
}

func (zl zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	zl.sugar.Errorw(msg, keysAndValues...)
}

func (zl zapLogger) Sync() error {
	return zl.sugar.Sync()
}
