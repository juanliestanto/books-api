package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	logChan chan string
}

func NewLogger() *Logger {
	return &Logger{logChan: make(chan string)}
}

func (logger *Logger) Start() {
	for mssg := range logger.logChan {
		fmt.Println(mssg)
	}
}

func (logger *Logger) Log(mssg string) {
	logger.logChan <- fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), mssg)
}
