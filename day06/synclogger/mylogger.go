package synclogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// Logger logger interface 日志接口
type Logger interface {
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

// LogLevel log level 日志级别类型
type LogLevel uint64

// different level of log messages 具体日志级别
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARN
	ERROR
	FATAL
)

// 字符串转为日志级别
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}

}

// 日志级别转为字符串
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case UNKNOWN:
		return "UNKNOWN"
	}
	return " "
}

// 获取报错信息的文件名、函数名、报错行数
func getInfo(n int) (funcName string, fileName string, lineNumber int) {
	pc, file, lineNumber, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
