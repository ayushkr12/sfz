package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var EnableTimestamp = true // Default is true
var DisableWarn = false    // If true, disables Warn logs
var DisableDebug = true    // If true, disables Debug logs

// CustomFormatter formats logs with short levels and optional timestamps
type CustomFormatter struct {
	logrus.TextFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if EnableTimestamp {
		entry.Message = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + entry.Message
	}

	return f.TextFormatter.Format(entry)
}

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)

	log.SetFormatter(&CustomFormatter{
		TextFormatter: logrus.TextFormatter{
			ForceColors:      true,
			DisableQuote:     true,
			DisableTimestamp: true, // disable built-in one; we use custom
		},
	})
}

// === Exported Logger API ===

func Info(msg string) {
	log.Info(msg)
}

func Warn(msg string) {
	if DisableWarn {
		return
	}
	log.Warn(msg)
}

func Error(msg string) {
	log.Error(msg)
}

func Debug(msg string) {
	if DisableDebug {
		return
	}
	log.Debug(msg)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}
