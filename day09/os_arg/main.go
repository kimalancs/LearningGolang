package main

import (
	"fmt"
	"os"
)

//os.Arg
// os.Args是一个存储命令行参数的字符串切片
// 它的第一个元素是执行文件的名称
// 写简单脚本工具时，可以获取命令行执行程序时输入的参数


func main() {
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}