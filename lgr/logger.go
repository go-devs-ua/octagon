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
	l.log.Debugf("===========================Request start=========================\n")
	defer l.log.Debugf("===========================Request end===========================\n")

	l.log.Debugf("URI: %v\n", req.RequestURI)
	l.log.Debugf("Method: %v\n", req.Method)
	l.log.Debugf("Headers: %v\n", req.Header)

	body, err := req.GetBody()
	if err != nil {
		l.log.Debugf("Failed getting request body: %+v\n", err)
		return
	}

	l.log.Debugf("Body: %v\n", body)
}
