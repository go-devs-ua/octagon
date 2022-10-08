package lgr

import (
	"log"

	"go.uber.org/zap"
)

// Logger represents logger.
type Logger struct{ log *zap.SugaredLogger }

// New initialize logger
func New() *Logger {
	// TODO: Extend config and make adjustments
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("error creating new logger: %v", err)
	}

	return &Logger{logger.Sugar()}
}

// Flush will flush any buffered log entries.
func (l *Logger) Flush() {
	if err := l.log.Sync(); err != nil {
		l.log.Error(err)
	}
}

func (l *Logger) Errorf(format string, val ...any) {
	l.log.Errorf(format, val...)
}

func (l *Logger) Debugf(format string, val ...any) {
	l.log.Debugf(format, val...)
}

func (l *Logger) Infof(format string, val ...any) {
	l.log.Infof(format, val...)
}

func (l *Logger) Warnf(format string, val ...any) {
	l.log.Warnf(format, val...)
}
