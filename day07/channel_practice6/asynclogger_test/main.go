package main

import (
	"github.com/kimalancs/LearningGolang/day07/channel_practice6/asynclogger"
)

// 异步记录方式的日志库

func main() {
		var log asynclogger.Logger
		log = asynclogger.NewFileLogger("warn", "./log", "xxx.log", 512*1024)

	for {
		log.Debug("Debug log to the file test")
		log.Trace("Trace log to the file test")
		log.Info("Info log to the file test")
		log.Warn("Warning log to the file test")
		id := 1111
		name := "allen"
		log.Error("Error log to the file test, id=%d, name=%s", id, name)
		log.Fatal("Error log to the file test")
	}
}
