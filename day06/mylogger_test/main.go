package main

import (
	"github.com/kimalancs/LearningGolang/day06/mylogger"
)



func main() {
	log := mylogger.NewConsoleLog("info")

	log.Debug("Debug log test")
	log.Trace("Trace log test")

	id := 1001
	name := "kim"
	log.Info("Info log test, id: %d, name: %s", id, name)

	log.Warn("Warn log test")
	log.Error("Error log test")
	log.Fatal("Error log test")

	fileLog := mylogger.NewFileLogger("info", "./log","xxx.log", 10*1024*1024)

	fileLog.Debug("Debug log to the file test")
	fileLog.Info("Info log to the file test")
	
	id = 1111
	name = "allen"
	fileLog.Warn("Warning log to the file test, id=%d, name=%s", id, name)
	fileLog.Error("Error log to the file test, id=%d, name=%s", id, name)
	fileLog.Fatal("Error log to the file test")


}