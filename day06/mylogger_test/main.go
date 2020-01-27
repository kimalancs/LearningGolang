package main

import (
	"github.com/kimalancs/LearningGolang/day06/mylogger"
)

var log mylogger.Logger
// ConsoleLogger类型是值接收者实现Logger接口时，才可以直接把NewConsoleLogger构造函数的结果赋值给Logger接口变量
// 因为NewConsoleLogger构建的结果就是ConsoleLogger类型，只有值接收者实现接口，才可以把ConsoleLogger类型赋值给Logger接口变量
// 如果改为指针接收者实现Logger接口，使用的时候就要把*ConsoleLogger赋值给Logger接口变量，要写成log= &mylogger.NewConsoleLogger("info")

// FileLogger类型是指针接收者实现Logger接口，但是由于NewFileLogger构造函数返回的就是*FileLogger，所以其结果可以直接赋值给Logger接口变量

func main() {
	log = mylogger.NewConsoleLogger("info") // 构建结果就是ConsoleLogger类型

	log.Debug("Debug log test")
	log.Trace("Trace log test")

	id := 1001
	name := "kim"
	log.Info("Info log test, id: %d, name: %s", id, name)

	log.Warn("Warn log test")
	log.Error("Error log test")
	log.Fatal("Error log test")

	log = mylogger.NewFileLogger("warn", "./log", "xxx.log", 512*1024) // 构造结果是*FileLogger类型

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
