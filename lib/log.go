package lib

import (
	"os"

	"github.com/op/go-logging"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// Logger represents a logger strut
type Logger struct {
	*logging.Logger
}

// GetLogger configure and returns logger struct
func GetLogger() (l *Logger) {
	var log = &Logger{
		Logger: logging.MustGetLogger("fly"),
	}
	var backend = logging.NewLogBackend(os.Stderr, "", 0)
	var formated = logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formated)
	return log
}
