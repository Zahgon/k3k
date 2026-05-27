package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(debug bool, format string) *zap.Logger { _ = "STUB: not implemented"; return nil }

func newEncoder(format string) zapcore.Encoder {
	_ = "STUB: not implemented"
	return *new(zapcore.Encoder)
}
