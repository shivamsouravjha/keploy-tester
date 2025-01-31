package utils

import (
	"go.uber.org/zap"
)

func init() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, _ := zapConfig.Build()
	zap.ReplaceGlobals(logger)
	zap.L().Info("Logger initialized in utils") // This should print
}
