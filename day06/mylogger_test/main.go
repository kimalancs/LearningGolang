package main

import (
	"github.com/kimalancs/LearningGolang/day06/mylogger"
)



func main() {
	log := mylogger.NewLog("info")

	log.Debug("Debug log test")
	log.Trace("Trace log test")

	id := 1001
	name := "kim"
	log.Info("Info log test, id: %d, name: %s", id, name)
	
	log.Warn("Warn log test")
	log.Error("Error log test")
	log.Fatal("Error log test")

}