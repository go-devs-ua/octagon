package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger struct{ log *zap.SugaredLogger }

func New() *Logger {
	// TODO: Extend config and make adjustments
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("error creating new logger: %v", err)
	}

	return &Logger{logger.Sugar()}
}

func (w *Logger) Flush() {
	if err := w.log.Sync(); err != nil {
		w.log.Error(err)
	}
}

func (w *Logger) Errorf(format string, val ...any) {
	w.log.Errorf(format, val)
}

func (w *Logger) Debugf(format string, val ...any) {
	w.log.Debugf(format, val)
}

func (w *Logger) Infof(format string, val ...any) {
	w.log.Infof(format, val)
}

func (w *Logger) Warnf(format string, val ...any) {
	w.log.Warnf(format, val)
}
