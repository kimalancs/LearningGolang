package main

import (
	"fmt"
	"log"
	"os"
)

/*
日志库
使用接口的方式实现一个既可以向终端写日志也可以向文件写日志的简易日志库

Go内置log包实现了简单的日志服务，但功能有限，例如无法满足记录不同级别日志的情况，实际开发时根据需要选择使用第三方的日志库，如logrus、zap等
log包定义了Logger类型，提供一些格式化输出的方法


需求分析

日志要分等级
logger.Trace()
logger.Debug()
logger.Warning()
logger.Info()
logger.Error()
logger.Fatal()

支持开关控制，什么等级的日志要写入日志文件

支持往不同地方输出日志，比如打印到标准输出和打印到日志文件

完整的日志记录要包含时间、行号、文件名、日志级别、日志信息

日志文件要切割




*/

func main() {
	log.Println("test") // 打印到标准输出

	fileObj, err := os.OpenFile("./xxx.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	log.SetOutput(fileObj) // 设置日志写到文件中

	log.Println("test") // 设置SetOutput之后，日志会打印到日志文件

}
