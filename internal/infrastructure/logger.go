package infrastructure

import (
	"io"
	"log"
	"os"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func LogWriter(logPath string, format string, i ...interface{}) {

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%s", err)
	}

	defer file.Close()

	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Printf(format, i...)
}

func (l *Logger) LogError(format string, i ...interface{}) {
	LogWriter("./internal/log/error.log", format, i)
}

func (l *Logger) LogAccess(format string, i ...interface{}) {
	LogWriter("./internal/log/access.log", format, i)
}
