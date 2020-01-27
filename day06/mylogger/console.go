package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// ConsoleLogger log
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLog Logger构造函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c *ConsoleLogger) log(lv LogLevel, format string, args ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, args...)
		now := time.Now()
		funcName, fileName, lineNumber := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNumber, msg)
	}
}

func (c *ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

// Debug debug log
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DEBUG, format, args...)
}

// Trace trace log
func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	c.log(TRACE, format, args...)
}

// Info Info log
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(INFO, format, args...)
}

// Warn Warn log
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WARN, format, args...)
}

// Error Error log
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ERROR, format, args...)
}

// Fatal fatal log
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FATAL, format, args...)
}
