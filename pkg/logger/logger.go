package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var DisableWarn = false // If true, disables Warn logs
var DisableDebug = true // If true, disables Debug logs

// CustomFormatter formats logs with short levels and optional timestamps
type CustomFormatter struct {
	logrus.TextFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var EnableTimestamp bool

	if !DisableDebug {
		EnableTimestamp = true
	}

	if EnableTimestamp {
		entry.Message = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + entry.Message
	}

	entry.Message += "\n"

	return f.TextFormatter.Format(entry)
}

func init() {
	log = logrus.New()
	// Send all logs to stderr to avoid polluting stdout (e.g., when piping stdout to clip.exe)
	log.SetOutput(os.Stderr)
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
	os.Exit(1)
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
