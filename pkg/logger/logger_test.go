package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	DisableWarn = true
	Warn("Test warning message")
	Info("Test debug message")
}
