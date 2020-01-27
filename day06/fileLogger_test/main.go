package main

import (
	"fmt"
	"os"
)

// 1. 文件对象的类型
// 2. 获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	fmt.Printf("%T\n", fileObj)     // 文件类型
	fileInfo, err := fileObj.Stat() // 文件信息
	if err != nil {
		fmt.Printf("get file info failed, %v\n", err)
		return
	}
	fmt.Printf("文件大小：%d\n", fileInfo.Size())

}
