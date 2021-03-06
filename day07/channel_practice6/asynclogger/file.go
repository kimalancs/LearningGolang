package asynclogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

var (
	// MaxSize 日志通道缓冲区的容量大小
	MaxSize int = 50000
)

// 往log文件写入日志

// FileLogger log to the file
type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// NewFileLogger FileLogger构造函数
func NewFileLogger(levelStr, fp, fn string, maxFileSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxFileSize,
		logChan:     make(chan *logMsg, MaxSize), // 初始化logChan
	}
	err = f1.initFile() // 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return f1
}

// 创建文件类型实例时，初始化，打开实例
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err: %v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err: %v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	// 开启一个后台的goroutine去写日志
	go f.asynclogging()
	return nil
}

// Close close the log file 关闭日志文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// 切割文件
// 1. 按日期切割日志文件
// 每次计入日志前，判断当前写的文件的大小，超过设定的最大值，就关闭当前的文件，重命名后，再打开一个新的日志文件继续写入
// 2. 按照日志文件大小切割

// 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat() // 文件信息
	if err != nil {
		fmt.Printf("get file info failed, %v\n", err)
		return false
	}
	// 如果当前文件大小大于等于日志文件的最大值，就返回true
	return fileInfo.Size() >= f.maxFileSize
}

// 切割日志文件
func (f *FileLogger) spFile(file *os.File) (*os.File, error) {
	//切割日志
	// 获取文件信息
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, %v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak.%s", logName, nowStr) // 拼接一个备份日志文件文件名
	// 关闭日志文件
	file.Close()
	// 重命名，备份
	os.Rename(logName, newLogName)
	// 打开一个新日志文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return nil, err
	}
	// 4. 将打开的新文件赋值给f.fileObj
	return fileObj, nil
}

//
func (f *FileLogger) asynclogging() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.spFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.spFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
				// 如果日志级别大于等于ERROR，在err文件中再记录一遍
			}
		default:
			// 取不到日志，先休息500ms，让出CPU资源
			time.Sleep(500 * time.Millisecond)
		}

	}
}

// 记录日志到日志文件
func (f *FileLogger) log(lv LogLevel, format string, args ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, args...)
		now := time.Now()
		funcName, fileName, lineNumber := getInfo(3)
		// 先把日志发送到通道中
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line:      lineNumber,
		}
		select {
		case f.logChan <- logTmp:
		default: // 如果通道写满了，而且从通道中接收日志写入文件的goroutine挂掉了，就会出现阻塞，通过select保证日志无法写入时直接扔掉，保证业务代码顺畅执行
		}
	}
}

// 判断是否需要记录日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// Debug debug log
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DEBUG, format, args...)
}

// Trace trace log
func (f *FileLogger) Trace(format string, args ...interface{}) {
	f.log(TRACE, format, args...)
}

// Info Info log
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(INFO, format, args...)
}

// Warn Warn log
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WARN, format, args...)
}

// Error Error log
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ERROR, format, args...)
}

// Fatal fatal log
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FATAL, format, args...)
}
