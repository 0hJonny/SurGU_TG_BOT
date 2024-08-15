package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Level int

const (
	InfoLevel Level = iota
	WarningLevel
	PanicLevel
)

type Logger struct {
	Level Level
	Path  string
}

func (l *Logger) Info(msg string) {
	if l.Level <= InfoLevel {
		l.write(fmt.Sprintf("[INFO] %s %s", l.now(), msg))
	}
}

func (l *Logger) Warning(msg string) {
	if l.Level <= WarningLevel {
		l.write(fmt.Sprintf("[WARNING] %s %s", l.now(), msg))
	}
}

func (l *Logger) Panic(msg string) {
	if l.Level <= PanicLevel {
		l.write(fmt.Sprintf("[PANIC] %s %s", l.now(), msg))
		panic(msg)
	}
}

func (l *Logger) now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (l *Logger) write(msg string) {
	if l.Path == "" {
		l.Path = "./log"
	}

	if err := os.MkdirAll(l.Path, 0755); err != nil {
		fmt.Println("Can't create log dir:", err)
		return
	}

	filePath := filepath.Join(l.Path, "log.txt")

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Can't open log file:", err)
		return
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("Can't get log file info:", err)
		return
	}

	if fi.Size() > 4*1024*1024 {
		file.Close()
		newName := fmt.Sprintf("%s.%s", filePath, time.Now().Format("2006-01-02_15-04"))
		if err := os.Rename(filePath, newName); err != nil {
			fmt.Println("Can't rename log file:", err)
			return
		}
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Can't open log file:", err)
			return
		}
		defer file.Close()
	}

	if _, err := file.WriteString(fmt.Sprintf("%s\n", msg)); err != nil {
		fmt.Println("Can't write to log file:", err)
	}
}
