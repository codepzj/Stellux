package logger

import "go.uber.org/zap"

var (
	Logger *zap.Logger = zap.NewExample()
)
