package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	topicDebugging := zapcore.AddSync(io.Discard)
	topicErrors := zapcore.AddSync(io.Discard)

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	c := zapcore.NewTee(
		zapcore.NewCore(enc, topicErrors, highPriority),
		zapcore.NewCore(enc, consoleErrors, highPriority),
		zapcore.NewCore(enc, topicDebugging, lowPriority),
		zapcore.NewCore(enc, consoleDebugging, lowPriority),
	)

	return zap.New(c)
}
