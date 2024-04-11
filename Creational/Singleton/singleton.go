package main

import (
	"fmt"
	"sync"
)

type MyLogger struct {
	loglevel int
}

func (logger *MyLogger) Log(s string) {
	fmt.Println(logger.loglevel, ":", s)
}

func (logger *MyLogger) SetLogLevel(level int) {
	logger.loglevel = level
}

var logger *MyLogger

var once sync.Once

func getLoggerInstance() *MyLogger {
	once.Do(func() {
		fmt.Println("Creating new MyLogger instance")
		logger = &MyLogger{}
	})

	fmt.Println("Returning MyLogger instance")
	return logger
}
