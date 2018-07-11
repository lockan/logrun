package logrun

import (
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func getCallerInfo() log.Fields {
	// stack frame offset is 2 because:
	// offset 0 is this function,
	// offset 1 is the calling log wrapper function in this file
	// --> 2 is the consuming file/function that is calling the Log functions.
	_, fl, ln, ok := runtime.Caller(2)
	fields := log.Fields{}
	if !ok {
		fields = log.Fields{"caller": "unknown", "line": "n/a"}
	}
	fields = log.Fields{"caller": fl, "line": ln}
	return fields
}

// EnableJSON - toggles json logging format on
func EnableJSON() {
	log.SetFormatter(&log.JSONFormatter{})
}

// SetLevel - sets the minimum logging level
func SetLevel(level string) {

	switch strings.ToLower(level) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

//Debug - debug log trace
func Debug(msg interface{}) {
	fields := getCallerInfo()
	log.WithFields(log.Fields(fields)).Debug(msg)
}

//Info - info log trace
func Info(msg interface{}) {
	fields := getCallerInfo()
	log.WithFields(log.Fields(fields)).Info(msg)
}

// Warn - warn log trace
func Warn(msg interface{}) {
	fields := getCallerInfo()
	log.WithFields(log.Fields(fields)).Warn(msg)
}

//Error - error log trace
func Error(msg interface{}) {
	fields := getCallerInfo()
	log.WithFields(log.Fields(fields)).Error(msg)
}

// Fatal - fatal log trace
func Fatal(msg interface{}) {
	fields := getCallerInfo()
	log.WithFields(log.Fields(fields)).Fatal(msg)
}

// Panic - panic log trace, also causes panic
func Panic(msg interface{}) {
	fields := getCallerInfo()
	log.WithFields(log.Fields(fields)).Panic(msg)
}

// TestLogging - for use in unit tests, but ultimately not that useful here.
func TestLogging() {
	Debug("Test Debug")
	Info("Test Info")
	Warn("Test Warn")
	Error("Test Error")
	// We can't really test fatal or panic now can we?
}
