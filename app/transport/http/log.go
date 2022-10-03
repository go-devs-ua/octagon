package http

import (
	"io"
	"log"
	"os"
)

const logPath = "./log.txt"

type Log struct{}

func NewLogger() Logger {
	return &Log{}
}

// Printf print to the stdout and logfile
func (l *Log) Printf(fmt string, val ...any) {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening log file: %v", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error closing log file: %v", err)
		}
	}()

	log.SetOutput(io.MultiWriter(os.Stdout, file))
	log.SetFlags(log.LstdFlags)
	log.Printf(fmt, val...)
}
