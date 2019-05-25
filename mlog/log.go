package mlog

import (
	"go.uber.org/zap/zapcore"
)

// Type and function aliases from zap to limit the libraries scope into MM code
type Field = zapcore.Field
