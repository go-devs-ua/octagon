package lgr

import (
	"log"
	"net/http"

	"go.uber.org/zap"
)

// Logger represents logger.
type Logger struct{ log *zap.SugaredLogger }

// New initialize logger
func New() *Logger {
	// TODO: Extend config and make adjustments
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed creating new logger: %v", err)
	}

	return &Logger{logger.Sugar()}
}

// Flush will flush any buffered log entries.
func (l *Logger) Flush() {
	if err := l.log.Sync(); err != nil {
		l.log.Error(err)
	}
}

// Methods above will implement all needful logging behavior

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

// LogRequest will log http.Request parameters
func (l *Logger) LogRequest(req *http.Request) {
	body, err := req.GetBody()
	if err != nil {
		l.log.Debugf("Failed getting request body while logging: %+v\n", err)
	}

	l.log.Infow("Request",
		"URI", req.RequestURI,
		"Method", req.Method,
		"Header", req.Header,
		"Body", body,
	)
}

func (l *Logger) Infow(msg string, keyVal ...any) {
	l.log.Infow(msg, keyVal...)
}
