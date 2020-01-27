package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

func log(lv LogLevel, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	now := time.Now()
	funcName, fileName, lineNumber := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNumber, msg)
}


// Debug debug log
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.enable(DEBUG) {
		log(DEBUG, format, args...)
	}
}

// Trace trace log
func (l *Logger) Trace(format string, args ...interface{}) {
	if l.enable(TRACE) {
		log(TRACE, format, args...)
	}
}

// Info Info log
func (l *Logger) Info(format string, args ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format, args...)
	}
}

// Warn Warn log
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.enable(WARN) {
		log(WARN, format, args...)
	}
}

// Error Error log
func (l *Logger) Error(format string, args ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, args...)
	}
}

// Fatal fatal log
func (l *Logger) Fatal(format string, args ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL, format, args...)
	}
}
