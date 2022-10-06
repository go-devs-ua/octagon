package logger

import (
	"log"

	"go.uber.org/zap"
)

type Wrapper struct{ log *zap.SugaredLogger }

func NewLogger() *Wrapper {
	// TODO: Extend config and make adjustments
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("error creating new logger: %v", err)
	}

	return &Wrapper{logger.Sugar()}
}

func (w *Wrapper) Flush() {
	if err := w.log.Sync(); err != nil {
		w.log.Error(err)
	}
}

// Methods above will be implemented on transport layer according to http.Logger

func (w *Wrapper) Errorf(format string, val ...any) {
	w.log.Errorf(format, val)
}

func (w *Wrapper) Debugf(format string, val ...any) {
	w.log.Errorf(format, val)
}

func (w *Wrapper) Infof(format string, val ...any) {
	w.log.Errorf(format, val)
}

func (w *Wrapper) Warnf(format string, val ...any) {
	w.log.Errorf(format, val)
}
