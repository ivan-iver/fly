package lib

import (
	"os"

	"github.com/op/go-logging"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message} | %{shortfunc}`,
)

// Logger represents a logger strut
type Logger struct {
	*logging.Logger
	filename string
}

// GetLogger configure and returns logger struct
func GetLogger(logfile string) (log *Logger, err error) {
	log = &Logger{
		Logger:   logging.MustGetLogger("fly"),
		filename: logfile,
	}
	var backend = logging.NewLogBackend(os.Stderr, "", 0)
	var formated = logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formated)
	//err = setOutput(log)
	return
}

func setOutput(log *Logger) (err error) {
	f, err := os.OpenFile(log.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	//log.SetOutput(f)
	//log.Println("This is a test log entry")
	return
}
