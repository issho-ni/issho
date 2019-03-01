package issho

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.DebugLevel
	}

	customFormatter := &log.JSONFormatter{}
	customFormatter.TimestampFormat = time.RFC3339Nano

	log.SetFormatter(customFormatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(logLevel)
}
