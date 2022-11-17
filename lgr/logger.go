package lgr

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// List of levels for Logger.
const (
	InfoLevel  = "INFO"
	DebugLevel = "DEBUG"
)

// Logger represents logger.
type Logger struct{ log *zap.SugaredLogger }

// New initialize logger.
func New(logLevel string) (*Logger, error) {
	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		return nil, fmt.Errorf("error with logger level parsing: %w", err)
	}

	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("сan't build: %w", err)
	}

	return &Logger{logger.Sugar()}, nil
}

// Flush will flush any buffered log entries.
func (l *Logger) Flush() {
	if err := l.log.Sync(); err != nil {
		l.log.Error(err)
	}
}

// Methods above will implement all needful logging behavior.
func (l *Logger) Errorf(format string, val ...any) {
	l.log.Errorf(format, val...)
}

func (l *Logger) Errorw(format string, val ...any) {
	l.log.Errorw(format, val...)
}

func (l *Logger) Debugf(format string, val ...any) {
	l.log.Debugf(format, val...)
}

func (l *Logger) Debugw(format string, val ...any) {
	l.log.Debugw(format, val...)
}

func (l *Logger) Infof(format string, val ...any) {
	l.log.Infof(format, val...)
}

func (l *Logger) Warnf(format string, val ...any) {
	l.log.Warnf(format, val...)
}

func (l *Logger) Warnw(format string, val ...any) {
	l.log.Warnw(format, val...)
}

func (l *Logger) Infow(msg string, keyVal ...any) {
	l.log.Infow(msg, keyVal...)
}
