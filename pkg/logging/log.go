package logging

import (
	"log"
	"os"
)

type Level int

var (
	F *os.File

	DefaultPrefix       = ""
	DefaultCallerDepath = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)
