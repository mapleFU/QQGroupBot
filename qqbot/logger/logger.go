package logger

import (
	"go.uber.org/zap"
)

var SLogger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	SLogger = logger.Sugar()
}
