package service

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type formatter struct {
	fields log.Fields
	lf     log.Formatter
}

func (f *formatter) Format(e *log.Entry) ([]byte, error) {
	for k, v := range f.fields {
		if _, ok := e.Data[k]; !ok {
			e.Data[k] = v
		}
	}

	return f.lf.Format(e)
}

func setFormatter(service string) {
	var err error
	var logLevel log.Level

	if logLevel, err = log.ParseLevel(os.Getenv("LOG_LEVEL")); err != nil {
		logLevel = log.DebugLevel
	}

	customFormatter := &log.JSONFormatter{}
	customFormatter.TimestampFormat = time.RFC3339Nano

	f := &formatter{
		fields: log.Fields{
			"service":   service,
			"span.kind": "server",
			"system":    "system",
		},
		lf: customFormatter,
	}

	log.SetFormatter(f)
	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel)
}
