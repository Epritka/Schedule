package zap

import "go.uber.org/zap"

func New(isDebug bool) *zap.SugaredLogger {
	zapOptions := []zap.Option{
		zap.AddCallerSkip(1),
	}
	l, _ := zap.NewProduction(
		zapOptions...,
	)
	if isDebug {
		l, _ = zap.NewDevelopment(
			zapOptions...,
		)
	}
	return l.Sugar()
}
