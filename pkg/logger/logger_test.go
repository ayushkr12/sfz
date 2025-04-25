package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	DisableWarn = true
	EnableTimestamp = false
	Warn("Test warning message")
	Info("Test debug message")
}
