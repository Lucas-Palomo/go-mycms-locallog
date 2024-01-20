package pkg

import (
	"fmt"
	"log"
	"os"
	"time"
)

type logLevel int

const (
	INIT    logLevel = iota
	INFO    logLevel = iota
	WARNING logLevel = iota
	ERROR   logLevel = iota
	FATAL   logLevel = iota
	PANIC   logLevel = iota
)

func (level logLevel) Name() string {
	levels := []string{
		"INFO",
		"WARNING",
		"ERROR",
		"FATAL",
		"PANIC",
	}

	return levels[level]
}

type LogFile struct {
	File *os.File
}

// Write a log based on levels, allowed log types are:
//
//   - INFO
//
//   - WARNING
//
//   - ERROR
//
//   - FATAL
//
//   - PANIC
//
//   - The log level FATAL is used to log and call the os.Exit(1) method
//
//   - The log level PANIC is used to log and call the panic() method
func (lf LogFile) Write(level logLevel, content any) {

	line := fmt.Sprintf("%s [%s] %+v\n", time.Now().Format(time.UnixDate), level.Name(), content)
	if _, err := lf.File.WriteString(line); err != nil {
		panic(err)
	}

	if level == PANIC {
		log.Panicln(content)
	}

	if level == FATAL {
		log.Fatalln(content)
	}
}

func New(name string) LogFile {
	filename := fmt.Sprintf("%s.log", name)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	return LogFile{File: file}
}
