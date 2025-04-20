package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var EnableTimestamp = true // Default is true

// CustomFormatter formats logs with short levels and optional timestamps
type CustomFormatter struct {
	logrus.TextFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Map level to short format
	level := strings.ToUpper(entry.Level.String())
	switch entry.Level {
	case logrus.InfoLevel:
		level = "INF"
	case logrus.WarnLevel:
		level = "WRN"
	case logrus.ErrorLevel:
		level = "ERR"
	case logrus.DebugLevel:
		level = "DBG"
	case logrus.TraceLevel:
		level = "TRC"
	case logrus.FatalLevel:
		level = "FTL"
	case logrus.PanicLevel:
		level = "PNC"
	}

	// Optional timestamp
	timestamp := ""
	if EnableTimestamp {
		timestamp = entry.Time.Format("2006-01-02 15:04:05") + " "
	}

	// Update message
	entry.Message = "[" + level + "] " + timestamp + entry.Message

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
	log.Warn(msg)
}

func Error(msg string) {
	log.Error(msg)
}

func Debug(msg string) {
	log.Debug(msg)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}
