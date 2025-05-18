package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	DisableWarn = true
	DisableDebug = false
	Warn("Test warning message")
	Debug("Test debug message")
	Info("Test debug message")
}
